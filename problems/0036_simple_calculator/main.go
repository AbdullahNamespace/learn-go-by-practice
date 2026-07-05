// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 36: simple_calculator.go
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
// Write a program that reads two numbers and an operation type (+, -, *, /) from
// the user, then performs the arithmetic operation and prints the result.
//
// AR:
// اكتب برنامجًا يقرأ عددين ونوع عملية (+، -، *، /) من المستخدم، ثم ينفذ العملية
// الحسابية ويطبع النتيجة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Number1 = 10, Number2 = 5, OpType = '+'
// Output: Result = 15
//
// Example 2:
// Input:  Number1 = 10, Number2 = 0, OpType = '/'
// Output: Cannot divide by zero!
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Operation type must be +, -, *, or / | يجب أن يكون نوع العملية +، -، *، أو /
// • Handle division by zero properly | التعامل مع القسمة على صفر بشكل صحيح
// • Use a custom type/enum for operation type | استخدم نوعًا مخصصًا لنوع العملية
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type OperationType byte
// func readOpType() OperationType
// func calculate(number1, number2 float64, opType OperationType) float64
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
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readOpType() OperationType {
	for {
		input, err := readString("Please enter operation type '-', '+', '/', '*'")

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}
		switch input {
		case "+":
			return '+'
		case "-":
			return '-'
		case "/":
			return '/'
		case "*":
			return '*'
		default:
			printError("Invalid input please enter valid operation in '-', '+', '/', '*'")
		}
	}
}

func (c *Calculate) Read() {
	c.number1 = readNumber("Please enter number 1")
	c.number2 = readNumber("Please enter number 2")
	c.operation = readOpType()
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculate(c Calculate) float64 {
	var result float64 = 0

	switch c.operation {
	case '+':
		result = c.number1 + c.number2

	case '-':
		result = c.number1 - c.number2

	case '*':
		result = c.number1 * c.number2

	case '/':
		{
			if c.number2 == 0 {
				printError("Cannot divide by zero")
				break
			}
			result = c.number1 / c.number2
		}
	}

	return result
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(c Calculate) {
	fmt.Printf("%f %q %f = %f\n", c.number1, c.operation, c.number2, calculate(c))
}

// ======================
//         MAIN
// ======================

func main() {
	var c Calculate
	c.Read()
	printResult(c)
}

// ======================
//   UTILITY FUNCTIONS
// ======================

type OperationType byte

type Calculate struct {
	number1   float64
	number2   float64
	operation OperationType
}

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
