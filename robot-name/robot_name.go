package robotname


import (
	"math/rand"
	"strings"
)


const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"
var names map[string]bool


type Robot struct {
	name string
}


func init() {
	names = make(map[string]bool)
}

func genName() string {
	var builder strings.Builder
	for i:=0; i<5; i++ {
		if i<2 {
			builder.WriteByte(chars[rand.Intn(len(chars))])
		} else {
			builder.WriteByte(digits[rand.Intn(len(digits))])
		}
	}
	return builder.String()
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		var name string
		for {
			name = genName()
			if !names[name] {
				break
			}
		}
		names[name] = true
		r.name = name
	}
	return r.name, nil
}


func (r *Robot) Reset() {
	r.name = ""
}
