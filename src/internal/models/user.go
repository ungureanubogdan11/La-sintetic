package models

import "time"

type Role string

const (
	RoleMember Role = "member"
	RoleAdmin  Role = "admin"
)

type Position string

const (
	PositionGoalkeeper Position = "goalkeeper"
	PositionField      Position = "field"
)

type User struct {
	ID                uint64 `gorm:"primaryKey"`
	Name              string `gorm:"not null"`
	Email             string `gorm:"not null"`
	PasswordHash      string `gorm:"not null"`
	Role              Role   `gorm:"type:user_role;not null;default:member"`
	Bio               *string
	AvatarURL         *string
	PreferredPosition *Position `gorm:"type:player_position"`
	IsPublic          bool      `gorm:"not null;default:true"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
