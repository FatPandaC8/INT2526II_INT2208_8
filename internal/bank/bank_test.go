package bank

import "testing"

func TestClassifyRisk(t *testing.T) {

	tests := []struct {
		name         string
		creditScore  int
		expectedRisk string
	}{
		{
			name:         "300 => High",
			creditScore:  300,
			expectedRisk: "High",
		},
		{
			name:         "500 => High",
			creditScore:  500,
			expectedRisk: "High",
		},
		{
			name:         "501 => Medium",
			creditScore:  501,
			expectedRisk: "Medium",
		},
		{
			name:         "700 => Medium",
			creditScore:  700,
			expectedRisk: "Medium",
		},
		{
			name:         "701 => Low",
			creditScore:  701,
			expectedRisk: "Low",
		},
		{
			name:         "850 => Low",
			creditScore:  850,
			expectedRisk: "Low",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := ClassifyRisk(tt.creditScore)

			if result != tt.expectedRisk {

				t.Errorf(
					"expected %s, got %s",
					tt.expectedRisk,
					result,
				)
			}
		})
	}
}

func TestLoanDecision(t *testing.T) {

	tests := []struct {
		name        string
		age         int
		income      float64
		creditScore int
		employment  string
		expected    string
	}{

		// Validation Test Cases

		{
			name:        "TC01 - age < 18",
			age:         17,
			income:      100.0,
			creditScore: 650,
			employment:  "C",
			expected:    "Invalid Input",
		},
		{
			name:        "TC02 - age = 18",
			age:         18,
			income:      100.0,
			creditScore: 650,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC03 - age = 19",
			age:         19,
			income:      100.0,
			creditScore: 650,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC04 - age = 64",
			age:         64,
			income:      100.0,
			creditScore: 650,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC05 - age = 65",
			age:         65,
			income:      100.0,
			creditScore: 650,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC06 - age > 65",
			age:         66,
			income:      100.0,
			creditScore: 650,
			employment:  "C",
			expected:    "Invalid Input",
		},

		// income boundaries

		{
			name:        "TC07 - income < 5.0",
			age:         30,
			income:      4.9,
			creditScore: 650,
			employment:  "C",
			expected:    "Invalid Input",
		},
		{
			name:        "TC08 - income = 5.0",
			age:         30,
			income:      5.0,
			creditScore: 650,
			employment:  "C",
			expected:    "REJECT",
		},
		{
			name:        "TC09 - income = 5.1",
			age:         30,
			income:      5.1,
			creditScore: 650,
			employment:  "C",
			expected:    "REJECT",
		},
		{
			name:        "TC10 - income = 499.9",
			age:         30,
			income:      499.9,
			creditScore: 650,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC11 - income = 500.0",
			age:         30,
			income:      500.0,
			creditScore: 650,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC12 - income > 500.0",
			age:         30,
			income:      500.1,
			creditScore: 650,
			employment:  "C",
			expected:    "Invalid Input",
		},

		// credit score boundaries

		{
			name:        "TC13 - credit score < 300",
			age:         30,
			income:      100.0,
			creditScore: 299,
			employment:  "C",
			expected:    "Invalid Input",
		},
		{
			name:        "TC14 - credit score = 300",
			age:         30,
			income:      100.0,
			creditScore: 300,
			employment:  "C",
			expected:    "REJECT",
		},
		{
			name:        "TC15 - credit score = 301",
			age:         30,
			income:      100.0,
			creditScore: 301,
			employment:  "C",
			expected:    "REJECT",
		},
		{
			name:        "TC16 - credit score = 849",
			age:         30,
			income:      100.0,
			creditScore: 849,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC17 - credit score = 850",
			age:         30,
			income:      100.0,
			creditScore: 850,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC18 - credit score > 850",
			age:         30,
			income:      100.0,
			creditScore: 851,
			employment:  "C",
			expected:    "Invalid Input",
		},

		// employment validation

		{
			name:        "TC19 - employment = C",
			age:         30,
			income:      100.0,
			creditScore: 650,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC20 - employment = F",
			age:         30,
			income:      100.0,
			creditScore: 650,
			employment:  "F",
			expected:    "MANUAL REVIEW",
		},
		{
			name:        "TC21 - invalid employment",
			age:         30,
			income:      100.0,
			creditScore: 650,
			employment:  "X",
			expected:    "Invalid Input",
		},

		// Business Logic Test Cases

		{
			name:        "TC22 - High Risk",
			age:         30,
			income:      100.0,
			creditScore: 450,
			employment:  "C",
			expected:    "REJECT",
		},
		{
			name:        "TC23 - Medium Risk and income < 15",
			age:         30,
			income:      10.0,
			creditScore: 650,
			employment:  "C",
			expected:    "REJECT",
		},
		{
			name:        "TC24 - Low Risk + Contract + income < 15",
			age:         30,
			income:      10.0,
			creditScore: 750,
			employment:  "C",
			expected:    "MANUAL REVIEW",
		},
		{
			name:        "TC25 - Freelance and income < 15",
			age:         30,
			income:      10.0,
			creditScore: 750,
			employment:  "F",
			expected:    "REJECT",
		},
		{
			name:        "TC26 - income >= 15 and Contract",
			age:         30,
			income:      20.0,
			creditScore: 650,
			employment:  "C",
			expected:    "APPROVE",
		},
		{
			name:        "TC27 - income >= 15 and Freelance",
			age:         30,
			income:      20.0,
			creditScore: 650,
			employment:  "F",
			expected:    "MANUAL REVIEW",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := LoanDecision(
				tt.age,
				tt.income,
				tt.creditScore,
				tt.employment,
			)

			if result != tt.expected {

				t.Errorf(
					"expected %s, got %s",
					tt.expected,
					result,
				)
			}
		})
	}
}

// Combinatory
func TestExhaustiveInvalidInputs(t *testing.T) {

	ages := []int{
		17,
		18,
		30,
		65,
		66,
	}

	incomes := []float64{
		4.9,
		5.0,
		15.0,
		500.0,
		500.1,
	}

	scores := []int{
		299,
		300,
		650,
		850,
		851,
	}

	employments := []string{
		"C",
		"F",
		"X",
	}

	for _, age := range ages {

		for _, income := range incomes {

			for _, score := range scores {

				for _, emp := range employments {

					result := LoanDecision(
						age,
						income,
						score,
						emp,
					)

					isValid := IsValidInput(
						age,
						income,
						score,
						emp,
					)

					if !isValid &&
						result != "Invalid Input" {

						t.Errorf(
							"invalid input should return Invalid Input: age=%d income=%f score=%d emp=%s got=%s",
							age,
							income,
							score,
							emp,
							result,
						)
					}
				}
			}
		}
	}
}