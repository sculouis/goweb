package pojo

import "time"

type Todo struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"Name"`
	Done      bool   `json:"Done"`
	Selected  bool   `json:"Selected"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
