package bank

const (
	Invalid = "Invalid Input"
	Reject = "REJECT"
	Manual = "MANUAL REVIEW"
	Approve = "APPROVE"

	High = "High"
	Medium = "Medium"
	Low = "Low"
)

func ClassifyRisk(creditScore int) string {

	if creditScore >= 300 && creditScore <= 500 {
		return High
	}

	if creditScore >= 501 && creditScore <= 700 {
		return Medium
	}

	return Low
}

func IsValidInput(
	age int,
	income float64,
	creditScore int,
	employment string,
) bool {

	if age < 18 || age > 65 {
		return false
	}

	if income < 5.0 || income > 500.0 {
		return false
	}

	if creditScore < 300 || creditScore > 850 {
		return false
	}

	if employment != "C" && employment != "F" {
		return false
	}

	return true
}

func LoanDecision(age int, income float64, creditScore int, employment string) string {
	if !IsValidInput(
		age,
		income,
		creditScore,
		employment,
	) {
		return Invalid
	}

	risk := ClassifyRisk(creditScore)

	if risk == High {
		return Reject
	}

	if income < 15.0 {
		if risk == Medium {
			return Reject
		}

		if employment == "F" {
			return Reject
		}

		return Manual
	}

	if employment == "F" {
		return Manual
	}

	return Approve
}
