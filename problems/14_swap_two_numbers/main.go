// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 14: swap_two_numbers.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Variable Manipulation | معالجة المتغيرات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads two integer numbers from the user and swaps their
// values using a temporary variable. The program should read two integers,
// display the original values before swapping, perform the swap operation using
// a temporary variable (demonstrating the classic three-way swap algorithm),
// and then display the swapped values. In Go, this can be implemented using
// pointers or tuple-style assignment.
//
// AR:
// اكتب برنامجًا يقرأ عددين صحيحين من المستخدم ويبدل قيمهما باستخدام متغير مؤقت.
// يجب على البرنامج قراءة عددين صحيحين، وعرض القيم الأصلية قبل التبديل، وإجراء
// عملية التبديل باستخدام متغير مؤقت (توضيح خوارزمية التبديل الثلاثية الكلاسيكية)،
// ثم عرض القيم المبدلة. في Go، يمكن تنفيذ ذلك باستخدام المؤشرات أو الإسناد المتعدد.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  A = 10, B = 20
// Output: Before swap: Number1 = 10, Number2 = 20
//         After swap: Number1 = 20, Number2 = 10
//   Why:    Values successfully swapped using temporary variable
//
// Example 2:
// Input:  A = 7, B = 7
// Output: Before swap: Number1 = 7, Number2 = 7
//         After swap: Number1 = 7, Number2 = 7
//   Why:    Equal values remain the same after swap
//
// Example 3 (Edge Case):
// Input:  A = -5, B = 10
// Output: Before swap: Number1 = -5, Number2 = 10
//         After swap: Number1 = 10, Number2 = -5
//   Why:    Handles negative numbers correctly
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Numbers can be any valid int integers | الأعداد يمكن أن تكون أي أعداد صحيحة int صالحة
// • Must use temporary variable for swapping | يجب استخدام متغير مؤقت للتبديل
// • Display values before and after swap | عرض القيم قبل وبعد التبديل
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumbers() (int, int)
// func swap(a, b *int)
// func printValues(label string, num1, num2 int)
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
	number1 := readNumber("Please enter number 1")
	number2 := readNumber("Please enter number 2")

	return number1, number2
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func swap(number1, number2 *int) {
	temp := *number1
	*number1 = *number2
	*number2 = temp

	// بديل: باستخدام الإسناد المتعدد في Go
	// *number1, *number2 = *number2, *number1

	// بديل: تبديل بدون متغير مؤقت (خوارزمية رياضية)
	// *number1 = *number1 + *number2
	// *number2 = *number1 - *number2
	// *number1 = *number1 - *number2
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printValues(label string, num1, num2 int) {
	fmt.Printf("%s : number1 = %d, and number2 = %d\n", label, num1, num2)
}

// ======================
//         MAIN
// ======================

func main() {
	number1, number2 := readNumbers()

	printValues("Before swap", number1, number2)

	swap(&number1, &number2)

	printValues("After swap", number1, number2)
}
