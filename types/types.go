package types

type ModFile struct {
	GoRelease    GoRelease
	Module       Module
	Dependencies []Dependency
	Replacements []Replacement
}

type GoRelease struct {
	Major string
	Minor string
	Name  string // Full Release Name
}

type Module struct {
	Name string
}

type Dependency struct {
	Name    string
	Version string
	Hash    string
}

type Replacement struct {
	Name         string
	ReplacedWith string
	Version      string
	Hash         string
}
