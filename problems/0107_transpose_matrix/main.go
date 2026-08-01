// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 107: transpose_matrix.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices | المصفوفات الثنائية
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that creates a 3x3 matrix filled with sequential numbers
// from 1 to 9, prints it, computes its transpose (where element [row][col]
// of the source becomes element [col][row] of the result), and prints the
// transposed matrix.
//
// العربية:
// اكتب برنامجًا يُنشئ مصفوفة 3x3 مملوءة بأرقام متسلسلة من 1 إلى 9، يطبعها،
// يحسب مقلوبها (Transpose) (حيث يصبح العنصر [row][col] في المصدر هو العنصر
// [col][row] في النتيجة)، ويطبع المصفوفة المقلوبة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[1 2 3] [4 5 6] [7 8 9]]
// Output: Original matrix:
//         1 2 3
//         4 5 6
//         7 8 9
//
//         Transposed matrix:
//         1 4 7
//         2 5 8
//         3 6 9
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The source matrix must remain unmodified after calling transpose |
//   يجب أن تبقى المصفوفة المصدرية دون تعديل بعد استدعاء القلب
// • The result must be a newly allocated matrix, not an alias of the
//   source | يجب أن تكون النتيجة مصفوفة مُخصَّصة حديثًا وليست اسمًا بديلًا
//   للمصدر
// • Works for any square matrix (rows == cols); non-square is out of
//   scope here | يعمل مع أي مصفوفة مربعة (rows == cols)؛ غير المربعة خارج
//   نطاق هذه المسألة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// const MatrixSize = 3
// func newMatrix() [][]int
// func fillMatrixWithOrderedNumbers(matrix [][]int)
// func printMatrix(matrix [][]int)
// func transposeMatrix(matrix [][]int) [][]int
//
// ────────────────────────────────────────────────────────────────────────────
package main

import "fmt"

// ======================
//   UTILITY
// ======================

const MatrixSize = 3

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

func fillMatrixWithOrderedNumbers(matrix [][]int) {
	counter := 1
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = counter
			counter++
		}
	}
}

func isSquareMatrix(matrix [][]int) bool {
	length := len(matrix)
	for i := range matrix {
		if len(matrix[i]) != length {
			return false
		}
	}
	return true
}

func transposeMatrix(matrix [][]int) [][]int {
	transpose := newMatrix()
	if isSquareMatrix(matrix) {
		for i := range matrix {
			for j := range matrix[i] {
				transpose[i][j] = matrix[j][i]
			}
		}
	} else {
		fmt.Println("non-square is out of scope here we return empty matrix")
	}
	return transpose
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

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
	fillMatrixWithOrderedNumbers(matrix)
	printMatrix(matrix)
	transpose := transposeMatrix(matrix)
	fmt.Println("=========")
	printMatrix(transpose)
}
