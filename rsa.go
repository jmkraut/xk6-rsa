package rsa

import (
	"os"

	"github.com/golang-jwt/jwt"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/rsa", new(Rsa))
}

// Rsa is the custom API type.
type Rsa struct{}

// SignFromString returns a signed JWT, or an error otherwise.
func (r *Rsa) SignFromString(claims map[string]any, privateKey string) (string, error) {
	return sign(claims, []byte(privateKey))
}

// SignFromFilePath reads from a file path, and if the path is valid returns a signed JWT, or an error otherwise.
func (r *Rsa) SignFromFilePath(claims map[string]any, privateKeyPath string) (string, error) {
	privateKeyFile, err := os.ReadFile(privateKeyPath)

	if err != nil {
		return "", err
	}

	return sign(claims, privateKeyFile)
}

func sign(claims map[string]any, privateKeyFile []byte) (string, error) {
	parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)

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
