package theme

import "image/color"

type Colors struct {
	Gray, DarkGray color.NRGBA
}

type Theme struct {
	Colors Colors
}

var AppTheme = Theme{
	Colors: Colors{
		Gray:     color.NRGBA{R: 155, G: 155, B: 155, A: 255},
		DarkGray: color.NRGBA{R: 100, G: 100, B: 100, A: 255},
	},
}
