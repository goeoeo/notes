package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	// 如果命令行设置了 cpuprofile
	if *cpuprofile != "" {
		// 根据命令行指定文件名创建 profile 文件
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		// 开启 CPU profiling
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
}
