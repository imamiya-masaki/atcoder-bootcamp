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
	k := scanInt()
	p := scanInts(n)
	ruiseki := make([]int, n)
	ruiseki[0] = p[0]
	for i:= 1; i<n;i++ {
		ruiseki[i] = p[i]+ruiseki[i-1]
	}
	herasu := 0
	max := 0
	for i:=k-1; i<n; i++ {
		if max < ruiseki[i]-herasu {
			max = ruiseki[i]-herasu
		}
		// fmt.Println(i,k, ruiseki)
		herasu = ruiseki[i-(k-1)]
	}
	val := k-1
	max = max - (k-1)
	kitaiti := float64(0)
	keisuu := 10000000
	for i:= 1; i<= max ; i++ {
		kitaiti += float64(i*keisuu/max)
	}
	// kitaiti += float64(val)
	kitaiti /= float64(keisuu)
	kitaiti += float64(val)
	fmt.Println(kitaiti)
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