package user

import (
	"fmt"
	"github.com/Jm-Zion/troc-bike-go/app"
)


type AccountType string

const (
	FREE_ACCOUNT AccountType = "FREE"
)

func (e AccountType) String() string {
	accounts := [...]string{"FREE"}
 
	x := string(e)
	for _, v := range accounts {
		if v == x {
			return x
		}
	}
 
	return ""
}

func GetAccountByType(accountType AccountType) (*Account, error){
	accountString := accountType.String()
	acc := &Account{
		Type: &accountString,
	}
	err := app.PGMain().Model(acc).Where("type = ?",accountType).Select()
	if err != nil {
		fmt.Errorf("No account found for type %v", accountType)
		return nil, err
	}
	return acc, nil
}