package main

import "strings"

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListStategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type Mka