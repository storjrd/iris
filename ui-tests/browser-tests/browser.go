package uitest

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"storj.io/common/testcontext"
)

// Test defines common services for uitests.
type Test func(t *testing.T, browser *rod.Browser)

type zapWriter struct {
	*zap.Logger
}

func (log zapWriter) Write(data []byte) (int, error) {
	log.Logger.Info(string(data))
	return len(data), nil
}

// Run starts a new UI test.
func Browser(t *testing.T, test Test) {
	showBrowser := os.Getenv("STORJ_TEST_SHOW_BROWSER") != ""

	logLauncher := zaptest.NewLogger(t).Named("launcher")

	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	launch := launcher.New().
		Headless(!showBrowser).
		Leakless(false).
		Devtools(false).
		NoSandbox(true).
		UserDataDir(ctx.Dir("browser")).
		Logger(zapWriter{Logger: logLauncher})

	if browserBin := os.Getenv("STORJ_TEST_BROWSER"); browserBin != "" {
		launch = launch.Bin(browserBin)
	}

	defer launch.Cleanup()

	url, err := launch.Launch()
	require.NoError(t, err)

	logBrowser := zaptest.NewLogger(t).Named("rod")

	browser := rod.New().
		Timeout(time.Minute).
		ControlURL(url).
		SlowMotion(300 * time.Millisecond).
		Logger(utils.Log(func(msg ...interface{}) {
			logBrowser.Info(fmt.Sprintln(msg...))
		})).
		Context(ctx).
		WithPanic(func(v interface{}) { require.Fail(t, "check failed", v) })
	defer ctx.Check(browser.Close)

	require.NoError(t, browser.Connect())

	test(t, browser)
}
