package main

import (
	"fmt"
	"runtime"
)

func main() {
	ConfigRuntime()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()                   // 获取当前机器的 CPU 核心数
	runtime.GOMAXPROCS(nuCPU)                   // 设置 Go 运行时的最大操作系统线程数 = CPU 核心数
	fmt.Printf("Running with %d CPUs\n", nuCPU) // 打印使用的 CPU 数量
}
