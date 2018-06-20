package model

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"github.com/zzayne/go-blog/config"
)

// User 用户
type User struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt"`
	ActivatedAt  *time.Time `json:"activatedAt"`
	Name         string     `json:"name"`
	Pass         string     `json:"-"`
	Email        string     `json:"-"`
	Sex          uint       `json:"sex"`
	Location     string     `json:"location"`
	Introduce    string     `json:"introduce"`
	Phone        string     `json:"-"`
	Score        uint       `json:"score"`
	ArticleCount uint       `json:"articleCount"`
	CommentCount uint       `json:"commentCount"`
	CollectCount uint       `json:"collectCount"`
	Signature    string     `json:"signature"` //个人签名
	Role         int        `json:"role"`      //角色
	AvatarURL    string     `json:"avatarURL"` //头像
	CoverURL     string     `json:"coverURL"`  //个人主页背景图片URL
	Status       int        `json:"status"`
}

// 根据登录名获得用户
func (u *User) GetUserByName(name string) (user User, err error) {
	err = DB.Where("name = ?", name).First(&user).Error
	return user, err
}

func (u *User) GetUserByID(id int) (user User, err error) {
	err = DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// 验证密码是否正确
func (u *User) CheckPassword(password string) (result bool) {
	if password == "" || u.Pass == "" {
		return false
	}
	// 15276076190ca1b616d7c942b79adcafe06a12355e
	hash := u.EncryptPassword(password, u.Salt())
	return hash == u.Pass
}

// Salt 每个用户都有一个不同的盐
func (user *User) Salt() string {
	var userSalt string
	if user.Pass == "" {
		userSalt = strconv.Itoa(int(time.Now().Unix()))
	} else {
		userSalt = user.Pass[0:10]
	}
	return userSalt
}

// EncryptPassword 给密码加密
func (user *User) EncryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	hash = salt + password + config.AppConfig.PassSalt
	hash = salt + fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	return hash
}

const (
	// UserRoleNormal 普通用户
	UserRoleNormal = 1

	// UserRoleAdmin 管理员
	UserRoleAdmin = 2
)
