package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type GeneralResponse struct {
	Success bool `json:"success"`
}
