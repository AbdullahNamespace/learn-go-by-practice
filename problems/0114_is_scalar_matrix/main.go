// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 114: is_scalar_matrix.go
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
// Write a program that builds a predefined 3x3 scalar matrix (the same
// constant value repeated on the main diagonal, 0s everywhere else),
// prints it, then checks whether it actually qualifies as a scalar matrix,
// and prints the result.
//
// العربية:
// اكتب برنامجًا يبني مصفوفة قياسية (Scalar) محددة مسبقًا بحجم 3x3 (نفس
// القيمة الثابتة مكررة على القطر الرئيسي، وأصفار في كل مكان آخر)، يطبعها،
// ثم يتحقق مما إذا كانت تستحق فعلاً وصف المصفوفة القياسية، ويطبع النتيجة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[9 0 0] [0 9 0] [0 0 9]]
// Output: YES: Matrix is scalar.
//
// Example 2:
// Input:  Matrix = [[9 0 0] [0 5 0] [0 0 9]]
// Output: No: Matrix is NOT scalar.
//   Why:    Matrix[1][1] = 5 but the diagonal value must stay consistent
//           with matrix[0][0] = 9
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The expected scalar value is taken from matrix[0][0], not hardcoded |
//   القيمة القياسية المتوقعة تُؤخذ من matrix[0][0]، وليست ثابتة بالكود
// • All non-diagonal elements must equal 0 | كل العناصر غير القطرية يجب أن تساوي 0
// • Every diagonal element must equal that same scalar value | كل عنصر
//   قطري يجب أن يساوي تلك القيمة القياسية نفسها
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func scalarMatrix() [][]int
// func isScalarMatrix(matrix [][]int) bool
// func printMatrix(matrix [][]int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import "fmt"

// ======================
//     INPUT FUNCTIONS
// ======================

func scalarMatrix() [][]int {
	return [][]int{{9, 0, 0}, {0, 9, 0}, {0, 0, 9}}
}

func notScalarMatrix() [][]int {
	return [][]int{{9, 0, 0}, {0, 1, 0}, {0, 0, 8}}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func isScalarMatrix(matrix [][]int) bool {
	if len(matrix) == 0 {
		return false
	}

	for i := range matrix {
		if len(matrix) != len(matrix[i]) {
			return false
		}
		for j := range matrix[i] {
			expected := 0
			if i == j {
				expected = matrix[0][0]
			}
			if matrix[i][j] != expected {
				return false
			}
		}
	}
	return true
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
	scalarMatrix1 := scalarMatrix()
	notScalarMatrix1 := notScalarMatrix()
	printMatrix(scalarMatrix1, "\n\nscalar Matrix")
	if isScalarMatrix(scalarMatrix1) {
		fmt.Println("YES: Matrix is scalar.")
	} else {
		fmt.Println("No: Matrix is NOT scalar.")
	}

	printMatrix(notScalarMatrix1, "\n\nNot scalar Matrix")
	if isScalarMatrix(notScalarMatrix1) {
		fmt.Println("YES: Matrix is scalar.")
	} else {
		fmt.Println("No: Matrix is NOT scalar.")
	}
}
