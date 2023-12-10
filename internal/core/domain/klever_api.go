package domain

import "math/big"

type AddressDetails struct {
	Page               int      `json:"page"`
	TotalPages         int      `json:"totalPages"`
	ItemsOnPage        int      `json:"itemsOnPage"`
	Address            string   `json:"address"`
	Balance            string   `json:"balance"`
	TotalReceived      string   `json:"totalReceived"`
	TotalSent          string   `json:"totalSent"`
	UnconfirmedBalance string   `json:"unconfirmedBalance"`
	UnconfirmedTxs     int      `json:"unconfirmedTxs"`
	Txs                int      `json:"txs"`
	Txids              []string `json:"txids"`
}

type TransationDetails struct {
	TxId          string `json:"txId"`
	Version       int    `json:"version"`
	BlockHash     string `json:"blockHash"`
	BlockHeight   int    `json:"blockHeight"`
	Confirmations int    `json:"confirmations"`
	BlockTime     int    `json:"blockTime"`
	Size          int    `json:"size"`
	Vsize         int    `json:"vsize"`
	Value         string `json:"value"`
	ValueIn       string `json:"valueIn"`
	Fees          string `json:"fees"`
	Hex           string `json:"hex"`
	Vin           []Vin  `json:"vin"`
	VOut          []VOut `json:"vout"`
}

type Vin struct {
	TxId      string   `json:"txId"`
	Vout      big.Int  `json:"vout"`
	Sequence  big.Int  `json:"sequence"`
	N         big.Int  `json:"n"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
	Value     string   `json:"value"`
}

type VOut struct {
	Value     string   `json:"value"`
	N         int      `json:"n"`
	Spent     bool     `json:"spent"`
	Hex       string   `json:"hex"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
}

type UTXODetails struct {
	Txid          string `json:"txid"`
	Vout          int    `json:"vout"`
	Value         string `json:"value"`
	Height        int    `json:"height"`
	Confirmations int    `json:"confirmations"`
}
