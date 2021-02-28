package main

import "master/src"

func init() {

}

func Load() interface{} {
	return new(src.HelloObject)
}
