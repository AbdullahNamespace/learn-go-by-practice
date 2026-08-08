// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 113: is_identity_matrix.go
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
// Write a program that builds a predefined 3x3 identity matrix (1s on the
// main diagonal, 0s everywhere else), prints it, then checks whether it
// actually qualifies as an identity matrix — every diagonal element (where
// row == col) must equal 1, and every non-diagonal element must equal 0 —
// and prints the result.
//
// العربية:
// اكتب برنامجًا يبني مصفوفة وحدة (Identity) محددة مسبقًا بحجم 3x3 (آحاد على
// القطر الرئيسي، وأصفار في كل مكان آخر)، يطبعها، ثم يتحقق مما إذا كانت
// تستحق فعلاً وصف مصفوفة الوحدة — كل عنصر قطري (حيث row == col) يجب أن
// يساوي 1، وكل عنصر غير قطري يجب أن يساوي 0 — ويطبع النتيجة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[1 0 0] [0 1 0] [0 0 1]]
// Output: YES: Matrix is identity.
//
// Example 2:
// Input:  Matrix = [[1 2 3] [4 5 6] [7 8 9]]
// Output: No: Matrix is NOT identity.
//   Why:    Matrix[0][1] = 2 but a non-diagonal cell must be 0
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Diagonal elements (row == col) must all equal 1 | العناصر القطرية
//   (row == col) يجب أن تساوي جميعها 1
// • All non-diagonal elements must equal 0 | كل العناصر غير القطرية يجب أن
//   تساوي 0
// • Must return false immediately on the first violation (early exit) |
//   يجب إرجاع false فورًا عند أول مخالفة (خروج مبكر)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func identityMatrix() [][]int
// func isIdentityMatrix(matrix [][]int) bool
// func printMatrix(matrix [][]int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import "fmt"

// ======================
//     INPUT FUNCTIONS
// ======================

func identityMatrix() [][]int {
	return [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
}

func notIdentityMatrix() [][]int {
	return [][]int{{1, 0, 0}, {0, 9, 0}, {0, 0, 0}}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func isIdentityMatrix(matrix [][]int) bool {
	for i := range matrix {
		if len(matrix) != len(matrix[i]) {
			return false
		}
		for j := range matrix[i] {
			expected := 0
			if i == j {
				expected = 1
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
	identityMatrix1 := identityMatrix()
	notIdentityMatrix1 := notIdentityMatrix()
	printMatrix(identityMatrix1, "\n\nidentity Matrix")
	if isIdentityMatrix(identityMatrix1) {
		fmt.Println("YES: Matrix is identity.")
	} else {
		fmt.Println("No: Matrix is NOT identity.")
	}

	printMatrix(notIdentityMatrix1, "\n\nNot Identity Matrix")
	if isIdentityMatrix(notIdentityMatrix1) {
		fmt.Println("YES: Matrix is identity.")
	} else {
		fmt.Println("No: Matrix is NOT identity.")
	}
}
