package repository

import (
	"gorm.io/gorm"
)

type ContentRepo struct {
	db *gorm.DB
}

type Content struct {
	ID         string `gorm:"primaryKey"`
	Writer     string
	WriteScope string
	ViewScope  string
	gorm.Model
}

func (cr *ContentRepo) GetAll() ([]*Content, error) {
	contents := []*Content{}
	res := cr.db.Find(&contents)
	return contents, res.Error
}

func (cr *ContentRepo) Get() {

}

func (cr *ContentRepo) GetListByWriter(ID string) {

}

func (cr *ContentRepo) Delete(ID string) {

}

func (cr *ContentRepo) Update() {

}
