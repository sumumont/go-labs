package files

import (
	"github.com/dustin/go-humanize"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ResourceFile struct {
	Items     []*ResourceFile `json:"items"`
	RootPath  string          `json:"-"`
	FullPath  string          `json:"fullPath"`
	Path      string          `json:"path"`
	Name      string          `json:"name"`
	Size      int64           `json:"size"`
	HumanSize string          `json:"humanSize"` //人类友好的单位
	Extension string          `json:"extension"`
	ModTime   time.Time       `json:"modified"`
	Mode      os.FileMode     `json:"mode"`
	IsDir     bool            `json:"isDir"`
	IsSymlink bool            `json:"isSymlink"`
}

func NewRf(rootPath, subPath string) *ResourceFile {
	return &ResourceFile{
		RootPath: rootPath,
		Path:     subPath,
	}
}

func (r *ResourceFile) Fill() error {
	path := filepath.Join(r.RootPath, r.Path)
	fi, err := os.Stat(path)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil
		}
		return err
	}
	r.Items = nil
	r.Name = fi.Name()
	r.HumanSize = ""
	r.Extension = ""
	r.ModTime = fi.ModTime()
	r.Mode = fi.Mode()
	r.IsDir = fi.IsDir()
	r.IsSymlink = IsSymlink(fi.Mode())
	if !fi.IsDir() { //如果是文件，则结束了
		r.Extension = filepath.Ext(fi.Name())
		r.Size = fi.Size()
		r.HumanSize = humanize.Bytes(uint64(fi.Size()))
		return nil
	}

	//如果是文件夹 则递归读取
	var fs []os.DirEntry
	if fs, err = os.ReadDir(path); err == nil {
		for _, f := range fs {
			subPath := filepath.Join(r.Path, f.Name())
			s := NewRf(r.RootPath, subPath)
			err = s.Fill()
			if err != nil {
				return err
			}
			r.Items = append(r.Items, s)
			r.Size = r.Size + s.Size
		}
		r.HumanSize = humanize.Bytes(uint64(r.Size))
	} else {
		return err
	}

	return nil
}
func IsSymlink(mode os.FileMode) bool {
	return mode&os.ModeSymlink != 0
}
