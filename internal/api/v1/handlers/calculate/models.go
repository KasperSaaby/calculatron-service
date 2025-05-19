package calculate

type Request struct {
	OperationType string    `json:"operation_type"`
	Operands      []float64 `json:"operands"`
	Precision     int       `json:"precision"`
}

type Response struct {
	Result      float64 `json:"result"`
	Precision   int     `json:"precision"`
	OperationID string  `json:"operationId"`
	Timestamp   string  `json:"timestamp"`
}
