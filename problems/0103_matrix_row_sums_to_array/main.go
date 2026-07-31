// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 103: matrix_row_sums_to_array.go
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
// 100, prints it, then computes the sum of each row and collects those sums
// into a separate 1D slice (instead of printing them immediately), and
// finally prints that slice of row sums.
//
// This builds directly on Problem 102 by reusing rowSum. Rather than
// mutating a caller-provided output slice by reference, the idiomatic Go
// approach is for the collecting function to allocate and return a new
// []int itself. Returning the result is generally preferred in Go over
// mutating an out-parameter, since it keeps the function's inputs read-only
// and its effect visible directly in the return value.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة 3x3 بأرقام عشوائية بين 1 و100، يطبعها، ثم يحسب
// مجموع كل صف ويجمع تلك المجاميع في شريحة أحادية البُعد منفصلة (بدلاً من
// طباعتها فورًا)، وأخيرًا يطبع شريحة مجاميع الصفوف تلك.
//
// يبني هذا مباشرة على المسألة 102 بإعادة استخدام rowSum. بدلاً من تعديل
// شريحة إخراج يوفرها المستدعي بالمرجع، فإن النهج الاعتيادي في غو هو أن
// تُخصص الدالة الجامعة شريحة []int جديدة وتُرجعها بنفسها. إرجاع النتيجة
// عمومًا أُفضّل في غو من تعديل معامل إخراج، لأنه يُبقي مدخلات الدالة للقراءة
// فقط ويجعل أثرها ظاهرًا مباشرة في القيمة المُرجَعة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[10 20 30] [5 5 5] [1 2 3]]
// Output: [60 15 6]
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The returned slice length must equal the matrix's row count | يجب أن
//   يتساوى طول الشريحة المُرجَعة مع عدد صفوف المصفوفة
// • Must reuse rowSum internally, not reimplement the summing logic | يجب
//   إعادة استخدام rowSum داخليًا وعدم إعادة كتابة منطق الجمع
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
// func rowSum(matrix [][]int, row int) int
// func matrixRowSums(matrix [][]int) []int
// func printMatrix(matrix [][]int)
// func printArray(arr []int)
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

func rowSum(matrix [][]int, row int) int {
	sum := 0
	for _, value := range matrix[row] {
		sum += value
	}
	return sum
}

func matrixRowSums(matrix [][]int) []int {
	sumMatrix := make([]int, 0, len(matrix))
	for i := range matrix {
		sumMatrix = append(sumMatrix, rowSum(matrix, i))
	}
	return sumMatrix
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

func printArray(arr []int) {
	fmt.Printf("\n%d\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	matrix := newMatrix()
	fillMatrixWithRandomNumbers(matrix)
	printMatrix(matrix)
	printArray(matrixRowSums(matrix))
}
