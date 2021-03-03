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
	T := scanInt()
	t := scanInts(n)
	sum := 0
	oyu := 0
	minus := 0
	for i:= 0; i< n; i++ {
		// ラストは特殊処理
		take := t[i]
		// fmt.Println(i, oyu, take, sum)
		if oyu > take {
			oyu = take + T
			continue
		} else {
			sum += oyu - minus
			minus = take
			oyu = T + take
		}
	}
	if oyu-minus > 0 {
		sum += oyu-minus
	}
	fmt.Println(sum)
}

var sc = bufio.NewScanner(os.Stdin)
func relu (a int) int{
	if a < 0 {
		return 0
	} else {
		return a
	}
}
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