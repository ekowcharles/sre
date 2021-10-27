package main

import "os"

type Payload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AppSetting struct {
	Version         string `json:"version"`
	LoadBalancerUrl string `json:"-"`
}

type Handler struct {
	as *AppSetting
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
