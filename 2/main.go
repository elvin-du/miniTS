package main

import "log"

var source = `
//一个函数的声明，这个函数很简单，只打印"Hello World!"
function sayHello(){
    println("Hello World!");
}
//调用刚才声明的函数
sayHello();
`

func init() {
	log.SetFlags(log.Llongfile)
}
func main() {
	compileAndRun()
}

func compileAndRun() {
	//词法分析
	tokenizer := NewTokenizer(NewCharStream(source))

	//语法分析，构建AST
	parser := NewParser(tokenizer)
	prog := parser.ParseProg()
	prog.Dump("  ")

	//语义分析，进行函数消解
	NewRefResolver().VisitProg(prog)
	log.Println("语义分析，进行函数引用消解")
	prog.Dump("  ")
	//程序运行

	NewIntepreter().Start(prog)
}
