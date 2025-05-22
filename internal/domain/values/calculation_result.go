package values

import "time"

type CalculationResult struct {
	Result      float64
	Precision   int
	OperationID OperationID
	Timestamp   time.Time
}
