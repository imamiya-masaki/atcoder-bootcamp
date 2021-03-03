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
	h := scanInt()
	w := scanInt()
	s := make([][]string, h)
	grids := make([][]int, h)
	lines := make([]string, h)
	for i:= 0; i<h; i++ {
		lines[i] = scan()
		s[i] = strings.Split(lines[i], "")
		grids[i] = make([]int, w)
	}
	for i:= 0; i<h; i++ {
		for j:= 0; j<w; j++ {
			if s[i][j] == "#" {
				grids[i][j]++
				if i > 0 {
					grids[i-1][j]++
				}
				if i < h-1 {
					grids[i+1][j]++
				}
				if j > 0 {
					grids[i][j-1]++
				}
				if j < w-1 {
					grids[i][j+1]++
				}
			}
		}
	}
	flag := true
	for i:= 0; i<h; i++ {
		if !flag {
			break
		}
		for j:=0; j<w; j++ {
			if s[i][j] == "#" && grids[i][j] <= 1 {
				// gridPrintln(grids)
				// gridPrintlnString(s)
				// fmt.Println(i,j)
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
}

var sc = bufio.NewScanner(os.Stdin)

func gridPrintln (memo [][]int) {
	for i:= 0; i< len(memo); i++ {
		for j:=0; j< len(memo[i]); j++ {
			fmt.Print(memo[i][j])
		}
		fmt.Println("")
	}
}
func gridPrintlnString (memo [][]string) {
	for i:= 0; i< len(memo); i++ {
		for j:=0; j< len(memo[i]); j++ {
			fmt.Print(memo[i][j])
		}
		fmt.Println("")
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