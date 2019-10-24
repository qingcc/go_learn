package model

import (
	"time"
)

type Config struct {
	Name      string    `xorm:"not null varchar(20)"` //评论量
	Value     string    `xorm:"varchar(80)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (*Config) table() string {
	return "config"
}
