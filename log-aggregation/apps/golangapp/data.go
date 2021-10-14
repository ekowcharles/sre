package main

type Payload struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type App struct {
	Version string `json:"version"`
}
