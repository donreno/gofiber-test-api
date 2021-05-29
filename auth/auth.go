package auth

func DoAuth(username, pass string) bool {
	users := map[string]string{
		"username": "password",
	}

	realPass, found := users[username]

	return found && realPass == pass
}
