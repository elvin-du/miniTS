package main

type AstNode interface {
	//打印对象信息，prefix是前面填充的字符串，通常用于缩进显示
	//String()
	Kind() NodeKind
}
