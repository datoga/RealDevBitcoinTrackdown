package main

type ApiRespRawAddr struct {
	Transactions  ApiRespRawAddrTxs `json:"txs"`
	NTransactions int               `json:"n_tx"`
}

type ApiRespRawAddrTxs []ApiRespRawAddrTx

type ApiRespRawAddrTx struct {
	Inputs ApiRespRawAddrInputs `json:"inputs"`
	Outs   ApiRespRawAddrOuts   `json:"out"`
}

type ApiRespRawAddrInputs []ApiRespRawAddrInput

type ApiRespRawAddrInput struct {
	Sequence int `json:"sequence"`
}

type ApiRespRawAddrOuts []ApiRespRawAddrOut

type ApiRespRawAddrOut struct {
	Address string `json:"addr"`
	Value   int    `json:"value"`
}
