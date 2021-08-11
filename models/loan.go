package models

type Status int

const (
	New Status = iota
	Approved
	Rejected
	Cancelled
)

func (s Status) toString() string {
	switch s {
	case New:
		return "New"
	case Approved:
		return "Approved"
	case Rejected:
		return "Rejected"
	case Cancelled:
		return "Cancelled"
	default:
		return "unknown"

	}
}

type LoanRequest struct {
	customerName string `json:"customer_name"`
	phoneNo      string `json:"phone_no"`
	email        string `json:"email"`
	loanAmount   string `json:"loan_amount"`
	status       Status `json:"status"`
	creditScore  int    `json:"credit_score"`
}
