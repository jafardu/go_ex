/*
* @Description:
* @Author: jafar du
* @Date: 2022/5/28 22:27
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

/*
go build main.go
.\main.exe -cpu=true
go tool pprof .\cpu.pprof
top 4 查看cpu占用情况
list 函数名 查看函数中使用最多的代码逻辑
*/

func logicCode() {
	var c chan int //nil
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			//time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUPprof bool // 是否开启CPUprofile的标志位
	var isMemPprof bool // 是否开启内存profile的标志位

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		f1, err := os.Create("cpu.pprof") // 在当前路径下创建一个cpu.pprof文件
		if err != nil {
			fmt.Printf("create cpu pprof failed, err: %#v\n", err)
			return
		}
		pprof.StartCPUProfile(f1) // 往文件中记录cpu profile信息
		defer func() {
			pprof.StopCPUProfile()
		}()
	}

	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)

	if isMemPprof {
		f2, err := os.Create("mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err: %#v\n", err)
			return
		}
		pprof.WriteHeapProfile(f2)
		f2.Close()
	}

}
