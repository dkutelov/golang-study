package ui

import (
	"errors"
	"strconv"

	"fyne.io/fyne/v2"
)

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)

			if err != nil {
				return errors.New("must be a positive integer")
			}

			if width <= 0 {
				return errors.New("must be > 0")
			}
			return nil
		}
	})
}