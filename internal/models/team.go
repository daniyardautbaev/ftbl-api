package models

import "gorm.io/gorm"

type Team struct {
    gorm.Model
    Name    string   `gorm:"type:varchar(100);not null" json:"name"`
    City    string   `gorm:"type:varchar(100)" json:"city"`
    Coach   string   `gorm:"type:varchar(100)" json:"coach"`
    Players []Player `json:"players,omitempty"`
}
