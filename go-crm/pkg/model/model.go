package model

import (
	"github.com/jangirvipin/go-crm/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Company string `json:"company"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()

	if err := db.AutoMigrate(&Lead{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

func GetLeads() ([]Lead, error) {
	var leads []Lead
	if err := db.Find(&leads).Error; err != nil {
		return nil, err
	}
	return leads, nil
}

func GetLead(id int64) (*Lead, error) {
	if id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var lead Lead
	result := db.First(&lead, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lead, nil
}

func (lead *Lead) Create() *Lead {
	db.Create(&lead)
	return lead
}

func DeleteLead(id int64) (*Lead, error) {
	if id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var lead Lead
	result := db.First(&lead, id)
	if result.Error != nil {
		return nil, result.Error
	}
	db.Delete(&lead)
	return &lead, nil
}

func UpdateLead(lead *Lead) error {
	result := db.Save(lead)
	return result.Error
}
