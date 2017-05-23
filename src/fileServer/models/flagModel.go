package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"encoding/json"
	"math/rand"
)

/*************************************************************************************
			Flag
 *************************************************************************************/
type Flags struct {
	FlagInfo []FlagInfo
}

type FlagInfo struct {
	Username 	string	`json:"username"`
	Tmxidx 		string	`json:"tmxidx"`
	FlagPos		string  `json:"flagpos"`
}

type FlagInfoDb struct {
	Id		int
	TmxId 		string
	Uid		string		//zq debug
	FlagPos		string
	FindCount	int
	LikeCount	int
	UnlikeCount	int
	Cuid		int		//最后一个评论的id
}

func WriteNewFlag(flag FlagInfo) string{
	o := orm.NewOrm()
	var newFlag = new(FlagInfoDb)

	newFlag.TmxId = flag.Tmxidx
	newFlag.Uid = flag.Username	//zq debug
	newFlag.FlagPos = flag.FlagPos

	id, err := o.Insert(newFlag)
	if nil != err {
		fmt.Printf("create new flag id===%d  err===%v \n", id, err)
		return "error"
	}

	return "success"
}

func NewFlagEvent(data string) string{
	var result FlagInfo
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("json data err")
		return "error"
	}
	//
	//fmt.Println("json===", result)
	//fmt.Println("username====", result.Username)
	//fmt.Println("tmx id====", result.Tmxidx)
	//fmt.Println("flag pos====", result.FlagPos)

	return WriteNewFlag(result)
}
/*************************************************************************************************************************************************/



type RequestFlag struct {
	Username  	string		`json:"username"`
	Tmxids		[]string	`json:"tmxids"`
}

type ResponseFlags struct {
	FlagsInfo	[]*FlagInfoDb 	`json:"flagsinfo"`
}

func GetFlagsEvent(data string) string {

	var result RequestFlag
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("json data err")
		return "error"
	}

	flagInfo := getRandomFlagArrayByTMXArray(result.Tmxids, result.Username) //getRandomFlagByTMX(result.Tmxids[0], result.Username)
	if flagInfo != nil {
		var responseFlags ResponseFlags
		responseFlags.FlagsInfo = flagInfo
		jsonStr, err := json.Marshal(responseFlags)
		if err != nil {
			fmt.Println("json marshal error  flaginfo===", responseFlags)
			return "error"
		}
		return string(jsonStr)
	}

	return "error"
}

func getRandomFlagArrayByTMXArray(tmxIds []string, username string) []*FlagInfoDb {

	var num = len(tmxIds)
	flagInfos := make([]*FlagInfoDb, num)
	//var flagInfos [num]*FlagInfoDb
	for i := 0; i < len(tmxIds); i++ {
		//fmt.Printf("i=====%d    tmx id====%s   username===%s \n", i, tmxIds[i], username)
		flagInfos[i] = getRandomFlagByTMX(tmxIds[i], username)
	}
	//tmpArray := flagInfos[0:num]
	return flagInfos
}

func getRandomFlagByTMX(tmxId string, username string) *FlagInfoDb {
	o := orm.NewOrm()

	var flags []*FlagInfoDb
	num, err := o.QueryTable("FlagInfoDb").Filter("tmxId", tmxId).Exclude("uid", username).All(&flags)
	if nil != err {
		fmt.Println("select flags error")
		return nil
	}

	if num == 0 {
		fmt.Println("fuck flags num ====0")
		return nil
	}

	if len(flags) > 1 {
		flag := flags[rand.Intn(len(flags)-1)]
		if flag != nil {
			return flag
		}
	} else {
		return flags[0]
	}

	return nil
}
/*************************************************************************************************************************************************/


func ShareFlagEvent(data string) string {

	return "error"
}

