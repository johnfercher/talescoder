package models

type Asset struct {
	Id           []byte
	LayoutsCount int16
	Layouts      []*Layout
}
