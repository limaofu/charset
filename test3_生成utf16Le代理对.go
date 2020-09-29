// utf16 project main.go
package main

import (
	"fmt"
	"os"

	"github.com/limaofu/charset"
)

func main() {

	var str string = "我在𨱇f16BE.txt, o𬱖𬜬ab"
	ot, e := charset.Utf8ToUtf16LE([]byte(str))
	if e != nil {
		fmt.Println("error")
		return
	}
	f, ee := os.OpenFile("D:\\testDir2\\utf16BE.txt", os.O_RDWR|os.O_CREATE, 0666)
	if ee != nil {
		fmt.Println("open file error")
		return
	}
	defer f.Close()

	var bom [2]byte
	bom[0] = 0xFF
	bom[1] = 0xFE
	f.Write(bom[:])
	f.Write(ot)
	fmt.Printf("len ot : %d\n", len(ot))
	fmt.Println("完成")
}
