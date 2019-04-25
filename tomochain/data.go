package tomochain

// const (
// 	TIME_TO_DELETE = 18000
// )

type TokenAPI struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Decimals    int    `json:"decimals"`
	UsdID       string `json:"cmc_id"`
	TimeListing uint64 `json:"listing_time,omitempty"`
	CGId        string `json:"cg_id"`
	// DelistTime  uint64 `json:"delist_time,omitempty"`
}

type Token struct {
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Address    string `json:"address"`
	Decimal    int    `json:"decimals"`
	DelistTime uint64 `json:"delist_time"`
	CGId       string `json:"cg_id"`
	Priority   bool   `json:"priority"`
	TokenID    string `json:"token_id"`
}

func TokenAPIToToken(tokenAPI TokenAPI) Token {
	// if tokenAPI.DelistTime == 0 || uint64(time.Now().UTC().Unix()) <= TIME_TO_DELETE+tokenAPI.DelistTime {
	return Token{
		Name:    tokenAPI.Name,
		Symbol:  tokenAPI.Symbol,
		Address: tokenAPI.Address,
		Decimal: tokenAPI.Decimals,
		CGId:    tokenAPI.CGId,
		TokenID: tokenAPI.Symbol,
	}
}
