package supportingfunctions

import "net/http"

// CloseBodyHTTPRequest закрытие http.Body запроса с предварительной проверкой
func CloseBodyHTTPRequest(req *http.Request) {
	if req == nil || req.Body == nil {
		return
	}

	req.Body.Close()
}

// CloseBodyHTTPResponse закрытие http.Body ответа с предварительной проверкой
func CloseBodyHTTPResponse(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}

	res.Body.Close()
}
