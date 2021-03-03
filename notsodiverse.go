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

type number struct {
	number int
	count int
}
var memo map[int]int
func main() {
	scan_init()
	n := scanInt()
	k := scanInt()
	a := scanInts(n)
	memo = map[int]int{}
	features := 0
	for i:=0; i<n; i++ {
		if _,ok := memo[a[i]]; !ok {
			memo[a[i]] = 0
			features++
		}
		memo[a[i]]++
	}
	target := []number{}
	for key,val := range memo {
		target = append(target, number{key, val})
	}
	sort.Slice(target,func (i,j int) bool{
		return target[i].count > target[j].count
	})
	sum := 0
	for i:=k; i<len(target); i++ {
		sum += target[i].count
	}
	fmt.Println(sum)
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