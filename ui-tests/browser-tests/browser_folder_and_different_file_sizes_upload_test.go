package uitest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-rod/rod"
	"github.com/stretchr/testify/require"
)

func TestBrowserFolderAndDifferentFileSizesUpload(t *testing.T) {
	Browser(t, func(t *testing.T, browser *rod.Browser) {
		page := browser.MustPage("http://ui")

		// Must switch to login
		page.MustElementR("button", "Login to Storj").MustClick()

		// Login
		page.MustElement("#emailAddress").MustInput(os.Getenv("IRIS_EMAIL"))
		page.MustElement("#password").MustInput(os.Getenv("IRIS_PASSWORD"))
		page.MustElementR("button", "Login").MustClick()

		// Navigate into bucket and browser
		page.MustElementR("span", "go-rod-test").MustClick()
		page.MustElement("#passphrase").MustInput("go-rod-test")
		page.MustElementR("button", "Unlock").MustClick()

		// Verify that browser component has loaded and that the dropzone is present
		page.MustElementR("p", "Drop Files Here to Upload")

		// Create a Folder
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("testing")
		page.MustElementR("button", "Save Folder").MustClick()
		page.MustElementR("[aria-roledescription=folder]", "testing").MustClick()

		// Attempt to create a folder with spaces
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("   ")
		require.Equal(t, "true", page.MustElementR("button", "Save Folder").MustProperty("disabled").Str(), "Folder is not disabled on invalid folder name with spaces")
		require.Equal(t, "   ", page.MustElement("[placeholder=\"Name of the folder\"]").MustText(), "Folder input does not contain the empty invalid name")
		page.MustElementR("button", "Cancel").MustClick()

		// // Drag and drop folder with duplicate names ***No implementation has worked yet**
		// // storjlogo2, _ := filepath.Abs("./testdata/")
		// // mouse := page.Mouse
		// // mouse.MustMove(3, 3)
		// // mouse.MustDown("left")
		// // mouse.MustUp("left")

		// Upload a folder (folder upload doesn't work when headless)
		testData, _ := filepath.Abs("./testdata")
		testDataTxtFile, _ := filepath.Abs("./testdata/test0bytes.txt")

		if os.Getenv("STORJ_TEST_SHOW_BROWSER") == "" {
			// Create folder
			page.MustElementR("button", "New Folder").MustClick()
			page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("testData")
			page.MustElementR("button", "Save Folder").MustClick()

			// Navigate into folder and upload file
			page.MustElementR("[aria-roledescription=folder]", "testData").MustClick()
			page.MustElement("[href=\"/app/buckets/go-rod-test/browse/testing/testData/\"]")
			page.MustElementR("p", "Drop Files Here to Upload").MustText()

			// Attempt to create a folder with spaces
			page.MustElementR("button", "New Folder").MustClick()
			page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("   ")
			require.Equal(t, "true", page.MustElementR("button", "Save Folder").MustProperty("disabled").Str(), "Folder is not disabled on invalid folder name with spaces")
			require.Equal(t, "   ", page.MustElement("[placeholder=\"Name of the folder\"]").MustText(), "Folder input does not contain the empty invalid name")
			page.MustElementR("button", "Cancel").MustClick()

			wait := page.MustWaitRequestIdle()
			page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testDataTxtFile)
			wait()
			page.MustElementR("[aria-roledescription=file]", "test0bytes.txt").MustClick()
			require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the test0bytes.txt file")
		} else {
			wait2 := page.MustWaitRequestIdle()
			page.MustElement("input[aria-roledescription=folder-upload]").MustSetFiles(testData)
			page.MustElementR("[aria-roledescription=files-uploading-count]", "5 files waiting to be uploaded...")
			page.MustElementR("[aria-roledescription=files-uploading-count]", "1 file waiting to be uploaded...")
			wait2()
			page.MustElementR("table > tbody > tr:nth-child(2) > td", "testdata")
			page.MustElementR("[aria-roledescription=folder]", "testdata").MustClick()
			page.MustElementR("[aria-roledescription=file]", "test0bytes.txttest0bytes....").MustClick()
			require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The uploaded folder did not upload the files correctly")
		}

		page.MustElement("#close-modal").MustClick()
		page.MustElement("#navigate-back").MustClick()

		// Upload duplicate folder (folder upload doesn't work when headless)
		if os.Getenv("STORJ_TEST_SHOW_BROWSER") == "" {
			// Create folder
			page.MustElementR("button", "New Folder").MustClick()
			page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("testdata (1)")
			page.MustElementR("button", "Save Folder").MustClick()

			// Navigate into folder and upload file
			page.MustElementR("[aria-roledescription=folder]", "testdata \\(1\\)").MustClick()
			page.MustElement("[href=\"/app/buckets/go-rod-test/browse/testing/testdata (1)/\"]")
			page.MustElementR("p", "Drop Files Here to Upload")

			// Attempt to create a folder with spaces
			page.MustElementR("button", "New Folder").MustClick()
			page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("   ")
			require.Equal(t, "true", page.MustElementR("button", "Save Folder").MustProperty("disabled").Str(), "Folder is not disabled on invalid folder name with spaces")
			require.Equal(t, "   ", page.MustElement("[placeholder=\"Name of the folder\"]").MustText(), "Folder input does not contain the empty invalid name")
			page.MustElementR("button", "Cancel").MustClick()

			wait3 := page.MustWaitRequestIdle()
			page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testDataTxtFile)
			wait3()
			page.MustElementR("[aria-roledescription=file]", "test0bytes.txt").MustClick()
			require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the test0bytes.txt file")
		} else {
			wait4 := page.MustWaitRequestIdle()
			page.MustElement("input[aria-roledescription=folder-upload]").MustSetFiles(testData)
			page.MustElementR("[aria-roledescription=files-uploading-count]", "2 files waiting to be uploaded...")
			page.MustElementR("[aria-roledescription=files-uploading-count]", "1 file waiting to be uploaded...")
			wait4()
			page.MustElementR("table > tbody > tr:nth-child(1) > td", "..")
			page.MustElementR("[aria-roledescription=folder]", "testdata \\(1\\)").MustClick()
			page.MustElementR("[aria-roledescription=file]", "test0bytes.txttest0bytes....").MustClick()
			require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The uploaded folder did not upload the files correctly")
		}

		page.MustElement("#close-modal").MustClick()
		page.MustElement("#navigate-back").MustClick()

		// Upload a 0 byte file
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testDataTxtFile)
		page.MustElementR("[aria-roledescription=file]", "test0bytes.txt").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the 0 byte file")
		page.MustElement("#close-modal").MustClick()

		// Upload duplicate 0 byte file
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testDataTxtFile)
		page.MustElementR("[aria-roledescription=file]", "test0bytes \\(1\\).txt").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the duplicate file")
		page.MustElement("#close-modal").MustClick()

		// Upload a 50 MB file
		testFile := generateByteFile("testFile", 50)
		wait5 := page.MustWaitRequestIdle()
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testFile)
		wait5()
		page.MustElementR("[aria-roledescription=file]", "testFile").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the 50 MB file")
		page.MustElement("#close-modal").MustClick()

		// Attempt to upload a large file and cancel the upload after a few segments have been uploaded successfully
		testFile2 := generateByteFile("testFile2", 130)
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testFile2)
		require.Equal(t, " testing/testFile2", page.MustElement("[aria-roledescription=file-uploading]").MustText(), "The testFile2 file has not started uploading")
		page.MustElementR("[aria-roledescription=files-uploading-count]", "1 file waiting to be uploaded...")
		page.MustElementR("[aria-roledescription=progress-bar]", "1")
		page.MustElementR("button", "Cancel").MustClick()
		page.MustElementR("table > tbody > tr:nth-child(6) > td > span", "testFile")
		page.MustElementR("[aria-roledescription=file]", "testFile").MustClick()
		page.MustElement("#close-modal").MustClick()

		// Upload a 130MB file
		wait6 := page.MustWaitRequestIdle()
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testFile2)
		require.Equal(t, " testing/testFile2", page.MustElement("[aria-roledescription=file-uploading]").MustText(), "The testFile2 file has not started uploading")
		page.MustElementR("[aria-roledescription=files-uploading-count]", "1 file waiting to be uploaded...")
		wait6()
		page.MustElementR("[aria-roledescription=file]", "testFile2").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the 130MB file")
		page.MustElement("#close-modal").MustClick()

		// Navigate out of nested folder and delete everything
		page.MustElement("#navigate-back").MustClick()
		page.MustElement("button[aria-roledescription=dropdown]").MustClick()
		page.MustElementR("button", "Delete").MustClick()
		wait7 := page.MustWaitRequestIdle()
		page.MustElementR("button", "Yes").MustClick()
		wait7()

		// Delete created files
		removeFiles([]string{"testFile", "testFile2"})
	})
}
