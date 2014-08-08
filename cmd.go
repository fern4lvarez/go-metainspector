package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fern4lvarez/go-metainspector/metainspector"
)

// exit prints  a message and exit the current program
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// defaultURL returns "www.example.com"
// if there's no arguments
// or help is required
// and returns the first argument if exists
func defaultURL() string {
	if len(os.Args) > 1 {
		if a := os.Args[1]; a == "--h" {
			return "www.example.com"
		}
		return os.Args[1]
	}
	return "www.example.com"
}

// sprintIfExists prints a pretty line if value exists
func sprintIfExists(n string, v string) {
	if v != "" {
		fmt.Printf("----> %s: %s\n", n, v)
	}
}

// aprintIfExists prints pretty lines if the slice
// is not empty
func aprintIfExists(n string, v []string, a bool) {
	if len(v) != 0 {
		fmt.Printf("----> %s: ", n)
		for i := range v {
			if !a && i == 10 {
				fmt.Printf("...")
				break
			}
			fmt.Printf("%s ", v[i])
		}
		fmt.Printf("\n")
	}
}

// prettyspector prints nicely all results
func prettyspector(mi *metainspector.MetaInspector, a bool) {
	sprintIfExists("Title", mi.Title())
	sprintIfExists("Author", mi.Author())
	sprintIfExists("Description", mi.Description())
	sprintIfExists("Generator", mi.Generator())
	sprintIfExists("Charset", mi.Charset())
	sprintIfExists("Language", mi.Language())
	sprintIfExists("Feed URL", mi.Feed())
	aprintIfExists("Keywords", mi.Keywords(), a)
	aprintIfExists("Links", mi.Links(), a)
	aprintIfExists("Images", mi.Images(), a)
}

func main() {
	var url = flag.String("u", defaultURL(), "URL to metainspect.")
	var all = flag.Bool("all", false, "Show full results.")
	flag.Parse()
	mi, err := metainspector.New(*url)
	if err != nil {
		exit("Something went wrong. Please, try again.")
	}
	prettyspector(mi, *all)
}
