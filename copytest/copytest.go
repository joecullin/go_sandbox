package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

func move(newPath, oldPath string) (err error) {
	inputFile, err := os.Open(oldPath)
	if err != nil {
		fmt.Println("can't open source file!")
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(newPath)
	if err != nil {
		fmt.Println("can't create destination file!")
		return err
	}
	defer outputFile.Close()

	n, err := io.Copy(outputFile, inputFile)
	if err != nil {
		fmt.Println("can't write output!")
		return err
	}
	fmt.Printf("wrote %d bytes to %s\n", n, newPath)
	return nil
}

func main() {

	doCopy := false
	flag.BoolVar(&doCopy, "copytest", false, "run copy test")

	replaceSelf := false
	flag.BoolVar(&replaceSelf, "replaceself", false, "test copying new version over ourself")

	showVersion := false
	flag.BoolVar(&showVersion, "version", false, "show version")

	repeatVersion := false
	flag.BoolVar(&repeatVersion, "version-repeat", false, "repeatedly show version, every 5 seconds")

	flag.Parse()

	if showVersion {
		fmt.Println("Version=2")
	} else if repeatVersion {
		fmt.Println("printing version every 2 seconds forever...")
		ticker := time.Tick(2 * time.Second)
		for next := range ticker {
			fmt.Printf("Version=2 (tick %v)\n", next)
		}
	} else if replaceSelf {
		selfPath := filepath.Join(".", "bin", "copytest")
		newVersionPath := filepath.Join(".", "bin", "copytestV2")
		err := move(selfPath, newVersionPath)
		if err != nil {
			fmt.Printf("ERROR copying file: %s\n", err)
			os.Exit(1)
		}
		// if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		// 	fmt.Println("Updating permissions on copied file.")
		// 	if err = os.Chmod(selfPath, 0744); err != nil {
		// 		fmt.Printf("ERROR changing permissions on %s: %s\n", selfPath, err)
		// 		os.Exit(1)
		// 	}
		// }
		fmt.Println("done overwriting self!")

		env := os.Environ()
		execErr := syscall.Exec(selfPath, []string{"_", "--version-repeat"}, env)
		if execErr != nil {
			fmt.Println("Error re-starting.", execErr)
		}

		fmt.Println("done re-starting!")

		// cmd :=  exec.Command(selfPath, "--version")
		// var out strings.Builder
		// cmd.Stdout = &out
		// cmdErr := cmd.Run()
		// if cmdErr != nil {
		// 	fmt.Println("Error!", cmdErr)
		// }
		// fmt.Println("output from new command:", out.String())

	} else if doCopy {
		tasks := []struct {
			newPath    string
			oldPath    string
			executable bool
		}{
			{
				oldPath:    filepath.Join(".", "files", "test.txt"),
				newPath:    filepath.Join(".", "files", "test_copy.txt"),
				executable: true,
			},
		}

		errorCount, successCount := 0, 0
		for index, task := range tasks {
			fmt.Printf("\ntask %d of %d: moving %s to %s\n", index+1, len(tasks), task.oldPath, task.newPath)
			err := move(task.newPath, task.oldPath)
			if err != nil {
				fmt.Printf("ERROR for %s: %s\n", task.newPath, err)
				errorCount++
				continue
			}
			if task.executable && runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
				fmt.Println("Updating permissions on downloaded file.")
				if err = os.Chmod(task.newPath, 0744); err != nil {
					fmt.Printf("ERROR changing permissions on %s: %s\n", task.newPath, err)
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
	} else {
		fmt.Printf("unknown mode! Run with -h to see params.")
		os.Exit(1)
	}
}
