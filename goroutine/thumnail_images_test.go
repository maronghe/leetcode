package main

import (
	"log"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func makeThumbnail3(filenames []string) {
	ch = make(chan struct{})
	for _, f := range filenames {
		go func() {
			thumbna
			thumbnail.ImageFile(f) // error, 匿名共享变量，持续更新，可能当一个goroutine读取的时候f已经是一namefile最后一个了
			ch <- struct{}{}       // 用于计数
		}()
	}
	for range filenames {
		<-ch
	}
}

func makeThumbnail2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // error, do not wait result and returned
	}
}

// 串行处理效率低下，且不依赖于其他逻辑
func makeThumbnail(filenames []string) {
	for _, f := range filenames {
		//
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Printf("err : %+v", err)
		}
	}
}
