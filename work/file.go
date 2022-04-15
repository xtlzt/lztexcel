package work

import (
	"fmt"
	"os"
	"path"
	"time"
)

//文件结构体
type File struct {
	FileName    string
	FileExt     string
	FileSize    int64
	FileIsDir   bool
	FileMode    os.FileMode
	FileModTime time.Time
	Sys         interface{}
	FileMethod  bool //默认是小文件读取,false 小文件，true是大文件
}

func FileInfo(address string, FileReadersSize int64) *File {
	fileInfo, err := os.Stat(address)
	if err != nil {
		fmt.Println(err)
	}
	FileMethod := false
	if fileInfo.Size() > FileReadersSize {
		FileMethod = true
	}
	ext := path.Ext(address)
	return &File{
		FileName:    fileInfo.Name(),
		FileExt:     ext[1:],
		FileSize:    fileInfo.Size(),
		FileIsDir:   fileInfo.IsDir(),
		FileMode:    fileInfo.Mode(),
		FileModTime: fileInfo.ModTime(),
		Sys:         fileInfo.Sys(),
		FileMethod:  FileMethod,
	}
}

//判断上传文件类型
func (fl *File) checkFileType(fileType []string) {
	
}
