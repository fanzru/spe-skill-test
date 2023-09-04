package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SpeSkillTest struct{}

// Challenge 1: Narcissistic Number
func (s *SpeSkillTest) isNarcissisticNumber(num int) bool {
	numStr := strconv.Itoa(num)
	numDigits := len(numStr)
	sum := 0

	for _, digitChar := range numStr {
		digit, _ := strconv.Atoi(string(digitChar))
		sum += int(math.Pow(float64(digit), float64(numDigits)))
	}

	return sum == num
}

// Challenge 2: Parity Outlier
func (s *SpeSkillTest) findParityOutlier(numbers []int) int {
	evenCount := 0
	oddCount := 0
	var even, odd int

	for _, num := range numbers {
		if num%2 == 0 {
			evenCount++
			even = num
		} else {
			oddCount++
			odd = num
		}
	}

	if evenCount == 1 {
		return even
	} else if oddCount == 1 {
		return odd
	}

	return 0 // Return 0 if no outlier is found
}

// Challenge 3: Needle in a Haystack
func (s *SpeSkillTest) findNeedle(haystack []string, needle string) int {
	for i, str := range haystack {
		if strings.Contains(str, needle) {
			return i
		}
	}
	return -1 // Return -1 if needle is not found
}

// Challenge 4: The Blue Ocean Reverse
func (s *SpeSkillTest) blueOceanReverse(list1 []int, list2 []int) []int {
	// Create a map to store the elements from list2
	elementsInList2 := make(map[int]bool)
	for _, num := range list2 {
		elementsInList2[num] = true
	}

	// Create a new slice to store elements that are not in list2
	result := []int{}
	for _, num := range list1 {
		if !elementsInList2[num] {
			result = append(result, num)
		}
	}

	return result
}
func main() {

	// Implmentation
	speImpl := SpeSkillTest{}

	// Challenge 1
	fmt.Println("------------------------------------------")
	fmt.Println("Example")
	narcissistic := speImpl.isNarcissisticNumber(153)
	fmt.Println("Narcissistic Number - 153 :", narcissistic)
	narcissistic = speImpl.isNarcissisticNumber(111)
	fmt.Println("Narcissistic Number - 111 :", narcissistic)
	fmt.Println("------------------------------------------")

	// Challenge 2
	fmt.Println("Example")
	numbers := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	parityOutlier := speImpl.findParityOutlier(numbers)
	fmt.Println("Parity Outlier:", parityOutlier)
	numbers = []int{160, 3, 1719, 19, 11, 13, -21}
	parityOutlier = speImpl.findParityOutlier(numbers)
	fmt.Println("Parity Outlier:", parityOutlier)
	numbers = []int{11, 13, 15, 19, 9, 13, -21}
	parityOutlier = speImpl.findParityOutlier(numbers)
	fmt.Println("Parity Outlier:", parityOutlier)
	fmt.Println("------------------------------------------")

	// Challenge 3
	fmt.Println("Example")
	haystack := []string{"red", "blue", "yellow", "black", "grey"}
	needle := "blue"
	needleIndex := speImpl.findNeedle(haystack, needle)
	fmt.Println("List : ", haystack)
	fmt.Printf(`Index of %s in a Haystack : %d`, needle, needleIndex)
	fmt.Println()
	fmt.Println("------------------------------------------")

	// Challenge 4
	fmt.Println("Example")
	list1 := []int{1, 2, 3, 4, 6, 10}
	list2 := []int{1}
	fmt.Println("list 1 : ", list1)
	fmt.Println("list 2 : ", list2)
	reversed := speImpl.blueOceanReverse(list1, list2)
	fmt.Println("The Blue Ocean Reverse - 1:", reversed)
	list3 := []int{1, 5, 5, 5, 5, 3}
	list4 := []int{3}
	fmt.Println("list 3  : ", list3)
	fmt.Println("list 4  : ", list4)
	reversed = speImpl.blueOceanReverse(list3, list4)
	fmt.Println("The Blue Ocean Reverse - 2:", reversed)
	fmt.Println("------------------------------------------")
}
