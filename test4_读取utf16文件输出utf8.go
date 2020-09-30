// testUtf16toUtf8 project main.go
package main

import (
	//"bufio"
	"fmt"
	"os"

	"github.com/limaofu/charset"
)

func main() {
	f, _ := os.OpenFile("D:\\testDir2\\newt.txt", os.O_RDONLY, 0600)
	//fr := bufio.NewReader(f)
	//str, _ := fr.ReadString('\n')
	data := make([]byte, 10)
	f.Read(data)
	//fmt.Printf("%s\n", str)
	utf8b, _ := charset.Utf16LEToUtf8(data)
	fmt.Printf("%s\n", string(utf8b[:]))
}
