package somegoutil

import (
	"testing"
	"time"
)

func TestConQueue_Enqueue(t *testing.T) {
	type fields struct {
		buffer chan int
	}
	type args struct {
		item int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "enqueue a number",
			fields:  fields{buffer: make(chan int, 10)},
			args:    args{item: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConQueue[int]{
				buffer: tt.fields.buffer,
			}
			if err := r.Enqueue(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("ConQueue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConQueue_Dequeue(t *testing.T) {
	type fields struct {
		buffer chan int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		{
			name:    "dequeue a number",
			fields:  fields{buffer: make(chan int, 10)},
			want:    0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConQueue[int]{
				buffer: tt.fields.buffer,
			}
			got, got1 := r.Dequeue()
			if got != tt.want {
				t.Errorf("ConQueue.Dequeue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConQueue.Dequeue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestConQueue_EnDequeue(t *testing.T) {
	q := NewConQueue[string](2)
	item1 := "a"
	item2 := "b"
	in1 := q.Enqueue(item1)
	in2 := q.Enqueue(item2)
	if in1!=nil || in2!=nil {
		t.Errorf("ConQueue.Eequeue() failed")
	}
	in3 := q.Enqueue(item2)
	if in3 == nil {
		t.Errorf("ConQueue.Eequeue() should be failed, but got no error.")
	}
	out1, ok1 := q.Dequeue()
	if !ok1 || out1 != item1 {
		t.Errorf("ConQueue.Dequeue() got = %v, want %v", out1, item1)
	}
	out2, ok2 := q.Dequeue()
	if !ok2 || out2 != item2 {
		t.Errorf("ConQueue.Dequeue() got = %v, want %v", out2, item2)
	}

}

func TestConQueue_DequeueBlock(t *testing.T) {
	q := NewConQueue[string](2)
	item1 := "a"
	_ = q.Enqueue(item1)

	start := time.Now()
	q.DequeueWithBlock(1000)	//won't block
	q.DequeueWithBlock(10)		//should block for 10ms
	elapsed := time.Since(start)
	if elapsed < 10 * time.Millisecond {
		t.Errorf("Cost %v, shorter than expected.", elapsed)
	}
	if elapsed > 20 * time.Millisecond {
		t.Errorf("Cost %v, longer than expected.", elapsed)
	}
}

func TestConQueue_DequeueBlockAbort(t *testing.T) {
	q := NewConQueue[string](2)
	item1 := "a"
	go func ()  {
		time.Sleep(10 * time.Millisecond)
		_ = q.Enqueue(item1)
	} ()

	start := time.Now()
	out, ok := q.DequeueWithBlock(1000)	//block and then return
	elapsed := time.Since(start)
	if !ok || out != item1 {
		t.Errorf("ConQueue.DequeueWithBlock() got = %v (ok=%v), want = %v", out, ok, item1)
	}
	if elapsed > 500 * time.Millisecond {
		t.Errorf("Cost %v, longer than expected.", elapsed)
	}
}