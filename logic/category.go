package logic

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/databases"
	"github.com/qingcc/goblog/model"
	"github.com/qingcc/goblog/util"
	redisPool "github.com/qingcc/goblog/utils/redis"
	"strconv"
	"strings"
)

type CategoryLogic struct{}

var DefaultCategory = CategoryLogic{}

//region Remark: 列表 Author:Qing
func (self CategoryLogic) List(c *gin.Context) []*model.Category {
	redis_string := "categories:list"
	data := make([]*model.Category, 0)
	if ok, _ := util.Exists(redis_string); !ok {
		objLog := GetLogger(c)
		err := databases.Orm.Desc("id").Find(&data)
		if err != nil {
			objLog.Errorf("CategoryLogic find errof:", err)
			return nil
		}
		byte, _ := json.Marshal(data)
		util.Set(redis_string, string(byte), -1)
	} else {
		rep, _ := redis.Bytes(util.Get(redis_string))
		json.Unmarshal(rep, &data)
	}

	return data
}

//endregion

//region Remark: 获取分类列表 Author:Qing
func (self CategoryLogic) LevelList(c *gin.Context) []*model.Category {
	redis_string := "categories:list"
	showList := make([]*model.Category, 0)
	if ok, _ := util.Exists(redis_string); !ok {
		objLog := GetLogger(c)
		data := make([]*model.Category, 0)
		err := databases.Orm.Asc("pid").Asc("sort").Find(&data)
		if err != nil {
			objLog.Errorf("CategoryLogic find errof:", err)
		}

		showList := make([]*model.Category, 0, len(data))
		self.Data2List(&showList, data, 0, 1, 0, "|-")
		bytes, _ := json.Marshal(showList)
		util.Set(redis_string, string(bytes), -1)
	} else {
		rep, _ := redis.Bytes(util.Get(redis_string))
		json.Unmarshal(rep, &showList)
	}
	return showList
}

//endregion

//region Remark: 将数据排序 Author:Qing
func (self CategoryLogic) Data2List(showNodeList *[]*model.Category, dataList []*model.Category, parentId, curLevel, pos int64, html string) {
	for num := len(dataList); pos < int64(num); pos++ {
		item := dataList[pos]
		if item.Pid == parentId {
			newhtml := strings.Repeat(html, int(curLevel-1))
			item.TitleHtml = newhtml + item.Title
			item.TimeHtml = item.CreatedAt.Format("2006-01-02 15:04:05")
			*showNodeList = append(*showNodeList, item)
			self.Data2List(showNodeList, dataList, item.Id, curLevel+1, 0, html)
		}
	}
}

//endregion

//region Remark: 获取角色 Author:Qing
func (self CategoryLogic) FindOne(c *gin.Context, field string, val interface{}) (*model.Category, error) {
	objLog := GetLogger(c)

	item := &model.Category{}
	ok, err := databases.Orm.Where(field+" = ?", val).Get(item)
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
	}
	if !ok {
		return nil, err
	}
	return item, err
}

//endregion

//region Remark: 角色下是否有管理员 Author:Qing
func (self CategoryLogic) RoleHasAdmin(c *gin.Context, id int64, ids []string) bool {
	objLog := GetLogger(c)

	has1, err := databases.Orm.Table("admin").In("role_id", ids).Count()
	util.CheckErr(err)
	has2, err := databases.Orm.Table("admin").Where("role_id = ?", id).Count()
	if err != nil {
		objLog.Errorf("CategoryLogic find errof:", err)
	}
	fmt.Println(has1, has2)
	if has1 > 0 || has2 > 0 {
		return true
	}
	return false
}

//region Remark:递归取出无限级联
func GetCategoryByParentId(parent_id int64) *[]model.Category {
	data := new([]model.Category)
	err := databases.Orm.Where("pid = ?", parent_id).Desc("id").Find(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

func GetAllCategory(parent_id int64) []map[string]interface{} {
	data := GetCategoryByParentId(parent_id)
	list := make([]map[string]interface{}, 0)
	for _, value := range *data {
		item := make(map[string]interface{})
		item["id"] = value.Id
		item["title"] = value.Title
		item["num"] = GetArticleNum(value.Id)
		son := GetCategoryByParentId(value.Id)
		if len(*son) != 0 {
			_list := GetAllCategory(value.Id) //获取所有的同级分类放入一个切片中
			item["child"] = _list             //将该切片塞入父级的 'child' 字段中
			item["has_child"] = true
		} else {
			item["has_child"] = false
		}
		list = append(list, item) //将同级分类数据 append 到一个切片中
	}
	return list
}

func GetArticleNum(category_id int64) int64 {
	redis_key := "article_category_num:" + strconv.FormatInt(category_id, 10)
	if ok, _ := redisPool.Exists(redis_key); ok {
		n, _ := redis.String(redisPool.Get(redis_key))
		num, _ := strconv.ParseInt(n, 10, 64)
		return num
	} else {
		ids := append(getCategoryIds(category_id), category_id)
		num, _ := databases.Orm.Table("article").In("category_id", ids).Count()
		redisPool.Set(redis_key, num, -1)
		return num
	}
}

func getCategoryIds(category_id int64) []int64 {
	ids := make([]int64, 0)
	err := databases.Orm.Table("category").Where("pid = ?", category_id).Cols("id").Find(&ids)
	util.CheckErr(err)
	if len(ids) == 0 {
		return ids
	}
	for _, id := range ids {
		ids = append(ids, getCategoryIds(id)...)
	}
	return ids
}
