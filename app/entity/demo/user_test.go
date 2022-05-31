package demo

import (
	"fmt"
	"github.com/go-home-admin/home/bootstrap/services/database"
	"testing"
)

func TestNewOrmUsers(t *testing.T) {
	// 读取一条数据
	user, has := NewOrmUsers().WhereId(1).First()
	fmt.Println(user, has)

	// 读取列表数据 select * from users where id > 1 and create_ed >= now() limit 16
	users := NewOrmUsers().WhereIdGt(1).WhereCreatedAtGte(database.Now()).Limit(15).Get()
	fmt.Println(users)
	// 结果的辅助函数
	ids := users.GetIdList()
	fmt.Println(NewOrmUsers().WhereIdIn(ids).Get())
	fmt.Println(users.GetIdMap())
}
