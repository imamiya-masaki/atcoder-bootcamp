package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scan_init()
	n := scanInt()
	X := scanInt()
	x := make([]int, n)
	for i:= 0; i<n; i++ {
		x[i] = scanInt()
	}
	x = append(x, X)
	sort.Ints(x)
	aida := make([]int, n)
	for i:=1; i <= n; i++ {
		// fmt.Println(x[i], x[i-1], x[i] - x[i-1])
		aida[i-1] = x[i] - x[i-1]
	}
	take := aida[0]
	for i:= 1; i<len(aida); i++ {
		take = gcd(take, aida[i])
	}
	fmt.Println(take)
}

func gcd(a int, b int) int {
		// aの方が小さくなるように設定する
		// ユークリッドの互除法で最大公約数を出す
		if a>b {
			//swap
			tmp := a
			a = b
			b = tmp
		}
		if a == 0 {
			return b
		}
		r := b % a
		if r != 0 {
			return gcd(r,a)
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