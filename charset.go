// charset project charset.go
// Author: Cof-Lee
package charset

import (
	"errors"
)

//input utf8's []byte, output []uint32's Runes
func Utf8GetRune(b []byte) ([]uint32, error) {
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
