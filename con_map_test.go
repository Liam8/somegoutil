package somegoutil

import "testing"

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
			args: args{"a","123"},
		},
		{
			name: "exist 2",
			args: args{"ab","fefe"},
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