package glob

import "path/filepath"
import "strings"
import "errors"
import "os"

type matchEntry struct {
	path  string
	index int
}

func Glob(root string, pattern string) (matches []string, e error) {
	if strings.Index(pattern, "**") < 0 {
		return filepath.Glob(filepath.Join(root, pattern))
	}

	segments := strings.Split(pattern, string(os.PathSeparator))

	workingEntries := []matchEntry{
		matchEntry{path: root, index: 0},
	}

	for len(workingEntries) > 0 {
		var temp []matchEntry
		for _, entry := range workingEntries {
			workingPath := entry.path
			index := entry.index
			segment := segments[entry.index]

			if segment == "**" {
				// add all subdirectories and move yourself one step further into pattern
				entry.index++

				subDirectories, err := getAllSubDirectories(entry.path)

				if err != nil {
					return nil, err
				}

				for _, name := range subDirectories {
					path := filepath.Join(workingPath, name)

					newEntry := matchEntry{
						path:  path,
						index: index,
					}

					temp = append(temp, newEntry)
				}

			} else {
				// look at all results
				// if we're at the end of the pattern, we found a match
				// else add it to a working entry
				path := filepath.Join(workingPath, segment)
				results, err := filepath.Glob(path)

				if err != nil {
					return nil, err
				}

				for _, result := range results {
					if index+1 < len(segments) {
						newEntry := matchEntry{
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

func isDir(path string) (val bool, err error) {
	fi, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func getAllSubDirectories(path string) (dirs []string, err error) {
	if dir, err := isDir(path); err != nil || !dir {
		return nil, errors.New("Not a directory " + path)
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
		if dir, err := isDir(path); err == nil && dir {
			dirs = append(dirs, file)
		}
	}

	return
}
