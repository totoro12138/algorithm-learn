package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/url"
	"os"
	"time"
)

type Address struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Name string `json:"name"`

	System   bool `gorm:"default:false" json:"system"`
	Password bool `gorm:"default:false" json:"password"`

	Clod bool `gorm:"default:false"`

	ChainCoin  string `gorm:"primary_key;size:64" json:"chain_coin"`
	Address    string `gorm:"primary_key;size:128" json:"address"`
	Memo       string `json:"memo"`
	PrivateKey string `gorm:"type:longtext" json:"-"`
	PublicKey  string `gorm:"type:longtext" json:"-"`
}

type BTCTransaction struct {
	CreatedAt time.Time
	TxHash    string `gorm:"primary_key;size:128"`
	To        string `gorm:"index" gorm:"primary_key;size:128"`

	Value     int64
	Direction int `gorm:"default:0"`

	AssetCoin string
}

type OmniTransaction struct {
	CreatedAt time.Time
	TxHash    string `gorm:"primary_key;size:128"`
	To        string `gorm:"index" gorm:"primary_key;size:128"`

	PropertyID   int    `json:"property_id"`
	PropertyName string `json:"property_name";gorm:"size:128"`

	From      string `json:"from"`
	Amount    string `json:"amount"`
	AssetCoin string `json:"asset_coin"`
	Direction int    `gorm:"default:0"`
}

type InviteFriendInfo struct {
	Address     string `json:"address"`
	OmniToCount string `json:"omni_to_count"`
	BTCToCount  string `json:"btc_to_count"`
}

func main() {
	SyncDB()
	var data []InviteFriendInfo
	var address []*Address
	DB.Find(&address)
	if len(address) == 0 {
		panic("address len is 0")
	}
	addres := make([]string, len(address))
	for i, v := range address {
		addres[i] = v.Address
	}

	DB.Model(&BTCTransaction{}).Select("`to` as address,count(`to`) as btc_to_count").Where("`to` in (?)", addres).Group("`to`").Scan(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println("data 1", len(data))
	DB.Model(&OmniTransaction{}).Select("`to` as address,count(`to`) as omni_to_count").Where("`to` in (?)", addres).Group("`to`").Scan(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println("data 2", len(data))

	var scanData []InviteFriendInfo
	DB.Model(&OmniTransaction{}).Select("`to` as address,count(`to`) as omni_to_count").Where("`to` in (?)", addres).Group("`to`").Scan(&scanData)
	for _, v := range scanData {
		fmt.Println(v)
	}
	fmt.Println("scanData", len(scanData))
}

var DB *gorm.DB

func SyncDB() {
	createDB()
	Connect()
	DB.
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate()
}

/**
数据库链接
*/
func Connect() {
	//db_host := beego.AppConfig.String("db_host")
	//db_port := beego.AppConfig.String("db_port")
	//db_user := beego.AppConfig.String("db_user")
	//db_pass := beego.AppConfig.String("db_pass")
	//db_name := beego.AppConfig.String("db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true",
		//db_user,
		"root",
		//db_pass,
		"6621423",
		//db_host,
		"127.0.0.1",
		//db_port,
		"3306",
		//db_name,
		"wallet",
		url.QueryEscape("Asia/Shanghai"))

	var err error

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("detabase connect error:", err)
		os.Exit(0)
	}
	DB.SingularTable(true)
	//DB.DB().SetMaxOpenConns(200)
	//DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(20)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetConnMaxLifetime(1 * time.Second)

	if beego.AppConfig.String("runmode") == "dev" {
		//DB = DB.Debug()
	}

}

func createDB() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&loc=%s&parseTime=true",
		db_user,
		db_pass,
		db_host,
		db_port, url.QueryEscape("Asia/Shanghai"))
	sqlstring := fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8mb4 COLLATE utf8mb4_general_ci",
		db_name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()

}
