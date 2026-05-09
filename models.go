package main

import (
	"html/template"
	"time"
)

type Option struct {
	Idx    int
	Letter string
	Text   string
	Score  int
}

type Question struct {
	ID      int
	EN      string
	TL      string
	Options []Option
}

type Session struct {
	Name            string
	CurrentQuestion int
	TotalScore      int
	MaxScore        int
	CreatedAt       time.Time
}

type Result struct {
	DashOffset int
	Name       string
	Score      int
	Max        int
	Percent    int
	Label      string
	Desc       string
	Emoji      string
	Color      template.CSS
}
