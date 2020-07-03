package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type custTheme struct {
}

func (c custTheme) BackgroundColor() color.Color {
	return theme.DarkTheme().BackgroundColor()
}

func (c custTheme) ButtonColor() color.Color {
	return theme.DarkTheme().ButtonColor()
}

func (c custTheme) HyperlinkColor() color.Color {
	return theme.DarkTheme().HyperlinkColor()
}

func (c custTheme) TextColor() color.Color {
	return theme.DarkTheme().TextColor()
}

func (c custTheme) PlaceHolderColor() color.Color {
	return theme.DarkTheme().PlaceHolderColor()
}

func (c custTheme) PrimaryColor() color.Color {
	return theme.DarkTheme().PrimaryColor()
}

func (c custTheme) FocusColor() color.Color {
	return theme.DarkTheme().FocusColor()
}

func (c custTheme) ScrollBarColor() color.Color {
	return theme.DarkTheme().ScrollBarColor()
}

func (c custTheme) TextSize() int {
	return theme.DarkTheme().TextSize()
}

func (c custTheme) TextFont() fyne.Resource {
	font, err := fyne.LoadResourceFromPath("irohamaru-mikami-Regular.ttf")
	if err != nil {
		fmt.Println("It was not possible to load the font. Err: ", err)
	}
	return font
}

func (c custTheme) TextBoldFont() fyne.Resource {
	font, err := fyne.LoadResourceFromPath("irohamaru-mikami-Regular.ttf")
	if err != nil {
		fmt.Println("It was not possible to load the font. Err: ", err)
	}
	return font
}

func (c custTheme) TextItalicFont() fyne.Resource {
	font, err := fyne.LoadResourceFromPath("irohamaru-mikami-Regular.ttf")
	if err != nil {
		fmt.Println("It was not possible to load the font. Err: ", err)
	}
	return font
}

func (c custTheme) TextBoldItalicFont() fyne.Resource {
	font, err := fyne.LoadResourceFromPath("irohamaru-mikami-Regular.ttf")
	if err != nil {
		fmt.Println("It was not possible to load the font. Err: ", err)
	}
	return font
}

func (c custTheme) TextMonospaceFont() fyne.Resource {
	font, err := fyne.LoadResourceFromPath("irohamaru-mikami-Regular.ttf")
	if err != nil {
		fmt.Println("It was not possible to load the font. Err: ", err)
	}
	return font
}

func (c custTheme) Padding() int {
	return theme.DarkTheme().Padding()
}

func (c custTheme) IconInlineSize() int {
	return theme.DarkTheme().IconInlineSize()
}

func (c custTheme) ScrollBarSize() int {
	return theme.DarkTheme().ScrollBarSize()
}

func (c custTheme) DisabledButtonColor() color.Color {
	return theme.DarkTheme().DisabledButtonColor()
}

func (c custTheme) DisabledIconColor() color.Color {
	return theme.DarkTheme().DisabledIconColor()
}

func (c custTheme) DisabledTextColor() color.Color {
	return theme.DarkTheme().DisabledTextColor()
}

func (c custTheme) HoverColor() color.Color {
	return theme.DarkTheme().HoverColor()
}

func (c custTheme) IconColor() color.Color {
	return theme.DarkTheme().IconColor()
}

func (c custTheme) ScrollBarSmallSize() int {
	return theme.DarkTheme().ScrollBarSmallSize()
}

func (c custTheme) ShadowColor() color.Color {
	return theme.DarkTheme().ShadowColor()
}
