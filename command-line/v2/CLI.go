package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in *bufio.Scanner
}
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in: bufio.NewScanner(in),
	}
}

// 给对应的名字record win
func (c *CLI) PlayPoker() {
	name := c.extractWinner(c.readLine())
	c.playerStore.RecordWin(name)
}

// 根据命令行内容抽取名字
func (c *CLI) extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

// 读取命令行内容
func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}