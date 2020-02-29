package parser

import (
	"fmt"
	"strings"

	"github.com/mrueg/go-deconstruct/types"
	"github.com/rsc/goversion/version"
)

func parseModuleInfo(moduleInfo string) (types.Module, []types.Dependency, []types.Replacement, error) {
	parsedModuleInfo := strings.Split(moduleInfo, "\n")
	module := types.Module{}
	dependencies := []types.Dependency{}
	replacements := []types.Replacement{}
	for _, item := range parsedModuleInfo {

		// Record the main module
		if strings.HasPrefix(item, "mod\t") {
			tok := strings.Split(strings.TrimPrefix(item, "mod\t"), "\t")
			module = types.Module{Name: tok[0]}
		}

		// Record a dependency
		if strings.HasPrefix(item, "dep\t") {
			tok := strings.Split(strings.TrimPrefix(item, "dep\t"), "\t")
			dependency := types.Dependency{}
			if tok[0] != "" {
				switch len(tok) {
				case 3:
					dependency = types.Dependency{Name: tok[0], Version: tok[1], Hash: tok[2]}
				case 2:
					dependency = types.Dependency{Name: tok[0], Version: tok[1], Hash: ""}
				default:
					return module, dependencies, replacements, fmt.Errorf("Unknown Dependency %s", item)
				}
				dependencies = append(dependencies, dependency)
			}
		}

		// Record a replacement
		if strings.HasPrefix(item, "=>\t") {
			tok := strings.Split(strings.TrimPrefix(item, "=>\t"), "\t")
			replacement := types.Replacement{Name: dependencies[len(dependencies)-1].Name, ReplacedWith: tok[0], Version: tok[1], Hash: tok[2]}
			replacements = append(replacements, replacement)
		}
	}
	return module, dependencies, replacements, nil
}

func parseGoRelease(release string) (types.GoRelease, error) {
	vsn := release
	if strings.HasPrefix(vsn, "go") {
		vsn = strings.TrimPrefix(vsn, "go")
	}
	rel := strings.Split(vsn, ".")

	goRelease := types.GoRelease{Major: rel[0], Minor: rel[1], Name: release}
	return goRelease, nil
}

func GetInfoFromBinary(path string) (types.ModFile, error) {
	var release types.GoRelease
	var module types.Module
	var dependencies []types.Dependency
	var replacements []types.Replacement
	var modFile types.ModFile
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
	modFile = types.ModFile{GoRelease: release, Module: module, Dependencies: dependencies, Replacements: replacements}
	return modFile, nil
}
