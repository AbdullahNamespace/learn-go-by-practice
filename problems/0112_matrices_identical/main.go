// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 112: matrices_identical.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices & Logic | المصفوفات الثنائية والمنطق
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills two separate 3x3 matrices with random numbers
// between 1 and 10, prints both, then checks whether the two matrices are
// truly identical — every corresponding cell holds the exact same value —
// with an early exit as soon as any mismatch is found, and prints the result.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفتين منفصلتين 3x3 بأرقام عشوائية بين 1 و10، يطبع
// كلتيهما، ثم يتحقق مما إذا كانت المصفوفتان متطابقتين حقًا — كل خلية
// متناظرة تحمل نفس القيمة بالضبط — مع خروج مبكر فور العثور على أي عدم
// تطابق، ويطبع النتيجة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix1 = [[1 2 3] [4 5 6] [7 8 9]], Matrix2 = same values
// Output: YES: both matrices are typical.
//
// Example 2:
// Input:  Matrix1 = [[1 2 3] [4 5 6] [7 8 9]], Matrix2 differs at [0][0]
// Output: No: matrices are NOT typical.
//   Why:    The very first cell comparison already fails, so the check
//           exits immediately without scanning the rest
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Both matrices must have the same dimensions | يجب أن تتطابق أبعاد
//   المصفوفتين
// • Comparison must stop at the first mismatch (early exit) | يجب أن يتوقف
//   الفحص عند أول عدم تطابق (خروج مبكر)
// • This performs true element-wise equality, unlike the sum-based
//   approximation in Problem 111 | هذا تساوٍ حقيقي عنصرًا بعنصر، بخلاف
//   التقريب القائم على المجموع في المسألة 111
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
// func areIdenticalMatrices(matrix1, matrix2 [][]int) bool
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

func areIdenticalMatrices(matrix1, matrix2 [][]int) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}
	for i := range matrix1 {
		if len(matrix1[i]) != len(matrix2[i]) {
			return false
		}
		for j := range matrix1[i] {
			if matrix1[i][j] != matrix2[i][j] {
				return false
			}
		}
	}
	return true
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printMatrix(matrix [][]int) {
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
	matrix1 := newMatrix()
	matrix2 := newMatrix()

	fillMatrixWithRandomNumbers(matrix1)
	fillMatrixWithRandomNumbers(matrix2)
	fmt.Println("Matrix 1")
	printMatrix(matrix1)
	fmt.Println("Matrix 2")
	printMatrix(matrix2)
	if areIdenticalMatrices(matrix1, matrix2) {
		fmt.Println("YES: both matrices are equal.")
	} else {
		fmt.Println("NO: both matrices are not equal.")
	}
}
