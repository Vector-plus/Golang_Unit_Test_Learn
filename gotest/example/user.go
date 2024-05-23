package example

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	UserName string `gorm:"type:varchar(20),not null,userName" json:"userName"`
	Password string `gorm:"type:varchar(20),not null,password" json:"password"`
	Age      int    `gorm:"type:varchar(20),not null,age" json:"age"`
}

var ctx = context.Background()

type SQLDB struct {
	db *gorm.DB
}

type Rdb struct {
	Rdb *redis.Client
}

func NewUserDb(db *gorm.DB) *SQLDB {
	return &SQLDB{db: db}
}

func (db *SQLDB) FindByUserId(uid int) (User, error) {
	var user User
	err := db.db.Where("ID = ?", uid).Find(&user).Error
	return user, err
}

func (rdb *Rdb) GetAllUser(key string) ([]string, error) {
	re, err := rdb.Rdb.SMembers(ctx, key).Result()
	return re, err
}

func GetUserInfo(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("返回错误：%d", resp.StatusCode)
	}

	data, _ := io.ReadAll(resp.Body)

	return string(data), nil
}
