package models

import "time"

type Course struct {
	ID          	uint         `json:"id" gorm:"primaryKey"`
	Title       	string       `json:"title" gorm:"not null;size:255"`
	Description		string       `json:"description" gorm:"type:text"`
	Price	       	float64      `json:"price" gorm:"not null;default:0"`

	CreatorID   	uint         `json:"creator_id" gorm:"index;not null"`
	Creator     	User         `gorm:"foreignKey:CreatorID"`
	Participants 	[]UserCourse `gorm:"foreignKey:CourseID"`

	CreatedAt   	time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   	time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}