package wlog

import "github.com/daviddengcn/go-colortext"

// Color holds the code and brigtness of the color
type Color struct {
	Code   ct.Color
	Bright bool
}

var (
	//BrightRed creates a bright red color
	BrightRed = Color{ct.Red, true}
	//BrightBlue creates a bright blue color
	BrightBlue = Color{ct.Blue, true}
	//BrightYellow creates a bright yellow color
	BrightYellow = Color{ct.Yellow, true}
	//Red creates a red color
	Red = Color{ct.Red, false}
	//Blue creaets a blue color
	Blue = Color{ct.Blue, false}
	//Yellow creates a yellow color
	Yellow = Color{ct.Yellow, false}
	//BrightGreen creates a bright green color
	BrightGreen = Color{ct.Green, true}
	//BrightCyan creates a bright cyan color
	BrightCyan = Color{ct.Cyan, true}
	//BrightMagenta creates a bright magenta color
	BrightMagenta = Color{ct.Magenta, true}
	//Green creates a green color
	Green = Color{ct.Green, false}
	//Cyan creates a cyan color
	Cyan = Color{ct.Cyan, false}
	//Magenta creates a magenta color
	Magenta = Color{ct.Magenta, false}
	//None does not change the color
	None = Color{ct.None, false}
)