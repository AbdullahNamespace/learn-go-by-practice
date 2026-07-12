// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 82: copy_array_in_reverse_order.go
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
// Write a program that fills an array with random numbers between 1 and 100,
// then copies all its elements into a second array in reverse order (the last
// element of the source becomes the first element of the destination, and so
// on). Finally, it prints both the original and the reversed array.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية بين 1 و 100، ثم ينسخ جميع عناصرها
// إلى مصفوفة ثانية بترتيب معكوس (يصبح آخر عنصر في المصدر أول عنصر في الوجهة،
// وهكذا). أخيرًا، يطبع كلًا من المصفوفة الأصلية والمصفوفة المعكوسة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Number of elements = 4
// Output: Array 1 elements: 10 20 30 40
//         Array 2 elements after copy: 40 30 20 10
//
// Example 2:
// Input:  Number of elements = 1
// Output: Array 1 elements: 55
//         Array 2 elements after copy: 55
//   Why:    A single-element array reversed is identical to the original
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number of elements must be positive (> 0) | عدد العناصر يجب أن يكون موجبًا
// • Both arrays must be of the same length | يجب أن تكون المصفوفتان بنفس الطول
// • Random numbers are in the range [1, 100] | الأرقام العشوائية تكون في المدى [1، 100]
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(min, max int) int
// func fillArrayWithRandomNumbers(length int) []int
// func copyArrayInReverseOrder(source []int) []int
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

func fillArrayWithRandomNumbers(length int) []int {
	arr := make([]int, length)

	for key := range length {
		arr[key] = randomNumber(1, 100)
	}

	return arr
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func copyArrayInReverseOrder(source []int) []int {
	arr := make([]int, len(source))

	for key := range source {
		arr[key] = source[len(source)-1-key]
	}

	return arr
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array Elements is %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter a numbers of elements")

	arr := fillArrayWithRandomNumbers(lengthOfArray)

	printArray(arr)

	rArr := copyArrayInReverseOrder(arr)

	printArray(rArr)
}
