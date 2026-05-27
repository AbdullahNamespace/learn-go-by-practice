// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 5: driver_hiring_with_recommendation.go
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
// Write a program that reads candidate information (age, driving license status,
// and recommendation status) and determines if they are hired or rejected for
// a driver position. A candidate is hired if:
// - They are older than 21 AND have a driving license, OR
// - They have a recommendation (regardless of age or license)
// The program includes robust input validation with error handling and accepts
// multiple input formats for boolean values (1/0, yes/no, y/n).
//
// العربية:
// اكتب برنامجًا يقرأ معلومات المرشح (العمر، حالة رخصة القيادة، وحالة التوصية)
// ويحدد ما إذا كان مقبولًا أو مرفوضًا لوظيفة سائق. يتم قبول المرشح إذا:
// - كان عمره أكبر من 21 سنة ولديه رخصة قيادة، أو
// - لديه توصية (بغض النظر عن العمر أو الرخصة)
// يتضمن البرنامج التحقق القوي من المدخلات مع معالجة الأخطاء ويقبل
// تنسيقات متعددة للقيم المنطقية (1/0، نعم/لا، y/n).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  age = 25, has_driving_license = 1, has_recommendation = 0
// Output: Hired
//
// Example 2:
// Input:  age = 18, has_driving_license = 0, has_recommendation = 1
// Output: Hired
//
// Example 3 (Edge Case):
// Input:  age = 20, has_driving_license = 1, has_recommendation = 0
// Output: Rejected
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Age must be between 1 and 120 | يجب أن يكون العمر بين 1 و 120
// • Boolean inputs accept: 1/0, yes/no, y/n (case insensitive) | المدخلات المنطقية تقبل: 1/0، نعم/لا، y/n
// • Invalid inputs trigger re-prompt until valid input is provided | المدخلات غير الصحيحة تطلب الإدخال مجددًا
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type CandidateInfo struct { Age uint8; HasDrivingLicense bool; HasRecommendation bool }
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

var reader = bufio.NewReader(os.Stdin)

// ======================
//     INPUT FUNCTIONS
// ======================

type CandidateInfo struct {
	Age               uint8
	HasDrivingLicense bool
	HasRecommendation bool
}

func readString(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readNumber(prompt string) int {
	for {
		input, err := readString(prompt)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		return value
	}
}

func readBoolean(prompt string) bool {
	for {
		input, err := readString(fmt.Sprintf("%s [y/n]: ", prompt))
		if err != nil {
			fmt.Println("Invalid input. Please enter again.")
			continue
		}
		input = strings.ToLower(input)

		switch input {
		case "y", "yes", "1", "t", "true":
			{
				return true
			}
		case "n", "no", "0", "f", "false":
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

func readAge() uint8 {
	for {
		input := readNumber("Please enter Candidate's age (1-120): ")
		if input > 0 && input <= 120 {
			return uint8(input)
		}
		fmt.Println("Age must be between 1 and 120")
	}
}

func readInfo() CandidateInfo {
	age := readAge()
	hasLicense := readBoolean("Does candidate have driving license")
	hasRec := readBoolean("Does candidate have recommendation")
	return CandidateInfo{
		Age:               age,
		HasDrivingLicense: hasLicense,
		HasRecommendation: hasRec,
	}
}

// ======================
//
//	PROCESSING FUNCTIONS
//
// ======================
func (c *CandidateInfo) isAccepted() bool {
	return c.Age >= 21 && c.HasDrivingLicense || c.HasRecommendation
}

// ======================
//
//	OUTPUT FUNCTIONS
//
// ======================
func (c *CandidateInfo) printResult() {
	if c.isAccepted() {
		fmt.Println("accepts")
	} else {
		fmt.Println("Rejected")
	}
}

// ======================
//         MAIN
// ======================

func main() {
	c := readInfo()
	c.printResult()
}
