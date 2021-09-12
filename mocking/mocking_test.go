package main

import (
	"bytes"
	"reflect"
	"testing"
)

/*如果你的模拟代码变得很复杂，或者你需要模拟很多东西来测试一些东西，那么你应该 倾听 那种糟糕的感觉，并考虑你的代码。通常这是一个征兆：
你正在进行的测试需要做太多的事情
把模块分开就会减少测试内容
它的依赖关系太细致
考虑如何将这些依赖项合并到一个有意义的模块中
你的测试过于关注实现细节
最好测试预期的行为，而不是功能的实现*/

// 测试sleep执行次数的spy, 通过实现一个sleep的接口，来进行测试。main函数使用另外的struct来实现接口。
// 这样就保证运行countdown函数时，主函数是实际情景，而测试环境只是模拟。不需要sleep 1s
type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

// 测试sleep和print执行顺序的spy, 通过实现Sleeper.sleep，io.write的接口，来进行测试。main函数使用另外的struct来实现接口。
// 这样就保证运行countdown函数时，主函数是实际情景，而测试环境只是模拟。
const write = "write"
const sleep = "sleep"
type CountdownOperationsSpy struct {
    Calls []string
}
func (s *CountdownOperationsSpy) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

func TestCountdown(t *testing.T) {
	t.Run("test count times", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
	
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		}
	})

	t.Run("test order of func", func(t *testing.T) {
		spySleepPtr := &CountdownOperationsSpy{}
		Countdown(spySleepPtr, spySleepPtr)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPtr.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPtr.Calls)
		}
	})

}


