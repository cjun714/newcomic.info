package db

import (
	"errors"
	"strconv"

	"newcomic.info/model"
)

func GetComicInfoList(offset, limit int, search string) ([]model.ComicInfo, error) {
	var list []model.ComicInfo
	var e error
	if search == "" {
		e = dbs.Limit(limit).Offset(offset).Find(&list).Error
	} else {
		search = "%" + search + "%"
		e = dbs.Where("name like ?", search).Limit(limit).Offset(offset).Find(&list).Error
	}

	return list, e
}

func GetComicCount(search string) (int, error) {
	count := 0
	var e error
	if search == "" {
		e = dbs.Model(&model.ComicInfo{}).Count(&count).Error
	} else {
		search = "%" + search + "%"
		e = dbs.Model(&model.ComicInfo{}).Where("name like ?", search).Count(&count).Error
	}
	return count, e
}

func AddFavorite(id int) error {
	return updateComicInfo(id, "favorite", true)
}

func DeleteFavorite(id int) error {
	return updateComicInfo(id, "favorite", false)
}

func AddDownload(id int) error {
	return updateComicInfo(id, "download", true)
}

func DeleteDownload(id int) error {
	return updateComicInfo(id, "download", false)
}

func updateComicInfo(id int, field string, val interface{}) error {
	info := &model.ComicInfo{ID: id}
	count := 0
	e := dbs.Model(info).Update(field, val).Count(&count).Error

	if count == 0 {
		return errors.New("no record with id:" + strconv.Itoa(id))
	}

	return e
}

func SearchComic(field string, value interface{}) ([]model.ComicInfo, error) {
	var list []model.ComicInfo
	e := dbs.Where(field+"= ?", value).Find(&list).Error
	return list, e
}
