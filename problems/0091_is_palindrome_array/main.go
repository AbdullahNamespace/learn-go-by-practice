// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 91: is_palindrome_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Logic | المصفوفات والمنطق
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills an array with a predefined set of values that form
// a palindrome (e.g., 10, 20, 30, 30, 20, 10), prints the array, then checks
// whether the array is a palindrome — meaning it reads the same forward and
// backward — and prints the result.
//
// The palindrome check works by comparing the element at position i with its
// mirrored counterpart at position (length - 1 - i), for i ranging from 0 up
// to the middle of the array. If any pair does not match, the array is not a
// palindrome; the function can return false immediately without checking the
// rest (early exit). If the loop completes without any mismatch, the array is
// a palindrome.
//
// In Go, this function should operate directly on a slice ([]int) rather than
// a fixed-capacity array, since slices carry their own length via len(arr) and
// there is no need to pass a separate length parameter as in C++. The loop
// only needs to iterate up to len(arr)/2 for efficiency, since checking beyond
// the midpoint re-compares pairs that have already been checked — although a
// full-length loop (like the original C++ version) also works correctly, just
// redundantly. Both approaches are acceptable, but the half-length loop is the
// more idiomatic and efficient Go approach.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بمجموعة قيم محددة مسبقًا تُشكّل عددًا متناظرًا
// (Palindrome) (مثل: 10, 20, 30, 30, 20, 10)، يطبع المصفوفة، ثم يتحقق مما
// إذا كانت المصفوفة متناظرة — أي تُقرأ بنفس الطريقة من الأمام والخلف —
// ويطبع النتيجة.
//
// يعمل فحص التناظر عن طريق مقارنة العنصر في الموضع i بنظيره المعكوس في
// الموضع (الطول - 1 - i)، لقيم i من 0 وحتى منتصف المصفوفة. إذا لم يتطابق
// أي زوج، فالمصفوفة ليست متناظرة؛ يمكن للدالة إرجاع false فورًا دون التحقق
// من الباقي (خروج مبكر). إذا اكتملت الحلقة دون أي عدم تطابق، فالمصفوفة
// متناظرة.
//
// في لغة غو، يجب أن تعمل هذه الدالة مباشرة على شريحة ([]int) بدلاً من
// مصفوفة بسعة ثابتة، لأن الشرائح تحمل طولها الخاص عبر len(arr) ولا حاجة
// لتمرير معامل طول منفصل كما في C++. تحتاج الحلقة فقط للمرور حتى
// len(arr)/2 لتحقيق الكفاءة، لأن التحقق بعد منتصف المصفوفة يعيد مقارنة
// أزواج تم فحصها بالفعل — رغم أن الحلقة الكاملة (كما في نسخة C++ الأصلية)
// تعمل أيضًا بشكل صحيح، لكنها زائدة عن الحاجة. كلا النهجين مقبولان، لكن
// الحلقة نصف الطول هي النهج الأكثر كفاءة واعتيادًا في غو.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array = [10, 20, 30, 30, 20, 10]
// Output: Array Elements: 10 20 30 30 20 10
//         Yes array is Palindrome
//   Why:    The array reads the same forward and backward
//
// Example 2:
// Input:  Array = [1, 2, 3, 4]
// Output: Array Elements: 1 2 3 4
//         NO array is NOT Palindrome
//   Why:    arr[0]=1 does not equal arr[3]=4
//
// Example 3 (Edge Case - Odd Length):
// Input:  Array = [5, 8, 5]
// Output: Array Elements: 5 8 5
//         Yes array is Palindrome
//   Why:    The middle element (8) doesn't need a pair; only the outer pair (5, 5) is compared
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The array can contain any set of integers (fixed or random). | يمكن أن تحتوي المصفوفة على أي مجموعة من الأعداد الصحيحة (ثابتة أو عشوائية).
// • The check should stop as soon as a mismatch is found (early exit). | يجب أن يتوقف الفحص فور العثور على عدم تطابق (خروج مبكر).
// • An array of length 0 or 1 is trivially a palindrome. | مصفوفة بطول 0 أو 1 تُعتبر متناظرة بشكل بديهي.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func fillArray() []int
// func printArray(arr []int)
// func isPalindromeArray(arr []int) bool
//
// ────────────────────────────────────────────────────────────────────────────
package main

import "fmt"

// ======================
//   PROCESSING FUNCTIONS
// ======================

func fillArray() []int {
	return []int{8, 5, 8}
}

func isPalindromeArray(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(arr []int) {
	if isPalindromeArray(arr) {
		fmt.Println("Yes array is Palindrome")
	} else {
		fmt.Println("NO array is NOT Palindrome")
	}
}

func printArray(arr []int) {
	fmt.Printf("Array is %v\n", arr)
}

// ======================
//	MAIN
// ======================

func main() {
	arr := fillArray()
	printArray(arr)
	printResult(arr)
}
