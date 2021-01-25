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
	list, e := db.GetComicInfoList(1, 20)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, list)
}

func GetComicDetail(c echo.Context) error {
	log.I("sdfsdsdf")
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
