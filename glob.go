package glob

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MatchEntry struct {
	path  string
	index int
}

func Glob(root string, pattern string) (matches []string, e error) {
	if strings.Index(pattern, "**") < 0 {
		return filepath.Glob(filepath.Join(root, pattern))
	}

	segments := strings.Split(pattern, string(os.PathSeparator))

	workingEntries := []MatchEntry{
		{path: root, index: 0},
	}

	for len(workingEntries) > 0 {
		var temp []MatchEntry
		for _, entry := range workingEntries {
			workingPath := entry.path
			index := entry.index
			segment := segments[entry.index]

			if segment == "**" {
				// Add all sub dirs and move yourself one step further into pattern
				entry.index++

				subDirs, err := getAllSubDirs(entry.path)

				if err != nil {
					return nil, err
				}

				for _, name := range subDirs {
					path := filepath.Join(workingPath, name)

					newEntry := MatchEntry{
						path:  path,
						index: index,
					}

					temp = append(temp, newEntry)
				}

			} else {
				// look at all results
				// if we're at the end of the pattern,
				// we found a match
				// else add it to a working entry
				path := filepath.Join(workingPath, segment)
				results, err := filepath.Glob(path)

				if err != nil {
					return nil, err
				}

				for _, result := range results {
					if index+1 < len(segments) {
						newEntry := MatchEntry{
							path:  result,
							index: index + 1,
						}

						temp = append(temp, newEntry)
					} else {
						matches = append(matches, result)
					}
				}
				// delete ourself regardless
				entry.index = len(segments)
			}

			// check whether current entry is still valid
			if entry.index < len(segments) {
				temp = append(temp, entry)
			}
		}

		workingEntries = temp
	}

	return
}

func getAllSubDirs(path string) (dirs []string, err error) {
	dir, err := isDir(path)
	if err != nil {
		return nil, err
	}
	if !dir {
		return nil, fmt.Errorf("Glob: the path %s is not a dir ", path)
	}

	d, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	files, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		path := filepath.Join(path, file)
		dir, err := isDir(path)
		if err != nil {
			fmt.Printf("Glob: checking is dir error (%s) for path %s\n", err.Error(), path)
			continue
		}
		if dir {
			dirs = append(dirs, file)
		}
	}

	return
}

func isDir(path string) (val bool, err error) {
	fi, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}
