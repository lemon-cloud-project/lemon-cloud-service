package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type IOUtils struct {
}

var ioInstance *IOUtils
var ioOnce sync.Once

func IO() *IOUtils {
	ioOnce.Do(func() {
		ioInstance = &IOUtils{}
	})
	return ioInstance
}

func (i *IOUtils) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (i *IOUtils) CopyFile(src string, dst string) error {
	srcFile, errSrc := os.Open(src)
	if errSrc != nil {
		return errSrc
	}
	return i.CopyFileFromReader(srcFile, dst)
}

func (i *IOUtils) CopyFileFromReader(srcFile io.Reader, dst string) error {
	dirPath, _ := path.Split(dst)
	if !i.PathExists(dirPath) {
		dstDirErr := os.MkdirAll(dirPath, os.ModePerm)
		if dstDirErr != nil {
			return dstDirErr
		}
	}
	dstFile, errDstCreate := os.Create(dst)
	if errDstCreate != nil {
		return errDstCreate
	}
	defer dstFile.Close()
	_, errDestCopy := io.Copy(dstFile, srcFile)
	if errDestCopy != nil {
		return errDestCopy
	}
	return nil
}

func (i *IOUtils) CopyDir(src string, dest string) error {
	err := filepath.Walk(src, func(currentSrc string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		aimPath := strings.Replace(currentSrc, src, dest, 1)
		if !f.IsDir() {
			i.CopyFile(currentSrc, aimPath)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *IOUtils) JsonFileToStruct(jsonSrc string, obj interface{}) error {
	data, readErr := ioutil.ReadFile(jsonSrc)
	if readErr != nil {
		return readErr
	}
	parseErr := json.Unmarshal(data, obj)
	if parseErr != nil {
		return parseErr
	}
	return nil
}

func (i *IOUtils) StructToJsonFile(jsonSrc string, obj interface{}) error {
	data, parseErr := json.Marshal(obj)
	if parseErr != nil {
		return parseErr
	}
	return i.ReplaceStrToFile(string(data), jsonSrc)
}

func (i *IOUtils) GetRuntimePath(filename string) string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(dir, filename)
}

func (i *IOUtils) ReplaceStrToFile(content, path string) error {
	if i.PathExists(path) {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, wErr := file.WriteString(content)
	if wErr != nil {
		return wErr
	}
	return nil
}
