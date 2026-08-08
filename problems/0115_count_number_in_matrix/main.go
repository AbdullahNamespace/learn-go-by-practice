// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 115: count_number_in_matrix.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices & Searching | المصفوفات الثنائية والبحث
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that builds a predefined 3x3 matrix containing repeated
// values, prints it, asks the user for a number, then counts and prints how
// many times that number appears anywhere in the matrix.
//
// Reading the target number should reuse a validated readNumber-style input
// function (from earlier problems) rather than reading raw, unchecked input
// directly — skipping validation entirely means a non-numeric entry would
// go completely unhandled. Consistently validating all numeric input
// through one shared reader, as this whole problem series has done since
// Problem 84, avoids reintroducing that gap here.
//
// العربية:
// اكتب برنامجًا يبني مصفوفة محددة مسبقًا بحجم 3x3 تحتوي على قيم مكررة،
// يطبعها، يسأل المستخدم عن رقم، ثم يحسب ويطبع عدد مرات ظهور ذلك الرقم في
// أي مكان في المصفوفة.
//
// يجب أن تُعيد قراءة الرقم المستهدف استخدام دالة إدخال بنمط readNumber
// مُتحقق منها (من مسائل سابقة) بدلاً من قراءة إدخال خام غير مُتحقق منه
// مباشرة — تجاهل التحقق تمامًا يعني أن إدخالًا غير رقمي سيمر دون معالجة
// على الإطلاق. التحقق المتسق من كل إدخال رقمي عبر قارئ واحد مشترك، كما
// فعلت هذه السلسلة كاملة منذ المسألة 84، يتجنب إعادة إدخال هذه الثغرة هنا.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[9 1 12] [0 9 1] [0 9 9]], number to count = 9
// Output: Number 9 count in matrix is: 4
//
// Example 2:
// Input:  Matrix = [[9 1 12] [0 9 1] [0 9 9]], number to count = 100
// Output: Number 100 count in matrix is: 0
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Must scan every cell exactly once — this counts all occurrences, so no
//   early exit is possible | يجب فحص كل خلية مرة واحدة بالضبط — هذا يحسب
//   كل التكرارات، فلا يمكن الخروج المبكر
// • The searched number must be read via a validated numeric input function
//   | يجب قراءة الرقم المطلوب عبر دالة إدخال رقمي مُتحقق منها
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) int
// func sampleMatrixWithDuplicates() [][]int
// func countNumberInMatrix(matrix [][]int, number int) int
// func printMatrix(matrix [][]int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"strconv"
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

func readString(prompt string) (string, error) {
	fmt.Printf("%s : ", prompt)

	var input string

	_, err := fmt.Scan(&input)

	if err != nil {
		return "", err
	}

	return input, nil
}

func readNumber(prompt string) int {
	for {
		input, err := readString(prompt)
		if err != nil {

			printError("Invalid input please enter again!")
			continue
		}

		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Printf("Input input please enter a valid number!")
			continue
		}

		return number
	}
}

func sampleMatrixWithDuplicates() [][]int {
	return [][]int{{1, 2, 5}, {1, 2, 5}, {1, 2, 5}}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func countNumberInMatrix(matrix [][]int, number int) int {
	count := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == number {
				count++
			}
		}
	}
	return count
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printMatrix(matrix [][]int, prompt string) {
	fmt.Printf("%s\n", prompt)
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
}

// ======================
//         MAIN
// ======================

func main() {
	matrix := sampleMatrixWithDuplicates()
	printMatrix(matrix, "Matrix is : ")
	countNumber := readNumber("Please enter number to count in matrix")
	fmt.Printf("Number %d count in matrix is: %d", countNumber, countNumberInMatrix(matrix, countNumber))

}
