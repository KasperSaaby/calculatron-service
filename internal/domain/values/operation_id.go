package values

import "github.com/google/uuid"

type OperationID string

func NewOperationID() OperationID {
	return OperationID(uuid.NewString())
}

func (v OperationID) String() string {
	return string(v)
}
