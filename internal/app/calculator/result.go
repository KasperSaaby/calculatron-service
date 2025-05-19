package calculator

import "time"

type Result struct {
	Result      float64
	Precision   int
	OperationID string
	Timestamp   time.Time
}
