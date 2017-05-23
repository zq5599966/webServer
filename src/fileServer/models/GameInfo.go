package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
	"math/rand"
)


type UserInfo struct {
	Id 	int					//主键递增
	Uid 	string 	`orm:"unique" orm:"size(20)"`	//用户账户
	Uname   string  `orm:"size(30)"`		//用户昵称
	FacebookId string  `orm:"unique"`		//facebook id
	SetFlagCount int				//设置旗子总数
	FoundFlagCount int				//找到旗子总数
}


func init() {

	fmt.Println("db init==========================")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:Veewo$$$@/game_db?charset=utf8", 30)

	orm.RegisterModel(new(UserInfo), new(FlagInfoDb))

	orm.RunSyncdb("default", false, true)
}

func RegisetNewUser() *UserInfo{
	o := orm.NewOrm()
	now := time.Now().Unix()
	random1 := rand.Intn(99) + 100
	random2 := rand.Intn(99) + 1000

	user := new(UserInfo)
	user.Uid = fmt.Sprintf("u%d%d%d", now, random1, random2)
	user.Uname = fmt.Sprintf("guest%d%d%d", now, random1, random2)

	id, err := o.Insert(user)
	if nil != err {
		fmt.Printf("create new user error id==%d ===%v", id, err)
		return RegisetNewUser()
	} else{
		return user
	}
}


