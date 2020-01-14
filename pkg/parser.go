package pkg

import (
	"fmt"
	"strings"

	"github.com/rsc/goversion/version"
)

func parseModuleInfo(moduleInfo string) (Module, []Dependency, []Replacement, error) {
	parsedModuleInfo := strings.Split(moduleInfo, "\n")
	module := Module{}
	dependencies := []Dependency{}
	replacements := []Replacement{}
	for _, item := range parsedModuleInfo {

		// Record the main module
		if strings.HasPrefix(item, "mod\t") {
			tok := strings.Split(strings.TrimPrefix(item, "mod\t"), "\t")
			module = Module{tok[0]}
		}

		// Record a dependency
		if strings.HasPrefix(item, "dep\t") {
			tok := strings.Split(strings.TrimPrefix(item, "dep\t"), "\t")
			dependency := Dependency{}
			if tok[0] != "" {
				switch len(tok) {
				case 3:
					dependency = Dependency{tok[0], tok[1], tok[2]}
				case 2:
					dependency = Dependency{tok[0], tok[1], ""}
				default:
					return module, dependencies, replacements, fmt.Errorf("Unknown Dependency %s", item)
				}
				dependencies = append(dependencies, dependency)
			}
		}

		// Record a replacement
		if strings.HasPrefix(item, "=>\t") {
			tok := strings.Split(strings.TrimPrefix(item, "=>\t"), "\t")
			replacement := Replacement{dependencies[len(dependencies)-1].Name, tok[0], tok[1], tok[2]}
			replacements = append(replacements, replacement)
		}
	}
	return module, dependencies, replacements, nil
}

func parseGoRelease(release string) (GoRelease, error) {
	vsn := release
	if strings.HasPrefix(vsn, "go") {
		vsn = strings.TrimPrefix(vsn, "go")
	}
	rel := strings.Split(vsn, ".")

	goRelease := GoRelease{rel[0], rel[1], release}
	return goRelease, nil
}

func GetInfoFromBinary(path string) (ModFile, error) {
	var release GoRelease
	var module Module
	var dependencies []Dependency
	var replacements []Replacement
	var modFile ModFile
	vsn, err := version.ReadExe(path)
	if err != nil {
		return modFile, fmt.Errorf("%s", err)
	}
	if vsn.Release != "" {
		release, err = parseGoRelease(vsn.Release)
		if err != nil {
			return modFile, fmt.Errorf("%s", err)

		}
	}
	if vsn.ModuleInfo != "" {
		module, dependencies, replacements, err = parseModuleInfo(vsn.ModuleInfo)
		if err != nil {
			return modFile, fmt.Errorf("Unable to parse module info: %s", err)
		}
	} else {
		return modFile, fmt.Errorf("No modules detected, or binary stripped")
	}
	modFile = ModFile{GoRelease: release, Module: module, Dependencies: dependencies, Replacements: replacements}
	return modFile, nil
}
