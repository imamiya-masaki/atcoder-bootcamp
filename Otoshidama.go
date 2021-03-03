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
	y := scanInt()
	for i:= 0; i<=n; i++ {
		man := i*10000
		for j:= 0; j<=n; j++ {
			gosen := j*5000
			get := y - (man + gosen)
			nokori := n - (i+j)
			// fmt.Println(nokori, get, get/1000, man, gosen)
			if get/1000 == nokori && nokori >= 0 {
				fmt.Println(i)
				fmt.Println(j)
				fmt.Println(get/1000)
				// fmt.Println(string(i) + " " + string(j) + " " + string(get/1000))
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
var clear []int
var clearFlag bool
var sc = bufio.NewScanner(os.Stdin)

var memoString map[string]bool

func tansaku (sel []int, targetMoney int) {

	// fmt.Println(sel, targetMoney, "tansaku")
	chSel := append([]int{}, sel...)
	if targetMoney == 0 {
		clear = chSel
		clearFlag = true
	}
	if targetMoney >= 4  && chSel[1] >= 1{
		nextSel := append([]int{}, chSel...)
		nextSel[0] += 5
		nextSel[1] -= 1
		tansaku(nextSel, targetMoney-4)
	}
	if targetMoney >= 9 && chSel[2] >= 1{
		nextSel := append([]int{}, chSel...)
		nextSel[0] += 10
		nextSel[2] -= 1
		tansaku(nextSel, targetMoney-9)
	}
}

func sortIntMemos (ints []int) bool {
	// 順不同で入れる
	sort.Ints(ints)
	return intMemos(ints)
}
func intMemos (ints []int) bool{
	// intsでもらった任意の数に対して,memo化しなければ保存するあればfalseを返す
	// sort.Ints(ints)
	strins := make([]string,len(ints))
	for i:= 0; i<len(ints); i++ {
		strins[i] = strconv.Itoa(ints[i])
	}
	text := strings.Join(strins,"")
	if _,ok := memoString[text]; !ok {
		memoString[text] = true
		return true
	} else {
		return false
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