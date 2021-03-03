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
	scan_init()
	// input.scan(), input.scan_int()
	h := scanInt()
	w := scanInt()
	taker := make([]string, h)
	// reserves := make([]string, w)
	a := make([][]string, h)
	for i:= 0; i<h; i++ {
		taker[i] = scan()
		a[i] = strings.Split(taker[i],"")
	}
	nexts := []string{}
	for i:=0; i<h; i++ {
		// fmt.Println(taker[i],"taker")
		for j:= 0; j<w; j++ {
			// fmt.Println("a", a[i][j])
			if a[i][j] == "#" {
				nexts = append(nexts, taker[i])
				break
			}
		}
	}
	a = make([][]string, len(nexts))
	for i:= 0; i<len(nexts); i++ {
		a[i] = strings.Split(nexts[i],"")
	}
	clearLines := []int{}
	for j:= 0; j<w; j++ {
		for i:=0; i<len(nexts); i++ {
			if a[i][j] == "#" {
				clearLines = append(clearLines, j)
				break
			}
		}
	}
	// fmt.Println(clearLines)
	for i:= 0; i<len(nexts); i++ {
		for j:=0; j<len(clearLines); j++ {
			taker := clearLines[j]
			// fmt.Println(taker, a[i])
			fmt.Print(a[i][taker])
		}
		fmt.Println()
	}
	// fmt.Print(a, taker, nexts, w)
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