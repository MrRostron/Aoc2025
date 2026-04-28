// package main
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open: %v", fileName)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var turns []string
	for scanner.Scan() {
		turns = append(turns, scanner.Text())
	}
	return turns
}

// Dial is a struct that represents a safes dial, it has a range of 0 - 99.
type Dial struct {
	// startPoint is the value the dial is pointing at. (default should be 50)
	startPoint int
	// paswd is incremented each time the dial reaches 0.
	paswd int
}

// for testing purposes remove once done.
// var dummy = []string{"R100", "L100", "L100"}

// parseTurn accepts a string containing a direction and amount of clicks the dial must be turned eg."L100"
// it returns the direction as a string and number of clicks as a int.
func parseTurn(turn string) (direction string, clicks int) {
	direction = string(turn[0])
	clicks, err := strconv.Atoi(turn[1:])
	if err != nil {
		return "", 0
	}
	return
}

// leftTurn accepts a int value which is the number of clicks the dial must be turned counter clockwise.
func (d *Dial) leftTurn(clicks int) {
	for range clicks {
		d.startPoint--
		if d.startPoint < 0 {
			d.startPoint = 99
		}
	}
}

// rightTurn accepts a int value which is the number of clicks the dial must be turned clockwise.
func (d *Dial) rightTurn(clicks int) {
	for range clicks {
		d.startPoint++
		if d.startPoint > 99 {
			d.startPoint = 0
		}
	}
}

func (d *Dial) evalPass() {
	if d.startPoint == 0 {
		d.paswd++
	}
}

func main() {
	d := Dial{
		startPoint: 50,
	}
	turns := readInput("input.txt")

	for _, v := range turns {
		dir, clicks := parseTurn(v)
		switch dir {
		case "L":
			d.leftTurn(clicks)
			d.evalPass()
		case "R":
			d.rightTurn(clicks)
			d.evalPass()
		}
	}
	fmt.Println(d.paswd)
}
