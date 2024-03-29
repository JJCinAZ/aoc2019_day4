// function.go
package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Part1(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Missing input", http.StatusNotAcceptable)
		return
	}
	start, end, err := parseInput(string(input))
	if err != nil {
		http.Error(w, "Invalid input: " + err.Error(), http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("Number of possible passwords is %d",
		getPossible(start, end, isPossible))))
}

func Part2(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Missing input", http.StatusNotAcceptable)
		return
	}
	start, end, err := parseInput(string(input))
	if err != nil {
		http.Error(w, "Invalid input: " + err.Error(), http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("Number of possible passwords is %d",
		getPossible(start, end, isPossible2))))
}

func getPossible(start, end int, check func(int) int) int {
	count := 0
	for i := start; i <= end; i++ {
		count += check(i)
	}
	return count
}

// Return 0 if not possible, 1 if possible
func isPossible2(x int) int {
	s := strconv.Itoa(x)
	var (
		lastC rune
		gs [10]int
		notIncreasing bool
		hasDouble bool
	)
	for pos, c := range s {
		if pos > 0 {
			if lastC == c {
				gs[c - '0']++
			}
			if c < lastC {
				notIncreasing = true
				break
			}
		}
		lastC = c
	}
	for i := range gs {
		if gs[i] == 1 {
			hasDouble = true
			break
		}
	}
	if hasDouble && !notIncreasing {
		return 1
	}
	return 0
}

// Return 0 if not possible, 1 if possible
func isPossible(x int) int {
	s := strconv.Itoa(x)
	var (
		lastC rune
		hasDouble bool
		notIncreasing bool
	)
	for pos, c := range s {
		if pos > 0 {
			if lastC == c {
				hasDouble = true
			}
			if c < lastC {
				notIncreasing = true
				break
			}
		}
		lastC = c
	}
	if hasDouble && !notIncreasing {
		return 1
	}
	return 0
}

func parseInput(input string) (start int, end int, err error) {
	a := strings.Split(input, "-")
	if len(a) != 2 {
		err = fmt.Errorf("Need ###-###")
		return
	}
	if start, err = strconv.Atoi(a[0]); err != nil {
		return
	}
	end, err = strconv.Atoi(a[1])
	return
}
