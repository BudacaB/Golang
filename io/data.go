package main

// DataContainer structure
type DataContainer struct {
	Name string
	Data []*Data
}

// Data structure
type Data struct {
	Title       string
	Link        string
	Description string
	Content     string
}
