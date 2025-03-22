package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	GroupId  uuid.UUID `gorm:"type:uuid" json:"group_id"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

type LeaderBoard struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	Name      string    `json:"name" `
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

func (leaderBoard *LeaderBoard) BeforeCreate(tx *gorm.DB) (err error) {
	leaderBoard.ID = uuid.New()
	return
}

type Game struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid"`
	Name          string    `json:"name"`
	LeaderBoardID uuid.UUID `gorm:"type:uuid" json:"leaderBoardID"`
}

func (game *Game) BeforeCreate(tx *gorm.DB) (err error) {
	game.ID = uuid.New()
	return
}

type Group struct {
	ID   uuid.UUID `gorm:"type:uuid"`
	Name string    `json:"name"`
	Icon string    `json:"icon"`
}

func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.ID = uuid.New()
	return
}
