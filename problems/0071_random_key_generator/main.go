// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 71: random_key_generator.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Strings & Random | النصوص والأرقام العشوائية
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a positive number from the user representing how many
// keys to generate. Each key consists of four groups of 4 uppercase letters separated
// by hyphens, in the format "XXXX-XXXX-XXXX-XXXX". The program uses random character
// generation based on ASCII values to produce each letter, then assembles and prints
// all requested keys with their sequence numbers.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم يمثل عدد المفاتيح المراد إنشاؤها.
// كل مفتاح يتكون من أربع مجموعات، كل مجموعة 4 حروف إنجليزية كبيرة، مفصولة
// بخطوط أفقية بالصيغة "XXXX-XXXX-XXXX-XXXX". يستخدم البرنامج توليد حروف
// عشوائية بناءً على قيم ASCII لتجميع كل مفتاح ثم يطبع جميع المفاتيح المطلوبة
// مع أرقامها التسلسلية.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  3
// Output:
// Key [1] : ABCD-EFGH-IJKL-MNOP
// Key [2] : QRST-UVWX-YZAB-CDEF
// Key [3] : GHIJ-KLMN-OPQR-STUV
//
// (Note: Actual output varies due to random generation)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input must be a positive integer (> 0). | المدخل يجب أن يكون عددًا صحيحًا موجبًا (> 0).
// • Each key group is exactly 4 uppercase letters (ASCII 65-90). | كل مجموعة في المفتاح تتكون بالضبط من 4 حروف إنجليزية كبيرة (ASCII 65-90).
// • Groups are separated by hyphens '-'. | المجموعات مفصولة بخط أفقي '-'.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) int
// func randomNumber(from, to int) int
// func getRandomCharacter(charType int) rune
// func generateWord(charType int, length int) string
// func generateKey() string
// func generateKeys(numberOfKeys int)
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

type CharType int

const (
	CapitalLetter = iota + 1
	SmallLetter
	SpecialCharacter
	Digit
)

func printError(prompt string) {
	fmt.Printf("X Error : %s\n", prompt)
}

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

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumberInRange(from, to int) int {
	return rand.IntN(to-from+1) + from
}

func getRandomCharacter(ct CharType) rune {
	switch ct {

	case CapitalLetter:
		return rune(randomNumberInRange(65, 90))

	case SmallLetter:
		return rune(randomNumberInRange(97, 122))

	case SpecialCharacter:
		return rune(randomNumberInRange(33, 47))

	case Digit:
		return rune(randomNumberInRange(48, 57))

	default:
		return 0

	}
}

func generateWord(ct CharType, length int) string {
	var sb strings.Builder

	for i := 0; i < length; i++ {
		sb.WriteRune(getRandomCharacter(ct))
	}

	return sb.String()
}

func generateKey(ct CharType, length int) string {
	var sb strings.Builder

	for i := range length {
		if i > 0 {
			sb.WriteString("-")
		}
		sb.WriteString(generateWord(ct, length))
	}

	return sb.String()
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printKeys(numberOfKeys int) {
	for i := 0; i < numberOfKeys; i++ {
		fmt.Printf("Key [%d] : %s\n", i+1, generateKey(CapitalLetter, 4))
	}
}

// ======================
//         MAIN
// ======================

func main() {
	numberOfKeys := readPositiveNumber("Please enter a number of keys")
	printKeys(numberOfKeys)
}
