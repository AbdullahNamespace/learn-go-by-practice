// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 33: grade_letter.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Conditional Logic | المنطق الشرطي
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that prompts the user to enter a grade within a specified range
// (0 to 100), ensures valid input, and determines the corresponding letter grade.
//
// AR:
// اكتب برنامجًا يطالب المستخدم بإدخال درجة ضمن نطاق محدد (0 إلى 100)، ويضمن
// صحة المدخلات، ويحدد الدرجة الحرفية المقابلة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  95
// Output: A
//   Why:    95 is greater than or equal to 90
//
// Example 2:
// Input:  45
// Output: F
//   Why:    45 is less than 50
//
// Example 3 (Boundary):
// Input:  50
// Output: E
//   Why:    50 is the lower bound for passing
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Grade must be between 0 and 100 inclusive | الدرجة يجب أن تكون بين 0 و 100 شاملة
// • Keep prompting until valid input is received | استمر في الطلب حتى يتم إدخال قيمة صحيحة
// • Grading scale: 90:A, 80:B, 70:C, 60:D, 50:E, <50:F | مقيح الدرجات حسب النسب المحددة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumberInRange(from, to int, message string) int
// func getGradeLetter(grade int) string
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
			printError("Invalid input please enter a  valid number!")
			continue
		}

		return inputNumber
	}
}

func readNumberInRange(from, to int, prompt string) int {
	for {
		input := readNumber(prompt)

		if isNumberInRange(input, from, to) {
			return input
		}

		printError(fmt.Sprintf("Invalid input please enter number in range from %d to %d", from, to))
	}
}

func readGrade(prompt string) int {
	return readNumberInRange(0, 100, prompt)
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func isNumberInRange(number, from, to int) bool {
	return number >= from && number <= to
}

func getGradeLetter(grade int) string {
	switch {
	case grade >= 90:
		return "A"

	case grade >= 80:
		return "B"

	case grade >= 70:
		return "C"

	case grade >= 60:
		return "D"

	case grade >= 50:
		return "E"

	default:
		return "F"
	}
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(grade int) {
	gradeLetter := getGradeLetter(grade)
	fmt.Println(gradeLetter)
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readGrade("Please enter grade"))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
