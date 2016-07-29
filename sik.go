package main

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/size"
)

func main() {
	app := spartan.New()

	layout := spartan.LinearLayout{
		Direction:  direction.Vertical,
		IsBordered: false,
		Title:      "sik",
	}
	layout.Height = size.MatchParent
	layout.Width = size.MatchParent

	contentLayout := spartan.LinearLayout{
		Direction:  direction.Vertical,
		IsBordered: false,
		Title:      "content",
	}
	contentLayout.Width = size.MatchParent
	contentLayout.Height = size.MatchParent

	helloBox := spartan.TextView{
		Text: "Hello, World!",
	}
	helloBox.Width = size.MatchParent
	helloBox.Height = 1
	helloBox.ForegroundColor = termbox.ColorWhite | termbox.AttrBold
	helloBox.BackgroundColor = termbox.Attribute(236) // https://upload.wikimedia.org/wikipedia/en/1/15/Xterm_256color_chart.svg

	layout.AddChild(&contentLayout)
	layout.AddChild(&helloBox)

	app.SetContent(&layout)
	app.Run()
}
