package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type input struct {
	inputStr    []string
	targetIndex int
}

func main() {
	input := NewInput()
	// input.scan(), input.scan_int()
	n := input.scan_int()
	a := make([]int, n)
	memo := map[int]int{}
	for i:= 0; i<n; i++ {
		a[i] = input.scan_int()
		for j:=-1; j<=1; j++ {
			if _,ok := memo[a[i]+j]; !ok {
				memo[a[i]+j] = 0
			}
			memo[a[i]+j]++
		}
	}
	max := 0
	for _,val := range memo {
		if max < val {
			max = val
		}
	}
	fmt.Println(max)
}

var sc = bufio.NewScanner(os.Stdin)

func NewInput() *input {
	l := new(input)
	l.inputStr = getInputs()
	l.targetIndex = 0
	return l
}
func (in *input) scan() string {
	if len(in.inputStr) == in.targetIndex {
		return ""
	}
	getVal := in.inputStr[in.targetIndex]
	in.targetIndex++
	return getVal
}
func (in *input) scan_int() int {
	intVal, _ := strconv.Atoi(in.scan())
	return intVal
}
func scan_init() {
	sc.Split(bufio.ScanWords)
}

var rdr = bufio.NewReaderSize(os.Stdin, 200000)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
func scan() string {
	sc.Scan()
	return sc.Text()
}

func readRawLine() (string, error) {
	buf := make([]byte, 0, 200000)
	var err error
	for {
		l, p, e := rdr.ReadLine()
		err = e
		if e != nil {
			break
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf), err
}
func readLine() string {
	get, _ := readRawLine()
	return get
}
func getInputs() []string {
	strs := []string{}
	for {
		get, err := readRawLine()
		if err != nil {
			break
		}
		if len(get) == 0 {
			break
		}
		get_split := strings.Split(get, " ")
		if len(get_split) == 0 {
			break
		}
		strs = append(strs, get_split...)
	}
	return strs
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