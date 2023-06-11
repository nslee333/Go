package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	time.Sleep(1 * time.Second)
	s.Calls++
}

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const countdownStart = 3

func main() {
	sleeper := &DefaultSleeper{}

	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprintln(out, finalWord)
}
