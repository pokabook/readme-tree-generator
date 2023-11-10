package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
)

type Counter struct {
	dirs  int
	files int
}

func (counter *Counter) index(path string) {
	stat, _ := os.Stat(path)
	if stat.IsDir() {
		counter.dirs += 1
	} else {
		counter.files += 1
	}
}

func (counter *Counter) output() string {
	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
}

func dirnamesFrom(base string) []string {
	file, err := os.Open(base)
	if err != nil {
		fmt.Println(err)
	}

	names, _ := file.Readdirnames(0)
	file.Close()

	sort.Strings(names)
	return names
}

func tree(counter *Counter, base string, prefix string) {
	names := dirnamesFrom(base)

	for index, originName := range names {
		var output string
		if originName[0] == '.' {
			continue
		}
		subpath := path.Join(base, originName)
		counter.index(subpath)
		link := strings.Replace("./"+subpath, " ", "%20", -1)
		if strings.Contains(originName, ".") {
			output = "📄[**" + strings.Split(originName, ".")[0] + "**]" + "(" + link + ")<br>"
		} else {
			output = "📂[**" + originName + "**]" + "(" + link + ")<br>"
		}

		if index == len(names)-1 {
			fmt.Println(prefix+"┗━", output)
			tree(counter, subpath, prefix+"ㅤㅤㅤ")
		} else {
			fmt.Println(prefix+"┣━", output)
			tree(counter, subpath, prefix+"┃ㅤㅤ")
		}
	}
}

func main() {
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		directory = "."
	}

	counter := new(Counter)
	output := "📦[**" + directory + "**]" + "(" + directory + ")<br>"
	fmt.Println(output)

	tree(counter, directory, "")
	fmt.Println(counter.output())
}
