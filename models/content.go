package models

import (
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	PageName string `json:"page_name" gorm:"uniqueIndex:idx_page_section_key"` // e.g., "home", "about", "services"
	Section  string `json:"section" gorm:"uniqueIndex:idx_page_section_key"`                 // e.g., "hero", "skills", "experience"
	Key      string `json:"key" gorm:"uniqueIndex:idx_page_section_key"`                     // e.g., "title", "description", "image_url"
	Value    string `json:"value"`                   // The actual content
}