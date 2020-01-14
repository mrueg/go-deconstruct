package pkg

import "testing"

func TestParseGoRelease(t *testing.T) {
	tests := []struct {
		input  string
		output GoRelease
		err    error
	}{
		{"go1.13.8", GoRelease{"1", "13", "go1.13.8"}, nil},
	}
	for _, test := range tests {
		goRelease, _ := parseGoRelease(test.input)
		if goRelease != test.output {
			t.Errorf("ParseGoRelease of (%s) was incorrect, got: %s, want: %s.", test.input, goRelease, test.output)
		}
	}
}

func TestParseModuleInfo(t *testing.T) {
	tests := []struct {
		input        string
		module       Module
		dependencies []Dependency
		replacements []Replacement
		err          error
	}{
		{"path\tgithub.com/mrueg/go-deconstruct\nmod\tgithub.com/mrueg/go-deconstruct\t(devel)\t\ndep\tgithub.com/rsc/goversion\tv1.2.0\th1:zVF4y5ciA/rw779S62bEAq4Yif1cBc/UwRkXJ2xZyT4=\ndep\tgithub.com/spf13/cobra\tv0.0.5\th1:f0B+LkLX6DtmRH1isoNA9VTtNUK9K8xYd28JNNfOv/s=\ndep\tgithub.com/spf13/pflag\tv1.0.3\th1:zPAT6CGy6wXeQ7NtTnaTerfKOsV6V6F8agHXFiazDkg=\n",
			Module{"github.com/mrueg/go-deconstruct"},
			[]Dependency{Dependency{"github.com/rsc/goversion", "v1.2.0", "h1:zVF4y5ciA/rw779S62bEAq4Yif1cBc/UwRkXJ2xZyT4="},
				Dependency{"github.com/spf13/cobra", "v0.0.5", "h1:f0B+LkLX6DtmRH1isoNA9VTtNUK9K8xYd28JNNfOv/s="},
				Dependency{"github.com/spf13/pflag", "v1.0.3", "h1:zPAT6CGy6wXeQ7NtTnaTerfKOsV6V6F8agHXFiazDkg="}},
			[]Replacement{},
			nil},
		{"path\tgithub.com/mrueg/go-deconstruct\nmod\tgithub.com/mrueg/go-deconstruct\t(devel)\t\ndep\tgithub.com/rsc/goversion\tv1.2.0\n=>\treplacement.com/rsc/goversion\tv2.2.0\th1:zVF4y5ciA/rw779S62bEAq4Yif1cBc/UwRkXJ2xZyT4=",
			Module{"github.com/mrueg/go-deconstruct"},
			[]Dependency{Dependency{"github.com/rsc/goversion", "v1.2.0", ""},
				Dependency{"github.com/spf13/cobra", "v0.0.5", "h1:f0B+LkLX6DtmRH1isoNA9VTtNUK9K8xYd28JNNfOv/s="},
				Dependency{"github.com/spf13/pflag", "v1.0.3", "h1:zPAT6CGy6wXeQ7NtTnaTerfKOsV6V6F8agHXFiazDkg="}},
			[]Replacement{Replacement{"github.com/rsc/goversion", "replacement.com/rsc/goversion", "v2.2.0", "h1:zVF4y5ciA/rw779S62bEAq4Yif1cBc/UwRkXJ2xZyT4="}},
			nil},
	}
	for _, test := range tests {
		module, dependencies, replacements, err := parseModuleInfo(test.input)
		if err != test.err {
			t.Errorf("ParseModuleInfo of (%s) was incorrect, got: %s, want: %s.", test.input, err, test.err)
		}
		if module != test.module {
			t.Errorf("ParseModuleInfo of (%s) was incorrect, got: %s, want: %s.", test.input, module, test.module)
		}
		for i, item := range dependencies {
			if item != test.dependencies[i] {
				t.Errorf("ParseModuleInfo of (%s) was incorrect, got: %s, want: %s.", test.input, item, test.dependencies[i])
			}
		}
		for i, item := range replacements {
			if item != test.replacements[i] {
				t.Errorf("ParseModuleInfo of (%s) was incorrect, got: %s, want: %s.", test.input, item, test.replacements[i])
			}
		}

	}

}
