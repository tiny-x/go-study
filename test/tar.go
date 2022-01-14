package main

import (
	"archive/tar"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	//打包的文件名
	fileTarget := "/tmp/picture.tar"
	// 需要打包的文件路径
	filesource := "/opt/chaosblade-1.3.1/blade"
	// 创建打包文件
	tarfile, err := os.Create(fileTarget)

	if err != nil {
		if err == os.ErrExist {
			if err := os.Remove(fileTarget); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}
	// 关闭文件
	defer tarfile.Close()
	// 写入文件句柄
	trawriter := tar.NewWriter(tarfile)
	err = ArchiveTar(filesource, trawriter)
	if err != nil {
		fmt.Println(err)
	}

}

func ArchiveTar(file string, writer *tar.Writer) error {

	return filepath.Walk(file, func(path string, fileInfo fs.FileInfo, err error) error {
		if fileInfo == nil {
			return err
		}
		if fileInfo.IsDir() {
			if path == path {
				return nil
			}
			header, err := tar.FileInfoHeader(fileInfo, "")
			if err != nil {
				return err
			}
			header.Name = filepath.Join(path, strings.TrimPrefix(path, path))
			if err = writer.WriteHeader(header); err != nil {
				return err
			}
			os.Mkdir(strings.TrimPrefix(path, fileInfo.Name()), os.ModeDir)
			return ArchiveTar(path, writer)
		}
		return func(originFile, path string, fileInfo fs.FileInfo, writer *tar.Writer) error {
			if file, err := os.Open(path); err != nil {
				return err
			} else {

				if header, err := tar.FileInfoHeader(fileInfo, ""); err != nil {
					return err
				} else {

					index := strings.LastIndex(originFile, "/")
					header.Name = strings.ReplaceAll(path, originFile[0:index+1], "")

					if err := writer.WriteHeader(header); err != nil {
						return err
					}

					if _, err = io.Copy(writer, file); err != nil {
						return err
					}
				}

			}
			return nil
		}(file, path, fileInfo, writer)
	})
}
