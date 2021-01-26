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
	info := &model.ComicInfo{ID: id}
	count := 0
	e := dbs.Model(info).Update("favorite", true).Count(&count).Error

	if count == 0 {
		return errors.New("no record with id:" + strconv.Itoa(id))
	}

	return e
}

func DeleteFavorite(id int) error {
	info := &model.ComicInfo{ID: id}
	count := 0
	e := dbs.Model(info).Update("favorite", false).Count(&count).Error

	if count == 0 {
		return errors.New("no record with id:" + strconv.Itoa(id))
	}

	return e
}
