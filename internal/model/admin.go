package model

import "time"

type Admin struct {
	ID        int       `xorm:"not null pk autoincr BIGINT(20) id"`
	LoginName string    `xorm:"not null unique VARCHAR(30) login_name"`
	NickName  string    `xorm:"null VARCHAR(30) nick_name"`
	PasswdMd5 string    `xorm:"not null VARCHAR(255) passwdmd5"`
	Email     string    `xorm:"null unique VARCHAR(100) email"`
	Locked    int       `xorm:"not null default 0 TINYINT locked"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted TIMESTAMP deleted_at"`
}

func (Admin) TableName() string {
	return "gm_admin"
}
