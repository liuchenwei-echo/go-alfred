package alfred

type Item struct {
	Uid string `json:"uid"`
	Type string `json:"type"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg string `json:"arg"`
}

type ItemList struct {
	Items []Item `json:"items"`
}
