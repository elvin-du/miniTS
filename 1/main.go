package main

import "log"

func init() {
	log.SetFlags(log.Llongfile)
}
func main() {
	compileAndRun()
}

func compileAndRun() {
	//词法分析

	//语法分析，构建AST

	//语义分析，进行函数消解

	//程序运行

}
