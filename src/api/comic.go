package api

import (
	"encoding/base64"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/cjun714/glog/log"
	"github.com/gen2brain/go-unarr"
	"github.com/labstack/echo"

	"newcomic.info/db"
	"newcomic.info/model"
)

const comicPath = "z:/comic"

var namePathMap = make(map[string]string, 2000)

func init() {
	log.I("read downloaded comic list:", comicPath)

	_, e := os.Stat(comicPath)
	if e != nil {
		log.E("read list failed:", e)
		return
	}

	e = filepath.Walk(comicPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil // skip dir
		}
		if !isComic(path) {
			return nil // skip non-comic
		}

		namePathMap[filepath.Base(path)] = path
		return nil
	})

	if e != nil {
		log.E(e)
	}
}

func isComic(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".cbt" || ext == ".cbz" || ext == ".cbr"
}

func GetComicInfos(c echo.Context) error {
	pageSize := 40

	page := 0
	var e error

	pageStr := c.Param("page")
	if pageStr != "" {
		if page, e = strconv.Atoi(pageStr); e != nil || page < 1 {
			return c.JSON(http.StatusBadRequest, nil)
		}
	}

	search := c.QueryParams().Get("search")
	search, _ = url.QueryUnescape(search)
	log.I(search)

	list, e := db.GetComicInfoList((page-1)*pageSize, pageSize, search)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	count, e := db.GetComicCount(search)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"count": count, "data": list})
}

func GetComicDetail(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	info := &model.ComicInfo{}
	if e := db.Get(id, info); e != nil {
		log.E(e)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, info)
}

func GetComicSamples(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	info := &model.ComicInfo{}
	if e := db.Get(id, info); e != nil {
		log.E(e)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	log.I("read samples:", info.Name)

	path := getComicPath(info.Name, info.Year)
	samples, e := readComicSamples(path)
	if e != nil {
		log.E(e)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, samples)
}

func AddFavorite(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if e := db.AddFavorite(id); e != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func DeleteFavorite(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if e := db.DeleteFavorite(id); e != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func AddDownload(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if e := db.AddDownload(id); e != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func DeleteDownload(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if e := db.DeleteDownload(id); e != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func readComicSamples(path string) ([]string, error) {
	// ar, e := unarr.NewArchive(path)
	ar, e := unarr.NewArchive("z:/mera.cbt")
	if e != nil {
		return nil, e
	}
	defer ar.Close()

	li, e := ar.List()
	if e != nil {
		return nil, e
	}
	sort.Strings(li)
	ret := make([]string, 15)
	for i := 0; i < 15; i++ {
		log.I(li[i])
		e = ar.EntryFor(li[i])
		if e != nil {
			return nil, e
		}
		byts, e := ar.ReadAll()
		if e != nil {
			return nil, e
		}
		ret[i] = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(byts)
	}
	return ret, nil
}

func getComicPath(name string, year int) string {
	yearStr := strconv.Itoa(year)
	// name="Batman #2"
	// comicName = "Batman #2 (2014) (Digital) (xx.com).cbt"
	// comicName = "Batman #02 (2014) (Digital) (xx.com).cbt"
	for comicName, path := range namePathMap {
		idx := strings.Index(comicName, name)
		if idx == -1 {
			continue
		}

		idx = strings.Index(comicName, yearStr)
		if idx == -1 {
			continue
		}

		return path
	}

	return ""
}
