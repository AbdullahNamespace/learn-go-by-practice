// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 8: pass_fail_checker.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Conditional Logic / Constants | المنطق الشرطي / الثوابت
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Create a program that reads an exam mark from the user and determines whether the student
// has passed or failed. The passing mark is 50 or above. Use Constants `Pass` and `Fail`
// to represent the status. The program should use `switch` to print the appropriate message
// ("You Passed" or "You Failed").
//
// العربية:
// إنشاء برنامج يقرأ علامة الامتحان من المستخدم ويحدد ما إذا كان الطالب ناجحاً أم راسباً.
// علامة النجاح هي 50 فما فوق. استخدم ثوابت `Pass` و `Fail` لتمثيل الحالة.
// يجب أن يستخدم البرنامج `switch` لطباعة رسالة النتيجة المناسبة ("You Passed" أو "You Failed").
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  65
// Output: "You Passed"
//
// Example 2:
// Input:  42
// Output: "You Failed"
//
// Example 3 (Edge Case):
// Input:  50
// Output: "You Passed"
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Passing grade is 50. | درجة النجاح هي 50.
// • Input is an integer. | المدخل عدد صحيح.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readMark() int
// func checkMark(mark int) string
// func printResults(mark int)
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
	PassingMark = 50
	Pass        = "Pass"
	Fail        = "Fail"
)

var reader = bufio.NewReader(os.Stdin)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) (string, error) {
	fmt.Printf("%s", fmt.Sprintf("%s :", prompt))
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
			continue
		}
		inputNumber, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		return inputNumber
	}
}

func readMark() int {
	for {
		input := readNumber("Please enter mark")
		if input >= 1 && input <= 100 {
			return input
		}
		fmt.Printf("\nInvalid input please enter a valid mark in (1-100)")
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func checkMark(mark int) string {
	if mark >= PassingMark {
		return Pass
	}
	return Fail
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResults(mark int) {
	fmt.Println("\n========================================")
	fmt.Println("📊 EXAM RESULT")
	fmt.Println("========================================")
	fmt.Printf("Mark:          %d\n", mark)
	fmt.Printf("Passing Mark:  %d\n", PassingMark)

	result := checkMark(mark)
	switch result {
	case Pass:
		fmt.Println("Status:        ✅ 🎉 You Passed!")
		fmt.Printf("Note:          Passed by %d points\n", mark-PassingMark)
	case Fail:
		fmt.Println("Status:        ❌ You Failed!")
		fmt.Printf("Note:          Need %d more points to pass\n", PassingMark-mark)
	}
	fmt.Println("========================================")
}

// ======================
//         MAIN
// ======================

func main() {
	fmt.Println("🎯 PASS/FAIL CHECKER")
	fmt.Println("-------------------------")

	mark := readMark()
	printResults(mark)
}
