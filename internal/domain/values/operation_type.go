package values

type OperationType string

const (
	OperationType_Add      OperationType = "add"
	OperationType_Subtract OperationType = "subtract"
	OperationType_Multiply OperationType = "multiply"
	OperationType_Divide   OperationType = "divide"
	OperationType_Power    OperationType = "power"
)

func (v OperationType) String() string {
	return string(v)
}
