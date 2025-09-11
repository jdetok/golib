package docd

// outside package, hold project name etc. struct that will be passed
type Docd struct {
	FileName string   `json:"-"`
	FullPath string   `json:"-"`
	Project  string   `json:"project_name"` // project name
	Version  string   `json:"project_version"`
	Packages []PkgDoc `json:"packages"`
}

type PkgDoc struct {
	Pkg     string `json:"package_name"`
	GoFiles string
}
