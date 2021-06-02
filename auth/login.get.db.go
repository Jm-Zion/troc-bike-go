package auth
import (
	"github.com/Jm-Zion/troc-bike-go/app"
	"fmt"
)
func doesLoginExists(identifier string) (*Login, error) {
	login := &Login{
		Login: identifier,
	}
	err := app.PGMain().Model(login).Where("login = ?",identifier).Select()
	if err != nil {
		fmt.Errorf("No user found for %v", identifier)
		return nil, err
	}
	return login, nil
}