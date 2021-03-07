package main

import (
	"github.com/dearkk/component/market"
	"master/src"
)

func Load() market.Load {
	return &src.Apple{}
}
