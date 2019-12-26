package tjdoe

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type copyInfo struct {
	from string
	to   string
	info os.FileInfo
}

/*
Programs in this file refers to https://github.com/otiai10/copy.
*/
func (tjdoe *TJDoe) copy(from, to string) error {
	info, err := os.Lstat(from)
	if err != nil {
		return err
	}
	return tjdoe.copyImpl(&copyInfo{from: from, to: to, info: info})
}

func (tjdoe *TJDoe) copyImpl(info *copyInfo) error {
	if info.info.Mode()&os.ModeSymlink != 0 {
		return copySymlinks(tjdoe, info)
	} else if info.info.IsDir() {
		return copyDirectory(tjdoe, info)
	}
	return copyFile(tjdoe, info)
}

func copyDirectory(tjdoe *TJDoe, info *copyInfo) error {
	originalMode := info.info.Mode()
	newPath, err := makeDirectories(tjdoe, info.to)
	if err != nil {
		return err
	}
	defer os.Chmod(newPath, originalMode)
	children, err := ioutil.ReadDir(info.from)
	if err != nil {
		return err
	}
	for _, child := range children {
		childFrom := filepath.Join(info.from, child.Name())
		childTo := filepath.Join(newPath, child.Name())
		if err := tjdoe.copy(childFrom, childTo); err != nil {
			return err
		}
	}
	return nil
}

func isExistDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func updateLine(tjdoe *TJDoe, line string) string {
	for _, mapping := range tjdoe.mapping {
		if strings.Index(line, mapping.fromID) >= 0 {
			line = strings.ReplaceAll(line, mapping.fromID, mapping.toID)
		}
	}
	return line
}

func updateName(tjdoe *TJDoe, name string) string {
	for _, mapping := range tjdoe.mapping {
		if strings.Index(name, mapping.fromID) >= 0 {
			return strings.ReplaceAll(name, mapping.fromID, mapping.toID)
		}
	}
	return name
}

func isYear(name string) bool {
	if len(name) != 4 || !isNumeric(name) {
		return false
	}
	year, _ := strconv.Atoi(name)
	now := time.Now()
	return year > 1970 && year <= now.Year()+1
}

func updateBase(tjdoe *TJDoe, to string) string {
	dir := filepath.Dir(to)
	base := updateName(tjdoe, filepath.Base(to))
	if isYear(base) {
		base = "0000"
	}
	return filepath.Join(dir, base)
}

func makeDirectories(tjdoe *TJDoe, path string) (string, error) {
	dir := filepath.Dir(path)
	if isExistDir(path) {
		return path, nil
	}
	dir, err := makeDirectories(tjdoe, dir)
	if err != nil {
		return "", err
	}
	base := updateBase(tjdoe, filepath.Base(path))
	newPath := filepath.Join(dir, base)
	return newPath, os.Mkdir(filepath.Join(dir, base), os.ModePerm)
}

func copyContent(tjdoe *TJDoe, writer io.Writer, reader io.Reader) error {
	in := bufio.NewScanner(reader)
	for in.Scan() {
		line := updateLine(tjdoe, in.Text())
		writer.Write([]byte(line))
		writer.Write([]byte("\n"))
	}
	return in.Err()
}

func copyFile(tjdoe *TJDoe, info *copyInfo) error {
	parent, err := makeDirectories(tjdoe, filepath.Dir(info.to))
	if err != nil {
		return err
	}
	dest := updateBase(tjdoe, filepath.Join(parent, filepath.Base(info.to)))
	writer, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer writer.Close()
	reader, err := os.Open(info.from)
	if err != nil {
		return err
	}
	defer reader.Close()
	err = copyContent(tjdoe, writer, reader)
	return err
}

func copySymlinks(tjdoe *TJDoe, info *copyInfo) error {
	src, err := os.Readlink(info.from)
	if err != nil {
		return err
	}
	return os.Symlink(src, info.to)
}
