package models

import (
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

type CalculationResult struct {
	Result      float64
	Precision   int
	OperationID values.OperationID
	Timestamp   time.Time
}
