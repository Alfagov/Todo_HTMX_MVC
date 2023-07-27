package models

import "strings"

type NavBarLinks struct {
	Links []Link
}

type Link struct {
	Name string
	URL  string
}

func NewNavBarLinks(links []Link) NavBarLinks {

	for i, link := range links {
		if link.URL == "" {
			links[i].URL = strings.ToLower(link.Name)
		}
	}
	return NavBarLinks{links}
}

var DefaultNavBarLinks = []Link{
	{"Home", "/"},
	{"Products", ""},
	{"About", ""},
}
