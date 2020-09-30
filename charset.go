// charset project charset.go
// Author: Cof-Lee
// Date: 2020-09-30
package charset

import (
	"errors"
)

//input utf8's []byte, output []uint32's Runes
func Utf8GetRunes(b []byte) ([]uint32, error) {
	var cur int = 0
	var u uint32 = 0
	ulen := len(b)
	var rune []uint32
	for cur < ulen {
		if b[cur] >= 0 && b[cur] < 128 {
			u = uint32(b[cur])
			cur++
		} else if b[cur] > 191 && b[cur] < 224 {
			u = uint32(b[cur] & 0x1f)
			u = (u << 6) | uint32(b[cur+1]&0x3f)
			cur += 2
		} else if b[cur] > 223 && b[cur] < 240 {
			u = uint32(b[cur] & 0x0f)
			u = (u << 6) | uint32(b[cur+1]&0x3f)
			u = (u << 6) | uint32(b[cur+2]&0x3f)
			cur += 3
		} else if b[cur] > 239 && b[cur] < 248 {
			u = uint32(b[cur] & 0x07)
			u = (u << 6) | uint32(b[cur+1]&0x3f)
			u = (u << 6) | uint32(b[cur+2]&0x3f)
			u = (u << 6) | uint32(b[cur+3]&0x3f)
			cur += 4
		} else {
			rune = append(rune, u)
			return rune, errors.New("Not_A_Unicode_Rune")
		}
		rune = append(rune, u)
	}
	return rune, nil
}

//input utf8's []byte, output first []uint32's single Rune
func Utf8GetFirstRune(b []byte) (uint32, error) {
	var cur int = 0
	var u uint32 = 0
	if b[cur] >= 0 && b[cur] < 128 {
		u = uint32(b[cur])
		cur++
	} else if b[cur] > 191 && b[cur] < 224 {
		u = uint32(b[cur] & 0x1f)
		u = (u << 6) | uint32(b[cur+1]&0x3f)
		cur += 2
	} else if b[cur] > 223 && b[cur] < 240 {
		u = uint32(b[cur] & 0x0f)
		u = (u << 6) | uint32(b[cur+1]&0x3f)
		u = (u << 6) | uint32(b[cur+2]&0x3f)
		cur += 3
	} else if b[cur] > 239 && b[cur] < 248 {
		u = uint32(b[cur] & 0x07)
		u = (u << 6) | uint32(b[cur+1]&0x3f)
		u = (u << 6) | uint32(b[cur+2]&0x3f)
		u = (u << 6) | uint32(b[cur+3]&0x3f)
		cur += 4
	} else {
		return u, errors.New("Not_A_Unicode_Rune")
	}
	return u, nil
}

//input single uint32's Rune, output utf8's []byte
func Uint32ToUtf8(rune uint32) ([]byte, error) {
	if rune >= 0 && rune < 128 {
		u := make([]byte, 1)
		u[0] = byte(rune)
		return u, nil
	} else if rune > 127 && rune < 2048 {
		u := make([]byte, 2)
		u[1] = byte(0x00000080 | (rune & 0x0000003f))
		u[0] = byte(0x000000c0 | ((rune >> 6) & 0x0000001f))
		return u, nil
	} else if rune > 2047 && rune < 65536 {
		u := make([]byte, 3)
		u[2] = byte(0x00000080 | (rune & 0x0000003f))
		u[1] = byte(0x00000080 | ((rune >> 6) & 0x0000003f))
		u[0] = byte(0x000000e0 | ((rune >> 12) & 0x0000000f))
		return u, nil
	} else if rune > 65535 && rune < 1114112 {
		u := make([]byte, 4)
		u[3] = byte(0x00000080 | (rune & 0x0000003f))
		u[2] = byte(0x00000080 | ((rune >> 6) & 0x0000003f))
		u[1] = byte(0x00000080 | ((rune >> 12) & 0x0000003f))
		u[0] = byte(0x000000f0 | ((rune >> 18) & 0x00000007))
		return u, nil
	} else {
		u := make([]byte, 1)
		return u, errors.New("Not_A_Unicode_Rune")
	}
}

