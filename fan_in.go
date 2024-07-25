package somegoutil

import "time"

func FanIn[T any](input1, input2 <-chan T) <-chan T {
	c := make(chan T)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func FanInWithTimeout[T any](input1, input2 <-chan T, timeout time.Duration) <-chan T {
	c := make(chan T)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			case <-time.After(timeout):
				close(c)
				return
			}
		}
	}()
	return c
}
