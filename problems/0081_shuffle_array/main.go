// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 81: shuffle_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Random | المصفوفات والعشوائية
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a positive number N from the user, fills an array
// with sequential numbers from 1 to N, prints the array before shuffling, then
// shuffles the array elements randomly by repeatedly swapping two randomly
// chosen positions, and finally prints the array after shuffling.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا N من المستخدم، يملأ مصفوفة بأرقام متسلسلة من
// 1 إلى N، يطبع المصفوفة قبل الخلط، ثم يخلط عناصر المصفوفة عشوائيًا عن طريق
// تبديل موقعين عشوائيين بشكل متكرر، وأخيرًا يطبع المصفوفة بعد الخلط.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Number of elements = 5
// Output: Array elements before shuffle: 1 2 3 4 5
//         Array elements after shuffle:  3 1 5 2 4
//   Why:    Elements are the same set {1..5} but in a randomized order
//
// Example 2:
// Input:  Number of elements = 3
// Output: Array elements before shuffle: 1 2 3
//         Array elements after shuffle:  2 3 1
//
// (Note: Actual shuffled output varies due to randomness)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The number of elements must be a positive integer greater than 0. | يجب أن يكون عدد العناصر عددًا صحيحًا موجبًا أكبر من 0.
// • The array is filled with values from 1 to N sequentially before shuffling. | تُملأ المصفوفة بقيم من 1 إلى N بشكل متسلسل قبل الخلط.
// • Shuffling must produce a different order (not guaranteed for very small N). | يجب أن ينتج الخلط ترتيبًا مختلفًا (غير مضمون لقيم N الصغيرة جدًا).
// • Use Go slices []int for dynamic sizing based on user input. | استخدم شرائح غو []int للحجم الديناميكي بناءً على مدخلات المستخدم.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) int
// func fillArrayWith1ToN(length int) []int
// func randomNumber(min, max int) int
// func shuffleArray(arr []int)
// func printArray(prompt string, arr []int)
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
	fmt.Printf("X Error  : %s\n", prompt)
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

func fillArrayWith1ToN(length int) []int {
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = (i + 1)
	}

	return arr
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func shuffleArray(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		random := randomNumber(0, i)
		arr[i], arr[random] = arr[random], arr[i]
	}
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(prompt string, arr []int) {
	fmt.Printf("Array elements %s : %v\n", prompt, arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter numbers of elements")
	arr := fillArrayWith1ToN(lengthOfArray)
	printArray("before shuffle", arr)
	shuffleArray(arr)
	printArray("after shuffle ", arr)
}