//input utf8's []byte, output utf16's []byte (BigEndian), support surrogate pair
func Utf8ToUtf16BE(input []byte) ([]byte, error) {
	var output []byte
	rune, e := Utf8GetRunes(input) //rune is []uint32
	if e != nil {
		return output, e
	}
	for i := 0; i < len(rune); i++ {
		if rune[i] < 0x10000 { //base 0
			headbyte := byte(rune[i] >> 8)
			tailbyte := byte(rune[i])
			output = append(output, headbyte, tailbyte)
		} else if rune[i] > 0xFFFF && rune[i] < 0x100001 { // surrogate pair
			var runei uint32 = 0
			runei = rune[i] - 0x10000
			lead_head := byte(((runei >> 18) & 0x00000003) | 0x000000D8)
			lead_tail := byte(runei >> 10)
			trail_head := byte(((runei >> 8) & 0x00000003) | 0x000000DC)
			trail_tail := byte(runei)
			output = append(output, lead_head, lead_tail, trail_head, trail_tail)
		} else {
			return output, errors.New("Not_A_Unicode_Rune")
		}
	}
	return output, nil
}

//input utf8's []byte, output utf16's []byte (LittleEndian), support surrogate pair
func Utf8ToUtf16LE(input []byte) ([]byte, error) {
	var output []byte
	rune, e := Utf8GetRunes(input) //rune is []uint32
	if e != nil {
		return output, e
	}
	for i := 0; i < len(rune); i++ {
		if rune[i] < 0x10000 { //base 0
			headbyte := byte(rune[i] >> 8)
			tailbyte := byte(rune[i])
			output = append(output, tailbyte, headbyte)
		} else if rune[i] > 0xFFFF && rune[i] < 0x100001 { // surrogate pair
			var runei uint32 = 0
			runei = rune[i] - 0x10000
			lead_head := byte(((runei >> 18) & 0x00000003) | 0x000000D8)
			lead_tail := byte(runei >> 10)
			trail_head := byte(((runei >> 8) & 0x00000003) | 0x000000DC)
			trail_tail := byte(runei)
			output = append(output, lead_tail, lead_head, trail_tail, trail_head)
		} else {
			return output, errors.New("Not_A_Unicode_Rune")
		}
	}
	return output, nil
}

//input utf16's []byte (BigEndian)/with BOM, output utf8's []byte
func Utf16BEToUtf8(input []byte) ([]byte, error) {
	var output []byte
	if (len(input) % 2) != 0 {
		return output, errors.New("input bytes' num is odd")
	}
	var i int = 0
	if input[0] == 0xFE || input[0] == 0xFF {
		i = 2
	}
	for i < len(input) {
		if input[i] >= 0xD8 && input[i] < 0xDD { //surrogate pair
			rune := uint32(input[i]&0x03)<<8 | uint32(input[i+1])
			rune = rune<<10 | (uint32(input[i+2]&0x03)<<8 | uint32(input[i+3]))
			bytes, e := Uint32ToUtf8(rune + 0x10000)
			if e != nil {
				return output, e
			}
			output = append(output, bytes...)
			i += 4
		} else { //单一码元base 0
			rune := uint32(input[i])<<8 | uint32(input[i+1])
			bytes, e := Uint32ToUtf8(rune)
			if e != nil {
				return output, e
			}
			output = append(output, bytes...)
			i += 2
		}
	}
	return output, nil
}

//input utf16's []byte (LittleEndian)/with BOM, output utf8's []byte
func Utf16LEToUtf8(input []byte) ([]byte, error) {
	var output []byte
	if (len(input) % 2) != 0 {
		return output, errors.New("input bytes' num is odd")
	}
	var i int = 0
	if input[0] == 0xFE || input[0] == 0xFF {
		i = 2
	}
	for i < len(input) {
		if input[i+1] >= 0xD8 && input[i+1] < 0xDD { //surrogate pair
			rune := uint32(input[i+1]&0x03)<<8 | uint32(input[i])
			rune = rune<<10 | (uint32(input[i+3]&0x03)<<8 | uint32(input[i+2]))
			bytes, e := Uint32ToUtf8(rune + 0x10000)
			if e != nil {
				return output, e
			}
			output = append(output, bytes...)
			i += 4
		} else { //单一码元base 0
			rune := uint32(input[i+1])<<8 | uint32(input[i])
			bytes, e := Uint32ToUtf8(rune)
			if e != nil {
				return output, e
			}
			output = append(output, bytes...)
			i += 2
		}
	}
	return output, nil
}
