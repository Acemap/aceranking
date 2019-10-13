package controller

type Response struct {
	Status  bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RespondError(message string) Response {
	return Response{
		Status:  false,
		Message: message,
		Data:    nil,
	}
}

func RespondData(data interface{}) Response {
	return Response{
		Status:  true,
		Message: "",
		Data:    data,
	}
}
