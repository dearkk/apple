module master

go 1.15

require (
	github.com/dearkk/component v0.0.0-20210210064905-f4726ef201a2
	github.com/emicklei/go-restful v2.9.6+incompatible
	github.com/sirupsen/logrus v1.6.0
	gorm.io/gorm v1.21.2
)

replace github.com/dearkk/component v0.0.0-20210210064905-f4726ef201a2 => ../component
