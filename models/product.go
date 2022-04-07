package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}
