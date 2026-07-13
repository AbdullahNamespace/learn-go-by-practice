// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 95: negative_count_in_array.go
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
// Write a program that fills an array with random numbers in the range
// [-100, 100], prints the array, then counts and prints how many elements
// are strictly negative (less than zero).
//
// This problem is the direct counterpart to Problem 94. The counting
// condition here is (number < 0), which is the strict complement of the
// "positive" condition (number >= 0) used previously. Together, these two
// counts should always add up to the total array length:
// positiveCount(arr) + negativeCount(arr) == len(arr), since every number is
// either >= 0 or < 0, with no overlap and no gap. This invariant is a useful
// mental check (and could even be used as a simple correctness test) when
// verifying that both counting functions were implemented correctly.
//
// In Go, it's worth highlighting that boundary conditions like "is zero
// negative?" must be handled with a strict inequality (< 0, not <= 0) to
// avoid double-counting zero in both the positive and negative counts. This
// kind of off-by-one / boundary-inclusion mistake is one of the most common
// sources of subtle bugs in counting and filtering logic, so it deserves
// deliberate attention here.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية في النطاق [-100، 100]، يطبع
// المصفوفة، ثم يحسب ويطبع عدد العناصر السالبة تمامًا (أقل من صفر).
//
// هذه المشكلة هي النظير المباشر للمشكلة 94. شرط العد هنا هو (number < 0)،
// وهو المكمل الصارم لشرط "الموجب" (number >= 0) المستخدم سابقًا. معًا،
// يجب أن يتساوى مجموع هذين العددين دائمًا مع الطول الكلي للمصفوفة:
// positiveCount(arr) + negativeCount(arr) == len(arr)، لأن كل رقم إما
// >= 0 أو < 0، دون تداخل أو فجوة. هذه الخاصية الثابتة (invariant) هي فحص
// ذهني مفيد (ويمكن استخدامها حتى كاختبار صحة بسيط) عند التحقق من أن كلتا
// دالتي العد نُفذتا بشكل صحيح.
//
// في لغة غو، يستحق الأمر التنويه أن حالات الحدود مثل "هل الصفر سالب؟"
// يجب معالجتها بمتباينة صارمة (< 0، وليس <= 0) لتجنب احتساب الصفر مرتين
// في كل من عددي الموجب والسالب. هذا النوع من أخطاء الحدود الفاصلة
// (off-by-one) هو أحد أكثر مصادر الأخطاء الدقيقة شيوعًا في منطق العد
// والتصفية، لذا يستحق اهتمامًا متعمدًا هنا.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array = [-15, 22, -33, 40, 0]
// Output: Array Elements: -15 22 -33 40 0
//         Negative Numbers count is: 2
//   Why:    -15 and -33 are < 0
//
// Example 2 (All Positive):
// Input:  Array = [1, 3, 5]
// Output: Array Elements: 1 3 5
//         Negative Numbers count is: 0
//
// Example 3 (Edge Case - Zero is NOT Negative):
// Input:  Array = [0]
// Output: Array Elements: 0
//         Negative Numbers count is: 0
//   Why:    0 < 0 is false, so zero does NOT count toward the negative total
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random numbers are in the range [-100, 100]. | الأرقام العشوائية في النطاق [-100، 100].
// • Zero is NOT counted as negative (condition: number < 0). | الصفر لا يُحتسب سالبًا (الشرط: number < 0).
// • The array length must be a positive integer (> 0). | طول المصفوفة يجب أن يكون عددًا صحيحًا موجبًا.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(min, max int) int
// func fillArrayWithRandomNumbers(lengthOfArray int) []int
// func negativeCount(arr []int) int
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
		arr[key] = randomNumber(-100, 100)
	}

	return arr
}

func negativeCount(arr []int) int {
	count := 0

	for _, value := range arr {
		if value < 0 {
			count++
		}
	}

	return count
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array is :  %v\n", arr)
}

func printResult(arr []int) {
	printArray(arr)
	fmt.Printf("Negative Numbers count is: %d\n", negativeCount(arr))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(fillArrayWithRandomNumbers(readPositiveNumber("Please enter number of elements")))
}
