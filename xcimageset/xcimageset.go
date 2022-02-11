package xcimageset

import (
	"strings"
)

type XCImageSet struct {
	Images []XCImage
	Info   XCImageInfo
}

type XCImage struct {
	Filename string
	Idiom    string
	Scale    string
}

type XCImageInfo struct {
	Author  string
	Version string
}

func (imgs *XCImageSet) AssetName() string {
	if imgs.Images[0].Filename != "" {
		return strings.Split(imgs.Images[0].Filename, ".jpg")[0]
	}
	if imgs.Images[1].Filename != "" {
		return strings.Split(imgs.Images[1].Filename, ".jpg")[0]
	}
	if imgs.Images[2].Filename != "" {
		return strings.Split(imgs.Images[2].Filename, ".jpg")[0]
	}
	panic("No file name to extract")
}
