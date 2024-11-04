package CLI

import (
	"fmt"
	"os"
	"path/filepath"
)

// Which используется для определения местоположения исполняемых файлов.
func Which(args []string) {
	if len(args) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, file := range args[1:] {

		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			// Does it exist?
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() {
					// Is it executable?
					if mode&0111 != 0 {
						fmt.Println(fullPath)
					}
				}
			}
		}
	}
}
