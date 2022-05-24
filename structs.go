package unattended_install_client

// BuildArgs contains required parameters for building an ISO.
type BuildArgs struct {
	Distribution string        `json:"dist" desc:"distribution name e.g. debian, ubuntu, centos, rocky, opensuse"`
	Version      string        `json:"version" desc:"distribution version e.g. 9, 10 (debian), 16.04, 18.04, 20.04 (ubuntu), 7, 8 (centos), stable, 15.X, tumbleweed (opensuse) etc."`
	Hostname     string        `json:"hostname" desc:"hostname"`
	Networks     []NetworkArgs `json:"networks" desc:"networks"`
}

// BuildOpts contains optional parameters for building an ISO.
type BuildOpts struct {
	Locale          string   `json:"locale" desc:"set locale string"`
	Keyboard        string   `json:"keyboard" desc:"set keymap string"`
	Password        string   `json:"password" desc:"set root password using a sha-512 hash for linux (e.g. mkpasswd -m sha-512)"`
	SSHPasswordAuth bool     `json:"sshpasswordauth" desc:"enable/disable ssh password authentication"`
	SSHKeys         []string `json:"sshkeys" desc:"list of public ssh keys added to authorized_keys"`
	Timezone        string   `json:"timezone" desc:"timezone"`
	Arch            string   `json:"arch" desc:"architecture e.g. x86_64"`
	Packages        []string `json:"packages" desc:"a list of packages added to the base installation"`
}

// NetworkArgs contains network information for building bootable ISO.
type NetworkArgs struct {
	DHCP       bool     `json:"dhcp" desc:"enable IP configuration via dhcp"`
	Domain     string   `json:"domain,omitempty" desc:"network specific domain"`
	MAC        string   `json:"mac,omitempty" desc:"interface specific mac address"`
	IPNet      string   `json:"ipnet,omitempty" desc:"IP cidr e.g. 192.168.0.200/16"`
	Gateway    string   `json:"gateway,omitempty" desc:"network gateway ip address"`
	DNS        []string `json:"dns,omitempty" desc:"optional dns servers"`
	NoInternet bool     `json:"nointernet,omitempty" desc:"optional parameter if network has not internet access it can't be used for installation"`
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

// Os contains data about an operation system.
type Os struct {
	Architecture string `json:"arch" desc:"architecture e.g. x86_64"`
	DisplayName  string `json:"displayname"`
	Distribution string `json:"dist"`
	Version      string `json:"version" desc:"distribution version e.g. 9, 10 (debian), 16.04, 18.04, 20.04 (ubuntu), 7, 8 (centos), stable, 15.X, tumbleweed (opensuse) etc."`
}
