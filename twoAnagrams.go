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
	s := scan()
	t := scan()
	s_sp := strings.Split(s, "")
	t_sp := strings.Split(t,"")
	sort.Slice(s_sp, func (i,j int) bool {return s_sp[i] < s_sp[j]})
	sort.Slice(t_sp, func (i,j int) bool {return t_sp[i] > t_sp[j]})
	// fmt.Println(s_sp, t_sp)
	size := len(s_sp)
	if len(t_sp) > size {
		size = len(t_sp)
	}
	bool := true
	for i:= 0; i<size+1; i++ {
		if len(t_sp) == i {
			bool = false
			break
		}
		if len(s_sp) == i {
			break
		}
		if s_sp[i] < t_sp[i] {
			break
		}
		if s_sp[i] > t_sp[i] {
			bool =false
			break
		}
	}
	if bool {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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