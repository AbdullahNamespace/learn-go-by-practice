// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 44: day_of_week.go
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
// Write a program that prompts the user to enter a day number (1-7) and prints
// the corresponding day name.
//
// AR:
// اكتب برنامجًا يطالب المستخدم بإدخال رقم يوم (1-7) ويطبع اسم اليوم المقابل.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  1
// Output: Saturday
//   Why:    1 corresponds to Saturday in this specific enum
//
// Example 2:
// Input:  5
// Output: Wednesday
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Day number must be between 1 and 7 inclusive | رقم اليوم يجب أن يكون بين 1 و 7
// • Use a custom type for DayOfWeek | استخدم نوعًا مخصصًا لأيام الأسبوع
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type DayOfWeek int
// func readDayOfWeek() DayOfWeek
// func getDayOfWeek(day DayOfWeek) string
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

func readDayOfWeek() DayOfWeek {
	return DayOfWeek(readNumberInRange(1, 7, "Please enter number of day"))
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func isNumberInRange(number, from, to int) bool {
	return number >= from && number <= to
}

func getDayOfWeek(day DayOfWeek) string {
	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	return days[day-1]
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(day DayOfWeek) {
	fmt.Printf("%s\n", getDayOfWeek(day))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readDayOfWeek())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

type DayOfWeek int

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
