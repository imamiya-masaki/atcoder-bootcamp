package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan_init()
	n := scanInt()
	a := scanInts(n)
	memo := map[int]int{}
	memoKeisan := map[int]int{}
	for i:= 0; i<n; i++ {
		if _,ok := memo[a[i]]; !ok {
			memo[a[i]] = 0
		}
		memo[a[i]]++
	}
	sum := 0
	for key,val := range memo {
		if val >= 2 {
			keisan := val*(val-1)/2
			sum += val*(val-1)/2
			memoKeisan[key] = keisan
		} else {
			memoKeisan[key] = 0
		}
	}
	for i:=0; i<n; i++ {
		take := sum
		take -= memoKeisan[a[i]]
		if memo[a[i]] > 1 {
			take += (memo[a[i]]-1)*(memo[a[i]]-2)/2
		}
		fmt.Println(take)
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scan_init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}
func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
func scanInts(n int) []int {
	take := make([]int, n)
	for i := 0; i < n; i++ {
		take[i] = scanInt()
	}
	return take
}
func scan() string {
	sc.Scan()
	return sc.Text()
}

var rdr = bufio.NewReaderSize(os.Stdin, 200000)

func readLine() string {
	buf := make([]byte, 0, 200000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
func readLineToNumber() int {
	S := readLine()
	number, _ := strconv.Atoi(S)
	return number
}
func readLineToSlice() []string {
	S := readLine()
	list := strings.Split(S, "")
	return list
}
func readLineToNumberSlice(memo map[string]int) []int {
	// err時は-1を代入
	S := readLine()
	intList := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if val, ok := memo[string(S[i])]; ok {
			intList[i] = val
		} else {
			intList[i] = -1
		}
	}
	return intList
}