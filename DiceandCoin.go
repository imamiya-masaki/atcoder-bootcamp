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
	sum := float64(0)
	sums := make([]float64, n)
	for i:=1; i<=n; i++ {
		take := i
		count := 0
		two := 2
		if k > take {
			count++
			for i := 0; i < k; i++ {
				if take*two >= k {
					// fmt.Println(take*two, k, "え")
					break
				}
				count++
				two *= 2
			}
		}
		kakuritu := math.Pow(2,float64(count))*float64(n)
		sums[i-1] = kakuritu
		// fmt.Println(kakuritu, count, k, take, two)
		// sum += kakuritu
	}
	for i:= 0; i<n; i++ {
		// fmt.Println(float64(int(math.Pow(10,10))/sums[i]), sums[i])
		sum += float64(math.Pow(10,15)/sums[i])
	}
	fmt.Println(sum/math.Pow(10,15))
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