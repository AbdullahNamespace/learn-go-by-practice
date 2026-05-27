// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 4: hiring_without_recommendation.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Conditional Logic | المنطق الشرطي
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads candidate information (age and driving license status)
// and determines if they are hired or rejected for a driver position.
// A candidate is hired only if they are older than 21 AND have a driving license.
// The program uses a struct to organize candidate data and separates concerns
// into different functions (input, validation, output).
//
// العربية:
// اكتب برنامجًا يقرأ معلومات المرشح (العمر وحالة رخصة القيادة)
// ويحدد ما إذا كان مقبولًا أو مرفوضًا لوظيفة سائق.
// يتم قبول المرشح فقط إذا كان عمره أكبر من 21 سنة ولديه رخصة قيادة.
// يستخدم البرنامج struct لتنظيم بيانات المرشح ويفصل المسؤوليات
// إلى دوال مختلفة (الإدخال، التحقق، الإخراج).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  age = 25, has_driving_license = 1 (Yes)
// Output: Hired
//
// Example 2:
// Input:  age = 20, has_driving_license = 1 (Yes)
// Output: Rejected
//
// Example 3 (Edge Case):
// Input:  age = 22, has_driving_license = 0 (No)
// Output: Rejected
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Age must be a valid positive integer (uint8) | يجب أن يكون العمر عددًا صحيحًا موجبًا
// • Driving license input must be 1 (Yes) or 0 (No) | يجب أن يكون إدخال الرخصة 1 (نعم) أو 0 (لا)
// • Candidate must meet BOTH conditions to be hired | يجب أن يستوفي المرشح كلا الشرطين ليتم قبوله
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type CandidateInfo struct { Age uint8; HasDrivingLicense bool }
// func readInfo() CandidateInfo
// func isAccepted(info CandidateInfo) bool
// func printResult(info CandidateInfo)
//
// ────────────────────────────────────────────────────────────────────────────

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CandidateInfo struct {
	Age               uint8
	HasDrivingLicense bool
}

var reader = bufio.NewReader(os.Stdin)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readBool(prompt string) bool {
	for {
		input, err := readString(prompt)
		if err != nil {
			fmt.Println("Invalid Input Please Enter Again!")
			continue
		}
		input = strings.ToUpper(input)
		switch input {
		case "Y", "YES", "T", "TRUE", "1":
			{
				return true
			}
		case "N", "NO", "F", "FALSE", "0":
			{
				return false
			}

		default:
			{
				fmt.Println("Please enter 'y' for yes or 'n' for no.")
			}
		}
	}
}

func readNumber(prompt string) int {
	for {
		input, err := readString(prompt)

		if err != nil {
			fmt.Println("invalid input Please enter again!")
			continue
		}
		inputNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input Please enter a number!")
			continue
		}
		return inputNumber
	}
}

func NewCandidateInfo() *CandidateInfo {
	return &CandidateInfo{
		Age:               uint8(readNumber("Please Enter Your Age: ")),
		HasDrivingLicense: readBool("Do You have driver license enter? (y,n)"),
	}
}

// ======================
//
//	PROCESSING FUNCTIONS
//
// ======================
func (p *CandidateInfo) isAccepted() bool {
	return p.Age > 21 && p.HasDrivingLicense
}

// ======================
//
//	OUTPUT FUNCTIONS
//
// ======================
func (p *CandidateInfo) printResult() {
	if p.isAccepted() {
		fmt.Printf("\nHired\n")
	} else {
		fmt.Printf("\nRejected\n")
	}
}

// ======================
//         MAIN
// ======================

func main() {
	person := NewCandidateInfo()
	person.printResult()
}
