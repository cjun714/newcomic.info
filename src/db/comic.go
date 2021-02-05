package db

import (
	"errors"
	"strconv"

	"newcomic.info/model"
)

func GetComicInfoList(offset, limit int) ([]model.ComicInfo, error) {
	var list []model.ComicInfo
	e := dbs.Limit(limit).Offset(offset).Find(&list).Error

	return list, e
}

func GetComicCount() (int, error) {
	count := 0
	e := dbs.Model(&model.ComicInfo{}).Count(&count).Error
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
