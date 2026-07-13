// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 93: even_count_in_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Math | المصفوفات والرياضيات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills an array with random numbers between 1 and 100,
// prints the array, then counts and prints how many even numbers exist in it.
//
// This problem is the mirror image of Problem 92: instead of checking
// (number % 2 != 0), the condition becomes (number % 2 == 0). Structurally,
// the counting function is identical — same loop, same accumulator pattern —
// only the boolean condition inside the if-statement changes.
//
// This is a good opportunity to notice a recurring theme in this whole series:
// once you understand the "iterate + conditionally increment a counter" shape,
// you can solve odd-count, even-count, positive-count, and negative-count all
// with the same skeleton. In Go, a more advanced (optional) refactor would be
// to write one generic countIf(arr []int, predicate func(int) bool) int
// function and pass different predicate closures for each case — but for this
// learning stage, writing each counting function explicitly reinforces the
// pattern more clearly.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية بين 1 و 100، يطبع المصفوفة،
// ثم يحسب ويطبع عدد الأرقام الزوجية الموجودة فيها.
//
// هذه المشكلة هي الصورة المرآتية للمشكلة 92: بدلاً من التحقق من
// (number % 2 != 0)، يصبح الشرط (number % 2 == 0). من الناحية الهيكلية،
// دالة العد متطابقة — نفس الحلقة، نفس نمط التجميع — فقط الشرط المنطقي
// داخل عبارة if يتغير.
//
// هذه فرصة جيدة لملاحظة موضوع متكرر في هذه السلسلة بأكملها: بمجرد فهم
// شكل "المرور + زيادة عداد شرطيًا"، يمكنك حل مسائل عد الفردي والزوجي
// والموجب والسالب كلها بنفس الهيكل. في غو، إعادة هيكلة أكثر تقدمًا
// (اختيارية) ستكون كتابة دالة عامة واحدة countIf(arr []int,
// predicate func(int) bool) int وتمرير دوال predicate مختلفة لكل حالة —
// لكن في هذه المرحلة التعليمية، كتابة كل دالة عد بشكل صريح يُرسّخ النمط
// بشكل أوضح.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array = [15, 22, 33, 40, 7]
// Output: Array Elements: 15 22 33 40 7
//         Even Numbers count is: 2
//   Why:    22 and 40 are even
//
// Example 2 (All Odd):
// Input:  Array = [1, 3, 5]
// Output: Array Elements: 1 3 5
//         Even Numbers count is: 0
//
// Example 3 (All Even):
// Input:  Array = [2, 4, 6, 8]
// Output: Array Elements: 2 4 6 8
//         Even Numbers count is: 4
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random numbers are in the range [1, 100]. | الأرقام العشوائية في النطاق [1، 100].
// • A number is even if number % 2 == 0. | العدد زوجي إذا كان باقي قسمته على 2 يساوي صفر.
// • The array length must be a positive integer (> 0). | طول المصفوفة يجب أن يكون عددًا صحيحًا موجبًا.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(lengthOfArray int) []int
// func evenCount(arr []int) int
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

func isEvenNumber(number int) bool {
	return number%2 == 0
}

func randomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func fillArrayWithRandomNumbers(lengthOfArray int) []int {
	arr := make([]int, lengthOfArray)

	for key := range lengthOfArray {
		arr[key] = randomNumber(1, 100)
	}

	return arr
}

func evenCount(arr []int) int {
	count := 0

	for _, value := range arr {
		if isEvenNumber(value) {
			count++
		}
	}

	return count
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array is : %v\n", arr)
}

func printResult(arr []int) {
	printArray(arr)
	fmt.Printf("Even Numbers count is: %d\n", evenCount(arr))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(fillArrayWithRandomNumbers(readPositiveNumber("Please enter number of elements")))
}
