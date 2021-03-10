// +build !tool

package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/cjun714/glog/log"
	"newcomic.info/db"
)

func main() {
	e := db.Init()
	if e != nil {
		panic(e)
	}
	defer db.Close()

	list, e := readlist(os.Args[1])
	if e != nil {
		panic(e)
	}

	var wg sync.WaitGroup
	for _, str := range list {
		set, e := db.SearchComic("cover", str)
		if e != nil {
			panic(e)
		}
		if len(set) == 0 {
			log.I("can't find records:", str)
			continue
		}
		htmlpath := "z:/pages/" + set[0].PageURL
		url, e := getImageURL(htmlpath)
		if e != nil {
			panic(e)
		}
		if strings.HasPrefix(url, "/") {
			url = "https://newcomic.info/" + url
		}
		log.I(url)
		dir := "z:/"
		targetPath := dir + filepath.Base(url)
		wg.Add(1)
		go func(src, target string) {
			defer wg.Done()
			e = downloadImage(url, targetPath)
			if e != nil {
				log.E(e)
			}
		}(url, targetPath)
	}
	wg.Wait()
}

func readlist(path string) ([]string, error) {
	f, e := os.Open(path)
	if e != nil {
		log.E("failed opening file: %s", e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var strs []string

	for scanner.Scan() {
		str := scanner.Text()
		str = strings.Replace(str, ".jpg", ".webp", -1)
		str = strings.Replace(str, ".png", ".webp", -1)
		strs = append(strs, str)
	}

	return strs, nil
}

func getImageURL(path string) (string, error) {
	byts, e := ioutil.ReadFile(path)
	if e != nil {
		return "", e
	}
	r := bytes.NewReader(byts)
	doc, e := goquery.NewDocumentFromReader(r)
	if e != nil {
		return "", e
	}

	url, _ := doc.Find(".newcomic-m-img img").First().Attr("src")
	return url, nil
}

func downloadImage(url string, targetPath string) error {
	resp, e := http.Get(url)
	if e != nil {
		return e
	}
	defer resp.Body.Close()

	f, e := os.Create(targetPath)
	if e != nil {
		return e
	}
	defer f.Close()

	_, e = io.Copy(f, resp.Body)
	if e != nil {
		return e
	}

	return nil
}
