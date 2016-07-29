package main

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/gravity"
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

	statusBar := spartan.LinearLayout{
		Direction:  direction.Horizontal,
		IsBordered: false,
	}
	statusBar.Height = 1
	statusBar.Width = size.MatchParent
	statusBar.ForegroundColor = termbox.ColorWhite | termbox.AttrBold
	statusBar.BackgroundColor = termbox.Attribute(236) // https://upload.wikimedia.org/wikipedia/en/1/15/Xterm_256color_chart.svg

	clock := spartan.TextView{
		Text: "12:00 PM",
	}
	clock.Width = 20
	clock.Height = size.MatchParent
	clock.LayoutGravity = gravity.Right

	statusBar.AddChild(&clock)

	layout.AddChild(&contentLayout)
	layout.AddChild(&statusBar)

	app.SetContent(&layout)
	app.Run()
}
