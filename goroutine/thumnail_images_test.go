package main

import (
	"log"
	"os"
	"sync"
	"testing"

	"github.com/adonovan/gopl.io/tree/master/ch8/thumbnail" // TODO fixme ,it's too late
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

// get total values of thumbnail
func makeThumbnail6(filenames <-chan string) int64 {

	// 1. create total size receiver
	sizes := make(chan int64)
	// 2. create a counter
	var wg sync.WaitGroup
	// 3. write the item size
	for f := range filenames {
		wg.Add(1) //  counter++
		go func(f string) {
			defer wg.Done() // counter--
			thum, err := thumbnail.ImageGroup(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thum)
			sizes <- info.Size()
		}(f)

	}
	// 4. wait and close
	go func() {
		wg.Waite()
		close(sizes)
	}()
	// 5. get the total
	var rs int64
	// if chan used to foreach  := range sizes
	// if chan used to blocking get value  <- sizes
	for size := range sizes {
		rs += size
	}
	return size
}

func makeThumbnail5(filenames []string) (thumbnails []string, err error) {
	type item struct {
		tf  string // thumbnailfile
		err error
	}
	ch := make(chan item)
	for _, k := range filenames {
		go func(f string) {
			var it item
			it.tf, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(k)
	}

	for range thumbnails {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbnails = append(thumbnails, it.tf)
	}
	return thumbnails, nil
}

/*
func makeThumbnail4(filenames []string) error {
	ch := make(chan error)
	for _, file := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			ch <- err
		}(f)
	}

	for range filenames {
		if err := <-ch; err != nil {
			return err // error : 会goroutine泄露，导致goroutine不能正确关闭
		}
	}

}

func makeThumbnail3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func() {
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

*/
