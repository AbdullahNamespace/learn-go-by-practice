// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 25: read_age_until_valid.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Validation / Loops | التحقق من الصحة / الحلقات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that repeatedly prompts the user to enter an age until
// a valid age within the range [18, 45] is provided.
// The program should:
// 1. Repeatedly read an age from the user
// 2. Validate if the age is between 18 and 45 (inclusive)
// 3. Continue prompting until a valid age is entered
// 4. Display the valid age once accepted
//
// This demonstrates input validation with retry logic, a common pattern in
// interactive programs. It ensures the program only proceeds with valid data.
//
// AR:
// اكتب برنامجًا يطلب من المستخدم بشكل متكرر إدخال عمر حتى
// يتم تقديم عمر صالح ضمن النطاق [18, 45].
// يجب على البرنامج:
// 1. قراءة عمر من المستخدم بشكل متكرر
// 2. التحقق من أن العمر بين 18 و 45 (شامل)
// 3. الاستمرار في الطلب حتى يتم إدخال عمر صالح
// 4. عرض العمر الصالح بمجرد قبوله
//
// هذا يوضح التحقق من صحة الإدخال مع منطق إعادة المحاولة، نمط شائع في
// البرامج التفاعلية. يضمن أن البرنامج يتقدم فقط ببيانات صحيحة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1 (First Input Valid):
// Input:  Age = 30
// Output: Your Age is: 30
//   Why:  30 is within the valid range on first try
//
// Example 2 (Multiple Attempts):
// Input:  Age = 10 (invalid) → Age = 50 (invalid) → Age = 25 (valid)
// Output: Your Age is: 25
//   Why:  Program keeps asking until valid age is entered
//
// Example 3 (Edge Case - Lower Bound):
// Input:  Age = 17 (invalid) → Age = 18 (valid)
// Output: Your Age is: 18
//   Why:  18 is the minimum valid age
//
// Example 4 (Edge Case - Upper Bound):
// Input:  Age = 46 (invalid) → Age = 45 (valid)
// Output: Your Age is: 45
//   Why:  45 is the maximum valid age
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Age must be a positive integer | يجب أن يكون العمر عدداً صحيحاً موجباً
// • Valid range is [18, 45] (inclusive) | النطاق الصالح هو [18, 45] (شامل)
// • Program must loop until valid input is received | يجب أن يكرر البرنامج حتى يتم إدخال صالح
// • Should handle invalid input gracefully | يجب معالجة الإدخال غير الصحيح بشكل لطيف
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readAge() int
// func validateNumberInRange(number, from, to int) bool
// func readUntilAgeBetween(from, to int) int
// func printResult(age int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const (
	ValidFromAge = 18
	ValidToAge   = 45
)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) (string, error) {
	fmt.Printf("%s : ", prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readNumber(prompt string) int {
	for {
		input, err := readString(prompt)

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

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

func readAge() int {
	return readPositiveNumber("Please enter your age")
}

func readUntilAgeBetween(from, to int) int {
	for {
		input := readAge()

		if validateNumberInRange(input, from, to) {
			return input
		}
		printError(fmt.Sprintf("Invalid input please enter age between %d and %d", from, to))
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func validateNumberInRange(number, from, to int) bool {
	return number >= from && number <= to
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(age int) {
	fmt.Printf("Your Age is: %d\n", age)
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readUntilAgeBetween(ValidFromAge, ValidToAge))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
