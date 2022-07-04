package uiiclient

// OS contains data about an operation system.
type OS struct {
	Architecture string `json:"arch" desc:"architecture e.g. x86_64"`
	DisplayName  string `json:"displayname"`
	Distribution string `json:"dist"`
	Version      string `json:"version" desc:"distribution version e.g. 9, 10 (debian), 16.04, 18.04, 20.04 (ubuntu), 7, 8 (centos), stable, 15.X, tumbleweed (opensuse) etc."`
}
