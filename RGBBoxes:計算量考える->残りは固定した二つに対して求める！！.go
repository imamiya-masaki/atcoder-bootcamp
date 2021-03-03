package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sum int
func main() {
	scan_init()
	get := scanInts(3)
	n := scanInt()
	// sort.Ints(get)
	count := 0
	for i:= 0; i<=n; i++ {
		for j:= 0; j<=n; j++ {
			sum := get[0]*i + get[1]*j
			keisan := n - sum
			if keisan < 0 {
				continue
			}
			if keisan == 0 {
				count++
				continue
			}
			if keisan % get[2] == 0 {
				count++
				continue
			}
		}
	}
	fmt.Println(count)
}
var memo map[string]bool
var sc = bufio.NewScanner(os.Stdin)

func loop (count int , get []int, n int, selected []int) {
	if count > n {
		return
	}
	take := mergeInts(selected)
	if _,ok := memo[take]; !ok {
		memo[take] = true
	} else {
		return
	}
	if count == n {
		// fmt.Println(count, get ,n, selected)
		sum++
		return
	}
	for i:= 0; i<len(get); i++ {
		take := append([]int{}, selected...)
		take[i]++
		loop(count+get[i],get,n, take)
	}
}

func mergeInts (get []int) string{
	sakura := make([]string,len(get))
	// sort.Ints(get)
	for i:= 0; i< len(get); i++ {
		sakura[i] = strconv.Itoa(get[i])
	}
	return strings.Join(sakura,"")
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