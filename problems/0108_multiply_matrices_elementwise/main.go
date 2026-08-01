// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 108: multiply_matrices_elementwise.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices & Math | المصفوفات الثنائية والرياضيات
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills two separate 3x3 matrices with random numbers
// between 1 and 10, prints both, then builds a third matrix where each cell
// is the product of the corresponding cells in the two source matrices
// (result[i][j] = matrix1[i][j] * matrix2[i][j]), and finally prints that
// result matrix.
//
// This operation is element-wise multiplication (also called the Hadamard
// product), NOT the row-by-column dot-product multiplication taught in
// linear algebra. The naming is worth being deliberate about: a function
// called multiplyMatricesElementWise makes the operation unambiguous,
// whereas a generic name like multiplyMatrix could easily be mistaken for
// true matrix multiplication by a future reader, which is a completely
// different algorithm with different dimension-compatibility rules.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفتين منفصلتين 3x3 بأرقام عشوائية بين 1 و10، يطبع
// كلتيهما، ثم يبني مصفوفة ثالثة تكون فيها كل خلية ناتج ضرب الخليتين
// المتناظرتين في المصفوفتين المصدريتين (result[i][j] = matrix1[i][j] *
// matrix2[i][j])، وأخيرًا يطبع مصفوفة النتيجة تلك.
//
// هذه العملية هي الضرب عنصرًا بعنصر (يُسمى أيضًا حاصل ضرب هادامارد)، وليست
// ضرب النقطة صفًا بعمود المُدرَّس في الجبر الخطي. تسمية الدالة تستحق عناية
// متعمدة: دالة باسم multiplyMatricesElementWise تجعل العملية لا لبس فيها،
// بينما اسم عام مثل multiplyMatrix يمكن أن يُخطئ به قارئ مستقبلي ويظنه ضرب
// مصفوفات حقيقيًا، وهو خوارزمية مختلفة تمامًا بقواعد توافق أبعاد مختلفة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix1 = [[2 3 4] [5 6 7] [8 9 1]]
//         Matrix2 = [[1 2 3] [4 5 6] [7 8 9]]
// Output: Results:
//         2 6 12
//         20 30 42
//         56 72 9
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random values are in the range [1, 10] for both matrices | القيم
//   العشوائية في المدى [1، 10] لكلتا المصفوفتين
// • Both source matrices must have identical dimensions | يجب أن تتطابق
//   أبعاد المصفوفتين المصدريتين
// • This is element-wise multiplication, not linear-algebra matrix
//   multiplication | هذا ضرب عنصر بعنصر، وليس ضرب مصفوفات بالمفهوم الجبري
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
// func multiplyMatricesElementWise(matrix1, matrix2 [][]int) [][]int
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

func multiplyMatricesElementWise(matrix1, matrix2 [][]int) [][]int {
	multiply := newMatrix()
	for i := range matrix1 {
		for j := range matrix1[i] {
			multiply[i][j] = matrix1[i][j] * matrix2[i][j]
		}
	}
	return multiply
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
	matrix1 := newMatrix()
	matrix2 := newMatrix()
	fillMatrixWithRandomNumbers(matrix1)
	fillMatrixWithRandomNumbers(matrix2)
	printMatrix(matrix1)
	printMatrix(matrix2)
	multiply := multiplyMatricesElementWise(matrix1, matrix2)
	printMatrix(multiply)
}
