package somegoutil

import (
	"fmt"
	"testing"
	"time"
)

func Test_FanIn(t *testing.T) {
	c1 := make(chan string)
	go func() {
		c1 <- "a"
	}()
	c2 := make(chan string)
	go func() {
		c2 <- "b"
	}()
	c3 := FanIn(c1, c2)
	var res []string
	res = append(res, <-c3)
	res = append(res, <-c3)

	fmt.Println(res)
	if !contains(res, "a") {
		t.Errorf("not receive a")
	}
	if !contains(res, "b") {
		t.Errorf("not receive b")
	}
}

func Test_FanIn2(t *testing.T) {
	c1 := make(chan int)
	go func() {
		c1 <- 1
	}()
	c2 := make(chan int)
	go func() {
		c2 <- 2
	}()
	c3 := FanIn(c1, c2)
	var res []int
	res = append(res, <-c3)
	res = append(res, <-c3)
	fmt.Println(res)
	if !contains(res, 1) {
		t.Errorf("not receive 1")
	}
	if !contains(res, 2) {
		t.Errorf("not receive 2")
	}
}

func Test_FanInWithTimeout(t *testing.T) {
	c1 := make(chan string)
	go func() {
		c1 <- "a"
	}()
	c2 := make(chan string)
	go func() {
		c2 <- "b"
	}()
	c3 := FanInWithTimeout(c1, c2, 50 * time.Millisecond)
	var res []string
	res = append(res, <-c3)
	res = append(res, <-c3)

	fmt.Println(res)
	if !contains(res, "a") {
		t.Errorf("not receive a")
	}
	if !contains(res, "b") {
		t.Errorf("not receive b")
	}
	<-c3
}

// contains checks if a slice contains a specific element (generic version)
func contains[T comparable](slice []T, element T) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}
