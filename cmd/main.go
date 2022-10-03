package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	minisort "github.com/yudgxe/mini-sort"
	"github.com/yudgxe/mini-sort/config"
	"github.com/yudgxe/mini-sort/cstrings"
)

func main() {

	flag := flag.NewFlagSet("mini-sort", flag.ExitOnError)
	getConfing := config.Export(flag)

	if err := flag.Parse(os.Args[2:]); err != nil {
		log.Fatalln(err)
	}

	config := getConfing()
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	t := minisort.New(config)
	t.Read(file, " ")
	fmt.Println(config)

	if config.C {
		if i := t.Check(); i != -1 {
			fmt.Printf("mini-sort: %s: %d: %s", path, i+1, cstrings.BuildFromWords(t.GetRow(i)...))
		}
	} else {
		t.Sort()
		t.Write(os.Stdout)
	}
}
