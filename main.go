package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkE(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	maxCalories, sumThreeCalories := day1()
	fmt.Printf("Top #1: %v\n", maxCalories)
	fmt.Println("Max of top #3:", sumThreeCalories)
}

func day1() (int, int) {
	data, err := os.ReadFile("day1/input.txt")
	checkE(err)
	dataStr := strings.Split(string(data), "\n\n")
	var elvesCalories []int
	for _, elf := range dataStr {
		caloriesSplitted := strings.Split(elf, "\n")
		var elfCalories int
		for _, calorie := range caloriesSplitted {
			calorieInt, err := strconv.Atoi(calorie)
			checkE(err)
			elfCalories += calorieInt
		}

		elvesCalories = append(elvesCalories, elfCalories)
	}
	sort.Slice(elvesCalories, func(i, j int) bool {
		return elvesCalories[i] > elvesCalories[j]
	})
	maxCalories := elvesCalories[0]
	topThreeCalories := elvesCalories[:3]
	var sumThreeCalories int
	for _, calories := range topThreeCalories {
		sumThreeCalories += calories
	}
	return maxCalories, sumThreeCalories
}
