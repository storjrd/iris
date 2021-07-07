package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	u := launcher.New().Bin("/usr/bin/google-chrome").MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage("http://ui")

	page.MustWaitLoad().MustScreenshot("/output/screenshot.png")
}
