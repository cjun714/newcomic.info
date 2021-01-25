// +build !tool

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"newcomic.info/db"
	"newcomic.info/log"
	"newcomic.info/model"
)

func main() {
	if e := db.Init(); e != nil {
		panic(e)
	}

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

	doc.Find(".newcomic-short").Each(func(i int, s *goquery.Selection) {
		var mi model.ComicInfo

		// read pages + size
		fileInfo := s.Find(".newcomic-mask-top").Clone().Children().Remove().End().Text()
		fileInfo = strings.ReplaceAll(fileInfo, "\n", "")
		fileInfo = strings.Trim(fileInfo, "	")
		log.I(fileInfo)

		// read tags
		s.Find(".newcomic-mask-top a").Each(func(i int, s *goquery.Selection) {
			tag := s.Text()
			mi.Tags = mi.Tags + "|" + tag
		})

		s.Find(".newcomic-mask-bottom a").Each(func(i int, s *goquery.Selection) {
			title, _ := s.Attr("title")
			url, _ := s.Attr("href")
			mi.Name = title

			loadDetailPage(url)
		})

		// img
		imgURL, _ := s.Find("img").Attr("src")
		mi.CoverURL = imgURL

		infoList = append(infoList, mi)
	})

	return infoList, e
}

func loadDetailPage(path string) {

}
