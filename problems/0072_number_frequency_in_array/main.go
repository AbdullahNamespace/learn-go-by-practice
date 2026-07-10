// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 72: number_frequency_in_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Searching | المصفوفات والبحث
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads an array of integers from the user (first the number of
// elements, then each element), then reads a positive number to check. The program
// prints the original array and counts how many times the specified number appears
// (its frequency) within the array.
//
// العربية:
// اكتب برنامجًا يقرأ مصفوفة أعداد صحيحة من المستخدم (أولًا عدد العناصر، ثم كل عنصر)،
// ثم يقرأ عددًا موجبًا للتحقق منه. يطبع البرنامج المصفوفة الأصلية ويحسب كم مرة
// يظهر الرقم المحدد (تكراره) داخل المصفوفة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:
//   Number of elements: 5
//   Elements: 10 20 10 30 10
//   Number to check: 10
// Output:
//   Original array: 10 20 10 30 10
//   Number 10 is repeated 3 time(s)
//
// Example 2:
// Input:
//   Number of elements: 3
//   Elements: 1 2 3
//   Number to check: 7
// Output:
//   Original array: 1 2 3
//   Number 7 is repeated 0 time(s)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Array length must be a positive integer (> 0). | عدد عناصر المصفوفة يجب أن يكون عددًا صحيحًا موجبًا (> 0).
// • The number to check must be a positive integer. | الرقم المراد التحقق منه يجب أن يكون عددًا صحيحًا موجبًا.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) int
// func readArray(arr []int) int
// func printArray(arr []int)
// func timesRepeated(number int, arr []int) int
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"strconv"
)

// ======================
//   UTILITY
// ======================

func printError(prompt string) {
	fmt.Printf("X Error : %s\n", prompt)
}

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) string {
	fmt.Printf("%s : ", prompt)

	var input string

	fmt.Scan(&input)

	return input
}

func readNumber(prompt string) int {
	for {
		input := readString(prompt)

		inputNumber, err := strconv.Atoi(input)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) int {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

func readArray(arr []int) {
	for key := range arr {
		arr[key] = readPositiveNumber(fmt.Sprintf("Elements [%d]", key+1))
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func timesRepeated(number int, arr []int) int {
	count := 0

	for _, value := range arr {
		if value == number {
			count++
		}
	}

	return count
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter The length of array")
	array := make([]int, lengthOfArray)
	readArray(array)
	checkNumber := readPositiveNumber("Please enter the number you want to check")
	fmt.Printf("Original array %v\n", array)
	fmt.Printf("%d is repeated %d time(s)\n", checkNumber, timesRepeated(checkNumber, array))
}
