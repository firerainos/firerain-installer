package api


type Item struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Packages []Package
}
