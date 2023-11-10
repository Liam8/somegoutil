package somegoutil

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "exist 1",
			args: args{"a", "123"},
		},
		{
			name: "exist 2",
			args: args{"ab", "fefe"},
		},
	}
	m := NewConMap[string, string]()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Put(tt.args.key, tt.args.value)
			v, exist := m.Get(tt.args.key)
			if !exist {
				t.Errorf("Expect exist = true but got %v", exist)
			}
			if v != tt.args.value {
				t.Errorf("Expect %v but got %v", tt.args.value, v)
			}
		})
	}
}

func BenchmarkPut(b *testing.B) {
	m := NewConMap[int32, string]()
	for i := 0; i < b.N; i++ {
		m.Put(rand.Int31(), "avalue")
	}
}

func BenchmarkGet(b *testing.B) {
	m := NewConMap[string, string]()
	m.Put("akey", "avalue123")
	for i := 0; i < b.N; i++ {
		m.Get("akey")
	}
}

func TestString(t *testing.T) {
	m := NewConMap[string, string]()
	m.Put("akey", "avalue123")
	str := fmt.Sprintf("%v", m)
	exp := "map[akey:avalue123]"
	if str != exp {
		t.Errorf("Expect %v but got %v", exp, str)
	}
}