// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 102: matrix_row_sums.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices & Math | المصفوفات الثنائية والرياضيات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills a 3x3 matrix with random numbers between 1 and
// 100, prints it, then computes and prints the sum of each individual row.
//
// The row-sum logic isolates a single responsibility — "sum the values of
// row i" — into its own small function that takes the matrix and a row
// index and returns an int. Keeping this calculation separate from the
// printing loop (which just calls it once per row) mirrors the same
// separation-of-concerns pattern used for 1D arrays in earlier problems
// (e.g. oddCount, evenCount): one pure function does the math, one output
// function handles formatting, and neither knows about the other's job.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة 3x3 بأرقام عشوائية بين 1 و100، يطبعها، ثم يحسب
// ويطبع مجموع كل صف من الصفوف على حدة.
//
// يعزل منطق "مجموع الصف" مسؤولية واحدة — "اجمع قيم الصف رقم i" — في دالة
// صغيرة مستقلة تأخذ المصفوفة ورقم الصف وتُرجع int. فصل هذا الحساب عن حلقة
// الطباعة (التي تستدعيه مرة واحدة لكل صف فقط) يُطابق نفس نمط فصل المسؤوليات
// المستخدم مع المصفوفات أحادية البُعد في مسائل سابقة (مثل oddCount،
// evenCount): دالة نقية تقوم بالحساب، ودالة إخراج منفصلة تتولى التنسيق، ولا
// تعرف إحداهما عن عمل الأخرى.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[10 20 30] [5 5 5] [1 2 3]]
// Output: Row 1 Sum = 60
//         Row 2 Sum = 15
//         Row 3 Sum = 6
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random values are in the range [1, 100] | القيم العشوائية في المدى [1، 100]
// • Row numbers are printed 1-based (Row 1, Row 2, ...) even though the
//   underlying index is 0-based | أرقام الصفوف تُطبع بدءًا من 1 رغم أن
//   الفهرسة الداخلية تبدأ من 0
// • rowSum must not have any printing side effects (pure calculation only)
//   | يجب ألا يكون لدالة rowSum أي أثر جانبي في الطباعة (حساب خالص فقط)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// const MatrixSize = 3
// func randomNumber(from, to int) int
// func newMatrix() [][]int
// func fillMatrixWithRandomNumbers(matrix [][]int)
// func rowSum(matrix [][]int, row int) int
// func printEachRowSum(matrix [][]int)
// func printMatrix(matrix [][]int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"math/rand/v2"
)

// ======================
//   UTILITY
// ======================

const MatrixSize = 3

func randomNumber(min, max int) int {
	if min > max {
		max, min = min, max
	}
	return rand.IntN(max-min+1) + min
}

// ======================
//     INPUT FUNCTIONS
// ======================

func newMatrix() [][]int {
	matrix := make([][]int, MatrixSize)
	for i := range matrix {
		matrix[i] = make([]int, MatrixSize)
	}
	return matrix
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func fillMatrixWithRandomNumbers(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = randomNumber(1, 100)
		}
	}
}

func rowSum(matrix [][]int, row int) int {
	sum := 0
	if row >= 0 && row < len(matrix) {
		for _, value := range matrix[row] {
			sum += value
		}
	}
	return sum
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printEachRowSum(matrix [][]int) {
	for i := range matrix {
		fmt.Printf("\nRow %d Sum = %d", i+1, rowSum(matrix, i))
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, column := range row {
			fmt.Printf("%d\t", column)
		}
		fmt.Println()
	}
}

// ======================
//         MAIN
// ======================

func main() {
	matrix := newMatrix()
	fillMatrixWithRandomNumbers(matrix)
	printMatrix(matrix)
	printEachRowSum(matrix)
}
