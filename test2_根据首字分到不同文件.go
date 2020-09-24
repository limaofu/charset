func main() {
	f, e := os.OpenFile("D:\\cjk_0_base_a.dict.yaml", os.O_RDONLY, 0)
	if e != nil {
		return
	}
	defer f.Close()
	fb, eb := os.OpenFile("D:\\cjk_0_base_a_new.dict.yaml", os.O_RDWR|os.O_CREATE, 0666)
	if eb != nil {
		return
	}
	defer fb.Close()
	fbn, ebn := os.OpenFile("D:\\cjk_0_base_a.dict_IS_not.yaml", os.O_RDWR|os.O_CREATE, 0666)
	if ebn != nil {
		return
	}
	defer fbn.Close()
	var i, j int32 = 0, 0
	iread := bufio.NewReader(f)
	iw := bufio.NewWriter(fb)
	iwn := bufio.NewWriter(fbn)
	for {
		s, e := iread.ReadString('\n')
		if e == io.EOF {
			break
		}
		j++
		r,_ := charset.Utf8GetFirstRune([]byte(s))
		if r >= 0x00 && r < 0xffff {
			fb.WriteString(s)
			i++
		} else {
			fbn.WriteString(s)
		}

	}
	iw.Flush()
	iwn.Flush()
	fmt.Printf("i属于0的 is: %d\nj读取的行数 is: %d\n", i, j)
}