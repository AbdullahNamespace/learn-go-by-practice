// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 42: task_duration_in_seconds.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Structs & Time | الهياكل والوقت
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a task duration in days, hours, minutes, and seconds,
// and converts the entire duration into total seconds.
//
// AR:
// اكتب برنامجًا يقرأ مدة مهمة بالأيام والساعات والدقائق والثواني، ويحول المدة
// بالكامل إلى إجمالي الثواني.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Days=1, Hours=1, Minutes=1, Seconds=1
// Output: Task Duration In Seconds: 90061
//   Why:    (1*86400) + (1*3600) + (1*60) + 1 = 90061
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • All duration components must be positive integers | جميع مكونات المدة يجب أن تكون أعداد صحيحة موجبة
// • Use a struct to group the duration | استخدم هيكلاً (struct) لتجميع المدة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type TaskDuration struct { NumberOfDays, NumberOfHours, NumberOfMinutes, NumberOfSeconds int }
// func readTaskDuration() TaskDuration
// func taskDurationInSeconds(td TaskDuration) int
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
			printError("Invalid input please enter a valid number")
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

func readTaskDuration() TaskDuration {
	return TaskDuration{
		NumberOfDays:    readPositiveNumber("Please enter number of days"),
		NumberOfHours:   readPositiveNumber("Please enter number of hours"),
		NumberOfMinutes: readPositiveNumber("Please enter number of minutes"),
		NumberOfSeconds: readPositiveNumber("Please enter number of seconds"),
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func minutesToSeconds(minutes int) int {
	return minutes * SecondsInMinute
}

func hoursToMinutes(hours int) int {
	return hours * MinutesInHour
}

func DaysToHours(days int) int {
	return days * HoursInDays
}

func taskDurationInSeconds(td TaskDuration) int {
	// hours   = DaysToHours(td.days) + td.hours
	// minutes = hoursToMinutes(hours) + td.minutes
	// seconds = minutesToSeconds(minutes)

	hours := td.NumberOfHours + DaysToHours(td.NumberOfDays)
	minutes := td.NumberOfMinutes + hoursToMinutes(hours)
	seconds := td.NumberOfSeconds + minutesToSeconds(minutes)
	return seconds
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(td TaskDuration) {
	fmt.Printf("Task Duration In Seconds: %d\n", taskDurationInSeconds(td))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readTaskDuration())
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
)

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
