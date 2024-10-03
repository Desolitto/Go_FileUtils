package filefinder

import (
	"fmt"
	"os"
	"path/filepath"
)

func shouldPrint(path string, info os.FileInfo, fileFlag, dirFlag, symlinkFlag *bool, extFlag *string) {
	isDir := info.IsDir()
	isRegular := info.Mode().IsRegular()
	isSymlink := info.Mode()&os.ModeSymlink != 0

	if (*dirFlag && isDir) || (*fileFlag && isRegular) || (*symlinkFlag && isSymlink) || (!*fileFlag && !*dirFlag && !*symlinkFlag) {
		if isSymlink {
			target, err := os.Readlink(path)
			if err != nil {
				fmt.Printf("%s -> [broken]\n", path)
				return
			}
			if !filepath.IsAbs(target) {
				target = filepath.Join(filepath.Dir(path), target)
			}
			if _, err := os.Stat(target); os.IsNotExist(err) {
				fmt.Printf("%s -> [broken]\n", path)
				return
			}
			fmt.Printf("%s -> %s\n", path, target)
		} else if isRegular {
			if *extFlag != "" {
				if filepath.Ext(path) == "."+*extFlag {
					fmt.Println(path)
				}
			} else {
				fmt.Println(path)
			}
		} else if isDir {
			fmt.Println(path)
		}
	}
}

func WalkDirectory(root string, fileFlag, dirFlag, symlinkFlag *bool, extFlag *string) {
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				fmt.Printf("Permission denied: %s\n", path)
				return nil
			}
			return err
		}

		info, err := d.Info()
		if err != nil {
			if os.IsPermission(err) {
				fmt.Printf("Permission denied to get info: %s\n", path)
				return nil
			}
			return err
		}

		shouldPrint(path, info, fileFlag, dirFlag, symlinkFlag, extFlag)

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
