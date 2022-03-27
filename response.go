package clientix

type AddResponse struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}

type ListResponse struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
