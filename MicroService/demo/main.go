package main

import "github.com/mfslog/goPractice/MicroService/demo/Echo"

var VERSION string = "0.0.1"

func main() {

	Echo.Main(VERSION)
}
