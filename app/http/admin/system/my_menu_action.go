package system

import (
	"encoding/json"
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/entity/mysql"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// MyMenu   登陆
func (receiver *Controller) MyMenu(req *admin.MyMenuRequest, ctx http.Context) (*admin.MyMenuResponse, error) {
	menus := make(map[int64]*admin.Menu)
	list := mysql.NewOrmAdminMenu().Order("id").Get()

	for _, menu := range list {
		meta := &admin.Meta{}
		if menu.Meta != nil {
			json.Unmarshal([]byte(*menu.Meta), meta)
		}
		menus[menu.Id] = &admin.Menu{
			Path:      *menu.Path,
			Name:      menu.Name,
			Meta:      meta,
			Component: menu.Component,
			Children:  make([]*admin.Menu, 0),
		}

		if menu.ParentId != 0 {
			m := menus[menu.Id]
			p := menus[menu.ParentId]
			p.Children = append(p.Children, m)
		}
	}
	got := make([]*admin.Menu, 0)
	for _, menu := range list {
		if menu.ParentId == 0 {
			got = append(got, menus[menu.Id])
		}
	}

	return &admin.MyMenuResponse{
		Menu: got,
		Permissions: []string{
			"list.add",
			"list.edit",
			"list.delete",
			"user.add",
			"user.edit",
			"user.delete",
		},
	}, nil
}

// GinHandleMyMenu gin原始路由处理
// http.Get(/system/menu/my)
func (receiver *Controller) GinHandleMyMenu(ctx *gin.Context) {
	req := &admin.MyMenuRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MyMenu(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
