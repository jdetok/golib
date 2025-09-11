package docd

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
