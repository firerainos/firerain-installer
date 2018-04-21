package api

type Package struct {
	ItemID      uint   `json:"itemID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
