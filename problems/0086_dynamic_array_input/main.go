// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 86: dynamic_array_input.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Loops | المصفوفات والحلقات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that repeatedly asks the user to enter a number and appends
// it to a dynamically growing slice. After each entry, the program asks the
// user whether they want to add another number (yes/no). The loop continues
// until the user chooses to stop. Finally, the program prints the total number
// of elements entered and the full array contents.
//
// العربية:
// اكتب برنامجًا يطلب من المستخدم بشكل متكرر إدخال رقم ويضيفه إلى شريحة (slice)
// تنمو ديناميكيًا. بعد كل إدخال، يسأل البرنامج المستخدم إذا كان يريد إضافة رقم
// آخر (نعم/لا). تستمر الحلقة حتى يختار المستخدم التوقف. أخيرًا، يطبع البرنامج
// إجمالي عدد العناصر المُدخلة ومحتويات المصفوفة كاملة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Number: 10, Add more? Yes
//         Number: 20, Add more? Yes
//         Number: 30, Add more? No
// Output: Array Length: 3
//         Array elements: 10 20 30
//
// Example 2:
// Input:  Number: 5, Add more? No
// Output: Array Length: 1
//         Array elements: 5
//
// Example 3 (Edge Case - No Elements):
// Input:  (User immediately chooses No without adding)
// Output: Array Length: 0
//         Array elements:
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • No upper limit on the number of entries (use a dynamic slice, not a fixed-size array) | لا يوجد حد أقصى للإدخالات، استخدم شريحة ديناميكية وليس مصفوفة ثابتة الحجم
// • The "add more?" prompt must accept a clear yes/no format | يجب أن يقبل سؤال "إضافة المزيد؟" صيغة واضحة نعم/لا
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) int
// func addArrayElement() []int
// func printArray(arr []int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ======================
//   UTILITY
// ======================

func printError(prompt string) {
	fmt.Printf("X Error : %s\n", prompt)
}

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) string {
	fmt.Printf("%s : ", prompt)

	var input string

	fmt.Scan(&input)

	return input
}

func readNumber(prompt string) int {
	for {
		input := readString(prompt)

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

func readBoolean(prompt string) bool {
	for {
		input := readString(fmt.Sprintf("%s 'Yes or No'", prompt))

		input = strings.ToLower(input)

		switch input {
		case "y", "yes", "t", "true", "1":
			return true

		case "n", "no", "f", "false", "0":
			return false
		}
		printError("Invalid input you must enter 'Yes(y) or No(n)'")
	}
}

func addArrayElement() []int {
	arr := []int{}
	for {
		arr = append(arr, readPositiveNumber("Please enter number to add to array"))
		addMore := readBoolean("Do you want to add more")
		if !addMore {
			break
		}
	}
	return arr
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array Length: %d\n", len(arr))
	fmt.Printf("Array is : %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	arr := addArrayElement()

	printArray(arr)
}
