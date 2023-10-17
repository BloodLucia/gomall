package adminmodel

import "time"

type User struct {
	ID        int       `xorm:"not null pk autoincr BIGINT(20) id"`
	LoginName string    `xorm:"not null unique VARCHAR(50) login_name"`
	NickName  string    `xorm:"null VARCHAR(50) nick_name"`
	PasswdMd5 string    `xorm:"not null VARCHAR(255) passwdmd5"`
	Email     string    `xorm:"null unique VARCHAR(100) email"`
	Locked    int       `xorm:"not null default 0 TINYINT locked"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted TIMESTAMP deleted_at"`
}

type UserRegisterRequest struct {
	LoginName string `validate:"required|min_len:5|max_len:10" message:"required:login_name 登录名不能为空|min_len:login_name 登录名应为5-10个字符|max_len:login_name 登录名应为5-10个字符" json:"login_name"`
	Passwd    string `validate:"required|min_len:8|max_len:20" message:"required:passwd 密码不能为空|min_len:passwd 密码应为8-20个字符|max_len:passwd 密码应为8-20个字符" json:"passwd"`
}

type UserLoginRequest struct {
	LoginName string `validate:"required" message:"required:账号不能为空" json:"login_name"`
	Passwd    string `validate:"required" message:"required:密码不能为空" json:"passwd"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

func (u User) TableName() string {
	return "t_admin_users"
}
