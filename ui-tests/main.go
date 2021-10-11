package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func assertEquals(a string, b string, msg string) {
	if a != b {
		fmt.Printf("Failed on -> %s: \"%v\" != \"%v\"\n", msg, a, b)
		panic("Assertion Error")
	}
}

func assertContains(a string, b string, msg string) {
	if !strings.Contains(a, b) {
		fmt.Printf("Failed on -> %s: %v does not contain %v\n", msg, a, b)
		panic("Assertion Error")
	}
}

func waitToEnd(num int64) {
	time.Sleep(time.Duration(num) * time.Second)
}

func testHome(u string) {
	page := rod.New().ControlURL(u).MustConnect().MustPage("http://ui").Timeout(10 * time.Second)

	// Test hero title
	title := page.MustElement(".hero-title")
	assertEquals(title.MustText(), "Decentralized Cloud Storage is Here", "Home is not being displayed")

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

// 	page.MustWaitLoad().MustScreenshot("/output/login.png")
}

func testBrowser(u string) {
	page := rod.New().ControlURL(u).MustConnect().MustPage("http://ui").Timeout(10 * time.Second)

	// Must switch to login
	page.MustElement("button.btn-success").MustClick()

	// Login
	page.MustElement("#emailAddress").MustInput(os.Getenv("IRIS_EMAIL"))
	page.MustElement("#password").MustInput(os.Getenv("IRIS_PASSWORD"))
	page.MustElementR("button", "Login").MustClick()

	// Navigate into bucket and browser
	page.MustElementR("span", "go-rod-test").MustClick()
	page.MustElement("#passphrase").MustInput("go-rod-test")
	page.MustElementR("button", "Unlock").MustClick()

	// Verify that browser is empty
	dropZone := page.MustElementR("p", "Drop Files Here to Upload")
	assertEquals(dropZone.MustText(), "Drop Files Here to Upload", "Browser is not empty on login")

	// Attempt to create an invalid folder
	page.MustElement("button.btn.btn-light.btn-block").MustClick()
	folderInput := page.MustElement("td > input.form-control.input-folder")
	folderInput.MustInput("...")
	page.MustElement("td > button.btn.btn-primary.btn-sm.px-4").MustAttribute("disabled")
	assertEquals(page.MustElement("td > button.btn.btn-primary.btn-sm.px-4").MustProperty("disabled").Str(), "true", "Folder is not disabled on invalid folder name")
	assertEquals(folderInput.MustText(), "...", "Folder input does not contain the `...` invalid name")

	// Create a folder
	folderInput.SelectAllText()
	folderInput.MustInput("go-rod-test2")
	page.MustElement("td > button.btn.btn-primary.btn-sm.px-4").MustClick()
	assertEquals(page.MustElementR("a", "go-rod-test2").MustText(), "go-rod-test2", "The valid folder name `go-rod-test2` was not created")

	// Navigate into the folder and make sure it's empty
	page.MustElementR("a", "go-rod-test2").MustClick()
	assertEquals(page.MustElement("div.upload-help > p.drop-files-text.mt-4.mb-0").MustText(), "Drop Files Here to Upload", "The folder `go-rod-test2` is not empty upon creation")

	// Add a file and immediately stop the upload

	// Add a file into folder
	storjLogo, _ := filepath.Abs("./input/storjlogo.jpeg")
	page.MustElement("input[type=file]").SetFiles([]string{storjLogo})

	// Check to see if file was uploaded
	assertEquals(page.MustElementR("span", " storjlogo.jpeg").MustText(), " storjlogo.jpeg", "The file `storjlog.jpeg` was not uploaded successfully")

	// Click on the file name
	page.MustElementR("span", "storjlogo.jpeg").MustClick()
	assertContains(page.MustElement("img.preview.img-fluid").MustProperty("src").Str(), "storjlogo.jpeg", "The modal did not open on file click")

	// Share a file
	page.MustElementR("span", "Share").MustClick()
	assertEquals(page.MustElement("#generateShareLink").MustText(), "Copy Link", "The modal share functionality is not working")
	page.MustElement("div.col-6.col-lg-4.pr-5 > div.text-right > svg").MustClick()

	// Click on the hamburger and share
	page.MustElement("div.dropleft > button.btn.btn-white.btn-actions").MustClick()
	page.MustElementR("button.dropdown-item.action.p-3", "Share").MustClick()
	assertEquals(page.MustElement("#btn-copy-link").MustText(), "Copy Link", "The dropdown share functionality is not working")
	page.MustScreenshot("./output/shareModal.png")
	page.MustElement("#shareModal > div > div > div.modal-header.border-0 > button > span").MustClick()

	// Click on the hamburger and then details
	page.MustElement("div.dropleft > button.btn.btn-white.btn-actions").MustClick()
	page.MustElementR("button.dropdown-item.action.p-3", "Details").MustClick()
	assertContains(page.MustElement("img.preview.img-fluid").MustProperty("src").Str(), "storjlogo.jpeg", "The dropdown details functionality is not working")
	page.MustElement("div.col-6.col-lg-4.pr-5 > div.text-right > svg").MustClick()

	// Use the `..` to navigate out of the folder
	page.MustElement("td > a > a.px-2.font-weight-bold").MustClick()
	assertEquals(page.MustElement("span > a > a.file-name").MustText(), "go-rod-test2", "The `..` navigation is not working")

	// Add another folder
	page.MustElement("button.btn.btn-light.btn-block").MustClick()
	page.MustElement("td > input.form-control.input-folder").MustInput("go-rod-test3")
	page.MustElement("td > button.btn.btn-primary.btn-sm.px-4").MustClick()
	assertEquals(page.MustElementR("a", "go-rod-test3").MustText(), "go-rod-test3", "The second folder `go-rod-test3` was not created")

	// Add two files
	storjComponents, _ := filepath.Abs("./input/storjcomponents.png")
	page.MustElement("input[type=file]").SetFiles([]string{storjLogo})
	page.MustElement("input[type=file]").SetFiles([]string{storjComponents})
	assertEquals(page.MustElementR("span", " storjcomponents.png").MustText(), " storjcomponents.png", "The second file `storjcomponents.png` was not uploaded successfully")

	// Sort folders/files (by name, size, and date)
	assertEquals(page.MustElement("table > tbody > tr:nth-child(1) > td.w-50 > span > span > a > a").MustText(), "go-rod-test2", "The automatic sorting by name for folders is not working")
	assertEquals(page.MustElement("table > tbody > tr:nth-child(3) > td.w-50 > span").MustText(), " storjcomponents.png", "The automatic sorting by name for files is not working")
	page.MustElement("span > a.arrow > svg").MustClick()
	assertEquals(page.MustElement("table > tbody > tr:nth-child(1) > td.w-50 > span > span > a > a").MustText(), "go-rod-test3", "Sorting by name is not working for folders")
	assertEquals(page.MustElement("table > tbody > tr:nth-child(3) > td.w-50 > span").MustText(), " storjlogo.jpeg", "Sorting by name is not working for files")
	// sort by size and date still left to do

	// Navigate into folders and use the breadcrumbs to navigate out
	page.MustElementR("a", "go-rod-test3").MustClick()
	page.MustElement("div.d-inline > a > a.path").MustClick()
	assertEquals(page.MustElement("table > tbody > tr:nth-child(1) > td.w-50 > span > span > a > a").MustText(), "go-rod-test3", "Unable to get to root of browser by way of breadcrumbs")

	// Single folder select
	page.MustElement("table > tbody > tr:nth-child(1)").MustClick()
	assertContains(page.MustElement("table > tbody > tr:nth-child(1)").String(), ".selected-row", "The clicked folder row has not been selected properly")

	// Multifolder select **CURRENTLY NOT WORKING**
	// page.Keyboard.Down(input.Shift)
	// page.MustElement("table > tbody > tr:nth-child(2)").MustClick()
	// assertContains(page.MustElement("table > tbody > tr:nth-child(2)").String(), ".selected-row", "The second clicked folder row has not been selected properly")
	// assertContains(page.MustElement("table > tbody > tr:nth-child(1)").String(), ".selected-row", "The clicked folder row has not been selected properly")
	// page.Keyboard.MustUp(input.Shift)
	// could use some tests for command select as well

	// Multifolder unselect
	page.MustElementR("p", "Buckets").MustClick()
	secondFolderSelectionStatus, _, _ := page.MustElement("table > tbody > tr:nth-child(1)").Has(".selected-row")
	assertEquals(fmt.Sprint(secondFolderSelectionStatus), "false", "Multiple selected folders were not unselected successfully")

	// Single file select
	page.MustElement("table > tbody > tr:nth-child(3)").MustClick()
	assertContains(page.MustElement("table > tbody > tr:nth-child(3)").String(), ".selected-row", "Single file select is not working properly")

	// Multifile select


	// Multifile unselect
	page.MustElementR("p", "Buckets").MustClick()
	firstFileSelectionStatus, _, _ := page.MustElement("table > tbody > tr:nth-child(3)").Has(".selected-row")
	assertEquals(fmt.Sprint(firstFileSelectionStatus), "false", "Multiple selected files were not unselected successfully")

	// Select file and folders


	// Delete a folder by clicking on hamburger
	page.MustElement("table > tbody > tr:nth-child(1) > td.text-right > div > div > button").MustClick()
	page.MustElement("table > tbody > tr:nth-child(1) > td.text-right > div > div > div > button").MustClick()
	page.MustElement("table > tbody > tr:nth-child(1) > td.text-right > div > div > div > div > div > button.dropdown-item.trash.p-3.action").MustClick()
	waitToEnd(4)
	assertEquals(page.MustElement("table > tbody > tr:nth-child(1) > td.w-50 > span > span > a > a").MustText(), "go-rod-test2", "Folder deletion by way of hamburger is not working")

	// Cancel folder deletion by way of hamburger


	// Delete a folder by selecting and clicking on trashcan


	// Cancel folder deletion by way of hamburger


	// Add a duplicate file and check naming convetion
	page.MustElement("input[type=file]").SetFiles([]string{storjLogo})
	// assertEquals(page.MustElementR("span", "storjlogo (1).jpeg").MustText(), " storjlogo (1).jpeg", "The duplicate file `storjlogo (1).jpeg` was not uploaded successfully")

	// Delete a file by clicking on the hamburger


	// Cancel file deletion by way of hamburger


	// Delete a file by clicking on the row and clicking on the trashcan


	// Cancel file deletion by way of hamburger


	// Delete multiple folders by selection


	// Delete multiple files by selection


	// Empty out entire folder

}

func main() {
	u := launcher.New().Bin("/usr/bin/google-chrome").MustLaunch()

	testHome(u)

	testLogin(u)

	testBrowser(u)
}
