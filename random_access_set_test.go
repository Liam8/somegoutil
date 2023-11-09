package somegoutil

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "add number1",
			args: 123,
			want: true,
		},
		{
			name: "add number2",
			args: 456,
			want: true,
		},
	}
	set := NewRandSet[int]()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set.Add(tt.args)
			if got := set.Exists(tt.args); got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "remove number1",
			args: 123,
			want: false,
		},
		{
			name: "remove number2",
			args: 456,
			want: false,
		},
	}
	set := NewRandSet[int]()
	set.Add(123)
	set.Add(456)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set.Remove(tt.args)
			if got := set.Exists(tt.args); got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomPick(t *testing.T) {
	set := NewRandSet[int]()
	set.Add(123)
	set.Add(456)
	for i := 10; i > 0; i-- {
		v := set.RandomPick()
		if v != 123 && v != 456 {
			t.Errorf("RandomPick() returned unpected value = %v", v)
		}
	}
}
