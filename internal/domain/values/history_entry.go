package values

import "time"

type HistoryEntry struct {
	OperationID   OperationID
	OperationType OperationType
	Operands      []float64
	Result        float64
	Precision     int
	Timestamp     time.Time
}
