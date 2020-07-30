package pi

import (

	// "bytes"
	"fmt"
	// "io"
	// "os"
	"os/exec"
)

type Stat struct {
	Total float64
	Used  float64
	Free  float64
}

type Pi struct {
	Cpu  float64
	Mem  Stat
	Disk Stat
	Temp float64
}

func New() *Pi {
	return &Pi{}
}

func (c *Pi) GetCpu() {
	cmd := exec.Command("top")

	fmt.Println(cmd.Output())
	// ps := exec.Command("top", "-n1")
	// grep := exec.Command("awk", `'/Cpu\(s\):/ {print $2}'`)

	// r, w := io.Pipe() // 创建一个管道
	// defer r.Close()
	// defer w.Close()
	// ps.Stdout = w  // ps向管道的一端写
	// grep.Stdin = r // grep从管道的一端读

	// var buffer bytes.Buffer
	// grep.Stdout = &buffer // grep的输出为buffer

	// _ = ps.Start()
	// _ = grep.Start()
	// ps.Wait()
	// w.Close()
	// grep.Wait()
	// bs := make([]byte, 1024)
	// length, err := buffer.Read(bs)
	// fmt.Println(length, err, bs[:length])
	// io.Copy(os.Stdout, &buffer) // buffer拷贝到系统标准输出

}

func (c *Pi) GetMem() {

}

func (c *Pi) GetDisk() {

}

func (c *Pi) GetTemp() {

}
