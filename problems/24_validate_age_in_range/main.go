// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 24: validate_age_in_range.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Validation / Logic | التحقق من الصحة / المنطق
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that validates whether an age is within a specific range.
// The program should:
// 1. Read an age from the user
// 2. Validate if the age is between 18 and 45 (inclusive)
// 3. Display whether the age is valid or invalid
//
// This is a simple validation problem that checks if a number falls within
// a specified range. It's commonly used in form validation, eligibility checks,
// and data validation scenarios.
//
// AR:
// اكتب برنامجًا يتحقق من صحة عمر ضمن نطاق محدد.
// يجب على البرنامج:
// 1. قراءة عمر من المستخدم
// 2. التحقق من أن العمر بين 18 و 45 (شامل)
// 3. عرض ما إذا كان العمر صالحاً أو غير صالح
//
// هذه مشكلة تحقق بسيطة تفحص ما إذا كان رقم يقع ضمن
// نطاق محدد. يستخدم بشكل شائع في التحقق من النماذج، وفحوصات الأهلية،
// وسيناريوهات التحقق من صحة البيانات.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Age = 25
// Output: 25 is a valid age
//   Why:  25 is between 18 and 45 (inclusive)
//
// Example 2:
// Input:  Age = 50
// Output: 50 is an invalid age
//   Why:  50 is greater than 45
//
// Example 3 (Edge Case):
// Input:  Age = 45
// Output: 45 is a valid age
//   Why:  45 is the upper bound (inclusive)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Age must be a positive integer | يجب أن يكون العمر عدداً صحيحاً موجباً
// • Valid range is [18, 45] (inclusive) | النطاق الصالح هو [18, 45] (شامل)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readAge() int
// func validateNumberInRange(number, from, to int) bool
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

const (
	ValidFromAge = 18
	ValidToAge   = 45
)

var reader = bufio.NewReader(os.Stdin)

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
	if validateNumberInRange(age, ValidFromAge, ValidToAge) {
		fmt.Printf("%d is a valid age\n", age)
	} else {

		fmt.Printf("%d is an invalid age\n", age)
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readAge())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
