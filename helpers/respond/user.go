package respond

func MissingLogin() Respond {
	return Respond{
		Code:    3002,
		Message: "Username and password incorrect",
		Data:    nil,
	}
}
