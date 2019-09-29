// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The argutil package contains libraries for parsing commandline args.
package argutil

import (
	"fmt"
	"strings"
)

var ErrMultiVersion = fmt.Errorf("at most 1 version permitted")

// ParseDirVersion parses given string of the form dir@verion and returns dir
// and version.
func ParseDirVersion(dirVer string) (string, string, error) {
	if dirVer == "" {
		return "", "", nil
	}
	if !strings.Contains(dirVer, "@") {
		return dirVer, "", nil
	}
	parts := strings.Split(dirVer, "@")
	if len(parts) > 2 {
		return "", "", ErrMultiVersion
	}
	return parts[0], parts[1], nil
}

// ParseDirVersionWithDefaults parses given string of the form dir@version and
// returns dir and version with following defaults.
// if dir is missing, return current working directory
// if version is missing, return "master"
func ParseDirVersionWithDefaults(dirVer string) (string, string, error) {
	dir, version, err := ParseDirVersion(dirVer)
	if err != nil {
		return dir, version, err
	}
	if dir == "" {
		dir = "./"
	}
	if version == "" {
		version = "master"
	}
	return dir, version, nil
}

// ParseFieldPath parse a flag value into a field path
// TODO(pwittrock): Extract this into lib.kpt.dev
func ParseFieldPath(path string) ([]string, error) {
	// fixup '\.' so we don't split on it
	match := strings.ReplaceAll(path, "\\.", "$$$$")
	parts := strings.Split(match, ".")
	for i := range parts {
		parts[i] = strings.ReplaceAll(parts[i], "$$$$", ".")
	}

	// split the list index from the list field
	var newParts []string
	for i := range parts {
		if !strings.Contains(parts[i], "[") {
			newParts = append(newParts, parts[i])
			continue
		}
		p := strings.Split(parts[i], "[")
		if len(p) != 2 {
			return nil, fmt.Errorf("unrecognized path element: %s.  "+
				"Should be of the form 'list[field=value]'", parts[i])
		}
		p[1] = "[" + p[1]
		newParts = append(newParts, p[0], p[1])
	}
	return newParts, nil
}
