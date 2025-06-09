package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func move(destPath, sourcePath string) (err error) {

	// Windows gives "Access Denied" if we try to overwrite our the currently running exe.
	// But, it lets us do it in two steps:
	//   1. Rename the in-use file to ".bak"
	//   2. Copy the new file to the current exe.
	// I guess it's nice to have the ".bak" file too, for rollbacks in case of failure?
	// We can do the same on linux & mac for consistency. (We'll have to fix perms though.)

	// detect old file's current permissions before we rename it
	var perms os.FileMode
	perms = 0644
	if runtime.GOOS != "windows" {
		if fileInfo, err := os.Lstat(destPath); err != nil {
			fmt.Printf("can't get current permissions for dest file %s: %v\n", destPath, err)
			// (Keep going. Maybe it's a new file.)
		} else {
			perms = fileInfo.Mode().Perm()
			fmt.Printf("current permissions of dest file %s: %#o\n", destPath, perms)
		}
	}

	backupPath := destPath + ".bak"
	fmt.Printf("backing up file %s to %s!\n", destPath, backupPath)
	if err := os.Rename(destPath, backupPath); err != nil {
		fmt.Printf("can't back up file %s to %s!\n", destPath, backupPath)
		return err
	}

	inputFile, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("can't open source file!")
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
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
	fmt.Printf("copied %s to %s. wrote %d bytes.\n", sourcePath, destPath, n)

	if runtime.GOOS != "windows" {
		fmt.Println("Updating permissions on copied file.")
		if err = os.Chmod(destPath, perms); err != nil {
			fmt.Printf("ERROR changing permissions on %s: %s\n", destPath, err)
			os.Exit(1)
		}
	}

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
		selfName, sourceName := "copytest", "copytestV2"
		if runtime.GOOS == "windows" {
			selfName = selfName + ".exe"
			sourceName = sourceName + ".exe"
		}
		selfPath := filepath.Join(".", "bin", selfName)
		newVersionPath := filepath.Join(".", "bin", sourceName)
		if err := move(selfPath, newVersionPath); err != nil {
			fmt.Printf("ERROR copying file: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Done overwriting self! Now trying to restart...")

		// selfPath += "UMMMMM" // for testing failure

		restartParams := []string{"--version-repeat"}
		if runtime.GOOS != "windows" {
			// Exec is nicer when it's available.
			// (Replaces current process, so things like job control in shell still work.)
			params := append([]string{"_"}, restartParams...)
			env := os.Environ()
			if err := syscall.Exec(selfPath, params, env); err != nil {
				fmt.Printf("Error re-starting %s with params %v: %e\n", selfPath, params, err)
			}
			fmt.Println("done re-starting!")
		} else {
			// Windows: no syscall.exec.
			// So we'll start a new process then exit this one.
			cmd := exec.Command(selfPath, strings.Join(restartParams, " "))
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				fmt.Println("Error starting new process!", err)
			} else {
				fmt.Println("Started new process! Exiting this one")
				os.Exit(0)
			}
		}

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
			fmt.Printf("\ntask %d of %d: copying %s to %s\n", index+1, len(tasks), task.oldPath, task.newPath)
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
