package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shahejiuhuo/libary/e"
)

const dsn = "work:npWS1Iu5MCmYmG9U@tcp(104.160.43.85:3306)/shahejiuhuo?charset=utf8mb4&parseTime=True&loc=Local"

type Item struct {
	Id       int
	Name     string
	Icon     string
	Pictures []Picture
}

type Picture struct {
	Id     int
	ItemId int
	picUlr string
}

type inputGetList struct {
	ItemId string `form:"itemId" binding:"required"`
}

var DbCon *gorm.DB

func GetList(c *gin.Context) {
	var res = e.GetRrrReturn(e.SUCCESS)
	var inputParams inputGetList
	if err := c.ShouldBind(&inputParams); err != nil {
		c.JSON(200, e.GetRrrReturn(e.InvalidParams))
		return
	}

	initDb()
	var where = map[string]interface{}{
		"items.id":       inputParams.ItemId,
		"items.isDel":    1,
		"pictures.isDel": 1,
	}
	rows, err := DbCon.Table("items").Select("items.id as itemId, pictures.id as picId, pictures.picUrl").Joins("inner join pictures on pictures.itemId = items.id").Where(where).Rows()

	if err != nil {
		c.JSON(200, e.GetRrrReturn(e.FAILD))
		return
	}
	var resItemId int
	var resPicId int
	var resPicUrl string

	var resDatas = make([]map[string]interface{}, 0)
	for rows.Next() {
		var resData = make(map[string]interface{}, 0)
		rows.Scan(&resItemId, &resPicId, &resPicUrl)
		resData["itemId"] = resItemId
		resData["picId"] = resPicId
		resData["picUrl"] = resPicUrl
		resDatas = append(resDatas, resData)
	}

	res["data"] = resDatas
	c.JSON(200, res)
}

func initDb() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	DbCon = db
}
