package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func assertEquals(a string, b string) {
	if a != b {
		fmt.Printf("\"%v\" != \"%v\"\n", a, b)
		panic("Assertion Error")
	}
}

func main() {
	u := launcher.New().Bin("/usr/bin/google-chrome").MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage("http://ui")

	// Test hero title

	title := page.MustElement(".hero-title")
	assertEquals(title.MustText(), "Decentralized Cloud Storage is Here")

	// Test switch to login box
	page.MustElementR("h5", "Get Started")
	page.MustElement("button.btn-success").MustClick()
	page.MustElementR("h5", "Login")

	page.MustWaitLoad().MustScreenshot("/output/screenshot.png")
}
