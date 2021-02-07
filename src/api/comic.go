package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"newcomic.info/db"
	"newcomic.info/log"
	"newcomic.info/model"
)

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
	log.I(info.Name)

	return c.JSON(http.StatusOK, info)
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
