package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func scan(folder string){
	print("scan")
	fmt.Printf("found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filepath := getDotfilePath()
	addNewsSliceElementsToFile(filepath,repositories)
	fmt.Printf("\n\nSuccesfully added \n\n")

}




// stats generates a nice graph of your git contribution

func stats(email string){
	print("stats")
}


func main (){
	var folder string
	var email string
	flag.StringVar(&folder,"add","","add a  new folder to scan for git repository")
	flag.StringVar(&email,"email","your@gmail.com ","the email to scan")
	flag.Parse()
	if folder != ""{
		scan(folder)
		return
	}

	stats(email)

	
}

// scanGitFolders returns a list of subfolders of `folder` ending with `.git`.
// Returns the base folder of the repo, the .git folder parent.
// Recursively searches in the subfolders by passing an existing `folders` slice.
func scanGitFolders(folders []string,folder string) []string {

	folder := strings.TrimSuffix(folder,"/")
	f, err := os.Open(folder)
	if err := nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}


	var path string
	for _, file := range files {
		if file.IsDir(){
			path := folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path,"/")
				fmt.Println(path)
				folders = append(folders,path)
				continue
			}


			if file.Name() == "vendor" || file.Name() == "node_modules"{
				continue
			}

			folders = scanGitFolders(folders,path)
			}
		}
	}
	return folders
}
