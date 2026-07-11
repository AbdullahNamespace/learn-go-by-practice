// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 78: copy_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays | المصفوفات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills an array with random numbers, then copies all elements
// from this source array into a destination array. Finally, it prints both arrays
// to verify that the copy operation was successful.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية، ثم ينسخ جميع العناصر من هذه المصفوفة
// المصدرية إلى مصفوفة وجهة. أخيرًا، يطبع كلتا المصفوفتين للتحقق من نجاح عملية النسخ.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:
//   Number of elements: 4
// Output:
//   Array 1 elements:
//   10 20 30 40
//   Array 2 elements after copy:
//   10 20 30 40
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Both arrays must have the same capacity/length. | كلتا المصفوفتين يجب أن يكون لهما نفس السعة/الطول.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
// func copyArray(arrDestination []int, arrSource []int)
// func printArray(arr []int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"math/rand/v2"
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

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(from, to int) int {
	return rand.IntN(to-from+1) + from
}

func fillArrayWithRandomNumbers(arr []int) {
	for key := range arr {
		arr[key] = randomNumber(1, 100)
	}
}

func copyArray(arrDestination []int, arrSource []int) {
	copy(arrDestination, arrSource)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array Elements: %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter number of elements")

	arr := make([]int, lengthOfArray)

	fillArrayWithRandomNumbers(arr)
	printArray(arr)

	arrCopy := make([]int, lengthOfArray)

	copyArray(arrCopy, arr)
	fmt.Println("Array 2 elements after copy")
	printArray(arrCopy)
}
