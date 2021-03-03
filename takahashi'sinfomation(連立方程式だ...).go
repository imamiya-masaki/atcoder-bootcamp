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
	c := make([][]int, 3)
	a_s := make([]int, 3)
	a_s[0] = 0
	for i:=0; i<3; i++ {
		c[i] = scanInts(3)
	}
	b_s := make([]int, 3)
	for i:=0; i<3; i++ {
		// a_0からb_0 ~ b_2を決める
		b_s[i] = c[0][i] - a_s[0]
	}
	for i:=1; i<3; i++ {
		a_s[i] = c[i][1] - b_s[1]
	}
	flag := true
	for i:= 0; i<3; i++ {
		for j:=0; j<3; j++ {
			if a_s[i] + b_s[j] != c[i][j] {
				// fmt.Println(i,j, c[i][j], a_s[i])
				flag = false
				break
			}
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	// fmt.Println(c)
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