package uitest

import (
	"testing"
	"os"

	"github.com/go-rod/rod"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	Browser(t, func (t *testing.T, browser *rod.Browser)  {
		page := browser.MustPage("http://localhost:3000")

		// Must switch to login
		page.MustElementR("button", "Login to Storj").MustClick()
		page.MustElementR("h5", "Login")

		// Login and read bucket names
		page.MustElement("#emailAddress").MustInput(os.Getenv("IRIS_EMAIL"))
		page.MustElement("#password").MustInput(os.Getenv("IRIS_PASSWORD"))
		page.MustElementR("button", "Login").MustClick()
		title := page.MustElementR("h2", "Buckets")
		require.Equal(t, "Buckets", title.MustText(), "Login was unsuccessful")
	})
}
