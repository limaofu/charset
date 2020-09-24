// printFuhao project main.go
package main

import (
	"fmt"
	"os"

	"github.com/limaofu/charset"
)

func main() {
	f, e := os.OpenFile("D:\\testDir2\\zigen.yaml", os.O_APPEND, 0666)
	if e != nil {
		return
	}
	defer f.Close()
	var rune uint32 = 0
	f.Write([]byte("cjk笔画\r\n"))
	for rune = 0x31c0; rune <= 0x31ef; rune++ {
		d, e := charset.Uint32ToUtf8(rune)
		if e != nil {
			fmt.Println("uint32toutf8 error")
			return
		}
		f.Write(d)
		f.Write([]byte(string("\r\n")))
	}
	fmt.Println("完成")

}
