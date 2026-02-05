package models

import "time"

type UserCourse struct {
	ID            	uint      	`json:"id" gorm:"primaryKey"`
	CourseID      	uint      	`json:"course_id" gorm:"index;not null"`
	CreatorID     	uint      	`json:"creator_id" gorm:"index;not null"`
	ParticipantID 	uint      	`json:"participant_id" gorm:"index;not null"`

	Course      	Course		`gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	Creator     	User   		`gorm:"foreignKey:CreatorID"`
	Participant 	User   		`gorm:"foreignKey:ParticipantID"`

	CreatedAt     	time.Time 	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     	time.Time 	`json:"updated_at" gorm:"autoUpdateTime"`
}