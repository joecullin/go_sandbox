package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/* NEXT:
- look into error handling. How can I detect things like a 404 status ?
- try binary file
- set permissions on downloaded file
- timeout
	example from SO:
	client := http.Client{Timeout: 10 * time.Second}
	client.Get("http://example.com/")

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
	}{
		{
			downloadPath: "./files/home.html",
			url:          "https://www.joecullin.com/",
		},
		{
			downloadPath: "./files/404.txt",
			url:          "https://www.joecullin.com/afdasdfasdfdasfadfds",
		},
	}

	errorCount, successCount := 0, 0
	for index, task := range downloadTasks {
		fmt.Printf("\ntask %d of %d: downloading %s to %s\n", index+1, len(downloadTasks), task.url, task.downloadPath)
		err := download(task.url, task.downloadPath)
		if err != nil {
			fmt.Printf("ERROR for %s: %s\n", task.url, err)
			errorCount++
		} else {
			successCount++
		}
	}
	fmt.Printf("\nDone! %d good, %d bad.\n", successCount, errorCount)
	if errorCount > 0 {
		os.Exit(1)
	}
}
