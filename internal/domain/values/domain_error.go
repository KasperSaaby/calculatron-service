package values

import "fmt"

type Code string

const (
	Code_TwoOrMoreOperandsRequired Code = "two_or_more_operands_required"
	Code_TwoOperandsRequired       Code = "two_operands_required"
	Code_DivisionByZero            Code = "division_by_zero"
	Code_UnsupportedOperation      Code = "unsupported_operation"
)

type DomainError struct {
	Message string
	Code    Code
}

func NewDomainError(message string, code Code) *DomainError {
	return &DomainError{
		Message: message,
		Code:    code,
	}
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("(%s) %s", e.Code, e.Message)
}
