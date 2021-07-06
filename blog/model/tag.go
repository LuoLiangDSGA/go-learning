package model

import "gorm.io/gorm"

type Tag struct {
	Model    Model
	Name     string `json:name`
	CreateBy string `json:create_by`
	ModifyBy string `json:modify_by`
	state    int    `json:state`
}

func GetTags(current int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		err  error
		tags []Tag
	)
	if pageSize > 0 && current > 0 {
		err = sqlDB.Where(maps).Offset(current).Limit(pageSize).Find(tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTagTotal(maps interface{}) (count int64) {
	sqlDB.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func GetTag(id int) (Tag, error) {
	var (
		tag Tag
		err error
	)
	if err = sqlDB.FirstOrInit(&tag, map[string]interface{}{"id": id}).Error; err != nil {
		return tag, err
	}
	return tag, nil
}

func DeleteTag(id int) error {
	if err := sqlDB.Delete(&Tag{}, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTag(id int, data interface{}) error {
	if err := sqlDB.Model(&Tag{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
