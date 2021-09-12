package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord = "Go!"
	countdownStart = 3
)
type Sleeper interface {
    Sleep()
}

type ConfigurableSleeper struct {
    duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
    time.Sleep(o.duration)
}
// Countdown 应该在第一个打印之前 sleep，然后是直到最后一个前的每一个，例如：
// Sleep
// Print N
// Sleep
// Print N-1
// Sleep
// etc 测试时执行的顺序无法确定
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
        fmt.Fprintln(out, i)
    }
	sleeper.Sleep()
    fmt.Fprint(out, finalWord)
}
// func Countdown(out io.Writer, sleeper Sleeper) {
//     for i := countdownStart; i > 0; i-- {
//         sleeper.Sleep()
//     }

//     for i := countdownStart; i > 0; i-- {
//         fmt.Fprintln(out, i)
//     }

//     sleeper.Sleep()
//     fmt.Fprint(out, finalWord)
// }

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
    Countdown(os.Stdout, sleeper)
}