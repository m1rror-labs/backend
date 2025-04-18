package runtimes

type request struct {
	Code string `json:"code" binding:"required"`
}

type response struct {
	Error  string `json:"error"`
	Output string `json:"output"`
}
