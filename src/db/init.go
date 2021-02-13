package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"newcomic.info/model"
)

var dbs *gorm.DB

type mod struct {
	Table   interface{}
	InitSQL string
}

func Init() error {
	mods := []mod{
		{model.ComicInfo{}, ""},
	}

	var e error
	if e = openDB("sqlite3", "./data.db"); e != nil {
		return e
	}

	if e = initTables(mods); e != nil {
		return e
	}

	return nil
}

func openDB(db, url string) error {
	var e error
	if dbs, e = gorm.Open(db, url); e != nil {
		return e
	}

	dbs.Exec("PRAGMA foreign_keys = ON") // only for Sqlite3

	dbs.LogMode(true)
	dbs.DB().SetMaxIdleConns(10)
	dbs.DB().SetMaxOpenConns(100)
	dbs.SingularTable(true)

	return e
}

func initTables(mods []mod) error {
	for _, m := range mods {
		if !dbs.HasTable(m.Table) {
			if e := dbs.CreateTable(m.Table).Error; e != nil {
				return e
			}
		} else {
			if e := dbs.AutoMigrate(m.Table).Error; e != nil {
				return e
			}
		}
	}

	return nil
}

func Close() {
	if dbs == nil {
		return
	}

	if e := dbs.Close(); e != nil {
		log.Fatal("close db error:", e)
	}
}
