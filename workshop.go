/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// Lot armazena informação relativa aos lotes da cadeia de valor
type Lot struct {
	DocType string  `json:"docType"` // docType ("lot") tem de ser usado para se distinguir de outros documentos da state database
	ID      string  `json:"ID"`
	Product string  `json:"product"`
	Amount  float32 `json:"amount"`
	Unit    string  `json:"unit"`
	Owner   string  `json:"owner"`
}

// Activity armazena informação sobre as atividades da cadeia de valor
type Activity struct {
	DocType   string             `json:"docType"` // docType ("act") tem de ser usado para se distinguir de outros documentos da state database
	ID        string             `json:"ID"`
	InputLots map[string]float32 `json:"inputLots,omitempty" metadata:",optional"` // inputLots é opcional porque podemos ter atividades que apenas registam lotes que vêm de fora da cadeira de valor
	OutputLot Lot                `json:"outputLot"`
	Date      string             `json:"date"`
	Issuer    string             `json:"issuer"`
}
