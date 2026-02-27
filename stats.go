package main

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"

)

const daysInLastSixMonths = 183
const outOfRange = 99999

func processRepositories(email string) map[int]int {
	filepath := getDotFilePath()
	repos := ParseFileLinesToSlice(filepath)
	daysInMap := DaysInLastSixMonths
	commits := make(map[int]int,daysInMap)
	for i := daysInMap; i > 0; i--{
		commits[i] = 0 
	}
	for _, path := range repos {
		commits = fillCommits(email,path,commits)

	}
	return commits
}

// fillCommits given a repo found in 'path' ,get the commits and put them in the commits map,return it qhen completed

func fillCommits(email string,path string,commits map[int]int)map[int]int {
	//provide an instance to a git repo obj from path
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}


	// get the head refereccnce
	ref, err := repo.Head()
	if err != nil {
			panic(err)
	}

	//get commit histrory starting from head
	iterator, err := repo.Log(&git.LogOptions{from:ref.Hash()})
	if err != nil {
		panic(err)
	}

	//iterate the commits
	offset := calcOffset()
	err = iterator.ForEach(func(c *object.Commit)error){
		daysAgo := countDaysSinceDate(c.Author.When) + offset
		if c.Author.Email != email{
			return nil
		}

		if daysAgo != outOfRange {
			commits[daysAgo]++
		}
		return nil 
	})
	if err != nil {
		panic(err)
	}
	return commits


	
}


