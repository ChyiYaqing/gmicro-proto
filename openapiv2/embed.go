package openapiv2

import (
	// embed 功能来将静态文件潜入到编译后的二进制文件中
	"embed"
)

//go:embed OpenAPI/*
var OpenAPI embed.FS

/**
go:embed OpenAPI/* 编译器指令表示Go编译器要嵌入OpenAPI目录下的所有文件
var OpenAPI embed.FS 声明一个系统变量

这样可以将Swagger UI所有静态文件打包到你的程序中
运行时不需要依赖外部文件
*/
