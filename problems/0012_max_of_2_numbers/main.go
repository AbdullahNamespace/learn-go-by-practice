// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 12: max_of_2_numbers.go
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
// Write a program that reads two integer numbers from the user and determines
// which one is greater. The program should:
// - Read two integers from the user
// - Compare the two numbers
// - Display the maximum (larger) number
// The comparison should use simple if-else logic to determine the result.
//
// AR:
// اكتب برنامجًا يقرأ عددين صحيحين من المستخدم ويحدد أيهما أكبر. يجب على البرنامج:
// - قراءة عددين صحيحين من المستخدم
// - مقارنة العددين
// - عرض العدد الأكبر (الأقصى)
// يجب أن تستخدم المقارنة منطق if-else بسيط لتحديد النتيجة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Num1 = 10, Num2 = 5
// Output: The Maximum Number is: 10
//
// Example 2:
// Input:  Num1 = 25, Num2 = 100
// Output: The Maximum Number is: 100
//
// Example 3 (Equal Numbers):
// Input:  Num1 = 50, Num2 = 50
// Output: The Maximum Number is: 50
//
// Example 4 (Negative Numbers):
// Input:  Num1 = -10, Num2 = -5
// Output: The Maximum Number is: -5
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Numbers can be any valid int integers | الأعداد يمكن أن تكون أي أعداد صحيحة int صالحة
// • When numbers are equal, either can be returned | عندما تكون الأعداد متساوية، يمكن إرجاع أي منها
// • Input validation recommended for robustness | التحقق من المدخلات موصى به للمتانة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumbers() (int, int)
// func maxOf2Numbers(num1, num2 int) int
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
	fmt.Printf("%s :", prompt)

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

func readNumbers() (int, int) {
	num1 := readNumber("Please enter number 1")
	num2 := readNumber("Please enter number 2")

	return num1, num2
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func maxOf2Numbers(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2

}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printMaxResult(maxNum int) {
	fmt.Printf("\nThe Maximum Number is: %d \n", maxNum)
}

// ======================
//         MAIN
// ======================

func main() {
	fmt.Println("")
	printMaxResult(maxOf2Numbers(readNumbers()))
}
