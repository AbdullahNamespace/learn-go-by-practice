// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 11: average_with_pass_fail.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Conditional Logic | الرياضيات والمنطق الشرطي
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads three exam marks from the user, calculates their
// average, and determines whether the student has passed or failed. A student
// passes if the average is 50 or greater. The program should:
// - Read three integer marks from the user
// - Calculate the sum of the marks
// - Compute the average (as a floating-point number)
// - Check if average >= 50 (Pass) or < 50 (Fail)
// - Display the average and the pass/fail status
//
// AR:
// اكتب برنامجًا يقرأ ثلاث علامات امتحان من المستخدم، ويحسب متوسطها، ويحدد ما إذا
// كان الطالب قد نجح أو رسب. الطالب ينجح إذا كان المتوسط 50 أو أكثر. يجب على البرنامج:
// - قراءة ثلاث علامات صحيحة من المستخدم
// - حساب مجموع العلامات
// - حساب المتوسط (كرقم عشري)
// - التحقق من المتوسط >= 50 (نجاح) أو < 50 (رسوب)
// - عرض المتوسط وحالة النجاح/الرسوب
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1 (Pass):
// Input:  Mark1 = 60, Mark2 = 70, Mark3 = 80
// Output: Your Average is: 70.00
//         You Passed
//
// Example 2 (Fail):
// Input:  Mark1 = 30, Mark2 = 40, Mark3 = 45
// Output: Your Average is: 38.33
//         You Failed
//
// Example 3 (Boundary - Pass):
// Input:  Mark1 = 50, Mark2 = 50, Mark3 = 50
// Output: Your Average is: 50.00
//         You Passed
//
// Example 4 (Boundary - Fail):
// Input:  Mark1 = 49, Mark2 = 49, Mark3 = 50
// Output: Your Average is: 49.33
//         You Failed
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Marks should be integers (0-100 recommended) | العلامات أعداد صحيحة (يُنصح 0-100)
// • Pass threshold is 50 (inclusive) | عتبة النجاح هي 50 (شاملة)
// • Average is computed as float64 for precision | المتوسط يُحسب كـ float64 للدقة
// • Input validation recommended but not required | التحقق من المدخلات مُوصى به لكن غير مطلوب
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type PassFail string
// func readMarks() (int, int, int)
// func sumOf3Marks(mark1, mark2, mark3 int) int
// func calculateAverage(mark1, mark2, mark3 int) float64
// func checkAverage(average float64) PassFail
// func printResults(mark1, mark2, mark3 int)
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

type PassFail string

const (
	passThreshold = 50.0
	numMarks      = 3
)

// ======================
//        TYPES
// ======================

const (
	Pass PassFail = "PASS"
	Fail PassFail = "FAIL"
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

func readNumber(prompt string) float64 {
	for {
		input, err := readString(prompt)

		if err != nil {
			fmt.Println("Error: invalid input please enter again!")
			continue
		}

		inputNumber, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Error: invalid input please enter a number!")
			continue
		}
		return inputNumber
	}
}

func readMark(prompt string) float64 {
	for {
		input := readNumber(prompt)
		if input < 0 || input > 100 {
			fmt.Println("Error: invalid input please enter a valid mark in (1-100)")
			continue
		}

		return input
	}
}

func readMarks() (float64, float64, float64) {
	mark1 := readMark("Please enter mark 1")
	mark2 := readMark("Please enter mark 2")
	mark3 := readMark("Please enter mark 3")

	return mark1, mark2, mark3
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func sumOf3Marks(mark1, mark2, mark3 float64) float64 {
	return mark1 + mark2 + mark3
}

func calculateAverage(mark1, mark2, mark3 float64) float64 {
	return sumOf3Marks(mark1, mark2, mark3) / float64(numMarks)
}

func checkAverage(average float64) PassFail {
	if average >= passThreshold {
		return Pass
	}
	return Fail
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResults(mark1, mark2, mark3 float64) {
	sum := sumOf3Marks(mark1, mark2, mark3)
	average := calculateAverage(mark1, mark2, mark3)
	result := checkAverage(average)

	fmt.Println()
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("EXAM RESULTS ANALYSIS")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("Marks:     %f, %f, %f\n", mark1, mark2, mark3)
	fmt.Printf("Sum:       %f\n", sum)
	fmt.Printf("Average:   %.2f\n", average)
	fmt.Printf("Threshold: %.0f\n", passThreshold)
	fmt.Println(strings.Repeat("-", 40))

	switch result {
	case Pass:
		fmt.Println("Status:    PASSED")
		margin := average - passThreshold
		fmt.Printf("Margin:    +%.2f points above\n", margin)
	case Fail:
		fmt.Println("Status:    FAILED")
		deficit := passThreshold - average
		fmt.Printf("Deficit:   %.2f points below\n", deficit)
	}

	fmt.Println(strings.Repeat("=", 40))
}

// ======================
//         MAIN
// ======================

func main() {
	fmt.Println("AVERAGE WITH PASS/FAIL CHECKER")
	fmt.Println(strings.Repeat("-", 30))

	m1, m2, m3 := readMarks()
	printResults(m1, m2, m3)
}
