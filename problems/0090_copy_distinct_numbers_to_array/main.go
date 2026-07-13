// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 90: copy_distinct_numbers_to_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Searching | المصفوفات والبحث
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills a source array with random numbers between 1 and
// 100, then copies only the distinct (unique) numbers into a new
// destination slice, by checking — for every source element — whether it has
// already been added to the destination before appending it. Finally, print
// both the source array and the resulting array of distinct numbers.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة مصدر بأرقام عشوائية بين 1 و 100، ثم
//  يجب أن ينسخ البرنامج فقط الأرقام
// المميزة (غير المكررة) إلى شريحة وجهة جديدة، عن طريق التحقق — لكل عنصر في
// المصدر — مما إذا كان قد أُضيف بالفعل إلى الوجهة قبل إضافته. أخيرًا، اطبع
// المصفوفة المصدرية والمصفوفة الناتجة التي تحتوي على الأرقام المميزة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array 1 = [10 10 10 50 50 70 70 70 70 90]
// Output: Array 2 Distinct numbers: 10 50 70 90
//
// Example 2:
// Input:  Array 1 = [1 2 3 4 5]
// Output: Array 2 Distinct numbers: 1 2 3 4 5
//   Why:    All source elements are already unique
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Must reuse an "is number already in array" search function | يجب إعادة استخدام دالة "هل الرقم موجود بالفعل في المصفوفة"
// • The order of first appearance must be preserved in the destination | يجب الحفاظ على ترتيب أول ظهور في الوجهة
// • Array elements are positive integers (1 to 100) | عناصر المصفوفة أعداد صحيحة موجبة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func fillArrayWithRandomNumbers(arr []int)
// func isNumberInArray(number int, arr []int) bool
// func copyDistinctNumbersToArray(arrSource []int) []int
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

func fillArrayWithRandomNumbers(lengthOfArray int) []int {
	arr := make([]int, lengthOfArray)

	for key := range lengthOfArray {
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

func isNumberInArray(number int, arr []int) bool {
	for _, value := range arr {
		if value == number {
			return true
		}
	}

	return false
}

func copyDistinctNumbersToArray(arrSource []int) []int {
	var arr []int

	for _, value := range arrSource {
		if !isNumberInArray(value, arr) {
			arr = append(arr, value)
		}
	}

	return arr
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array element is %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter number of elements")
	arr := fillArrayWithRandomNumbers(lengthOfArray)

	fmt.Println("Array 1 elements")
	printArray(arr)

	distinctArr := copyDistinctNumbersToArray(arr)

	fmt.Println("Array 2 Distinct numbers")
	printArray(distinctArr)
}
