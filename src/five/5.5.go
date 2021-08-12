package main

import (
	"fmt"
	"github.com/astaxie/beego/client/orm"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


func checkError1(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Uid int	`orm:"PK;auto"`
	Name string `orm:"size(100)"`
	Profile *Profile `orm:"rel(one)"`
	Post []*Post `orm:"reverse(many)"`
}

type UserInfo struct {
	Uid	int	`orm:"PK;auto"`
	UserName string
	Department string
	Created time.Time
}

type Profile struct {
	Id int `orm:"PK;auto"`
	Age int16
	User *User `orm:"reverse(one)"`
}

type Post struct {
	Id int `orm:"PK;auto"`
	Title string
	User *User `orm:"rel(fk)"`
	Tags []*Tag	`orm:"rel(m2m)"`
}


type Tag struct {
	Id int `orm:"PK;auto"`
	Name string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	checkError1(err)

	err = orm.RegisterDataBase("default", "mysql", "root:micbay911@/escollect?charset=utf8")
	//orm.SetMaxOpenConns("default", 30)
	checkError1(err)

	orm.RegisterModel(new(Profile))
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(UserInfo))
	orm.RegisterModel(new(Tag))
	orm.RegisterModel(new(User))
	//orm.RunSyncdb("default", false, true)

}

func operateMysql() {
	o := orm.NewOrm()
	user := User{Name: "selina"}

	id, err := o.Insert(&user)
	fmt.Println(id, err)

	user.Name = "John"
	num, err := o.Update(&user)
	fmt.Println(num, err)

	u := User{Uid: user.Uid}
	err = o.Read(&u)
	fmt.Println(err)

	name, err := o.Delete(&u)
	checkError1(err)
	fmt.Println(name)

}

func insertData() {
	o := orm.NewOrm()
	var user UserInfo
	user.UserName = "Peter"
	user.Department = "商务部"
	user.Created = time.Now()

	id, err := o.Insert(&user)
	checkError1(err)
	fmt.Println(id)

	users := []UserInfo{
		{UserName: "GO", Department: "CS", Created: time.Now()},
		{UserName: "Python", Department: "CS", Created: time.Now()},
		{UserName: "C++", Department: "CS", Created: time.Now()},
	}

	successNums, err := o.InsertMulti(100, users)
	checkError1(err)
	fmt.Println(successNums)
}


func updateData() {
	o := orm.NewOrm()
	user := UserInfo{Uid: 1}

	if o.Read(&user) == nil {
		user.UserName = "康师傅"
		user.Department = "饮料"
		if num, err := o.Update(&user, "UserName", "Department"); err == nil {
			fmt.Println(num)
		}
	}

}

func queryData() {
	o := orm.NewOrm()
	userinfo := UserInfo{Uid: 2}

	err := o.Read(&userinfo)
	if err == orm.ErrNoRows {
		fmt.Println("查找不到此行")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(userinfo.Uid, userinfo.UserName)
	}

	var userinfo1 []*UserInfo
	qs := o.QueryTable("user_info")
	qs.Filter("UserName", "Peter")
	qs.Filter("Department", "白酒")
	num, err := qs.All(&userinfo1)
	checkError1(err)
	fmt.Println(num)

	for u, _ := range userinfo1 {
		fmt.Println(u)
	}
}

func (m *UserInfo) rawQuery() {
	o := orm.NewOrm()
	rs := o.Raw("select * from user_info where user_name=? and department=?", "Peter", "白酒")

	var users []*UserInfo
	num, err := rs.QueryRows(&users)
	checkError1(err)
	fmt.Println(num)
	for u, v := range users {
		fmt.Println(u, v)
	}
}

func main() {
	//operateMysql()

	//insertData()

	//updateData()

	//queryData()

	var userinfo UserInfo
	userinfo.rawQuery()

}