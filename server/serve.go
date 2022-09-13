package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type spice string

type spiceMix struct {
	Name   string  `json:"name"`
	Spices []spice `json:"spices"`
}

var mixes = []*spiceMix{
	{
		"Bland",
		[]spice{"salt", "pepper"},
	},
	{
		"Classic Steak",
		[]spice{"salt", "pepper", "garlix powder", "oregano"},
	},
	{
		"shawarma",
		[]spice{"salt", "pepper", "allspice", "cumin", "cinnamon", "paprika"},
	},
}

func RegisterHandlers() {
	http.HandleFunc("/mixes", GetAllSpiceMixes)
	http.HandleFunc("/grind", handleGrindCommand)
}

func Serve() error {
	fmt.Printf("listening on localhost:8080...\n")
	return http.ListenAndServe(":8080", nil)
}

func GetAllSpiceMixes(w http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(w).Encode(mixes)
	if err != nil {
		fmt.Printf("error encoding response: %s", err)
		_, _ = fmt.Fprintln(w, "server error")
	}
}

func handleGrindCommand(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello, World!")
}

/*

first 90 minutes is general technical
	30 minute prep and cloning repo and reading docs
	60 minute peer programming to see how I activiely problem solve, you can google and ask and do whatever
	10 at the end minute technical discussions. nobody ever finished this exercise. journey is more important than the destination.
	there will be a sql component

45 minute system design. whiteboard capability.

45 minute teamwork and collaboration interview. best practices. mistakes that I made.

*/
