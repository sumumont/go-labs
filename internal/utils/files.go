package utils

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type FileObject struct {
	Subpath string       //文件的相对路径  包含文件名
	Name    string       //文件名
	IsDir   bool         //是否文件夹
	Child   []FileObject //子文件 或 文件夹
}

func (receiver *FileObject) SourceFullPath(sourceRootPath string) string {
	return filepath.Join(sourceRootPath, receiver.Subpath)
}
func (receiver *FileObject) TargetFullPath(targetRootPath string) string {
	return filepath.Join(targetRootPath, receiver.Subpath)
}

func (receiver *FileObject) ReadChild(sourceRootPath string) error {
	sourcePath := receiver.SourceFullPath(sourceRootPath)
	isDir, err := FileIsDir(sourcePath)
	if err != nil {
		return err
	}
	if !isDir {
		return nil
	}

	files, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		return err
	}

	var children []FileObject
	for _, file := range files {
		child := &FileObject{}
		subpath := filepath.Join(receiver.Subpath, file.Name())
		child.Subpath = subpath
		child.Name = file.Name()
		if file.IsDir() {
			child.IsDir = true
			err = receiver.ReadChild(sourceRootPath)
			if err != nil {
				return err
			}
		}
		children = append(children, *child)
	}
	receiver.Child = children
	return nil
}

// 如果没有child 并且是文件夹,则直接copy整个文件夹
// 如果有child,则遍历子文件
func (receiver *FileObject) CopySelf(sourceRootPath string, targetRootPath string) error {
	sourcePath := receiver.SourceFullPath(sourceRootPath)
	targertPath := receiver.TargetFullPath(targetRootPath)
	if len(receiver.Child) == 0 {
		isDir, err := FileIsDir(sourcePath)
		if err != nil {
			return err
		}
		if isDir {
			err = CopyDir(sourceRootPath, targetRootPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(sourceRootPath, targetRootPath)
			if err != nil {
				return err
			}
		}
	} else {
		for _, fileObject := range receiver.Child {
			err := fileObject.CopySelf(sourcePath, targertPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func SplitKey(key string, split string) (string, string) {
	idx := strings.LastIndex(key, split)
	key1 := key[:idx]
	key2 := key[idx+1:]
	return key1, key2
}
