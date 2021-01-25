package db

func Save(obj interface{}) error {
	return dbs.Save(obj).Error
}

func Get(id int, obj interface{}) error {
	return dbs.First(obj, id).Error
}

func Update(obj interface{}) error {
	return dbs.Update(obj).Error
}

func UpdateField(obj interface{}, field string, value interface{}) error {
	return dbs.Model(obj).Update(field, value).Error
}

func Delete(obj interface{}) error {
	return dbs.Delete(obj).Error
}

func GetCount(obj interface{}) (int, error) {
	var count int
	e := dbs.Model(obj).Count(&count).Error

	return count, e
}
