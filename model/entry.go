package model

import (
	"diary_api/util/database"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserId  uint
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

func FindEntryById(id int64) (Entry, error) {
	var entry Entry
	err := database.Database.Where("id=?", id).Find(&entry).Error
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}
