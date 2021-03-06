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
	spl := strings.Split(s, "")
	min := 10000
	for i:= 'a'; i<='z'; i++ {
		count := 0
		next := append([]string{}, spl...)
		LOOP:
		for {
			flag := true
			if len(next) == 0 {
				count = 10000
				break
			}
			//if string(i) == "v" {
			//	fmt.Println(string(i), next)
			//}
			get := make([]string, len(next)-1)
			if next[0] != string(i) {
				flag = false
			}
			for c := 1; c < len(next); c++ {
				if next[c] != string(i) {
					flag = false
				}
				if next[c-1] == string(i) || next[c] == string(i) {
					get[c-1] = string(i)
				} else {
					get[c-1] = next[c-1]
				}
			}
			if flag {
				// fmt.Println(next, get, string(i))
				if min > count {
					min = count
				}
				break LOOP
			} else {
				next = append([]string{}, get...)
				count++
			}
		}
	}
	fmt.Println(min)
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