package ffHash

import (
	"fmt"
	"github.com/pkg/errors"
	"hash/crc32"
	"io"
	"os"
	"sync"
)

type fileHashTbl struct {
	Tbl   map[string][]string
	mutex sync.RWMutex
}

// Construct new fileHashTbl
func NewDataTbl() *fileHashTbl {
	var TblInstance fileHashTbl
	TblInstance.Tbl = make(map[string][]string)
	return &TblInstance
}

// Output identical files to stdout
func OutputDataTbl(datatbl fileHashTbl) error {
	if len(datatbl.Tbl) == 0 {
		return errors.New("no files in directory")
	}
	//output identical files
	for k, _ := range datatbl.Tbl {
		if len(datatbl.Tbl[k]) > 1 {
			fmt.Printf("Files hash: %x ", []byte(k))
			for _, v := range datatbl.Tbl[k] {
				fmt.Println(v)
			}
		}
	}
	//if ok return nil(no error)
	return nil
}

// Fill [file hash/file name] map in directory
func DirsFilesHashTotbl(dirname string, wg *sync.WaitGroup, datatbl *fileHashTbl) error {
	//Decrement sync counter
	defer wg.Done()
	//Read directory entries
	filesIndir, err := os.ReadDir(dirname)
	if err != nil {
		return errors.Wrap(err, "unable to read directory")
	}
	//iterate files
	for _, f := range filesIndir {
		//if not dir
		if f.IsDir() != true {
			filePath := dirname + f.Name()
			file, err := os.Open(filePath)
			if err != nil {
				return errors.Wrap(err, "unable to open file")
			}
			//hash crc32 for each file
			hash := crc32.NewIEEE()
			_, err = io.Copy(hash, file)
			if err != nil {
				file.Close()
				return errors.Wrap(err, "unable to generate file hash")
			}
			//mutex for safe access to resource
			datatbl.mutex.Lock()
			fileHash := string(hash.Sum(nil))
			datatbl.Tbl[fileHash] = append(datatbl.Tbl[fileHash], file.Name())
			datatbl.mutex.Unlock()
			file.Close()
		}
	}
	//if ok return nil(no error)
	return nil
}

func GetDirList(dirname string) ([]string, error) {
	//Dir list
	dirList := make([]string, 0)
	//Read directory entries
	filesIndir, err := os.ReadDir(dirname)
	if err != nil {
		return dirList, errors.Wrap(err, "unable to read directory")
	}
	//iterate files
	for _, f := range filesIndir {
		if f.IsDir() == true {
			dirPath := dirname + f.Name() + "/"
			dirList = append(dirList, dirPath)
			nestedDirList, err := GetDirList(dirPath)
			if err != nil {
				return dirList, errors.Wrap(err, "unable to read nested directory")
			}
			dirList = append(dirList, nestedDirList...)
		}
	}
	return dirList, nil
}

func DelIdenticalFiles(datatbl fileHashTbl) error {
	if len(datatbl.Tbl) == 0 {
		return errors.New("no files in directory")
	}
	//output identical files
	for k, _ := range datatbl.Tbl {
		if len(datatbl.Tbl[k]) > 1 {
			for i, v := range datatbl.Tbl[k] {
				if i > 0 {
					err := os.Remove(v)
					if err != nil {
						return errors.Wrap(err,"Cannot delete file")
					}
				}
			}
		}
	}
	//if ok return nil(no error)
	return nil
}
