package main

import ()

type Years map[string]Year

type Year struct {
	Name   string
	Months map[string]Month
}

type Month struct {
	Name string
	Days map[string]Day
}

type Day struct {
	Name    string
	Entries map[string]Entry
}

type Entry struct {
	Name string
}
