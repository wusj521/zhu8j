package models

import (
	//"fmt"
	"os"
	"path"
	//"sappo/saprfc"
	//"sappo/utils"
	//"strconv"
	//"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/lib/pq"
	//saprfc "github.com/sap/gorfc/gorfc"
)

const (
	_DB_NAME        = "data/zhubajie.db"
	_SQLITE3_DRIVER = "sqlite3"
)

/*
//字体首写字母大写
*/
//用户
type User struct {
	Id       int64
	Uname    string `orm "unique" "index"` //用户名 两条记录不能重复
	Unamemd5 string `orm "unique" "index"` //用户名 两条记录不能重复
	Pwd      string
	Tel      string `orm:"index"`
}

func RegisterDB() {
	// 检查数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	// 注册模型
	orm.RegisterModel(new(User))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	//	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

//以上是注册DB和自动建表
