package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const path = "hello.txt"

func Test1(t *testing.T) {
	ctt := ReadFile(path)
	fmt.Println(string(ctt))
}

func Test2(t *testing.T) {
	f := func(line string) {
		fmt.Printf("%s\n", line)
	}

	err := ReadFileStream(path, f)
	if err != nil {
		fmt.Printf("ReadFileStream err:%+v\n", err)
	}
}

func Test3(t *testing.T) {
	f := func(lineb []byte) {
		fmt.Printf("%s\n", string(lineb))
	}

	err := ReadFileSlice(path, f)
	if err != nil {
		fmt.Printf("ReadFileSlice err:%+v\n", err)
	}
}

// Fragment Upload : this suit for read big file, read slice
func ReadFileSlice(path string, handle func([]byte)) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println(2 << 10)
	s := make([]byte, 2<<10) // 2M
	for {
		switch nr, err := f.Read(s); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, " read file error:%+v", err)
			os.Exit(1)
		case nr == 0: // EOF
			return nil
		case nr > 0:
			handle(s[:nr])
		}
	}
	return nil
}

// this suit for read log file or formatted file
func ReadFileStream(path string, handle func(string)) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		// TODO compare the performance of the three type later
		/*
		   lineb, _, err := buf.ReadLine()
		   if err != nil {
		       if err == io.EOF {
		           return nil
		       }
		       return err
		   }
		   line := strings.TrimSpace(string(lineb))
		   handle(line)
		*/
		lineb, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		line := strings.TrimSpace(string(lineb))
		handle(line)
		/*
			line, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			line = strings.TrimSpace(line)
			handle(line)
		*/

	}
	return nil
}

// 1.suit for small file
func ReadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("read file err:%+v\n", err)
	}
	return content
}
