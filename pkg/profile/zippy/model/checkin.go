package zippy

type CheckinRequest struct {
	Tag    string `json:"tag"`
	Client bool   `json:"client"`
	Data   string `json:"data"`
}

type CheckinResponse struct {
	Tag    string `json:"tag"`
	Client bool   `json:"client"`
	Data   string `json:"data"`
}
