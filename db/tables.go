package db

import (
	"chessSvr/conf"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ChessUser struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	IsVip    int8   `json:"is_vip"`
}

type ChessRank struct {
	UserId int `json:"user_id"`
	Point  int `json:"point"`
	Rank   int `json:"rank"`
}

type ChessCard struct {
	Id          int    `json:"id"`
	CardName    string `json:"card_name"`
	CardLevel   int8   `json:"card_level"`
	Cost        int8   `json:"cost"`
	BasicAttack int16  `json:"basic_attack"`
	BasicHealth int16  `json:"basic_health"`
	SalePrice   int8   `json:"sale_price"`
	CardDesc    string `json:"card_desc"`
}

type CardPercent struct {
	Id           int  `json:"id"`
	TechLevel    int8 `json:"tech_level"`
	PercentOne   int8 `json:"percent_one"`
	PercentTwo   int8 `json:"percent_two"`
	PercentThree int8 `json:"percent_three"`
	PercentFour  int8 `json:"percent_four"`
	PercentFive  int8 `json:"percent_five"`
	CallNumber   int  `json:"call_number"`
}

var GDB *gorm.DB

func init() {
	connStr := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		conf.GConfig.DB.UserName, conf.GConfig.DB.PassWord, conf.GConfig.DB.DBName)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	GDB = db
}
