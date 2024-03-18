package model

import (
	"gorm.io/gorm"
	"time"
)

type Actor struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender" validate:"omitempty,oneof=male female"`
	Birthday time.Time `json:"birthday"`
	Films    []Film    `gorm:"many2many:film_actors;" json:"films"`
}

type Film struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"` // sort and search
	Description string    `json:"description"`
	Release     time.Time `json:"release"` // sort
	Rating      int       `json:"rating"`  // sort(default)
	Actors      []Actor   `gorm:"many2many:film_actors;" json:"actors"`
}

type CreateFilmDTO struct {
	Name        string    `json:"name"  validate:"gte=1,lte=150"` // sort and search
	Description string    `json:"description" validate:"lte=1000"`
	Release     time.Time `json:"release"`                        // sort
	Rating      int       `json:"rating" validate:"gte=0,lte=10"` // sort(default)
	Actors      []int     `json:"actors"`
}

type UpdateFilmDTO struct {
	ID          uint
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Release     time.Time `json:"release"`
	Rating      int       `json:"rating"`
	Actors      []int     `json:"actors"`
}

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `validate:"oneof=user admin"`
}
