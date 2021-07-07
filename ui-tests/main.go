package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"time"
)

func assertEquals(a string, b string) {
	if a != b {
		fmt.Printf("\"%v\" != \"%v\"\n", a, b)
		panic("Assertion Error")
	}
}

func main() {
	u := launcher.New().Bin("/usr/bin/google-chrome").MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage("http://ui").Timeout(5 * time.Second)

	// Test hero title

	title := page.MustElement(".hero-title")
	assertEquals(title.MustText(), "Decentralized Cloud Storage is Here")

	// Must be on registration by default
	page.MustElementR("h5", "Get Started")

	// Must switch to login
	page.MustElement("button.btn-success").MustClick()
	page.MustElementR("h5", "Login")

	// Must switch back
	page.MustElementR("a", "Sign Up").MustClick()
	page.MustElementR("h5", "Get Started")

	page.MustWaitLoad().MustScreenshot("/output/screenshot.png")
}
