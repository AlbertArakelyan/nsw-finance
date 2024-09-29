package main

import "image/color"

type Colors struct {
	gray, darkGray color.NRGBA
}

type Theme struct {
	colors Colors
}

var appTheme = Theme{
	colors: Colors{
		gray:     color.NRGBA{R: 155, G: 155, B: 155, A: 255},
		darkGray: color.NRGBA{R: 100, G: 100, B: 100, A: 255},
	},
}
