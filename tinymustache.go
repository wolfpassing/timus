package timus

import (
	"regexp"
	"strings"
)

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	SignPlus     = 0x2B
	SignMinus    = 0x2D
	SignMultiply = 0x2A
	SignDivide   = 0x2F
	SignPOpen    = 0x28
	SignPClose   = 0x29
	SignZero     = 0x00
)

func strpbrk(x []byte, st int) bool {
	for i := st; i < len(x); i++ {
		if x[i] == SignZero {
			return false
		}
		if x[i] == SignPOpen || x[i] == SignPClose || x[i] == SignMultiply || x[i] == SignPlus || x[i] == SignMinus || x[i] == SignDivide {
			return true
		}
	}
	return false
}

func atof(x []byte, st int) float64 {
	var z []byte
	for i := st; i < len(x); i++ {
		if x[i] == 0x00 {
			break
		}
		z = append(z, x[i])
	}
	f, err := strconv.ParseFloat(string(z), 64)
	if err != nil {
		return 0.0
	}
	return f
}

func parse(x []byte, st int) float64 {
	kl := 1
	p := st + 1
	if !strpbrk(x, st) {
		return atof(x, st)
	}
	if x[st] == SignPOpen {
		for (x[p] != SignZero) && (kl != 0) {
			if x[p] == SignPOpen {
				kl++
			}
			if x[p] == SignPClose {
				kl--
			}
			p++
		}
		if (kl == 0) && (x[p] == SignZero) {
			p--
			x[p] = SignZero
			return parse(x, st+1)
		}
	}
	for i := 0; i <= 1; i++ {
		kl = 0
		p = st
		for x[p] != SignZero {
			if x[p] == SignPOpen {
				kl++
			}
			if x[p] == SignPClose {
				kl--
			}
			if (kl == 0) && (i == 0) {
				if x[p] == SignMultiply {
					x[p] = SignZero
					return parse(x, st) * parse(x, p+1)
				}
				if x[p] == SignDivide {
					x[p] = SignZero
					return parse(x, st) / parse(x, p+1)
				}
				if x[p] == SignPlus {
					x[p] = SignZero
					return parse(x, st) + parse(x, p+1)
				}
				if x[p] == SignMinus {
					x[p] = SignZero
					return parse(x, st) - parse(x, p+1)
				}
			}
			p++
		}
	}
	return 0.0
}

func (m *TinyMustache)Evaluate(s string) float64 {
	//fmt.Println("Parse (", s, ")")
	s = strings.Replace(s," ","",-1)
	bs := []byte(s + "\x00\x00")
	return parse(bs, 0)
}

type TinyMustache struct {
	MustacheMap map[string]string
	PfMu        *regexp.Regexp
	RmMu        *regexp.Regexp
}

func NewMustache() *TinyMustache {
	obj := TinyMustache{}
	obj.MustacheMap = make(map[string]string)
	obj.PfMu = regexp.MustCompile("^{{2}[^{ }]+}{2}$")
	obj.RmMu = regexp.MustCompile("([a-zA-Z0-9_.,]+)")
	return &obj
}

func (m *TinyMustache) Merge(x map[string]string) {
	for key, value := range x {
		m.MustacheMap[key] = value
	}
}

func (m *TinyMustache) Add(key string, i interface{}) error {
	//Regex check if perfect key = ^\{{2}[^{ }]+\}{2}$
	if !m.PfMu.Match([]byte(key)) {
		key = m.RmMu.FindString(key)
		key = "{{" + key + "}}"
	}
	switch v := i.(type) {
	case string:
		m.MustacheMap[key] = v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		m.MustacheMap[key] = fmt.Sprintf("%d", v)
	case float32, float64:
		m.MustacheMap[key] = fmt.Sprintf("%0.15f", v)
	default:
		return errors.New("unknown type")
	}
	return nil
}

func (m *TinyMustache) Mustache(formular string) string {
	ss := formular
	for i, v := range m.MustacheMap {
		fmt.Println(i, v)
		re := regexp.MustCompile(i)
		ss = re.ReplaceAllString(ss, v)
	}
	return ss
}
