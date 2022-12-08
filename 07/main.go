package main

import (
	"fmt"
	"strconv"
	"strings"
)

type dir struct {
	Parent  *dir
	Name    string
	Subdirs []*dir
	Files   []file
}

type file struct {
	Size int64
	Name string
}

func (d *dir) AddDir(name string) {
	var sd dir
	sd.Name = name
	sd.Parent = d
	d.Subdirs = append(d.Subdirs, &sd)
}

func (d *dir) dirSize() int64 {
	var size int64
	size = 0
	for _, f := range d.Files {
		size = size + f.Size
	}

	for _, s := range d.Subdirs {
		size = size + s.dirSize()
	}

	if size < 100000 {
		total += size
		fmt.Println("name", d.Name, "size", size)

	}

	// option two, taking output from option one, grabbing the smallest by hand from the printout.
	//if size >= 2677139 {
	//	total += size
	//	fmt.Println("name", d.Name, "size", size)
	//
	//}
	
	return size
}

func (d *dir) print(indentLevel *int) {
	level := 1
	if indentLevel != nil {
		level = *indentLevel
	}


	fmt.Printf("%*s\n", level, d.Name)
	newLevel := level * 2
	for _, f := range d.Files {
		fmt.Printf("%*s ", newLevel, f.Name)
		fmt.Println(f.Size)
	}

	for _, sd := range d.Subdirs {
		//	fmt.Println(sd)
		sd.print(&newLevel)
	}
}

const isDebug = false
const dummyPayload = ``
const fullPayload = ``

var currentFs *dir
var total int64

func load1() string {
	return dummyPayload
}
func load2() string {
	return fullPayload
}

func main() {
	var input string
	isDummyPayload := false

	if isDummyPayload {
		input = load1()
	} else {
		input = load2()
	}
	objects := strings.Split(input, "\n")

	fmt.Println("First operation each row.")
	FirstCase(objects)

	fmt.Println("Second operation each row.")
	// SecondCase(objects)
}

func FirstCase(objects []string) {
	total = 0
	currentFs = &dir{Name: "top", Subdirs: []*dir{}, Files: []file{}}
	//topFs := *currentFs
	command := ""
	inCommand := false
	var output []string
	for _, obj := range objects {
		if obj[0:1] == "$" {
			if inCommand {
				currentFs = process(currentFs, command, output)
			}

			output = []string{}
			command = obj
			inCommand = true
		} else {
			output = append(output, obj)
		}
	}
	currentFs = process(currentFs, command, output)

	for {
		//	currentFs.print(nil)
		if currentFs.Parent == nil {
			break
		} else {
			currentFs = currentFs.Parent
		}
	}

	currentFs.print(nil)

	fmt.Println(currentFs.dirSize())

	fmt.Println("total:", total)
}

func process(fs *dir, cmd string, outputs []string) *dir {
	splitString := strings.Split(cmd, " ")[1]
	if splitString == "ls" {
		for _, output := range outputs {
			content := strings.Split(output, " ")
			val, err := strconv.ParseInt(content[0], 10, 0)
			if err != nil {
				fs.AddDir(content[1])
			} else {
				fs.Files = append(fs.Files, file{Name: content[1], Size: val})
			}
		}
	}

	if splitString == "cd" {
		dest := strings.Split(cmd, " ")[2]
		if dest == ".." {
			if fs.Parent != nil {
				fs = fs.Parent
			}
		} else {
			for _, sd := range fs.Subdirs {
				if sd.Name == dest {
					fs = sd
				}
			}
		}
	}
	return fs
}

