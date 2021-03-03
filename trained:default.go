package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan_init()
	// input.scan(), input.scan_int()
	n := scanInt()
	a := make([]int, n)
	for i:= 0; i<n; i++ {
		a[i] = scanInt() - 1
	}
	// b := make([]int, n)
	//for i := 0; i<n; i++ {
	//	b[a[i]] = i
	//}
	get := -1
	target := 0
	for i:=0; i<n; i++ {
		if (target == 1) {
			get = i
			break
		}
		target = a[target]
	}
	fmt.Println(get)
}

var sc = bufio.NewScanner(os.Stdin)

func scan_init() {
	sc.Split(bufio.ScanWords)
}
func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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