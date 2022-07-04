package uiiclient

type packageCombined struct {
	PackageArgs
	PackageOpts
}

// PackageArgs contains all mandatory parameters for requesting the package list.
type PackageArgs struct {
	Distribution string `json:"dist" desc:"distribution name e.g. debian, ubuntu, centos, rocky, opensuse"`
	Version      string `json:"version" desc:"distribution version e.g. 9, 10 (debian), 16.04, 18.04, 20.04 (ubuntu), 7, 8 (centos), stable, 15.X, tumbleweed (opensuse) etc."`
}

// PackageOpts contains all optional parameters for requesting the package list.
type PackageOpts struct {
	Architecture string `json:"arch" desc:"architecture e.g. x86_64"`
}

// PackageListResponse contains the result of reading a package list for a given OS.
type PackageListResponse struct {
	Distribution string   `json:"dist" desc:"distribution name e.g. debian, ubuntu, centos, rocky, opensuse"`
	Version      string   `json:"version" desc:"distribution version e.g. 9, 10 (debian), 16.04, 18.04, 20.04 (ubuntu), 7, 8 (centos), stable, 15.X, tumbleweed (opensuse) etc."`
	Architecture string   `json:"arch" desc:"architecture e.g. x86_64"`
	Packages     []string `json:"packages" desc:"a list of packages supported for this "`
}
