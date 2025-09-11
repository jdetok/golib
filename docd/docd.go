package docd

import (
	"fmt"
	"os"
)

/* PACKAGE INTENTION:
init a Docd struct at the start of a main function, pass it by reference to
generate documentation for each function it's passed to
* maybe add a CLI element
*/

// outside package, hold project name etc. struct that will be passed
// Docd functions will only execute if Generate is 1
type Docd struct {
	Generate bool     `json:"-"`
	FileName string   `json:"-"`
	FullPath string   `json:"-"`
	Project  string   `json:"project_name"` // project name
	Version  string   `json:"project_version,omitempty"`
	Packages []PkgDoc `json:"packages"`
}

type PkgDoc struct {
	Pkg     string    `json:"package_name"`
	GoFiles []FileDoc `json:"go_files"`
}

type FileDoc struct {
	GoFileName string    `json:"go_file"`
	Desc       string    `json:"file_desc,omitempty"`
	GoFuncs    []FuncDoc `json:"go_functions"`
}

type FuncDoc struct {
	FuncName string `json:"func_name"`
	Desc     string `json:"func_desc"`
}

// run at top of main
func InitDocd(path, fname, pname, version string) (*Docd, error) {
	d := Docd{}
	d.FullPath = path + "/" + fname

	// create dir if it doesn't exist
	if err := os.MkdirAll(path, 0777); err != nil {
		return nil, fmt.Errorf(
			"\n****\n**** error creating or opening directory at %s\n****", path)
	}

	f, err := os.OpenFile(d.FullPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf(
			"\n****\n**** can't open file at %s. CHECK IF DIRECTORY EXISTS\n****",
			d.FullPath)
	}

	nb, err := f.Write([]byte("testing"))
	if err != nil {
		return nil, fmt.Errorf(
			"\n****\n**** error writing to file at %s:\n****%e", d.FileName, err)
	}

	fmt.Printf("successfully wrote %d bytes to %s", nb, d.FileName)

	f.Close()

	return &d, nil
}
