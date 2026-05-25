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

func LoanDecision(age int, income float64, creditScore int, employment string) string {
	if age < 18 || age > 65 {
		return Invalid
	}

	if income < 5.0 || income > 500.0 {
		return Invalid
	}

	if creditScore < 300 || creditScore > 850 {
		return Invalid
	}

	if employment != "C" && employment != "F" {
		return Invalid
	}

	risk := ""
	if creditScore >= 300 && creditScore <= 500 {
		risk = High
	} else if creditScore >= 501 && creditScore <= 700 {
		risk = Medium
	} else {
		risk = Low
	}

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
