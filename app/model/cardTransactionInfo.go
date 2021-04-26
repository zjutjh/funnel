package model

type CardTransaction struct {
	ID              string `json:"id"`
	Account         string `json:"account"`
	CardType        string `json:"cardType"`
	TransactionType string `json:"transactionType"`
	Shop            string `json:"shop"`
	ShopPlace       string `json:"shopPlace"`
	Terminal        string `json:"terminal"`
	Transactions    string `json:"transactions"`
	Time            string `json:"time"`
	Wallet          string `json:"wallet"`
	Balance         string `json:"balance"`
}
