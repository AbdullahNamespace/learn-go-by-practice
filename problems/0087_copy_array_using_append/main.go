// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 87: copy_array_using_append.go
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
// Write a program that fills a source array with random numbers between 1 and
// 100, then builds a brand-new destination slice by appending each source
// element one at a time (instead of using Go's built-in `copy` function).
// This demonstrates growing a slice element-by-element. Finally, print both
// the source array and the newly built destination array.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة مصدر بأرقام عشوائية بين 1 و 100، ثم يبني شريحة
// وجهة جديدة تمامًا عن طريق إضافة كل عنصر من المصدر واحدًا تلو الآخر (بدلاً من
// استخدام دالة `copy` المدمجة في Go). هذا يوضح كيفية نمو الشريحة عنصرًا بعنصر.
// أخيرًا، اطبع المصفوفة المصدرية والمصفوفة الوجهة المُنشأة حديثًا.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Number of elements = 4
// Output: Array 1 elements: 10 20 30 40
//         Array 2 elements after copy: 10 20 30 40
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Must use the addArrayElement helper function for copying, not direct assignment or copy(). | يجب استخدام دالة addArrayElement المساعدة للنسخ، وليس التعيين المباشر أو copy().
// • The destination slice starts empty and grows through append calls. | تبدأ شريحة الوجهة فارغة وتنمو من خلال استدعاءات append.
// • All elements from the source must be copied (no filtering). | يجب نسخ جميع العناصر من المصفوفة المصدرية (بدون تصفية).
// • Random numbers are in the range [1, 100]. | الأرقام العشوائية في النطاق [1، 100].
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number of elements must be positive (> 0) | عدد العناصر يجب أن يكون موجبًا
// • Must not use Go's built-in `copy` function | يجب عدم استخدام دالة `copy` المدمجة
// • The resulting destination slice's length must equal the source's length | يجب أن يتساوى طول الشريحة الوجهة مع طول المصدر
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
// func copyArrayUsingAppend(arrSource []int) []int
// func printArray(arr []int)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(min, max int) int
// func fillArrayWithRandomNumbers(length int) []int
// func addArrayElement(number int, arr []int) []int
// func copyArrayUsingAddElement(source []int) []int
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

func randomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func fillArrayWithRandomNumbers(length int) []int {
	arr := make([]int, length)

	for key := range length {
		arr[key] = randomNumber(1, 100)
	}

	return arr
}

func addArrayElement(number int, arr []int) []int {
	return append(arr, number)
}

func copyArrayUsingAddElement(source []int) []int {
	var arr []int

	for _, value := range source {
		arr = addArrayElement(value, arr)
	}

	return arr
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array is : %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter numbers of elements")

	arr1 := fillArrayWithRandomNumbers(lengthOfArray)

	arr2 := copyArrayUsingAddElement(arr1)

	fmt.Println("Array 1 elements:")
	printArray(arr1)

	fmt.Println("Array 2 elements after copy:")
	printArray(arr2)
}
