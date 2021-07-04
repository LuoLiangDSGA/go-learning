package model

type Tag struct {
	Model
	Name string `json:name`
	CreateBy string `json:create_by`
	ModifyBy string `json:modify_by`
	state int `json:state`
}

func GetTags(current int, pageSize int, maps interface{}) (tags []Tag){
	sqlDB.Where(maps).Offset(current).Limit(pageSize).Find(tags)
	return
}

func GetTagTotal(maps interface {}) (count int64){
	sqlDB.Model(&Tag{}).Where(maps).Count(&count)

	return
}
