package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Timestamp time.Time
	Title     string
	Slug      string
	Category  string
	Excerpt   string
}

type DataResult struct {
	Total    int    `json:"recordsTotal"`
	Filtered int    `json:"recordsFiltered"`
	Data     []Post `json:"data"`
}
