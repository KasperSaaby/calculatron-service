package history

type GetHistoryEntriesResponse struct {
	Entries []Entry `json:"entries"`
}

type GetHistoryEntryResponse struct {
	Entry Entry `json:"entry"`
}

type Entry struct {
	OperationID   string    `json:"operation_id"`
	OperationType string    `json:"operation_type"`
	Operands      []float64 `json:"operands"`
	Result        float64   `json:"result"`
	Precision     int32     `json:"precision"`
	Timestamp     string    `json:"timestamp"`
}
