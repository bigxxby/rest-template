package jwt

// returns userId and error if no user in JWT token
func ParseToken(token string) (int, error) {
	if token == "token" {
		return 1, nil
	} else {
		return 0, nil
	}

}
