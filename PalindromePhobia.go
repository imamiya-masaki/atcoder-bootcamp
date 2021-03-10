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
	s := scan()
	boxes := make([]int, 3)
	for i:=0; i<len(s); i++ {
		switch s[i] {
		case 'a':
			boxes[0]++
		case 'b':
			boxes[1]++
		case 'c':
			boxes[2]++
		}
	}
	flag := false
	if abs(boxes[0] - boxes[1]) >1 {
		flag = true
	}
	if abs(boxes[1] - boxes[2]) >1 {
		flag = true
	}
	if abs(boxes[0] - boxes[2]) >1 {
		flag = true
	}
	if flag {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
func abs(a int) int{
	if a<0 {
		return -a
	} else {
		return a
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