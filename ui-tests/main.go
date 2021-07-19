package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"os"
	"time"
)

func assertEquals(a string, b string) {
	if a != b {
		fmt.Printf("\"%v\" != \"%v\"\n", a, b)
		panic("Assertion Error")
	}
}

func testHome(u string) {
	page := rod.New().ControlURL(u).MustConnect().MustPage("http://ui").Timeout(10 * time.Second)

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
}

func testLogin(u string) {
	page := rod.New().ControlURL(u).MustConnect().MustPage("http://ui").Timeout(10 * time.Second)

	// Must switch to login
	page.MustElement("button.btn-success").MustClick()
	page.MustElementR("h5", "Login")

	// Login and read bucket names
	page.MustElement("#emailAddress").MustInput(os.Getenv("IRIS_EMAIL"))
	page.MustElement("#password").MustInput(os.Getenv("IRIS_PASSWORD"))
	page.MustElement("button.btn-primary").MustClick()
	page.MustElement(".bucket-name")

	page.MustWaitLoad().MustScreenshot("/output/login.png")
}

func main() {
	u := launcher.New().Bin("/usr/bin/google-chrome").MustLaunch()

	testHome(u)

	testLogin(u)

}
