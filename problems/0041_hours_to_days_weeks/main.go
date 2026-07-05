// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 41: hours_to_days_weeks.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Time Conversion | تحويل الوقت
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a number of hours and converts them to days and weeks.
//
// AR:
// اكتب برنامجًا يقرأ عدد الساعات ويحولها إلى أيام وأسابيع.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  168
// Output: Total Hours = 168
//         Total Days = 7
//         Total Weeks = 1
//   Why:    168 / 24 = 7 days, 168 / 168 = 1 week
//
// Example 2:
// Input:  48
// Output: Total Hours = 48
//         Total Days = 2
//         Total Weeks = 0.2857...
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Hours must be a positive number | الساعات يجب أن تكون رقماً موجباً
// • 1 Day = 24 Hours, 1 Week = 24 * 7 Hours | يوم = 24 ساعة، أسبوع = 168 ساعة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) float64
// func hoursToDays(numberOfHours float64) float64
// func hoursToWeeks(numberOfHours float64) float64
// func daysToWeeks(numberOfDays float64) float64
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

func readNumber(prompt string) float64 {
	for {
		input, err := readString(prompt)

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

		inputNumber, err := strconv.ParseFloat(input, 64)

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) float64 {
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

func hoursToDays(numberOfHours float64) float64 {
	return numberOfHours / hoursInDay
}

func hoursToWeeks(numberOfHours float64) float64 {
	return daysToWeeks(hoursToDays(numberOfHours))
	//  other
	return numberOfHours / hoursInWeek
}

func daysToWeeks(numberOfDays float64) float64 {
	return numberOfDays / DaysInWeek
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(totalHours float64) {
	fmt.Printf("\nTotal Hours = %f", totalHours)
	fmt.Printf("\nTotal Days = %f", hoursToDays(totalHours))
	fmt.Printf("\nTotal Weeks = %f\n", hoursToWeeks(totalHours))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Please enter number of Hours"))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

const (
	hoursInDay  = 24
	DaysInWeek  = 7
	hoursInWeek = DaysInWeek * hoursInDay
)

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
