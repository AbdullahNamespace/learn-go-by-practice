// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 45: month_of_year.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Enums & Switch | المعدادات والتبديل
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that prompts the user to enter a month number (1-12) and prints
// the corresponding month name.
//
// AR:
// اكتب برنامجًا يطالب المستخدم بإدخال رقم شهر (1-12) ويطبع اسم الشهر المقابل.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  1
// Output: January
//
// Example 2:
// Input:  12
// Output: December
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Month number must be between 1 and 12 inclusive | رقم الشهر يجب أن يكون بين 1 و 12
// • Use a custom type for MonthOfYear | استخدم نوعًا مخصصًا لأشهر السنة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type MonthOfYear int
// func readMonthOfYear() MonthOfYear
// func getMonthOfYear(month MonthOfYear) string
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

func readNumberInRange(from, to int, prompt string) int {
	for {
		input := readNumber(prompt)

		if isNumberInRange(input, from, to) {
			return input
		}

		printError(fmt.Sprintf("Invalid input please enter a number in range (%d-%d)", from, to))
	}
}

func readMonthOfYear() MonthOfYear {
	return MonthOfYear(readNumberInRange(1, 12, "Please enter number of month"))
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func getMonthOfYear(month MonthOfYear) string {
	return time.Month(month).String()
}

func isNumberInRange(number, from, to int) bool {
	return number >= from && number <= to
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(month MonthOfYear) {
	fmt.Printf("%s\n", getMonthOfYear(month))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readMonthOfYear())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

type MonthOfYear int

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
