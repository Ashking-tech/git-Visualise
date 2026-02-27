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



func getBeginningOfTheDay(t time.Time)time.Time {
	year,month,day := t.Date()
	startOfDay := time.Date(year,month,day,0,0,0,0,t.location())
	return startOfDay
}

// countDaysSinceDate counts how many days passed since the passed date

countDaysSinceDate(date time.Time) int {
	days := 0
	now := getBeginningOfTheDay(time.Now())
	for date.Before(now){
		date = date.Add(time.Hour * 24)
		days++
		if days>daysInLastSixMonths{
			return outOfRange
		}
	}
	return days
}


//calcOffset determines and returns the amount of days passed since the passed 'date'
func calcOffset()int{
var offset int 
weekday := time.Now().weekday()
switch weekday{
	case time.Sunday:
		offset = 7
	case time.Monday:
		offset = 6
	case time.Tuesday:
		offset = 5
	case time.Wednesday:
		offset = 4
	case time.Thursday:
		offset = 3
	case time.Friday:
		offset = 2
	case time.Saturday:
		offset = 1

	}
	return offset
}

func printCommitStats(commits map[int]int){
	keys := sortMapIntoSlices(commits)
	cols := buildCols(keys, commits)
	printCells(cols)
}


func sortMapIntoSlices(m map[int]int) []int {
	// order map
	// To store the keys in slice in sorted order
	var keys []int
	for k := range m {
		keys = append(keys,k)
	}

	sort.Ints(keys)
	return keys
}