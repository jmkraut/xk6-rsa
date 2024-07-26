package rsa

import (
	"github.com/golang-jwt/jwt"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/rsa", new(Rsa))
}

// Rsa is the custom API type.
type Rsa struct {}

// Sign returns a signed JWT, or an error otherwise.
func (r *Rsa) Sign(claims map[string]any, privateKey string) (string, error) {
	parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))

	if err != nil {
		return "", err
	}

	var jwtClaims = jwt.MapClaims{}

	for claim, claimValue := range claims {
		jwtClaims[claim] = claimValue
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtClaims)

	// Sign the token with the private key
	signedToken, err := token.SignedString(parsedPrivateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
