package download

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileName  string
	extension string
)

func DownloadFile(folder string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http error: ", err)
	}
	defer resp.Body.Close()

	// Status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status code is different from 200")
	}

	// Build file name
	buildFileName(url)

	// Create the folder
	_, err = os.Stat(folder)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(folder, 0755)
		if errDir != nil {
			fmt.Println("Folder error: ", err)
		}
	}

	// Create the file
	out, err := os.Create(folder + "\\" + fileName)
	if err != nil {
		fmt.Println("File error: ", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func buildFileName(link string) error {
	fileUrl, err := url.Parse(link)
	if err != nil {
		fmt.Println("Parse error: ", err)
	}

	path := fileUrl.Path

	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]

	extension = filepath.Ext(fileName)

	return err
}
