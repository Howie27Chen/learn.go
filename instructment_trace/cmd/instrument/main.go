// instrument_trace/cmd/instrument/main.go

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"learn.go/instructment_trace/instrumenter"
	"learn.go/instructment_trace/instrumenter/ast"
)

var wrote bool

func init() {
	flag.BoolVar(&wrote, "w", false, "write result to (source) file instead of stdout")
}

func usage() {
	fmt.Println("instrument [-w] xxx.go")
	flag.PrintDefaults()
}

func main() {
	fmt.Println(os.Args)
	flag.Usage = usage
	flag.Parse() // 解析命令行参数

	if len(os.Args) < 2 {
		usage()
		return
	}

	var file string
	if len(os.Args) == 3 {
		file = os.Args[2]
	}

	if len(os.Args) == 2 {
		file = os.Args[1]
	}

	if filepath.Ext(file) != ".go" {
		usage()
		return
	}

	// 声明 instrument.Instrumenter 接口类型变量
	var ins instrumenter.Instrumenter
	// 创建以 ast 方式实现 Instrument 接口的 ast.instumenter 实例
	ins = ast.New("github.com/Howie27Chen/instrument_trace", "trace", "Trace")
	// 向 Go 源文件所有函数注入 Trace 函数
	newSrc, err := ins.Instrument(file)
	if err != nil {
		panic(err)
	}

	if newSrc == nil {
		fmt.Printf("no trace added for %s\n", newSrc)
		return
	}

	if !wrote {
		// 将代码生成打印到 stdout 上
		fmt.Println(string(newSrc))
		return
	}

	// 将生成的新代码内容写回原 Go 源文件
	if err = ioutil.WriteFile(file, newSrc, 0666); err != nil {
		fmt.Printf("write %s error: %v\n", file, err)
		return
	}
	fmt.Printf("instrument tarce for %s ok\n", file)
}
