// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 13: max_of_3_numbers.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Comparison & Conditional Logic | المقارنة والمنطق الشرطي
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads three integer numbers from the user and determines
// which one is the greatest. The program should read three integers (A, B, C),
// compare all three numbers using nested conditional logic, and display the
// maximum (largest) number among them. The comparison should handle all possible
// ordering scenarios including equal values and negative numbers.
//
// AR:
// اكتب برنامجًا يقرأ ثلاثة أعداد صحيحة من المستخدم ويحدد أيها الأكبر. يجب على
// البرنامج قراءة ثلاثة أعداد صحيحة (A, B, C)، مقارنة الأعداد الثلاثة باستخدام
// المنطق الشرطي المتداخل، وعرض العدد الأكبر (الأقصى) من بينها. يجب أن تتعامل
// المقارنة مع جميع سيناريوهات الترتيب الممكنة بما في ذلك القيم المتساوية
// والأعداد السالبة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  A = 10, B = 5, C = 8
// Output: The Maximum Number is: 10
//   Why:    10 is greater than both 5 and 8
//
// Example 2:
// Input:  A = 7, B = 7, C = 7
// Output: The Maximum Number is: 7
//   Why:    All numbers are equal, return any
//
// Example 3 (Edge Case):
// Input:  A = -5, B = -10, C = -3
// Output: The Maximum Number is: -3
//   Why:    -3 is the largest negative number
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Numbers can be any valid int integers | الأعداد يمكن أن تكون أي أعداد صحيحة int صالحة
// • Must use nested conditional statements (if-else) | يجب استخدام عبارات شرطية متداخلة
// • Should handle equal values and negative numbers | يجب التعامل مع القيم المتساوية والأعداد السالبة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumbers() (int, int, int)
// func maxOf3Numbers(a, b, c int) int
// func printResults(number1, number2, number3 int)
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
	fmt.Printf("\n%s : ", prompt)

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
			fmt.Println("Error: invalid input please enter again!")
			continue
		}
		inputNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: invalid input please enter a valid number!")
			continue
		}
		return inputNumber
	}
}

func readNumbers() (int, int, int) {
	number1 := readNumber("Please enter number 1")
	number2 := readNumber("Please enter number 2")
	number3 := readNumber("Please enter number 3")

	return number1, number2, number3
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func maxOf3Numbers(number1, number2, number3 int) int {
	if number1 > number2 {
		if number1 > number3 {
			return number1
		}
		return number3
	}
	if number2 > number3 {
		return number2
	}
	return number3

	// بديل: باستخدام متغير مؤقت
	max := number1

	if number2 > max {
		max = number2
	}
	if number3 > max {
		max = number3
	}
	return max
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResults(number1, number2, number3 int) {
	fmt.Printf("\nThe Maximum Number is: %d\n\n", maxOf3Numbers(number1, number2, number3))
}

// ======================
//         MAIN
// ======================

func main() {
	fmt.Println("max of 3 numbers")
	printResults(readNumbers())
}
