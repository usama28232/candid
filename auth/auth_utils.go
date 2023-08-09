package auth

import (
	"encoding/base64"
	"strings"

	"github.com/usama28232/candid/logging"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func ComparePassword(hashedPass, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(input))
}

func DecodeBase64(encoded string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return ""
	}
	return string(decodedBytes)
}

func BypassAuthFilter(data []string, r string) bool {
	logger := logging.GetLogger()
	base := strings.Split(r, "/")
	var basepath string
	if len(base) >= 1 {
		basepath = base[1]
	} else {
		basepath = ""
	}
	// fmt.Printf("r splits: %v\n", base)
	for _, value := range data {
		logger.Debugw("Authentication Filter", "url", r, "parsed", basepath, "comparing with", value)
		// WILDCARD CASE
		if strings.Contains(value, "*") {
			noauthArr := strings.Split(value, "/")
			if len(noauthArr) >= 1 {
				noAuthBase := noauthArr[1]
				return noAuthBase == basepath
			} else {
				return false
			}
		} else { // EXACT MATCH
			return value == r
		}
	}
	return false
}
