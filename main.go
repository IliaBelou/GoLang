package main

import (
	"bufio"
	"flag"
	"fmt"
	"main.go/ffHash"
	"os"
	"sync"
)

// Format: key = crc32 file hash, value: slice of file names with same hash

func main() {
	//Flags handler
	delFilesFlag := flag.Bool("delF", false, "delete identical files")
	flag.Parse()
	//Get data from stdin
	scanner := bufio.NewScanner(os.Stdin)
	//fmt.Print("Input path to directory, example './testDirectory/': ")
	scanner.Scan()
	dirname := scanner.Text()
	//Generate directory list
	dirList, err := ffHash.GetDirList(dirname)
	if err != nil {
		fmt.Println(err)
	}
	dirList = append(dirList, dirname)
	//Construct new data table
	datatbl := ffHash.NewDataTbl()
	//iterate over directory list to get [hash/files names] table
	var wg sync.WaitGroup
	wg.Add(len(dirList))
	for _, v := range dirList {
		go ffHash.DirsFilesHashTotbl(v, &wg, datatbl)
		//go (datatbl).ffHash.DirsFilesHashTotbl(v, &wg)
		if err != nil {
			fmt.Println(err)
		}
	}
	wg.Wait()
	//Output data to stdout
	err = ffHash.OutputDataTbl(*datatbl)
	if err != nil {
		fmt.Println(err)
	}
	//Delete files, optional
	if *delFilesFlag {
		fmt.Println("delete identical files...")
		err = ffHash.DelIdenticalFiles(*datatbl)
		if err != nil {
			fmt.Println(err)
		}
	}
}
