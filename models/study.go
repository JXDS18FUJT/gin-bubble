package models

import "bubble/dao"

// Study Model
type Study struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateStudy(study *Study) (err error) {
	err = dao.DB.Create(&study).Error
	return
}

func GetAllStudy() (studyList []*Study, err error) {
	if err = dao.DB.Find(&studyList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAStudy(id string) (study *Study, err error) {
	study = new(Study)
	if err = dao.DB.Debug().Where("id=?", id).First(study).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAStudy(study *Study) (err error) {
	err = dao.DB.Save(study).Error
	return
}

func DeleteAStudy(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Study{}).Error
	return
}
