// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 83: fill_array_with_generated_keys.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Strings & Random | المصفوفات والنصوص والعشوائية
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a positive number from the user representing how
// many license keys to generate. Each key consists of four groups of 4 random
// uppercase letters separated by hyphens, in the format "XXXX-XXXX-XXXX-XXXX".
// The program should store the generated keys inside a string array/slice, then
// print every key along with its index.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم يمثل عدد مفاتيح التفعيل المراد
// إنشاؤها. كل مفتاح يتكون من أربع مجموعات من 4 حروف إنجليزية كبيرة عشوائية
// مفصولة بشرطات، بالصيغة "XXXX-XXXX-XXXX-XXXX". يجب أن يخزن البرنامج المفاتيح
// المولّدة داخل مصفوفة/شريحة نصية، ثم يطبع كل مفتاح مع رقمه التسلسلي.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  How many keys? 2
// Output: Array[0] : ABCD-EFGH-IJKL-MNOP
//         Array[1] : QRST-UVWX-YZAB-CDEF
//
// (Note: Actual output varies due to random generation)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The number of keys must be a positive integer. | يجب أن يكون عدد المفاتيح عددًا صحيحًا موجبًا.
// • Each key is exactly 19 characters long (4+1+4+1+4+1+4). | كل مفتاح طوله بالضبط 19 حرفًا.
// • Only uppercase letters (A-Z) are used in the key. | تُستخدم الأحرف الكبيرة فقط (A-Z) في المفتاح.
// • Each run should produce different keys due to proper random seeding. | يجب أن ينتج كل تشغيل مفاتيح مختلفة بسبب تهيئة عشوائية صحيحة.
// • Use strings.Builder for efficient string concatenation. | استخدم strings.Builder لربط النصوص بكفاءة.
// • Groups are separated by a hyphen '-' | المجموعات مفصولة بشرطة '-'
// • Use an enum/type to represent character categories (letters/digits/special) | استخدم نوعًا مخصصًا لتمثيل فئات الأحرف
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type CharType int
// const (SmallLetter CharType = iota + 1; CapitalLetter; SpecialCharacter; Digit)
// func readPositiveNumber(prompt string) int
// func randomNumber(from, to int) int
// func getRandomCharacter(charType CharType) rune
// func generateWord(charType CharType, length int) string
// func generateKey() string
// func fillArrayWithKeys(arr []string)
// func printStringArray(arr []string)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

// ======================
//   UTILITY
// ======================

func printError(prompt string) {
	fmt.Printf("X Error : %s\n", prompt)
}

type CharType int

const (
	CapitalLetter CharType = iota + 1
	SmallLetter
	SpecialCharacter
	Digit
)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) string {
	fmt.Printf("%s : ", prompt)

	var input string

	fmt.Scan(&input)

	return input
}

func readNumber(prompt string) int {
	for {
		input := readString(prompt)

		inputNumber, err := strconv.Atoi(input)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) int {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

func fillArrayWithKeys(arr []string) {
	for key := range arr {
		arr[key] = generateKey(CapitalLetter, 4)
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func getRandomCharacter(ct CharType) rune {
	switch ct {
	case CapitalLetter:
		return rune(randomNumber(65, 90))

	case SmallLetter:
		return rune(randomNumber(97, 122))

	case SpecialCharacter:
		return rune(randomNumber(33, 47))

	case Digit:
		return rune(randomNumber(48, 57))

	default:
		return 0
	}
}

func generateWord(ct CharType, length int) string {
	var sb strings.Builder

	for range length {
		sb.WriteRune(getRandomCharacter(ct))
	}

	return sb.String()
}

func generateKey(ct CharType, length int) string {
	var sb strings.Builder

	for key := range length {
		if key > 0 {
			sb.WriteString("-")
		}
		sb.WriteString(generateWord(ct, 4))
	}

	return sb.String()
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []string) {
	for key := range arr {
		fmt.Printf("Array[%d] : %s\n", key, arr[key])
	}
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter a numbers of elements")

	arr := make([]string, lengthOfArray)

	fillArrayWithKeys(arr)

	printArray(arr)
}
