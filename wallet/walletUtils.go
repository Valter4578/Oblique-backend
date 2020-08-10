package wallet

import (
	"oblique/model"
	"strings"
)

// FindWallet finds a category by its title. Returns category and index
func FindWallet(title string) (model.Wallet, int) {
	var wallet model.Wallet
	var index int

	for i, w := range model.Wallets {
		if strings.ToLower(title) == strings.ToLower(w.Title) {
			wallet = w
			index = i
		}
	}

	return wallet, index
}
