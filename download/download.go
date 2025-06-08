package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

/* NEXT:
- timeout
	spin up a localhost web server to test this more easily
	example from SO:
	client := http.Client{Timeout: 10 * time.Second}
	client.Get("http://example.com/")
- explore headers, redirects, etc.
*/

func download(url, filePath string) (err error) {
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("can't create output file!")
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("can't fetch data!")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("can't write output!")
		return err
	}
	fmt.Printf("wrote %d bytes to %s\n", n, filePath)
	return nil
}

func main() {

	downloadTasks := []struct {
		url          string
		downloadPath string
		executable   bool
	}{
		{
			downloadPath: "./files/home.html",
			url:          "https://www.joecullin.com/",
		},
		// {
		// 	downloadPath: "./files/404.txt",
		// 	url:          "https://www.joecullin.com/afdasdfasdfdasfadfds",
		// },
		{
			downloadPath: "./files/mac_app",
			url:          "https://www.joecullin.com/go_test/app-amd64-mac",
			executable:   true,
		},
	}

	errorCount, successCount := 0, 0
	for index, task := range downloadTasks {
		fmt.Printf("\ntask %d of %d: downloading %s to %s\n", index+1, len(downloadTasks), task.url, task.downloadPath)
		err := download(task.url, task.downloadPath)
		if err != nil {
			fmt.Printf("ERROR for %s: %s\n", task.url, err)
			errorCount++
			continue
		}
		if task.executable && runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
			fmt.Println("Updating permissions on downloaded file.")
			if err = os.Chmod(task.downloadPath, 0744); err != nil {
				fmt.Printf("ERROR changing permissions on %s: %s\n", task.downloadPath, err)
				errorCount++
				continue
			}
		}

		successCount++
	}
	fmt.Printf("\nDone! %d good, %d bad.\n", successCount, errorCount)
	if errorCount > 0 {
		os.Exit(1)
	}
}
