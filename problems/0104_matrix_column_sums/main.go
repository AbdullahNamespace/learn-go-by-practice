// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 104: matrix_column_sums.go
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
// 100, prints it, then computes and prints the sum of each individual column.
//
// This is the column-oriented mirror of Problem 102's colSum logic: instead
// of fixing a row and summing across columns, it fixes a column and sums
// across rows. In Go, matrix[row][col] indexing means the outer loop for a
// column sum must iterate over the row index while holding the column index
// constant — a common source of confusion (row-sum vs. col-sum loops look
// almost identical but iterate in swapped order), so keeping colSum and
// rowSum as clearly separate, independently named functions (rather than
// one confusingly-parameterized function) is preferred for readability.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة 3x3 بأرقام عشوائية بين 1 و100، يطبعها، ثم يحسب
// ويطبع مجموع كل عمود من الأعمدة على حدة.
//
// هذه هي الصورة المرآتية العمودية لمنطق colSum في المسألة 102: بدلاً من
// تثبيت صف والجمع عبر الأعمدة، يُثبَّت عمود ويُجمَع عبر الصفوف. في غو،
// الفهرسة matrix[row][col] تعني أن الحلقة الخارجية لمجموع عمود يجب أن تمر
// على فهرس الصف بينما فهرس العمود ثابت — وهذا مصدر شائع للالتباس (حلقتا
// مجموع الصف ومجموع العمود تبدوان متطابقتين تقريبًا لكن ترتيب المرور فيهما
// معكوس)، لذا إبقاء colSum و rowSum كدالتين منفصلتين وواضح التسمية
// (بدلاً من دالة واحدة بمعامل مُربك) أفضل لسهولة القراءة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[10 20 30] [5 5 5] [1 2 3]]
// Output: Col 1 Sum = 16
//         Col 2 Sum = 27
//         Col 3 Sum = 38
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random values are in the range [1, 100] | القيم العشوائية في المدى [1، 100]
// • Column numbers are printed 1-based | أرقام الأعمدة تُطبع بدءًا من 1
// • colSum must not have any printing side effects | يجب ألا يكون لدالة
//   colSum أي أثر جانبي في الطباعة
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
// func colSum(matrix [][]int, col int) int
// func printMatrix(matrix [][]int)
// func printEachColSum(matrix [][]int)
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
		min, max = max, min
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

func colSum(matrix [][]int, col int) int {
	if col < 0 {
		return 0
	}
	sum := 0
	for i := range matrix {
		if col < len(matrix[i]) {
			sum += matrix[i][col]
		}
	}
	return sum
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printMatrix(matrix [][]int) {
	for _, valueX := range matrix {
		for _, valueY := range valueX {
			fmt.Printf("%d\t", valueY)
		}
		fmt.Println()
	}
}

func printEachColSum(matrix [][]int) {
	for i := range matrix {
		fmt.Printf("\nCol %d Sum = %d", i+1, colSum(matrix, i))
	}
}

// ======================
//         MAIN
// ======================

func main() {
	matrix := newMatrix()
	fillMatrixWithRandomNumbers(matrix)
	printMatrix(matrix)
	printEachColSum(matrix)
}
