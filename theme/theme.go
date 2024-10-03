package theme

import "image/color"

type Colors struct {
	Gray, DarkGray, LightGray, VeryLightGray color.NRGBA
}

type Theme struct {
	Colors Colors
}

var AppTheme = Theme{
	Colors: Colors{
		Gray:          color.NRGBA{R: 155, G: 155, B: 155, A: 255},
		DarkGray:      color.NRGBA{R: 100, G: 100, B: 100, A: 255},
		LightGray:     color.NRGBA{R: 200, G: 200, B: 200, A: 255},
		VeryLightGray: color.NRGBA{R: 240, G: 240, B: 240, A: 255},
	},
}
