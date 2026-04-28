// package main

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("could not open: %v", filepath)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := range data {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	var ids []string
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
	return ids
}

type ids struct {
	idSum     int
}

func convertInt(id string) int {
	id = strings.TrimSpace(id)
	if id == "" {
		log.Fatalf("empty id string")
	}
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("cannot convert '%s' to int: %v", id, err)
	}
	return s
}

func convertString(id int) string {
	return strconv.Itoa(id)
}

func (id *ids) parseInput(input []string) {
	for _, v := range input {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		s := strings.ReplaceAll(v, "-", " ")
		idRange := strings.Fields(s)
		start := convertInt(idRange[0])
		stop := convertInt(idRange[1])
		if start > stop {
			log.Fatalf("invalid range: start > stop in '%s'", v)
		}

		id.parseID(start, stop)
	}
}

func (id *ids) parseID(start, stop int) {
	for i := start; i <= stop; i++ {
		currentID := convertString(i)
		mid := len(currentID) >> 1
		l := currentID[:mid]
		r := currentID[mid:]
		if l == r {
			id.idSum += i
		}
	}
}

// var dummy = []string{"222", "3456", "3636", "7778"}

func main() {
	inputs := readInput("input.txt")
	id := ids{}
	id.parseInput(inputs)
	fmt.Println(id.idSum)
}
