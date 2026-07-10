// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 73: fill_array_with_random_numbers.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Random | المصفوفات والأرقام العشوائية
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that asks the user for the number of elements, then fills an array
// with random numbers between 1 and 100. Finally, it prints the array elements.
//
// العربية:
// اكتب برنامجًا يطلب من المستخدم إدخال عدد العناصر، ثم يملأ مصفوفة بأرقام
// عشوائية بين 1 و 100. أخيرًا، يطبع عناصر المصفوفة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:
//   Number of elements: 5
// Output:
//   Array Elements: 45 12 89 3 67
//
// (Note: Actual output varies due to random generation)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number of elements must be a positive integer (> 0). | عدد العناصر يجب أن يكون عددًا صحيحًا موجبًا (> 0).
// • Random numbers are in the range [1, 100]. | الأرقام العشوائية تكون في المدى [1، 100].
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
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

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	for _, value := range arr {
		fmt.Printf("%d\t", value)
	}
	fmt.Println()
}

// ======================
//         MAIN
// ======================

func main() {
	length := readPositiveNumber("Please enter number of elements:")
	arr := make([]int, length)
	fillArrayWithRandomNumbers(arr)
	printArray(arr)
}
