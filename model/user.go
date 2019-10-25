package model

import (
	"fmt"
	"time"
	"github.com/qingcc/blog_go/util"
)

type User struct {
	Id            int64     `xorm:"pk autoincr int(8)" json:"id"`
	Email         string    `xorm:"varchar(255)" json:"email"`
	Username      string    `xorm:"unique index varchar(255)" json:"username"`
	Url           string    `xorm:"not null varchar(255)" json:"url"`
	Key           string    `xorm:"varchar(64)" json:"key"`
	Password      string    `xorm:"varchar(255)" json:"password"`
	Token         string    `xorm:"varchar(255)" json:"token"`
	LastLoginIp   string    `xorm:"varchar(32)" json:"last_login_ip"`
	LastLoginTime string    `xorm:"varchar(255)" json:"last_login_time"`
	LoginCount    int64     `xorm:"int(8)" json:"login_count"`
	IsLock        bool      `xorm:"default false" json:"is_lock"`
	Describe      string    `json:"describe"`
	CreatedAt     time.Time `xorm:"created" json:"created_at"`
	UpdatedAt     time.Time `xorm:"updated" json:"updated_at"`
}

func (this *User) EncryptToken() map[string]interface{} {
	token, err := util.AesEncrypt{}.AesEncrypt([]byte(this.Token), []byte(this.Key[:8]+this.Token[8:]))

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	item := make(map[string]interface{})
	item["token"] = string(token)
	item["uid"] = this.Id*3 - 1
	return item

}
