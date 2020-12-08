package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatalf("Could not read file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	lines = deleteEmpty(lines) // Delete the empty lines (newline at the end of the file)
	PartOne(lines)
	PartTwo(lines)
}

// PartOne runs the first part of the puzzle
func PartOne(lines []string) {
	acc, _ := RunProgram(lines)
	logrus.Infof("Acc value for part 1: %d", acc)
}

// PartTwo runs the second part of the puzzle
func PartTwo(lines []string) {
	for i := 0; i < len(lines); i++ {
		linesFixed := make([]string, len(lines))
		copy(linesFixed, lines)
		cmd := strings.Split(linesFixed[i], " ")
		switch cmd[0] {
		case "acc":
			continue
		case "jmp":
			linesFixed[i] = fmt.Sprintf("nop %s", cmd[1])
			logrus.Debugf("jmp -> nop")
		case "nop":
			linesFixed[i] = fmt.Sprintf("jmp %s", cmd[1])
			logrus.Debugf("nop -> jmp")
		}
		if acc, loopDetected := RunProgram(linesFixed); !loopDetected {
			logrus.Infof("Acc value for part 2: %d", acc)
			break
		}
	}
}

func RunProgram(lines []string) (acc int64, loopDetected bool) {
	i := 0
	commandExecuted := make(map[int]bool)
	for i < len(lines) {
		line := lines[i]
		cmd := strings.Split(line, " ")
		switch cmd[0] {
		case "nop":
			i++
		case "jmp":
			i += int(parseInt(cmd[1]))
		case "acc":
			acc += parseInt(cmd[1])
			i++
		}
		if commandExecuted[i] {
			logrus.Debugf("Infinite loop detected at line %d", i+1)
			logrus.Debugf("Current value of acc: %d", acc)
			return acc, true
		}
		commandExecuted[i] = true
	}
	return acc, false
}

func parseInt(s string) int64 {
	rawValue, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		logrus.Fatalf("Could not parse integer %s: %v", s, rawValue)
	}
	return rawValue
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
