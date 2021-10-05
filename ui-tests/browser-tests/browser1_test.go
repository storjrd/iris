package uitest

import (
	"os"
	"fmt"
	"testing"

	"github.com/go-rod/rod"
	"github.com/stretchr/testify/require"
)

func TestBrowser(t *testing.T) {
	Browser(t, func (t *testing.T, browser *rod.Browser)  {
		page := browser.MustPage("http://localhost:3000")

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
		require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "Browser component has not loaded and dropzone is not visible")

		// Attempt to create an invalid folder
		page.MustElementR("button", "New Folder").MustClick()
		folderInput := page.MustElement("[placeholder=\"Name of the folder\"]")
		folderInput.MustInput("...")
		require.Equal(t, "true", page.MustElementR("button", "Save Folder").MustProperty("disabled").Str(), "Folder is not disabled on invalid folder name")
		require.Equal(t, "...", folderInput.MustText(), "Folder input does not contain the `...` invalid name")

		// Create a folder
		folderInput.SelectAllText()
		folderInput.MustInput("go-rod-test2")
		page.MustElementR("button", "Save Folder").MustClick()
		require.Equal(t, "go-rod-test2", page.MustElementR("[aria-roledescription=folder]", "go-rod-test2").MustText(), "The valid folder name `go-rod-test2` was not created")

		// Navigate into the folder and make sure the dropzone is visible
		page.MustElementR("[aria-roledescription=folder]", "go-rod-test2").MustClick()
		require.Equal(t, "go-rod-test2", page.MustElement("a[aria-current=\"page\"]").MustText(), "Navigating into the folder `go-rod-test2` has not been successful")
		require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "The folder `go-rod-test2` is not displaying the dropzone")

		// Attempt to create a new folder but cancel
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("Hello World!")
		page.MustElementR("button", "Cancel").MustClick()

		// Add a file into folder and check that dropzone is still visible
		img := generateImage("img.png")
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(img)
		require.Equal(t, " img.png", page.MustElementR("[aria-roledescription=file]", " img.png").MustText(), "The file `img.png` was not uploaded successfully")
		require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "The dropzone is not being displaying after a file upload")

		// Click on the file name
		page.MustElementR("[aria-roledescription=file]", "img.png").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=image-preview]").MustProperty("src").Str(), "img.png", "The modal did not open on file click")

		// Share a file
		page.MustElementR("span", "Share").MustClick()
		require.Equal(t, "Copy Link", page.MustElement("#generateShareLink").MustText(), "The modal share functionality is not working")
		page.MustElement("#close-modal").MustClick()

		// Click on the hamburger and share
		page.MustElement("button[aria-roledescription=dropdown]").MustClick()
		page.MustElementR("button", "Share").MustClick()
		require.Equal(t, "Copy Link", page.MustElement("#btn-copy-link").MustText(), "The dropdown share functionality is not working")
		page.MustElementR("span", "×").MustClick()

		// Click on the hamburger and then details
		page.MustElement("button[aria-roledescription=dropdown]").MustClick()
		page.MustElementR("button", "Details").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=image-preview]").MustProperty("src").Str(), "img.png", "The dropdown details functionality is not working")
		page.MustElement("#close-modal").MustClick()

		// Use the `..` to navigate out of the folder
		page.MustElement("#navigate-back").MustClick()
		require.Equal(t, "go-rod-test", page.MustElement("a[aria-current=\"page\"]").MustText(), "The `..` navigation is not working")

		// Add another folder
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("go-rod-test3")
		page.MustElementR("button", "Save Folder").MustClick()
		require.Equal(t, "go-rod-test3", page.MustElementR("a", "go-rod-test3").MustText(), "go-rod-test3", "The second folder `go-rod-test3` was not created")

		// Add two files
		img2 := generateImage("img2.png")
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(img2)
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(img)
		require.Equal(t, " img2.png", page.MustElementR("[aria-roledescription=file]", "img2.png").MustText(), "The second file `storjcomponents.png` was not uploaded successfully")
		require.Equal(t, " img.png", page.MustElementR("[aria-roledescription=file]", "img.png").MustText(), "The second file `img.png` was not uploaded successfully")

		// Sort folders/files (by name, size, and date)
		require.Equal(t, "go-rod-test2", page.MustElement("table > tbody > tr:nth-child(1) > td").MustText(), "The automatic sorting by name for folders is not working")
		require.Equal(t, " img.png", page.MustElement("table > tbody > tr:nth-child(3) > td").MustText(), "The automatic sorting by name for files is not working")
		page.MustElementR("th", "Name").MustClick()
		require.Equal(t, "go-rod-test3", page.MustElement("table > tbody > tr:nth-child(1) > td").MustText(), "Sorting by name is not working for folders")
		require.Equal(t, " img2.png", page.MustElement("table > tbody > tr:nth-child(3) > td").MustText(), "Sorting by name is not working for files")
		// sort by size and date still left to do

		// Single folder select
		page.MustElement("table > tbody > tr:nth-child(1)").MustClick()
		require.Contains(t, page.MustElement("table > tbody > tr:nth-child(1)").String(), ".selected-row", "The clicked folder row has not been selected properly")

		// Multifolder select **CURRENTLY NOT WORKING**
		// page.Keyboard.Down(input.Shift)
		// page.MustElement("table > tbody > tr:nth-child(2)").MustClick()
		// assertContains(page.MustElement("table > tbody > tr:nth-child(2)").String(), ".selected-row", "The second clicked folder row has not been selected properly")
		// assertContains(page.MustElement("table > tbody > tr:nth-child(1)").String(), ".selected-row", "The clicked folder row has not been selected properly")
		// page.Keyboard.MustUp(input.Shift)
		// could use some tests for command select as well

		// Multifolder unselect
		page.MustElementR("p", "Buckets").MustClick()
		require.Equal(t, "false", fmt.Sprint(page.MustElement("table > tbody > tr:nth-child(1)").MustHas(".selected-row")), "Multiple selected folders were not unselected successfully")

		// Single file select
		page.MustElement("table > tbody > tr:nth-child(3)").MustClick()
		require.Contains(t, page.MustElement("table > tbody > tr:nth-child(3)").String(), ".selected-row", "Single file select is not working properly")

		// Multifile select **CAN'T SIMULATE MULTIPLE FILE SELECT YET**

		// Multifile unselect
		page.MustElementR("p", "Buckets").MustClick()
		require.Equal(t, "false", fmt.Sprint(page.MustElement("table > tbody > tr:nth-child(3)").MustHas(".selected-row")), "Multiple selected files were not unselected successfully")

		// Select file and folders **CAN'T SIMULATE MULTIPLE FILE/FOLDERS SELECT YET**

		// Navigate into folders and use the breadcrumbs to navigate out
		page.MustElementR("a", "go-rod-test3").MustClick()
		page.MustElementR("a[href=\"/app/buckets/go-rod-test/browse/\"]", "go-rod-test").MustClick()
		require.Equal(t, "go-rod-test", page.MustElement("a[aria-current=\"page\"]").MustText(), "Unable to get to root of browser by way of breadcrumbs")

		// Cancel folder deletion by way of hamburger
		page.MustElement("button[aria-roledescription=dropdown]").MustClick()
		page.MustElementR("button", "Delete").MustClick()
		page.MustElementR("button", "No").MustClick()
		require.Equal(t, "go-rod-test3", page.MustElementR("a", "go-rod-test3").MustText(), "Canceling a folder deletion by way of hamburger did not work")

		// Delete a folder by clicking on hamburger
		page.MustElement("button[aria-roledescription=dropdown]").MustClick()
		page.MustElementR("button", "Delete").MustClick()
		page.MustElementR("button", "Yes").MustClick()
		require.Equal(t, "go-rod-test2", page.MustElementR("table > tbody > tr:nth-child(1) > td", "go-rod-test2").MustText(), "Folder deletion by way of hamburger is not working")

		// Cancel folder deletion by way of trashcan
		page.MustElement("tr[scope=\"row\"]").MustClick()
		page.MustElement("#header-delete").MustClick()
		page.MustElementR("button", "No").MustClick()
		require.Equal(t, "go-rod-test2", page.MustElement("a[href=\"/app/buckets/go-rod-test/browse/go-rod-test2/\"]").MustText(), "Folder cancellation by way of trashcan is not working")

		// Delete a folder by selecting and clicking on trashcan
		page.MustElement("tr[scope=\"row\"]").MustClick()
		page.MustElement("#header-delete").MustClick()
		page.MustElementR("button", "Yes").MustClick()
		require.Equal(t, " img2.png", page.MustElementR("table > tbody > tr:nth-child(1) > td", " img2.png").MustText(), "Deleting a folder by way of trashcan is not working")

		// Cancel file deletion by way of hamburger
		page.MustElement("button[aria-roledescription=dropdown]").MustClick()
		page.MustElementR("button", "Delete").MustClick()
		page.MustElementR("button", "No").MustClick()
		require.Equal(t, " img2.png", page.MustElement("table > tbody > tr:nth-child(1) > td").MustText(), "File deletion cancellation by way of hamburger is not working")

		// Delete a file by clicking on the hamburger
		page.MustElement("button[aria-roledescription=dropdown]").MustClick()
		page.MustElementR("button", "Delete").MustClick()
		page.MustElementR("button", "Yes").MustClick()
		require.Equal(t, " img.png", page.MustElementR("table > tbody > tr:nth-child(1) > td", " img.png").MustText(), "File deletion by way of hamburger is not working")

		// Cancel file deletion by way of trashcan
		page.MustElement("tr[scope=\"row\"]").MustClick()
		page.MustElement("#header-delete").MustClick()
		page.MustElementR("button", "No").MustClick()
		require.Equal(t, " img.png", page.MustElement("table > tbody > tr:nth-child(1) > td").MustText(), "File cancellation by way of trashcan is not working")

		// Delete a file by clicking on the row and clicking on the trashcan
		wait := page.MustWaitRequestIdle()
		page.MustElement("tr[scope=\"row\"]").MustClick()
		page.MustElement("#header-delete").MustClick()
		page.MustElementR("button", "Yes").MustClick()
		require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "Dropzone is not visible on last file deletion")
		wait()

		// Delete multiple folders by selection **SELECTION NOT WORKING**

		// Delete multiple files by selection **SELECTION NOT WORKING**

		// Empty out entire folder

		// Delete created files
		removeFiles([]string{"img.png", "img2.png"})
	})
}
