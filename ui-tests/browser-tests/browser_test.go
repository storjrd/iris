package uitest

import (
	"os"
	// "fmt"
	"testing"
	"path/filepath"

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

		// // Verify that browser is empty
		// require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "Browser is not empty on login")

		// // Attempt to create an invalid folder
		// page.MustElementR("button", "New Folder").MustClick()
		// folderInput := page.MustElement("[placeholder=\"Name of the folder\"]")
		// folderInput.MustInput("...")
		// require.Equal(t, "true", page.MustElementR("button", "Save Folder").MustProperty("disabled").Str(), "Folder is not disabled on invalid folder name")
		// require.Equal(t, "...", folderInput.MustText(), "Folder input does not contain the `...` invalid name")

		// // Create a folder
		// folderInput.SelectAllText()
		// folderInput.MustInput("go-rod-test2")
		// page.MustElementR("button", "Save Folder").MustClick()
		// require.Equal(t, "go-rod-test2", page.MustElementR("a", "go-rod-test2").MustText(), "The valid folder name `go-rod-test2` was not created")

		// // Navigate into the folder and make sure it's empty
		// page.MustElementR("a", "go-rod-test2").MustClick()
		// require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "The folder `go-rod-test2` is not empty upon creation")

		// // Add a file and immediately stop the upload **Currently not working because the file gets uploaded too quickly**
		// // storjBanner, _ := filepath.Abs("./input/sto-1.jpeg")
		// // page.MustElement("input[type=file]").SetFiles([]string{storjBanner})
		// // page.MustElementR("button", "cancel").MustClick()
		// // assertEquals(page.MustElementR("span", " sto-1.jpeg").MustText(), " storjlogo.jpeg", "Cancelling a file upload did not work properly.")

		// // Add a file into folder
		// storjLogo, _ := filepath.Abs("./input/storjlogo.jpeg")
		// page.MustElement("input[type=file]").MustSetFiles(storjLogo)
		// require.Equal(t, " storjlogo.jpeg", page.MustElementR("span.file-name", " storjlogo.jpeg").MustText(), "The file `storjlog.jpeg` was not uploaded successfully")

		// // Click on the file name
		// page.MustElementR("span", "storjlogo.jpeg").MustClick()
		// require.Contains(t, page.MustElement("img.preview.img-fluid").MustProperty("src").Str(), "storjlogo.jpeg", "The modal did not open on file click")

		// // Share a file
		// page.MustElementR("span", "Share").MustClick()
		// require.Equal(t, "Copy Link", page.MustElement("#generateShareLink").MustText(), "The modal share functionality is not working")
		// page.MustElement("#close-modal").MustClick()

		// // Click on the hamburger and share
		// page.MustElement("button[aria-haspopup=\"true\"]").MustClick()
		// page.MustElementR("button", "Share").MustClick()
		// require.Equal(t, "Copy Link", page.MustElement("#btn-copy-link").MustText(), "The dropdown share functionality is not working")
		// page.MustElementR("span", "×").MustClick()

		// // Click on the hamburger and then details
		// page.MustElement("button[aria-haspopup=\"true\"]").MustClick()
		// page.MustElementR("button", "Details").MustClick()
		// require.Contains(t, page.MustElement("img.preview.img-fluid").MustProperty("src").Str(), "storjlogo.jpeg", "The dropdown details functionality is not working")
		// page.MustElement("#close-modal").MustClick()

		// // Use the `..` to navigate out of the folder
		// page.MustElement("#navigate-back").MustClick()
		// require.Equal(t, "go-rod-test", page.MustElement("a[aria-current=\"page\"]").MustText(), "The `..` navigation is not working")

		// // Add another folder
		// page.MustElementR("button", "New Folder").MustClick()
		// page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("go-rod-test3")
		// page.MustElementR("button", "Save Folder").MustClick()
		// require.Equal(t, "go-rod-test3", page.MustElementR("a", "go-rod-test3").MustText(), "go-rod-test3", "The second folder `go-rod-test3` was not created")

		// // Add two files
		// storjComponents, _ := filepath.Abs("./input/storjcomponents.png")
		// page.MustElement("input[type=file]").MustSetFiles(storjLogo)
		// page.MustElement("input[type=file]").MustSetFiles(storjComponents)
		// require.Equal(t, " storjcomponents.png", page.MustElementR("span.file-name", "storjcomponents.png").MustText(), "The second file `storjcomponents.png` was not uploaded successfully")
		// require.Equal(t, " storjlogo.jpeg", page.MustElementR("span.file-name", "storjlogo.jpeg").MustText(), "The second file `storjlogo.jpeg` was not uploaded successfully")

		// // Add a duplicate file and check naming convetion **Not working yet**
		// // page.MustElement("input[type=file]").MustSetFiles(storjComponents)
		// // require.Equal(t, " storjcomponents (1).png", page.MustElementR("span.file-name", "storjcomponents (1).png").MustText(), "The duplicate file `storjcomponents (1).png (1).jpeg` was not uploaded successfully")
		// // fmt.Printf("After duplicate file upload")

		// // Sort folders/files (by name, size, and date)
		// require.Equal(t, "go-rod-test2", page.MustElement("table > tbody > tr:nth-child(1) > td").MustText(), "The automatic sorting by name for folders is not working")
		// require.Equal(t, " storjcomponents.png", page.MustElement("table > tbody > tr:nth-child(3) > td").MustText(), "The automatic sorting by name for files is not working")
		// page.MustElementR("th", "Name").MustClick()
		// require.Equal(t, "go-rod-test3", page.MustElement("table > tbody > tr:nth-child(1) > td").MustText(), "Sorting by name is not working for folders")
		// require.Equal(t, " storjlogo.jpeg", page.MustElement("table > tbody > tr:nth-child(3) > td").MustText(), "Sorting by name is not working for files")
		// // sort by size and date still left to do

		// // Single folder select
		// page.MustElement("table > tbody > tr:nth-child(1)").MustClick()
		// require.Contains(t, page.MustElement("table > tbody > tr:nth-child(1)").String(), ".selected-row", "The clicked folder row has not been selected properly")

		// // Multifolder select **CURRENTLY NOT WORKING**
		// // page.Keyboard.Down(input.Shift)
		// // page.MustElement("table > tbody > tr:nth-child(2)").MustClick()
		// // assertContains(page.MustElement("table > tbody > tr:nth-child(2)").String(), ".selected-row", "The second clicked folder row has not been selected properly")
		// // assertContains(page.MustElement("table > tbody > tr:nth-child(1)").String(), ".selected-row", "The clicked folder row has not been selected properly")
		// // page.Keyboard.MustUp(input.Shift)
		// // could use some tests for command select as well

		// // Multifolder unselect
		// page.MustElementR("p", "Buckets").MustClick()
		// require.Equal(t, "false", fmt.Sprint(page.MustElement("table > tbody > tr:nth-child(1)").MustHas(".selected-row")), "Multiple selected folders were not unselected successfully")

		// // Single file select
		// page.MustElement("table > tbody > tr:nth-child(3)").MustClick()
		// require.Contains(t, page.MustElement("table > tbody > tr:nth-child(3)").String(), ".selected-row", "Single file select is not working properly")

		// // Multifile select **CAN'T SIMULATE MULTIPLE FILE SELECT YET**


		// // Multifile unselect
		// page.MustElementR("p", "Buckets").MustClick()
		// require.Equal(t, "false", fmt.Sprint(page.MustElement("table > tbody > tr:nth-child(3)").MustHas(".selected-row")), "Multiple selected files were not unselected successfully")

		// // Select file and folders **CAN'T SIMULATE MULTIPLE FILE/FOLDERS SELECT YET**


		// // Navigate into folders and use the breadcrumbs to navigate out
		// page.MustElementR("a", "go-rod-test3").MustClick()
		// page.MustElementR("a[href=\"/app/buckets/go-rod-test/browse/\"]", "go-rod-test").MustClick()
		// require.Equal(t, "go-rod-test", page.MustElement("a[aria-current=\"page\"]").MustText(), "Unable to get to root of browser by way of breadcrumbs")

		// // Cancel folder deletion by way of hamburger
		// page.MustElement("button[aria-haspopup=\"true\"]").MustClick()
		// page.MustElementR("button", "Delete").MustClick()
		// page.MustElementR("button", "No").MustClick()
		// require.Equal(t, "go-rod-test3", page.MustElementR("a", "go-rod-test3").MustText(), "Canceling a folder deletion by way of hamburger did not work")


		// // Delete a folder by clicking on hamburger
		// page.MustElement("button[aria-haspopup=\"true\"]").MustClick()
		// page.MustElementR("button", "Delete").MustClick()
		// page.MustElementR("button", "Yes").MustClick()
		// require.Equal(t, "go-rod-test2", page.MustElementR("table > tbody > tr:nth-child(1) > td", "go-rod-test2").MustText(), "Folder deletion by way of hamburger is not working")


		// // Cancel folder deletion by way of trashcan
		// page.MustElement("tr[scope=\"row\"]").MustClick()
		// page.MustElement("#header-delete").MustClick()
		// page.MustElementR("button", "No").MustClick()
		// require.Equal(t, "go-rod-test2", page.MustElement("a[href=\"/app/buckets/go-rod-test/browse/go-rod-test2/\"]").MustText(), "Folder cancellation by way of trashcan is not working")


		// // Delete a folder by selecting and clicking on trashcan
		// page.MustElement("tr[scope=\"row\"]").MustClick()
		// page.MustElement("#header-delete").MustClick()
		// page.MustElementR("button", "Yes").MustClick()
		// require.Equal(t, " storjlogo.jpeg", page.MustElementR("table > tbody > tr:nth-child(1) > td > span", " storjlogo.jpeg").MustText(), "Deleting a folder by way of trashcan is not working")


		// // Cancel file deletion by way of hamburger
		// page.MustElement("button[aria-haspopup=\"true\"]").MustClick()
		// page.MustElementR("button", "Delete").MustClick()
		// page.MustElementR("button", "No").MustClick()
		// require.Equal(t, " storjlogo.jpeg", page.MustElement("table > tbody > tr:nth-child(1) > td > span").MustText(), "File deletion cancellation by way of hamburger is not working")


		// // Delete a file by clicking on the hamburger
		// page.MustElement("button[aria-haspopup=\"true\"]").MustClick()
		// page.MustElementR("button", "Delete").MustClick()
		// page.MustElementR("button", "Yes").MustClick()
		// require.Equal(t, " storjcomponents.png", page.MustElementR("table > tbody > tr:nth-child(1) > td > span", " storjcomponents.png").MustText(), "File deletion by way of hamburger is not working")

		// // Cancel file deletion by way of trashcan
		// page.MustElement("tr[scope=\"row\"]").MustClick()
		// page.MustElement("#header-delete").MustClick()
		// page.MustElementR("button", "No").MustClick()
		// require.Equal(t, " storjcomponents.png", page.MustElement("table > tbody > tr:nth-child(1) > td > span").MustText(), "File cancellation by way of trashcan is not working")

		// // Delete a file by clicking on the row and clicking on the trashcan
		// page.MustElement("tr[scope=\"row\"]").MustClick()
		// page.MustElement("#header-delete").MustClick()
		// page.MustElementR("button", "Yes").MustClick()
		// require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "Browser is not empty on last file deletion")

		// // Delete multiple folders by selection **SELECTION NOT WORKING**


		// // Delete multiple files by selection **SELECTION NOT WORKING**


		// // Empty out entire folder


		// Create Folder with special characters
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("Свобода")
		page.MustElementR("button", "Save Folder").MustClick()
		require.Equal(t, "Свобода", page.MustElementR("a.file-name", "Свобода").MustText(), "Folder with special characters was not created")

		// Navigate into folder and create another folder of the same name
		page.MustElementR("a.file-name", "Свобода").MustClick()
		require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "The folder `Свобода` is not empty upon creation")
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("Свобода")
		page.MustElementR("button", "Save Folder").MustClick()
		require.Equal(t, "Свобода", page.MustElementR("a.file-name", "Свобода").MustText(), "Folder with special characters was not created")

		// Navigate into nested folder and upload a video
		page.MustElementR("a.file-name", "Свобода").MustClick()
		require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "The folder `Свобода` is not empty upon creation")
		TardigradeExplainerVideo, _ := filepath.Abs("./input/TardigradeExplainerVideo.m4v")
		page.MustElement("input[type=file]").MustSetFiles(TardigradeExplainerVideo)
		require.Equal(t, " TardigradeExplainerVideo....", page.MustElementR("span.file-name", " TardigradeExplainerVideo....").MustText(), "The file `TardigradeExplainerVideo.m4v` was not uploaded successfully")
		page.MustElementR("span", " TardigradeExplainerVideo....").MustClick()
		require.Contains(t, page.MustElement("video.preview").MustProperty("src").Str(), "TardigradeExplainerVideo.m4v", "The modal did not open on video file click")
		page.MustElement("#close-modal").MustClick()

		// Upload an audio file
		sampleAudio, _ := filepath.Abs("./input/sampleAudio.mp3")
		page.MustElement("input[type=file]").MustSetFiles(sampleAudio)
		require.Equal(t, " sampleAudio.mp3", page.MustElementR("span.file-name", " sampleAudio.mp3").MustText(), "The file `sampleAudio.mp3` was not uploaded successfully")
		page.MustElementR("span", "sampleAudio.mp3").MustClick()
		require.Contains(t, page.MustElement("audio.preview").MustProperty("src").Str(), "sampleAudio.mp3", "The modal did not open on video file click")
		page.MustElement("#close-modal").MustClick()

		// Drag and drop folder with duplicate names **Haven't gotten it to work yet**
		// storjlogo2, _ := filepath.Abs("./input/storjlogo.jpeg")
		// mouse := page.Mouse
		// mouse.MustMove(3, 3)
		// mouse.MustDown("left")
		// mouse.MustUp("left")

		// Upload a folder
		dragAndDropFolder, _ := filepath.Abs("./input/dragAndDropFolder")
		page.MustElement("input[type=file]").MustSetFiles(dragAndDropFolder)
		require.Equal(t, " dragAndDropFolder", page.MustElementR("span.file-name", " dragAndDropFolder").MustText(), "The folder `dragAndDropFolder` was not uploaded successfully")
		page.MustElementR("span", "dragAndDropFolder").MustClick()
		require.Contains(t, page.MustElement("audio.preview").MustProperty("src").Str(), "dragAndDropFolder", "The modal did not open on video file click")
		page.MustElement("#close-modal").MustClick()
	})
}
