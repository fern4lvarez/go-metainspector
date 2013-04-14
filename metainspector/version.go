package metainspector

import (
	"io/ioutil"
)

// Version of the package
func Version() string {
	b, _ := ioutil.ReadFile("../VERSION")
	v := string(b)
	return v
}
