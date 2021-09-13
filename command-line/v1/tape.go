package poker

import (
	"os"
)

// 用来写入数据到文件

type tape struct {
	file *os.File
}

// 实现io.Write接口
func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
    return t.file.Write(p)
}