package main

type Schedule struct {
	Stop    string `json:"stop"`
	Route   string `json:"route"`
	Service string `json:"service"`
	Code    string `json:"code"`
	Times   string `json:"times"`
}

type Route struct {
	Id    string `json:"id"`
	Code  string `json:"code"`
	Shape string `json:"shape"`
	Name  string `json:"name"`
	Stops string `json:"stops"`
}

type RouteShape struct {
	Route  string `json:"route"`
	Points string `json:"points"`
}

type Stop struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Lat    string `json:"lat"`
	Lon    string `json:"lon"`
	Routes string `json:"routes"`
}
