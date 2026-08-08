// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 111: matrices_equal_by_sum.go
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
// Write a program that fills two separate 3x3 matrices with random numbers
// between 1 and 10, prints both, then determines whether the two matrices
// are "equal" using a sum-based comparison — comparing the total sum of
// Matrix1's elements against the total sum of Matrix2's elements — and
// prints the result.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفتين منفصلتين 3x3 بأرقام عشوائية بين 1 و10، يطبع
// كلتيهما، ثم يحدد ما إذا كانت المصفوفتان "متساويتان" باستخدام مقارنة
// قائمة على المجموع — مقارنة إجمالي مجموع عناصر Matrix1 مع إجمالي مجموع
// عناصر Matrix2 — ويطبع النتيجة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix1 sum = 45, Matrix2 sum = 45
// Output: YES: both matrices are equal (by sum).
//
// Example 2 (Misleading Case):
// Input:  Matrix1 = [[1 9 ...]] sum = 45, Matrix2 = [[5 5 ...]] sum = 45
// Output: YES: both matrices are equal (by sum).
//   Why:    Same total, even though the individual values differ completely —
//           this demonstrates exactly why sum-based equality is weak
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random values are in the range [1, 10] for both matrices | القيم
//   العشوائية في المدى [1، 10] لكلتا المصفوفتين
// • Must reuse matrixSum internally, not reimplement summation | يجب إعادة
//   استخدام matrixSum داخليًا وعدم إعادة كتابة منطق الجمع
// • This is a sum-based equality check, NOT true element-wise equality (see
//   Problem 112 for that) | هذا فحص تساوٍ قائم على المجموع، وليس تساويًا
//   حقيقيًا عنصرًا بعنصر (راجع المسألة 112 لذلك)
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
// func printMatrix(matrix [][]int)
// func matrixSum(matrix [][]int) int
// func areEqualBySum(matrix1, matrix2 [][]int) bool
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

func fillMatrixWithRandomNumbers(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = randomNumber(1, 10)
		}
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func matrixSum(matrix [][]int) int {
	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
	}
	return sum
}

func areEqualBySum(matrix1, matrix2 [][]int) bool {
	return matrixSum(matrix1) == matrixSum(matrix2)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printMatrix(matrix [][]int) {
	for i := range matrix {
		for _, value := range matrix[i] {
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
}

// ======================
//         MAIN
// ======================

func main() {
	matrix1 := newMatrix()
	matrix2 := newMatrix()
	fillMatrixWithRandomNumbers(matrix1)
	fillMatrixWithRandomNumbers(matrix2)
	fmt.Println("Matrix 1")
	printMatrix(matrix1)
	fmt.Println("Matrix 2")
	printMatrix(matrix2)
	if areEqualBySum(matrix1, matrix2) {
		fmt.Println("YES: both matrices are equal (by sum).")
	} else {
		fmt.Println("NO: both matrices are not equal (by sum).")
	}
}
