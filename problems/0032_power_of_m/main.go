// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 32: power_of_m.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Loops | الرياضيات والحلقات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a base number and an exponent (power) from the user,
// and computes the result using an iterative loop. Any number raised to the power
// of 0 is 1.
//
// AR:
// اكتب برنامجًا يقرأ عددًا أساسيًا وأُسًا (قوة) من المستخدم، ويحسب النتيجة
// باستخدام حلقة تكرارية. أي عدد مرفوع للقوة 0 يساوي 1.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Number = 2, M = 5
// Output: Result = 32
//   Why:    2 * 2 * 2 * 2 * 2 = 32
//
// Example 2:
// Input:  Number = 5, M = 0
// Output: Result = 1
//   Why:    Any number to the power of 0 is 1
//
// Example 3:
// Input:  Number = 3, M = 3
// Output: Result = 27
//   Why:    3 * 3 * 3 = 27
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Numbers can be valid integers | الأعداد يمكن أن تكون أعداد صحيحة صالحة
// • Must use a loop for calculation, not math.Pow | يجب استخدام حلقة للحساب، وليس math.Pow
// • Handle M = 0 edge case | التعامل مع حالة M = 0
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) int
// func powerOfM(number int, m int) int
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

// ======================
//   UTILITY
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}

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
			printError("Invalid input please enter again!")
			continue
		}

		return inputNumber
	}
}

func readNumbers() (int, int) {
	number := readNumber("Please enter number")
	m := readNumber("Please enter 'M'")
	return number, m
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func powerOfM(number int, m int) int {
	result := 1

	for i := 1; i <= m; i++ {
		result *= number
	}

	return result
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(number, m int) {
	fmt.Printf("\nThe Power number %d of %d is %d\n", number, m, powerOfM(number, m))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readNumbers())
}
