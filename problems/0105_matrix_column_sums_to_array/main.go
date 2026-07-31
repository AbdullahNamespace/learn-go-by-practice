// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 105: matrix_column_sums_to_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices & Arrays | المصفوفات الثنائية والأحادية
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills a 3x3 matrix with random numbers between 1 and
// 100, prints it, computes the sum of each column, collects those sums into
// a separate 1D slice, and finally prints that slice.
//
// This is the direct counterpart to Problem 103, following the exact same
// shape — reuse colSum instead of rowSum, and return a new []int rather than
// filling a caller-provided array. Recognizing that Problems 103 and 105
// share an identical structural skeleton (only the underlying sum function
// differs) is itself a useful Go lesson: it hints that a more general
// higher-order helper — e.g. a function that takes a "sum a line" function
// as a parameter and applies it across all lines — could unify both, though
// writing them separately first is a reasonable and clear starting point.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة 3x3 بأرقام عشوائية بين 1 و100، يطبعها، يحسب
// مجموع كل عمود، يجمع تلك المجاميع في شريحة أحادية البُعد منفصلة، وأخيرًا
// يطبع تلك الشريحة.
//
// هذه هي النظير المباشر للمسألة 103، وتتبع نفس الهيكل تمامًا — إعادة
// استخدام colSum بدلاً من rowSum، وإرجاع []int جديدة بدلاً من ملء مصفوفة
// يوفرها المستدعي. إدراك أن المسألتين 103 و105 تتشاركان نفس الهيكل تمامًا
// (فقط دالة الجمع الأساسية تختلف) هو درس مفيد بحد ذاته في غو: فهو يُلمّح
// إلى أن دالة مساعدة أعلى رتبة أكثر عمومية — مثل دالة تأخذ دالة "اجمع خطًا"
// كمعامل وتُطبّقها على كل الخطوط — يمكن أن توحّد الاثنتين، رغم أن كتابتهما
// منفصلتين أولًا نقطة بداية معقولة وواضحة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[10 20 30] [5 5 5] [1 2 3]]
// Output: Col Sums Array: [16 27 38]
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The returned slice length must equal the matrix's column count | يجب أن
//   يتساوى طول الشريحة المُرجَعة مع عدد أعمدة المصفوفة
// • Must reuse colSum internally | يجب إعادة استخدام colSum داخليًا
// • Random values are in the range [1, 100] | القيم العشوائية في المدى [1، 100]
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
// func matrixColSums(matrix [][]int) []int
// func printArray(arr []int)
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
	sum := 0
	for i := range matrix {
		if col >= 0 && col < len(matrix[i]) {
			sum += matrix[i][col]
		}
	}
	return sum
}

func matrixColSums(matrix [][]int) []int {
	colSums := make([]int, 0)
	for i := range matrix {
		colSums = append(colSums, colSum(matrix, i))
	}
	return colSums
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("%d\n", arr)
}

func printMatrix(matrix [][]int) {
	for _, valueI := range matrix {
		for _, valueJ := range valueI {
			fmt.Printf("%d\t", valueJ)
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
	printArray(matrixColSums(matrix))
}
