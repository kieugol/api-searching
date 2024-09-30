package respond

func Success(data interface{}, message string) Respond {
	return Respond{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

func MissingHeader() Respond {
	return Respond{
		Code:    1000,
		Message: "Missing request header",
		Data:    nil,
	}
}

func MissingParams() Respond {
	return Respond{
		Code:    1001,
		Message: "Missing params",
		Data:    nil,
	}
}

func NotFound() Respond {
	return Respond{
		Code:    1002,
		Message: "Not found",
		Data:    nil,
	}
}

func InternalServerError() Respond {
	return Respond{
		Code:    1003,
		Message: "Internal server error",
		Data:    nil,
	}
}
