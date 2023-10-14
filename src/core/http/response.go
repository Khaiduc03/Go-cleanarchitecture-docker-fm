package http

type HttpResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func NewHttpResponse(statusCode int, message string, data interface{}) HttpResponse {
	if data == nil {
		return HttpResponse{
			StatusCode: statusCode,
			Message:    message,
			Data:       nil,
		}
	}

	// check data is object
	if _, ok := data.(map[string]interface{}); ok {

		return HttpResponse{
			StatusCode: statusCode,
			Message:    message,
			Data:       data,
		}
	} else {

		return HttpResponse{
			StatusCode: statusCode,
			Message:    message,
			Data:       data,
		}
	}
}

func ConvertAttribute(o map[string]interface{}) map[string]interface{} {
	// check stop condition (object not child object)
	if o == nil {
		return nil
	}

	// convert attribute to camel case
	for key, value := range o {
		// check value is object
		if _, ok := value.(map[string]interface{}); ok {
			// convert attribute to camel case
			o[key] = ConvertAttribute(value.(map[string]interface{}))
		} else {
			// convert string to camel case
			o[key] = ConvertStringToCamelCase([]byte(value.(string)))
		}
	}

	return o
}

func ConvertStringToCamelCase(str []byte) []byte {
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = str[i] + 32
		}
	}

	return str
}
