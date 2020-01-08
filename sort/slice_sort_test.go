package main

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Year   int
	Length time.Duration
}

func Test(t *testing.T) {
	/*
	   arr := []string{"B", "D", "c", "C", "A", "!", "1"}
	   print(arr)
	   //sort.Strings(arr)
	   sort.Sort(sort.StringSlice(arr))
	   print(arr)
	   sort.Reverse(sort.StringSlice(arr))
	   print(arr)
	*/

	// 指针数组，对于交换数据频繁的操作性能要好一些
	var tracks = []*Track{
		{"Python", 2020, length("158s")},
		{"Java", 2009, length("1m18s")},
		{"GoLang", 2012, length("3m38s")},
		{"GoLang", 2011, length("3m38s")},
	}
	/*
	   printTracks(tracks)
	   sort.Sort(byTitle(tracks))
	   printTracks(tracks)
	   // 使用了组合，重写了方法
	   sort.Sort(sort.Reverse(byTitle(tracks)))
	   printTracks(tracks)
	*/
	printTracks(tracks)
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)

	vl := []int{3, 1, 4, 1}
	fmt.Println(vl)
	fmt.Println(sort.IntsAreSorted(vl))
	sort.Ints(vl)
	fmt.Println(sort.IntsAreSorted(vl))
	fmt.Println(vl)
	sort.Sort(sort.Reverse(sort.IntSlice(vl)))
	fmt.Println(sort.IntsAreSorted(vl))
	fmt.Println(vl)
}

// 1 define type
type byTitle []*Track

// 2 implement interface
func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// 1 define type
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

// 2 implement interface
func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func print(data interface{}) {
	fmt.Println(data)
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 6, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Year, t.Length)
	}
	tw.Flush()
}
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
