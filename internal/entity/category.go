package entity

import "time"

type Category struct {
	ID        int       `gorm:"column=id;primary_key"`
	Name      string    `gorm:"column=name"`
	CreatedAt time.Time `gorm:"column=created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column=updated_at;autoCreateTime;autoUpdateTime"`
}

func (category *Category) TableName() string {
	return "category"
}
