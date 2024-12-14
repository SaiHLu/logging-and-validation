package presenter

func DefaultSuccessJsonResponse(data any, message string) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
		"error":   nil,
	}
}

func DefaultErrorJsonResponse(message string, err any) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"message": message,
		"data":    nil,
		"error":   err,
	}
}
