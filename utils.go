package main

import "strings"

var shortcodes map[string]string = map[string]string{
	"islamabad":  "Isb",
	"rawalpindi": "Rwp",
}

func city_label(name string) string {
	name = strings.ToLower(name)
	if s, ok := shortcodes[name]; ok {
		return s
	}

	return name
}
