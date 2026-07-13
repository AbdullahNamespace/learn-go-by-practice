// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 92: odd_count_in_array.go
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
// prints the array, then counts and prints how many odd numbers exist in it.
//
// The counting logic iterates through every element and checks the modulo
// condition (number % 2 != 0). Each time this condition is true, an internal
// counter is incremented. After the loop finishes, the final counter value
// represents the total number of odd elements in the array.
//
// In Go, this kind of "count matching elements" pattern is a very common
// building block, and it's good practice to keep the counting logic in its
// own small, pure function (no side effects, just takes a slice and returns
// an int) so it can be reused, tested independently, and composed with other
// counting functions (even, positive, negative) that follow the exact same
// shape but with a different condition.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية بين 1 و 100، يطبع المصفوفة،
// ثم يحسب ويطبع عدد الأرقام الفردية الموجودة فيها.
//
// يمر منطق العد عبر كل عنصر ويتحقق من شرط باقي القسمة (number % 2 != 0).
// في كل مرة يكون هذا الشرط صحيحًا، يزداد عداد داخلي بمقدار واحد. بعد
// انتهاء الحلقة، تمثل القيمة النهائية للعداد إجمالي عدد العناصر الفردية
// في المصفوفة.
//
// في لغة غو، هذا النمط من "عد العناصر المطابقة" هو لبنة بناء شائعة جدًا،
// ومن الممارسات الجيدة إبقاء منطق العد في دالة صغيرة ونقية خاصة به (بدون
// آثار جانبية، تأخذ فقط شريحة وتُرجع int) حتى يمكن إعادة استخدامها، واختبارها
// بشكل مستقل، وتركيبها مع دوال عد أخرى (زوجي، موجب، سالب) تتبع نفس الشكل
// تمامًا لكن بشرط مختلف.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array = [15, 22, 33, 40, 7]
// Output: Array Elements: 15 22 33 40 7
//         Odd Numbers count is: 3
//   Why:    15, 33, and 7 are odd
//
// Example 2 (All Even):
// Input:  Array = [2, 4, 6, 8]
// Output: Array Elements: 2 4 6 8
//         Odd Numbers count is: 0
//
// Example 3 (All Odd):
// Input:  Array = [1, 3, 5]
// Output: Array Elements: 1 3 5
//         Odd Numbers count is: 3
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random numbers are in the range [1, 100]. | الأرقام العشوائية في النطاق [1، 100].
// • A number is odd if number % 2 != 0. | العدد فردي إذا كان باقي قسمته على 2 لا يساوي صفر.
// • The array length must be a positive integer (> 0). | طول المصفوفة يجب أن يكون عددًا صحيحًا موجبًا.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
// func oddCount(arr []int) int
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

func fillArrayWithRandomNumbers(lengthOfArray int) []int {
	arr := make([]int, lengthOfArray)

	for key := range lengthOfArray {
		arr[key] = randomNumber(1, 100)
	}

	return arr
}

func isOddNumber(number int) bool {
	return number%2 != 0
}

func oddCount(arr []int) int {
	count := 0

	for _, value := range arr {
		if isOddNumber(value) {
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
	fmt.Printf("Odd Numbers count is: %d\n", oddCount(arr))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(fillArrayWithRandomNumbers(readPositiveNumber("Please enter number of elements")))
}
