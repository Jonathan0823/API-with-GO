package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jon": {
		AuthToken: "456DEF",
		Username: "jon",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"jon": {
		Coins: 200,
		Username: "jon",
	},
}

func (m *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)
	var clientData = LoginDetails{}
	clientData, ok = mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (m *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(1 * time.Second)
	var clientData = CoinDetails{}
	clientData, ok = mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (m *mockDB) SetupDatabase() error {
	return nil
}