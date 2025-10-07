package models

import "gorm.io/gorm"

type Player struct {
    gorm.Model
    Name     string `gorm:"type:varchar(100);not null" json:"name"`
    Position string `gorm:"type:varchar(50)" json:"position"`
    Goals    int    `json:"goals"`
    Assists  int    `json:"assists"`
    Matches  int    `json:"matches"`
    TeamID   uint   `json:"team_id"`
}
