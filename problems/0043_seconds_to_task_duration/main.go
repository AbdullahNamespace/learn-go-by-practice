// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 43: seconds_to_task_duration.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Structs & Time | الهياكل والوقت
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a total number of seconds and converts it into a
// structured duration format (days, hours, minutes, seconds).
//
// AR:
// اكتب برنامجًا يقرأ إجمالي الثواني ويحوله إلى تنسيق مدة هيكلي (أيام، ساعات،
// دقائق، ثوانٍ).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  90061
// Output: 1:1:1:1
//   Why:    90061 seconds = 1 day, 1 hour, 1 minute, 1 second
//
// Example 2:
// Input:  3661
// Output: 0:1:1:1
//   Why:    3661 seconds = 0 days, 1 hour, 1 min, 1 sec
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Total seconds must be a positive integer | إجمالي الثواني يجب أن يكون عددًا صحيحًا موجبًا
// • Use integer division and modulo for conversion | استخدم القسمة الصحيحة وباقي القسمة للتحويل
// • Print in D:H:M:S format | الطباعة بصيغة يوم:ساعة:دقيقة:ثانية
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type TaskDuration struct { ... }
// func readPositiveNumber(message string) int
// func secondsToTaskDuration(totalSeconds int) TaskDuration
// func printTaskDurationDetails(td TaskDuration)
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

func secondsToTaskDuration(totalSeconds int) TaskDuration {
	var td TaskDuration
	remainder := 0

	td.NumberOfDays = totalSeconds / SecondsInDays
	remainder = totalSeconds % SecondsInDays

	td.NumberOfHours = remainder / SecondsInHour
	remainder = remainder % SecondsInHour

	td.NumberOfMinutes = remainder / SecondsInMinute

	td.NumberOfSeconds = remainder % SecondsInMinute

	return td
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printTaskDurationDetails(td TaskDuration) {
	fmt.Printf("\n%d:%d:%d:%d\n", td.NumberOfDays, td.NumberOfHours, td.NumberOfMinutes, td.NumberOfSeconds)
}

// ======================
//         MAIN
// ======================

func main() {
	printTaskDurationDetails(secondsToTaskDuration(readPositiveNumber("Please enter number of seconds")))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

type TaskDuration struct {
	NumberOfDays,
	NumberOfHours,
	NumberOfMinutes,
	NumberOfSeconds int
}

const (
	SecondsInMinute = 60
	MinutesInHour   = 60
	HoursInDays     = 24
	SecondsInHour   = SecondsInMinute * 60
	SecondsInDays   = SecondsInHour * 24
)

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
