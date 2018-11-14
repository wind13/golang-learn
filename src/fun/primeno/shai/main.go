package main

import (
	"fmt"
	"time"
)

const MAX = 100 // 00000

func main() {
	t := time.Now()
	fmt.Print(2, " ")
	var rs [MAX + 1]int
	for i := 0; i < MAX+1; i++ {
		rs[i] = 0
	}
	for i := 3; i < MAX+1; i += 2 {
		// 此三行可在算千万时减少 5 秒
		if rs[i] == 1 {
			continue
		}
		for j := i + i; j < MAX+1; j += i {
			rs[j] = 1
			// fmt.Print(j, " ")
		}
	}

	// /*
	for i := 3; i < MAX+1; i += 2 {
		if rs[i] == 0 {
			fmt.Print(i, " ")
		}
	}
	// */
	// aList := []int
	// 	for i in range(0, MAX+1):
	//     aList.append(0)

	// for i in range(3, MAX+1, 2):
	//     for j in range(i+i, MAX+1, i):
	//         aList[j] = 1

	// for i in range(3, MAX+1, 2):
	//     if(aList[i]==0):
	//         print(i, end=' ')
	e := time.Now()
	fmt.Println("time is ", e.Sub(t).Seconds())
	// 10万以内 0.16秒，1千万18秒 --> 13秒
}
