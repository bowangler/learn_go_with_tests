package main

import (
    "io"
    "fmt"
    "os"
    "time"
)

type Sleeper interface {
    Sleep()
}

type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}
const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart; i > 0; i-- {
        time.Sleep(1 * time.Second)
        fmt.Fprintln(out, i)
    }
    time.Sleep(1 * time.Second)
    fmt.Fprint(out, finalWord)
}

func main() {
    Countdown(os.Stdout)
}
