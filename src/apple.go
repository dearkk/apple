package src

import (
	"github.com/dearkk/component/market"
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

const routeTag = "账户管理"
const module = "apple"

type Apple struct {
}

type Login struct {
	Name string `json:"name"`
}

func (a *Apple) Start(db *gorm.DB, klog *log.Entry, param *[]market.Param, addRoute func(string, []market.Route)) {
	klog.Printf("Apple Start: %+v\n", *param)
	routes := []market.Route{
		{
			Name:   "登陆",
			Path:   "/login",
			Tag:    routeTag,
			Handle: a.login,
			Writes: Login{},
			Reads:  Login{},
		},
	}
	addRoute(module, routes)
}

func (a *Apple) login(request *restful.Request, response *restful.Response) {
	l := Login{
		Name: "login",
	}
	response.WriteHeaderAndEntity(http.StatusOK, &l)
}
