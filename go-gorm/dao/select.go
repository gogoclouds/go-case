package dao

import (
	"errors"

	"github.com/gogoclouds/go-gorm/model"
	"gorm.io/gorm"
)

func (dao *UserDao) FindOne() {
	var user *model.User

	// 获取第一条记录 - 主键升序
	dao.db.First(user) // select * from users order by id limit 1;

	// 获取最后一条记录 - 主键降序
	dao.db.Last(user) // select * from users order by id desc limit 1;

	// 获取第一条记录 - 没有排序
	result := dao.db.Take(user) // select * from users limit 1;

	// 数据为空的错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (dao *UserDao) FindByID(ID string) {
	var u *model.User
	dao.db.First(u, ID) // select * from users where id = ?;
	dao.db.First(u, "id = ?", ID)

	var result model.User
	dao.db.Model(&model.User{ID: ID}).First(result)

	dao.db.Find(u, []int{1, 2, 3}) // select * from users where id in (1,2,3);
}

func (dao *UserDao) FindAll() {
	var users []*model.User
	dao.db.Find(users) // select * from users;
}

func (dao *UserDao) Find_where() {
	var user *model.User
	// select * from users where name = 'yu.guan' order by id limit 1;
	dao.db.Where("name = ?", "yu.guan").First(user)

	var users []*model.User

	// select * from users where name <> 'fei.zhang';
	dao.db.Where("name <> ?", "fei.zhang").Find(users)

	// select * from users where name in ("yun.zhao", "chao.ma");
	dao.db.Where("name in ?", []string{"yun.zhao", "chao.ma"}).Find(users)

	// select * from users where like '%zhang';
	dao.db.Where("name like ?", "%zhang").Find(users)

	// select * from users where name = 'yueying.huang' and age = 23;
	dao.db.Where("name = ? AND age >= ?", "yueying.huang", 23)

	// select * from users where updated_at > '2022-11-10 22:19:00';
	dao.db.Where("updated_at > ?", "2022-11-10 22:19:00").Find(users)

	// select * from users where created_at between "2022-10-10 00:00:00" and "2022-11-10 22:19:00";
	dao.db.Where("created_at between ? and ?", "2022-10-10 00:00:00", "2022-11-10 22:19:00")
}

//  struct 零值不会作为查询条件，可以使用map
func (dao *UserDao) FindByStruct() {
	var u *model.User
	// select * from users where name = "chan.diao" and age = 20 order by id limit 1;
	dao.db.Where(&model.User{Name: "chan.diao", Age: 20}).First(u)

	var users []*model.User
	// select * from users where id in ("xxx", "xx", "xxxxx");
	dao.db.Where([]string{"xxx", "xx", "xxxxx"}).Find(users)

	// select * from users where name = 'qiao.da' and age = 0;
	dao.db.Where(map[string]any{"Name": "qiao.da", "Age": 0}).Find(users)

	// 指定字段查询
	// select * from user where name = 'qiao.xiao' and age = 0;
	dao.db.Where(&model.User{Name: "qiao.xiao"}, "name", "age").Find(users)
}

func (dao *UserDao) Find_not() {
	var u *model.User

	// select * from user where not name = 'shangxiang.sun' order by id limit 1;
	dao.db.Not("name = ?", "shangxiang.sun").First(u)

	var users []*model.User

	// select * from users where name not in ("mayi.si", "yu.xun")
	dao.db.Not(map[string]any{"name": []string{"mayi.si", "yu.xun"}}).Find(users)

	// select * from users where name <> 'zhi.cao' and age <> 18 order by id limit 1;
	dao.db.Not(model.User{Name: "zhi.cao", Age: 18}).First(u)

	// select * from users where id not in ("xxx", "xxx", "xxxxx") order by id limit 1;
	dao.db.Not([]string{"xxx", "xxx", "xxxxx"}).First(u)
}

func (dao *UserDao) Find_or() {

}