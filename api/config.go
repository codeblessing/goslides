package api

import (
	"encoding/json"
	"errors"
)

// Desired PDF page properties.
//
// # Miscellaneous
//
// At least one of `width` or `height` should be specified.
// If all `width`, `heigth` and `ratio` are specified then
// `ratio` takes precedence over `height`.
type Page struct {
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Ratio  float32 `json:"ratio"`
	Margin float32 `json:"margin"`
	Color  Color   `json:"background_color"`
}

// Desired font properties.
type Font struct {
	Family string `json:"family"`
	Size   uint   `json:"size"`
	Color  Color  `json:"color"`
}

// Color specified in RGB coordinates.
type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (_color *Color) UnmarshalJSON(source []byte) error {
	var rgb []uint8
	if err := json.Unmarshal(source, &rgb); err != nil {
		return err
	}
	if len(rgb) != 3 {
		return errors.New("invalid RGB value")
	}

	_color.Red = rgb[0]
	_color.Green = rgb[1]
	_color.Blue = rgb[2]

	return nil
}

// Desired properties for generated slides.
type Config struct {
	Page        Page `json:"page"`
	Font        Font `json:"font"`
	HintFont    Font `json:"hint_font"`
	LineSpacing float64
}
