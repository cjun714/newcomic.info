// +build !tool

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cjun714/glog/log"

	"newcomic.info/db"
	"newcomic.info/model"
)

func main() {
	if e := db.Init(); e != nil {
		panic(e)
	}
	defer db.Close()

	e := loadPages(os.Args[1])
	if e != nil {
		panic(e)
	}
}

func loadPages(src string) error {
	e := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		infos, e := parseIndexPage(path)
		if e != nil {
			return e
		}

		for _, info := range infos {
			if e := db.Save(&info); e != nil {
				log.E(info.Name, e)
			}
		}

		return nil
	})

	if e != nil {
		return e
	}

	return nil
}

func parseIndexPage(path string) ([]model.ComicInfo, error) {
	bs, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	r := bytes.NewReader(bs)
	doc, e := goquery.NewDocumentFromReader(r)
	if e != nil {
		return nil, e
	}

	infoList := make([]model.ComicInfo, 0, 1)

	dir := filepath.Dir(path)
	doc.Find(".newcomic-short").Each(func(i int, s *goquery.Selection) {
		var mi model.ComicInfo

		// read pages + size
		fileInfo := s.Find(".newcomic-mask-top").Clone().Children().Remove().End().Text()
		fileInfo = strings.ReplaceAll(fileInfo, "\n", "")
		fileInfo = strings.Trim(fileInfo, "	")
		// log.I(fileInfo)

		// read tags
		s.Find(".newcomic-mask-top a").Each(func(i int, s *goquery.Selection) {
			tag := s.Text()
			mi.Tags = mi.Tags + "|" + tag
		})

		s.Find(".newcomic-mask-bottom a").Each(func(i int, s *goquery.Selection) {
			mi.Name, _ = s.Attr("title")

			url, _ := s.Attr("href")
			dpath := dir + "/pages/" + filepath.Base(url)

			d, e := loadDetailPage(dpath)
			if e != nil {
				log.E(e)
				return
			}
			mi.PageURL = filepath.Base(url)
			mi.Pages = d.pages
			mi.Year = d.year
			mi.Size = d.size
			mi.Publisher = d.publisher
			mi.DownloadURL = d.downloadURL
		})

		// img
		imgURL, _ := s.Find("img").Attr("src")
		mi.Cover = filepath.Base(imgURL)
		mi.Cover = strings.Replace(mi.Cover, ".jpg", ".webp", -1)
		mi.Cover = strings.Replace(mi.Cover, ".jpeg", ".webp", -1)
		mi.Cover = strings.Replace(mi.Cover, ".png", ".webp", -1)

		infoList = append(infoList, mi)
	})

	return infoList, e
}

type comicDetail struct {
	downloadURL string
	publisher   string
	pages       int
	year        int
	size        int
}

func loadDetailPage(path string) (*comicDetail, error) {
	bs, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	r := bytes.NewReader(bs)
	doc, e := goquery.NewDocumentFromReader(r)
	if e != nil {
		return nil, e
	}

	detail := &comicDetail{}
	doc.Find(".newcomic-full-text li").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		switch i {
		case 0: // DC publisher
			detail.publisher = getPublisher(text)
		case 1: // Pages: 26
			detail.pages = getPages(text)
		case 2: // 2019 year
			detail.year = getYear(text)
		case 3: // English comics
		case 4: // Size: 18.6 mb.
			detail.size = getSize(text)
		case 5: // Tags: Whos Who G-Man
		default:
			// skip
		}
	})

	doc.Find(".newcomic-m-buttons a").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			url, _ := s.Attr("href")
			bname := filepath.Base(url)
			// handle name like 'https://florenfile.com/qsksoscx6ru9/somebook.cbr.html'
			if strings.HasSuffix(url, ".html") {
				url = strings.Trim(url, bname)
				detail.downloadURL = filepath.Base(url)
			} else {
				detail.downloadURL = bname
			}
		}
	})

	return detail, nil
}

func getPublisher(text string) string {
	s := strings.Split(text, " ")
	if len(s) > 0 {
		return s[0]
	}
	return ""
}

func getPages(text string) int {
	s := strings.Split(text, " ")
	if len(s) == 0 {
		return 0
	}
	i, e := strconv.Atoi(s[1])
	if e != nil {
		return 0
	}
	return i
}

func getYear(text string) int {
	s := strings.Split(text, " ")
	if len(s) == 0 {
		return 1900
	}
	i, e := strconv.Atoi(s[0])
	if e != nil {
		return 1900
	}
	return i
}

func getSize(text string) int {
	s := strings.Split(text, " ")
	if len(s) < 2 {
		return 0
	}
	f, e := strconv.ParseFloat(s[1], 32)
	if e != nil {
		return 0
	}
	return int(f)
}

func main2() {
	loadDetailPage("z:/test.html")
}
