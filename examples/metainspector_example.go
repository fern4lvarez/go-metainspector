package main

import (
	"fmt"
	mi "github.com/fern4lvarez/go-metainspector"
)

func main() {
	url := "http://www.cloudcontrol.com"

	MI, err := mi.NewMetaInspector(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("\nThe url is: %s, the scheme is %s, the host is %s, root url is %s. Title is %s, written in %s, author is %s, description is %s. Charset is %s.", MI.Url(), MI.Scheme(), MI.Host(), MI.RootURL(), MI.Title(), MI.Language(), MI.Author(), MI.Description(), MI.Charset())
		fmt.Printf("\nSubscribe now! ->%s", MI.Feed())
		fmt.Printf("\nThe links are: %#v", MI.Links())
		fmt.Printf("\nThe images are: %#v", MI.Images())
		fmt.Printf("\nThe keywords are: %#v", MI.Keywords())
		fmt.Printf("\nCompatibility: %#v", MI.Compatibility())
	}
}
