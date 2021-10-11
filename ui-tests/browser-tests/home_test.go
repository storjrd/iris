package uitest

import (
	"testing"

	"github.com/go-rod/rod"
	"github.com/stretchr/testify/require"
)

func TestHome(t *testing.T) {
	Browser(t, func(t *testing.T, browser *rod.Browser) {
		page := browser.MustPage("http://ui")

		// Test hero title
		title := page.MustElement(".hero-title")
		require.Equal(t, "Decentralized Cloud Storage is Here", title.MustText(), "Home is not being displayed")

		// Must be on registration by default
		page.MustElementR("h5", "Get Started")

		// Must switch to login
		page.MustElementR("button", "Login to Storj").MustClick()
		page.MustElementR("h5", "Login")

		// Must switch back
		page.MustElementR("a", "Sign Up").MustClick()
		h5Title := page.MustElementR("h5", "Get Started")
		require.Equal(t, "Get Started", h5Title.MustText())
	})
}
