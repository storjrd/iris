package uitest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-rod/rod"
	"github.com/stretchr/testify/require"
)

func TestBrowser2(t *testing.T) {
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

		// Sort by name
		page.MustElementR("th", "Name").MustClick()

		// Attempt to create a folder with spaces
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("   ")
		require.Equal(t, "true", page.MustElementR("button", "Save Folder").MustProperty("disabled").Str(), "Folder is not disabled on invalid folder name with spaces")
		require.Equal(t, "   ", page.MustElement("[placeholder=\"Name of the folder\"]").MustText(), "Folder input does not contain the empty invalid name")
		page.MustElementR("button", "Cancel").MustClick()

		// Create Folder with special characters
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("Свобода")
		page.MustElementR("button", "Save Folder").MustClick()
		require.Equal(t, "Свобода", page.MustElementR("[aria-roledescription=folder]", "Свобода").MustText(), "Folder with special characters was not created")

		// Navigate into folder and create another folder of the same name
		page.MustElementR("[aria-roledescription=folder]", "Свобода").MustClick()
		require.Equal(t, "Свобода", page.MustElement("[href=\"/app/buckets/go-rod-test/browse/Свобода/\"]").MustText(), "Navigating into the folder `Свобода` was unsuccessful")
		require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "The folder `Свобода` is not displaying the dropzone upon being clicked")
		page.MustElementR("button", "New Folder").MustClick()
		page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("Свобода")
		page.MustElementR("button", "Save Folder").MustClick()
		require.Equal(t, "Свобода", page.MustElementR("[aria-roledescription=folder]", "Свобода").MustText(), "Folder with special characters was not created")

		// upload a video
		wait := page.MustWaitRequestIdle()
		movie, _ := filepath.Abs("./testdata/movie.mp4")
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(movie)
		wait()
		require.Equal(t, " movie.mp4", page.MustElementR("[aria-roledescription=file]", " movie.mp4").MustText(), "The file `movie.mp4` was not uploaded successfully")
		require.Equal(t, "1.48 kB", page.MustElementR("[aria-roledescription=file-size]", "1.48 kB").MustText(), "The size of the file `movie.mp4` is not displaying")
		page.MustElement("[aria-roledescription=file-upload-date]")
		page.MustElementR("[aria-roledescription=file]", " movie.mp4").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=video-preview]").MustProperty("src").Str(), "movie.mp4", "The modal did not open on video file click")
		page.MustElement("#close-modal").MustClick()

		// Upload an audio file
		wait2 := page.MustWaitRequestIdle()
		audio, _ := filepath.Abs("./testdata/audio.mp3")
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(audio)
		wait2()
		require.Equal(t, " audio.mp3", page.MustElementR("[aria-roledescription=file]", " audio.mp3").MustText(), "The file `audio.mp3` was not uploaded successfully")
		page.MustElementR("[aria-roledescription=file]", " audio.mp3").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=audio-preview]").MustProperty("src").Str(), "audio.mp3", "The modal did not open on video file click")
		page.MustElement("#close-modal").MustClick()

		// // Drag and drop folder with duplicate names ***No implementation has worked yet**
		// // storjlogo2, _ := filepath.Abs("./testdata/")
		// // mouse := page.Mouse
		// // mouse.MustMove(3, 3)
		// // mouse.MustDown("left")
		// // mouse.MustUp("left")

		// Upload a folder (folder upload doesn't work when headless)
		generateByteFile("testFolder/testFolderFile1", 10)
		generateByteFile("testFolder/testFolderFile2", 10)
		testFolder, _ := filepath.Abs("./testdata/testFolder")

		if os.Getenv("STORJ_TEST_SHOW_BROWSER") == "" {
			// Create folder
			page.MustElementR("button", "New Folder").MustClick()
			page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("testFolder")
			page.MustElementR("button", "Save Folder").MustClick()
			require.Equal(t, "testFolder", page.MustElementR("[aria-roledescription=folder]", "testFolder").MustText(), "Folder created to simulate upload was not created")

			// Navigate into folder and upload file
			page.MustElementR("[aria-roledescription=folder]", "testFolder").MustClick()
			require.Equal(t, "testFolder", page.MustElement("[href=\"/app/buckets/go-rod-test/browse/Свобода/testFolder/\"]").MustText(), "Navigating into the folder `testFolder` was unsuccessful")
			require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "The folder `testFolder` is not displaying the dropzone upon being clicked")

			// Attempt to create a folder with spaces
			page.MustElementR("button", "New Folder").MustClick()
			page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("   ")
			require.Equal(t, "true", page.MustElementR("button", "Save Folder").MustProperty("disabled").Str(), "Folder is not disabled on invalid folder name with spaces")
			require.Equal(t, "   ", page.MustElement("[placeholder=\"Name of the folder\"]").MustText(), "Folder input does not contain the empty invalid name")
			page.MustElementR("button", "Cancel").MustClick()

			wait3 := page.MustWaitRequestIdle()
			testFolderFile1, _ := filepath.Abs("./testdata/testFolder/testFolderFile1")
			page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testFolderFile1)
			wait3()
			require.Equal(t, " testFolderFile1", page.MustElementR("[aria-roledescription=file]", " testFolderFile1").MustText(), "The file `testFolderFile1` was not uploaded successfully")
			page.MustElementR("[aria-roledescription=file]", "testFolderFile1").MustClick()
			require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the testFolderFile1 file")
		} else {
			wait4 := page.MustWaitRequestIdle()
			page.MustElement("input[aria-roledescription=folder-upload]").MustSetFiles(testFolder)
			require.Equal(t, "2 files waiting to be uploaded...", page.MustElementR("[aria-roledescription=files-uploading-count]", "2 files waiting to be uploaded...").MustText(), "The folder file upload countdown is not being shown")
			require.Equal(t, "1 file waiting to be uploaded...", page.MustElementR("[aria-roledescription=files-uploading-count]", "1 file waiting to be uploaded...").MustText(), "The folder is not uploading files one-by-one")
			wait4()
			require.Equal(t, "..", page.MustElement("table > tbody > tr:nth-child(1) > td > a > a").MustText(), "Folder has not finished uploading")
			require.Equal(t, "testFolder", page.MustElementR("[aria-roledescription=folder]", "testFolder").MustText(), "The folder `testFolder` was not uploaded successfully")
			page.MustElementR("[aria-roledescription=folder]", "testFolder").MustClick()
			page.MustElementR("[aria-roledescription=file]", "testFolderFile1testFolder...").MustClick()
			require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The uploaded folder did not upload the files correctly")
		}

		page.MustElement("#close-modal").MustClick()
		page.MustElement("#navigate-back").MustClick()

		// Upload duplicate folder (folder upload doesn't work when headless)
		if os.Getenv("STORJ_TEST_SHOW_BROWSER") == "" {
			// Create folder
			page.MustElementR("button", "New Folder").MustClick()
			page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("testFolder (1)")
			page.MustElementR("button", "Save Folder").MustClick()
			require.Equal(t, "testFolder (1)", page.MustElementR("[aria-roledescription=folder]", "testFolder \\(1\\)").MustText(), "Duplicate folder created to simulate upload was not created")

			// Navigate into folder and upload file
			page.MustElementR("[aria-roledescription=folder]", "testFolder \\(1\\)").MustClick()
			require.Equal(t, "testFolder (1)", page.MustElement("[href=\"/app/buckets/go-rod-test/browse/Свобода/testFolder (1)/\"]").MustText(), "Navigating into the folder `testFolder (1)` was unsuccessful")
			require.Equal(t, "Drop Files Here to Upload", page.MustElementR("p", "Drop Files Here to Upload").MustText(), "The folder `testFolder (1)` is not displaying the dropzone upon being clicked")

			// Attempt to create a folder with spaces
			page.MustElementR("button", "New Folder").MustClick()
			page.MustElement("[placeholder=\"Name of the folder\"]").MustInput("   ")
			require.Equal(t, "true", page.MustElementR("button", "Save Folder").MustProperty("disabled").Str(), "Folder is not disabled on invalid folder name with spaces")
			require.Equal(t, "   ", page.MustElement("[placeholder=\"Name of the folder\"]").MustText(), "Folder input does not contain the empty invalid name")
			page.MustElementR("button", "Cancel").MustClick()

			wait5 := page.MustWaitRequestIdle()
			testFolderFile2, _ := filepath.Abs("./testdata/testFolder/testFolderFile2")
			page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testFolderFile2)
			wait5()
			require.Equal(t, " testFolderFile2", page.MustElementR("[aria-roledescription=file]", " testFolderFile2").MustText(), "The file `testFolderFile2` was not uploaded successfully")
			page.MustElementR("[aria-roledescription=file]", "testFolderFile2").MustClick()
			require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the testFolderFile2 file")
		} else {
			wait6 := page.MustWaitRequestIdle()
			page.MustElement("input[aria-roledescription=folder-upload]").MustSetFiles(testFolder)
			require.Equal(t, "2 files waiting to be uploaded...", page.MustElementR("[aria-roledescription=files-uploading-count]", "2 files waiting to be uploaded...").MustText(), "The duplicate folder file upload countdown is not being shown")
			require.Equal(t, "1 file waiting to be uploaded...", page.MustElementR("[aria-roledescription=files-uploading-count]", "1 file waiting to be uploaded...").MustText(), "The duplicate folder is not uploading files one-by-one")
			wait6()
			require.Equal(t, "..", page.MustElement("table > tbody > tr:nth-child(1) > td > a > a").MustText(), "Duplicate folder has not finished uploading")
			require.Equal(t, "testFolder (1)", page.MustElementR("[aria-roledescription=folder]", "testFolder \\(1\\)").MustText(), "The duplicate folder `testFolder` was not uploaded successfully")
			page.MustElementR("[aria-roledescription=folder]", "testFolder \\(1\\)").MustClick()
			page.MustElementR("[aria-roledescription=file]", "testFolderFile1testFolder...").MustClick()
			require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The uploaded folder did not upload the files correctly")
		}

		page.MustElement("#close-modal").MustClick()
		page.MustElement("#navigate-back").MustClick()

		// Upload a 0 byte file
		zeroByte, _ := filepath.Abs("./testdata/test0bytes.txt")
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(zeroByte)
		require.Equal(t, " test0bytes.txt", page.MustElementR("[aria-roledescription=file]", " test0bytes.txt").MustText(), "The file 0 byte file was not uploaded successfully")
		page.MustElementR("[aria-roledescription=file]", "test0bytes.txt").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the 0 byte file")
		page.MustElement("#close-modal").MustClick()

		// Upload duplicate 0 byte file
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(zeroByte)
		require.Equal(t, " test0bytes (1).txt", page.MustElementR("[aria-roledescription=file]", "test0bytes \\(1\\).txt").MustText(), "The duplicate file was not uploaded successfully")
		page.MustElementR("[aria-roledescription=file]", "test0bytes \\(1\\).txt").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the duplicate file")
		page.MustElement("#close-modal").MustClick()

		// Upload a 50 MB file
		testFile := generateByteFile("testFile", 50)
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testFile)
		require.Equal(t, " testFile", page.MustElementR("[aria-roledescription=file]", " testFile").MustText(), "The 50 MB file was not uploaded successfully")
		page.MustElementR("[aria-roledescription=file]", "testFile").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the 50 MB file")
		page.MustElement("#close-modal").MustClick()

		// Attempt to upload a large file and cancel the upload after a few segments have been uploaded successfully
		testFile2 := generateByteFile("testFile2", 130)
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testFile2)
		require.Equal(t, " Свобода/testFile2", page.MustElement("[aria-roledescription=file-uploading]").MustText(), "The testFile2 app has not started uploading")
		require.Equal(t, "1 file waiting to be uploaded...", page.MustElementR("[aria-roledescription=files-uploading-count]", "1 file waiting to be uploaded...").MustText(), "The testFile2 app upload is not displaying the files waiting to be uploaded countdown")
		page.MustElementR("[aria-roledescription=progress-bar]", "1")
		wait7 := page.MustWaitRequestIdle()
		page.MustElementR("button", "Cancel").MustClick()
		wait7()
		require.Equal(t, " testFile", page.MustElement("table > tbody > tr:nth-child(5) > td").MustText(), "The testFile2 app was not deleted on cancel upload")
		page.MustElementR("[aria-roledescription=file]", "testFile").MustClick()
		page.MustElement("#close-modal").MustClick()

		// Upload a 130MB file
		wait8 := page.MustWaitRequestIdle()
		page.MustElement("input[aria-roledescription=file-upload]").MustSetFiles(testFile2)
		wait8()
		require.Equal(t, " testFile2", page.MustElementR("[aria-roledescription=file]", " testFile2").MustText(), "The 130 MB file was not uploaded successfully")
		page.MustElementR("[aria-roledescription=file]", "testFile2").MustClick()
		require.Contains(t, page.MustElement("[aria-roledescription=preview-placeholder]").String(), "svg", "The modal did not open upon clicking the 130MB file")
		page.MustElement("#close-modal").MustClick()

		// Navigate out of nested folder and delete everything
		page.MustElement("#navigate-back").MustClick()
		page.MustElement("button[aria-roledescription=dropdown]").MustClick()
		page.MustElementR("button", "Delete").MustClick()
		wait9 := page.MustWaitRequestIdle()
		page.MustElementR("button", "Yes").MustClick()
		wait9()

		// Delete created files
		removeFiles([]string{"testFile", "testFile2", "testFolder/testFolderFile1", "testFolder/testFolderFile2"})
	})
}
