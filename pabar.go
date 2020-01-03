package main

// 测试进度条
import (
	"github.com/qianlnk/pgbar"
	"time"
)

func main() {
	pgb := pgbar.New("upload file")
	b := pgb.NewBar("upload bar", 1)

	b.SetSpeedSection(10, 100)
	for {
		b.Add()
		hi()
		return
	}
}

func hi() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 100)
	}
}
