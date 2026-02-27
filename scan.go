package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func scan(folder string) {
	print("scan")
	fmt.Printf("found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filepath := getDotFilePath()
	addNewsSliceElementsToFile(filepath, repositories)
	fmt.Printf("\n\nSuccessfully added \n\n")

}

// stats generates a nice graph of your git contribution

func stats(email string) {
	print("stats")
}

// scanGitFolders returns a list of subfolders of `folder` ending with `.git`.
// Returns the base folder of the repo, the .git folder parent.
// Recursively searches in the subfolders by passing an existing `folders` slice.
func scanGitFolders(folders []string, folder string) []string {

	folder = strings.TrimSuffix(folder, "/")
	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string
	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}

			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}

			folders = scanGitFolders(folders, path)
		}
	}
	return folders
}

// recursiveScanFolder starts the recursive search of git repositories
// living in the `folder` subtree
func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

// getDotFilePath returns the dot file for repos list
// create it and enclosing folder if it doesnt exist

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)

	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"
	return dotFile
}

func addNewsSliceElementsToFile(filepath string, newRepos []string) {
	existingRepos := ParseFileLinesToSlice(filepath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringSliceToFile(repos, filepath)

}

func ParseFileLinesToSlice(filepath string) []string {
	f := openFile(filepath)
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
	return lines
}

func openFile(filepath string) *os.File {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	return f
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)

		}
	}
	return existing
}

// sliceContains returns true if slice contains value
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// dumpstringslicestofile writes content to file in path "filepath" (overwriting exisitng content )
func dumpStringSliceToFile(repos []string, filepath string) {
	content := strings.Join(repos, "\n")
	ioutil.WriteFile(filepath, []byte(content), 0755)
}
