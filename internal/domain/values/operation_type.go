package values

type OperationType string

const (
	OperationType_Add      OperationType = "add"
	OperationType_Subtract OperationType = "subtract"
)

func (v OperationType) String() string {
	return string(v)
}
