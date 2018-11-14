package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/lennon-guan/pipe"
)

func divideExactly(n int) func(int) bool {
	return func(s int) bool {
		return !(s > 1 && s > n && s/n > 1 && s%n == 0)
	}
}

func primeNo(s []int) []int {
	var r []int
	for i := 0; i < len(s); i++ {
		f := s[i]
		t := s[len(s)-1]
		if f > 1 && t/f > 1 {
			s = pipe.NewPipe(s).Filter(divideExactly(f)).ToSlice().([]int)
		} else {
			continue
			// r = append(r, f)
		}
	}
	return append(r, s...)
}

func genStr(n int, s string) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(s)
	}
	return buffer.String() // 拼接结果
}

func main() {
	t := time.Now()
	var ns [100]int
	for i := 0; i < len(ns); i++ {
		ns[i] = i + 1
	}
	var sl = ns[:]
	// fmt.Println("prime number", ns[78], ns[99])
	// fmt.Println("head: ", pipe.NewPipe(ns).Filter(divideExactly(2)).Filter(divideExactly(3)).Filter(divideExactly(5)).Filter(divideExactly(7)).ToSlice().([]int))
	os := primeNo(sl)
	fmt.Println("endy: ", os)

	/*
		const l = 30
		for _, o := range os {
			b := strconv.FormatInt(int64(o), 2)
			so := strconv.Itoa(o)
			lo := len(so)
			lb := len(b)
			e := l - lb - lo
			// fmt.Println(lo, e, lb)
			fmt.Println(so, genStr(e, " "), b)
		}
	*/
	// fmt.Println(strconv.FormatInt(4, 2))
	fmt.Println(len(os))
	e := time.Now()
	fmt.Println("time is ", e.Sub(t).Seconds())
	// 10万以内的素数计算，用时308秒
}
