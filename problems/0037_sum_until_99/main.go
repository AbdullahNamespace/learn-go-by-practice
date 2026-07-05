// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 37: sum_until_99.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Loops & Accumulators | الحلقات والمجمعات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads numbers from the user and sums them until the user
// enters -99 to stop.
//
// AR:
// اكتب برنامجًا يقرأ أرقامًا من المستخدم ويجمعها حتى يدخل المستخدم -99 للتوقف.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  10, 20, 30, -99
// Output: Result = 60
//   Why:    10 + 20 + 30 = 60 (-99 is not added)
//
// Example 2:
// Input:  -99
// Output: Result = 0
//   Why:    Loop stops immediately
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Stop condition is entering -99 | شرط التوقف هو إدخال -99
// • -99 should not be added to the sum | لا يجب إضافة -99 إلى المجموع
// • Use an infinite loop with a break condition | استخدم حلقة لا نهائية مع شرط كسر
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func sumNumbers() float64
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
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) (string, error) {
	fmt.Printf("%s :  ", prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func readNumber(prompt string) float64 {
	for {
		input, err := readString(prompt)

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

		inputNumber, err := strconv.ParseFloat(input, 64)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}
		return inputNumber
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func sumNumbers() float64 {
	var sum float64

	for {
		input := readNumber("")
		if input == -99 {
			break
		}
		sum += input
	}

	return sum
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(result float64) {
	fmt.Printf("Result = %f\n", result)
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(sumNumbers())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s", prompt)
}
