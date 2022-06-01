package dto

type Address struct {
	CoinType string
	AddrStr  string
}

func NewAddress(coinType, address string) *Address {
	return &Address{
		CoinType: coinType,
		AddrStr:  address,
	}
}

func (a *Address) ToMessage() *addressMessage {
	return &addressMessage{
		Subscriptions: map[string][]string{
			a.CoinType: {a.AddrStr},
		},
		Operation: AddSubscription,
	}
}
