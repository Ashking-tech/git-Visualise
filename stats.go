package main

import (
	"fmt"
	"sort"
	"time"
	"github.com/go-git/go-git/v5"

)

const daysInLastSixMonths = 183
const outOfRange = 99999
const weeksInLastSixMonths = 26

func processRepositories(email string) map[int]int {
	filepath := getDotFilePath()
	repos := ParseFileLinesToSlice(filepath)
	daysInMap := daysInLastSixMonths
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
	iterator, err := repo.Log(&git.LogOptions{From:ref.Hash()})
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

func countDaysSinceDate(date time.Time) int {
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
weekday := time.Now().Weekday()
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

//buildCols generates a map with rows and columns ready to be printed 

func buildCols(keys []int,commits map[int]int)map[int]column {
	cols := make(map[int]column)
	col := column{}
	for _,k := range keys {
		week := int(k / 7) 
		dayinWeek := k % 7
		if dayinWeek == 0 {
			col = column{}
		}

		col = append(col,commits[k])
		if dayinWeek == 6 {
			cols[week] = col
		}
	}
	return cols
}


func printCells(cols map[int]column){
	printMonths()
	for j := 6; j >=0; j-- {
		for i := weeksInLastSixMonths + 1; i >= 0; i-- {
			if i == weeksInLastSixMonths+1{
				printDayCol(j)

			}
			if col,ok := cols[i];ok {
				//specail case today
				if i == 0 && j == calcOffset()-1 {
					printCells(col[j],true)
					continue	
				}else{
					if len(col) > j {
						printCell(col[j],false)
						continue
					}
				}
			}
			printCell(0,false)
		}
		fmt.Printf("\n")
	}
}

func printMonths() {
	week := getBeginningOfTheDay(time.Now()).Add(-(daysInLastSixMonths * time.Hour * 24))
	month := week.Month()
	fmt.Printf("         ")
	for {
		if week.Month() != month {
			fmt.Printf("%s ", week.Month().String()[:3])
			month = week.Month()
		} else {
			fmt.Printf("    ")
		}

		week = week.Add(7 * time.Hour * 24)
		if week.After(time.Now()) {
			break
		}
	}
	fmt.Printf("\n")
}


//printDayCol given the day number (0 is sunday) so it prints the day name,alternating the rows(prints 2,4,6)

func printDayCol(day int){
	out := "    "
	switch day {
	case 1: 
		out = " Mon "
	case 3:
		out = " Wed "
	case 5:
		out = " Fri "
	}
	fmt.Printf(out)
}

// printCell given a cell value prints it with a different format
// based on the value amount, and on the `today` flag.
func printCell(val int, today bool) {
    escape := "\033[0;37;30m"
    switch {
    case val > 0 && val < 5:
        escape = "\033[1;30;47m"
    case val >= 5 && val < 10:
        escape = "\033[1;30;43m"
    case val >= 10:
        escape = "\033[1;30;42m"
    }

    if today {
        escape = "\033[1;37;45m"
    }

    if val == 0 {
        fmt.Printf(escape + "  - " + "\033[0m")
        return
    }

    str := "  %d "
    switch {
    case val >= 10:
        str = " %d "
    case val >= 100:
        str = "%d "
    }

    fmt.Printf(escape+str+"\033[0m", val)
}