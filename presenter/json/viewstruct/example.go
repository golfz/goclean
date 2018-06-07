package viewstruct

type Example struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Remain      uint16 `json:"remain"`
	CreatedDate string `json:"created-date"`
	Url         string `json:"url"`
	Rel         string `json:"rel"`
}
