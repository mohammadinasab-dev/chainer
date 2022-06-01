package models

type (
	// Types of transaction statuses
	ParsedDirection string
	ParsedStatus    string
	// Amount is a positive decimal integer string.
	// It is written in the smallest possible unit (e.g. Wei, Satoshis)
	ParsedAmount string
)
type ParsedTx struct {
	// Unique identifier
	ID string `json:"id"`
	// SLIP-44 coin index of the platform
	Coin uint `json:"coin"`
	// Address of the transaction sender
	From string `json:"from"`
	// Address of the transaction recipient
	To string `json:"to"`
	// Transaction fee (native currency)
	Fee string `json:"fee"`
	// Unix timestamp of the block the transaction was included in
	Date int64 `json:"date"`
	// Height of the block the transaction was included in
	Block uint64 `json:"block"`
	// Status of the transaction e.g: "completed", "pending", "error"
	Status string `json:"status"`
	// Empty if the transaction "completed" or "pending", else error explaining why the transaction failed (optional)
	Error string `json:"error,omitempty"`
	// Transaction nonce or sequence
	Sequence uint64 `json:"sequence"`
	// Type of metadata
	Type string `json:"type"`
	// Input addresses
	Inputs []TxOutput `json:"inputs,omitempty"`
	// Output addresses
	Outputs []TxOutput `json:"outputs,omitempty"`
	// Transaction Direction
	Direction Direction `json:"direction,omitempty"`
	// TokenTransfers
	TokenTransfers []TokenTransfer `json:"token_transfers,omitempty"`
	// Meta data object
	Memo string      `json:"memo"`
	Meta interface{} `json:"metadata"`
}

type ParsedTxOutput struct {
	Address string `json:"address"`
	Value   Amount `json:"value"`
}

// Transfer describes the transfer of currency native to the platform
type ParsedTransfer struct {
	Value    Amount `json:"value"`
	Symbol   string `json:"symbol"`
	Decimals uint   `json:"decimals"`
}

type ParsedTokenTransfer struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenID  string `json:"token_id"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
	From     string `json:"from"`
	To       string `json:"to"`
}
