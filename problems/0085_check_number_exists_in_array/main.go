// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 85: check_number_exists_in_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Searching | المصفوفات والبحث
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills an array with random numbers between 1 and 100,
// prints the array, then asks the user for a number to search for. The program
// should determine whether the number exists in the array or not (a boolean
// check, reusing the position-search logic), and print "Yes" or "No" accordingly.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية بين 1 و 100، يطبع المصفوفة، ثم يسأل
// المستخدم عن رقم للبحث عنه. يجب أن يحدد البرنامج ما إذا كان الرقم موجودًا في
// المصفوفة أم لا (فحص منطقي Boolean يعتمد على منطق البحث عن الموقع)، ويطبع
// "نعم" أو "لا" وفقًا لذلك.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array = [10 20 30], Number to search = 20
// Output: Yes it is found :-)
//
// Example 2:
// Input:  Array = [10 20 30], Number to search = 99
// Output: No, The number is not found :-(
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Array length must be a positive integer (> 0) | طول المصفوفة يجب أن يكون موجبًا
// • Must reuse findNumberPositionInArray internally | يجب إعادة استخدام دالة البحث عن الموقع داخليًا
// • Result must be a boolean value | يجب أن تكون النتيجة قيمة منطقية Boolean
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func fillArrayWithRandomNumbers(arr []int)
// func randomNumber(from, to int) int
// func findNumberPositionInArray(number int, arr []int) int
// func isNumberInArray(number int, arr []int) bool
// func printExistResult(exist bool)
// func printArray(arr []int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// ======================
//   UTILITY
// ======================

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
			printError("Invalid input please enter a valid number")
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

func fillArrayWithRandomNumbers(length int) []int {
	arr := make([]int, length)
	for key := range len(arr) {
		arr[key] = randomNumber(1, 100)
	}
	return arr
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(from, to int) int {
	return rand.IntN(to-from+1) + from
}

func findNumberPositionInArray(number int, arr []int) int {
	for key := range arr {
		if arr[key] == number {
			return key
		}
	}
	return -1
}

func isNumberInArray(number int, arr []int) bool {
	return findNumberPositionInArray(number, arr) != -1
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printExistResult(exist bool) {
	if exist {
		fmt.Println("Yes it is found :-)")
	} else {
		fmt.Println("No, The number is not found :-(")
	}
}

func printArray(arr []int) {
	fmt.Printf("Array = %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter numbers of elements")
	arr := fillArrayWithRandomNumbers(lengthOfArray)
	printArray(arr)

	searchNumber := readPositiveNumber("Please enter number to search in array")

	printExistResult(isNumberInArray(searchNumber, arr))
}
