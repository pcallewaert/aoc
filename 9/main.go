package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatalf("Could not read file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	numbers := parseNumbers(lines)
	resultPartOne := PartOne(numbers)
	PartTwo(resultPartOne, numbers)
}

// PartOne runs the first part of the puzzle
func PartOne(numbers []uint64) uint64 {
	for i := 25; i < len(numbers); i++ {
		if !isSumPossible(numbers[i-25:i], numbers[i]) {
			logrus.Infof("Sum not possible for %d", numbers[i])
			return numbers[i]
		}
	}
	return 0
}

// PartTwo runs the second part of the puzzle
func PartTwo(sum uint64, numbers []uint64) {
	totalNumbers := 2
	for true {
		logrus.Debugf("Trying %d numbers", totalNumbers)
		for index := totalNumbers; index < len(numbers); index++ {
			tmpSum := calculateSum(numbers[index-totalNumbers : index])
			if tmpSum == sum {
				logrus.Debugf("Found a matching serie with %d numbers", totalNumbers)
				min, max := calculateMinMax(numbers[index-totalNumbers : index])
				logrus.Infof("Weakness is %d", min+max)
				return
			}
		}
		totalNumbers++
	}
}

func calculateSum(slice []uint64) (sum uint64) {
	for _, v := range slice {
		sum += v
	}
	return
}

func calculateMinMax(slice []uint64) (min, max uint64) {
	for _, v := range slice {
		if min == 0 || v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

func parseNumbers(lines []string) []uint64 {
	var r []uint64
	for _, str := range lines {
		if str != "" {
			if parsed, err := strconv.ParseUint(str, 10, 64); err != nil {
				logrus.Errorf("Could not parse %s as uint64: %v", str, err)
			} else {
				r = append(r, parsed)
			}
		}
	}
	return r
}

func isSumPossible(slice []uint64, result uint64) bool {
	cs := combin.Permutations(len(slice), 2)
	for _, c := range cs {
		tmpResult := slice[c[0]] + slice[c[1]]
		if tmpResult == result {
			return true
		}
	}
	return false
}
