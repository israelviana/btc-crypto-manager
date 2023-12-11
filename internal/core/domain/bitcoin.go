package domain

type DetailsAddress struct {
	Address       string        `json:"address"`
	Balance       string        `json:"balance"`
	TotalTx       int           `json:"totalTx"`
	Total         Total         `json:"total"`
	BalanceDetail BalanceDetail `json:"balance_details"`
}

type BalanceDetail struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

type Total struct {
	Sent     string `json:"sent"`
	Received string `json:"received"`
}

type Send struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

type Utxos struct {
	TxId   string `json:"tx_id"`
	Amount string `json:"amount"`
}

type Transaction struct {
	Addresses []Address `json:"addresses"`
	Block     int       `json:"block"`
	TxID      string    `json:"txID"`
}

type Address struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}

type BitcoinRequest struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}
