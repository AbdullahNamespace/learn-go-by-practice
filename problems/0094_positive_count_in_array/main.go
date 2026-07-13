// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 94: positive_count_in_array.go
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
// [-100, 100] (i.e., including negative values, unlike previous array
// problems that only used positive ranges), prints the array, then counts
// and prints how many elements are "positive" — defined here as greater than
// or equal to zero (so zero itself counts as positive).
//
// The counting condition is (number >= 0). This is an important detail to
// notice and preserve faithfully: some definitions of "positive" exclude
// zero, but this specific problem statement treats zero as positive, which
// affects the boundary behavior when a generated random number happens to be
// exactly 0.
//
// In Go, since the random range now spans negative values, the randomNumber
// helper function itself does not need any special handling — it already
// supports negative "from" bounds as long as the formula
// rand.IntN(to-from+1) + from is used correctly. This reinforces that a
// well-designed random-range helper is naturally reusable across both
// positive-only and mixed-sign ranges without modification.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية في النطاق [-100، 100] (أي
// شاملةً القيم السالبة، بخلاف مسائل المصفوفات السابقة التي استخدمت نطاقات
// موجبة فقط)، يطبع المصفوفة، ثم يحسب ويطبع عدد العناصر "الموجبة" — والمُعرَّفة
// هنا بأنها أكبر من أو تساوي صفر (أي أن الصفر نفسه يُحتسب كموجب).
//
// شرط العد هو (number >= 0). هذه تفصيلة مهمة يجب ملاحظتها والحفاظ عليها
// بأمانة: بعض تعريفات "الموجب" تستثني الصفر، لكن نص هذه المسألة تحديدًا
// يعامل الصفر كموجب، مما يؤثر على سلوك الحد الفاصل عندما يكون الرقم
// العشوائي المُولّد يساوي 0 بالضبط.
//
// في لغة غو، بما أن النطاق العشوائي يمتد الآن ليشمل قيمًا سالبة، فإن دالة
// randomNumber المساعدة نفسها لا تحتاج لأي معالجة خاصة — فهي تدعم بالفعل
// حدود "from" سالبة طالما استُخدمت الصيغة rand.IntN(to-from+1) + from
// بشكل صحيح. هذا يُرسّخ أن دالة مساعدة للنطاق العشوائي مصممة جيدًا تكون
// قابلة لإعادة الاستخدام بشكل طبيعي عبر النطاقات الموجبة فقط والنطاقات
// مختلطة الإشارة دون تعديل.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array = [-15, 22, -33, 40, 0]
// Output: Array Elements: -15 22 -33 40 0
//         Positive Numbers count is: 3
//   Why:    22, 40, and 0 are >= 0
//
// Example 2 (All Negative):
// Input:  Array = [-1, -3, -5]
// Output: Array Elements: -1 -3 -5
//         Positive Numbers count is: 0
//
// Example 3 (Edge Case - Zero Counts as Positive):
// Input:  Array = [0]
// Output: Array Elements: 0
//         Positive Numbers count is: 1
//   Why:    0 >= 0 is true, so zero counts toward the positive total
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random numbers are in the range [-100, 100]. | الأرقام العشوائية في النطاق [-100، 100].
// • Zero is counted as positive (condition: number >= 0). | الصفر يُحتسب كموجب (الشرط: number >= 0).
// • The array length must be a positive integer (> 0). | طول المصفوفة يجب أن يكون عددًا صحيحًا موجبًا.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
// func positiveCount(arr []int) int
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

func positiveCount(arr []int) int {
	count := 0

	for _, value := range arr {
		if value >= 0 {
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
	fmt.Printf("Positive Numbers count is: %d\n", positiveCount(arr))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(fillArrayWithRandomNumbers(readPositiveNumber("positive enter number of elements")))
}
