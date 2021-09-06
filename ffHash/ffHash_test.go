package ffHash

import (
	"reflect"
	"testing"
)

func TestGetDirList(t *testing.T) {
	got, _ := GetDirList("./testDirectory/")
	want := []string {
		"./testDirectory/AnestedDirectory/",
		"./testDirectory/nestedDirectory/",
		"./testDirectory/nestedDirectory/nestedDirectoryLvl2/",
		"./testDirectory/nestedDirectory2/",
		"./testDirectory/nestedDirectory2/nestedDirectoryLvl2/",
	}
	if !reflect.DeepEqual(got,want) {
		t.Errorf("TestGetDirList test failed")
	}
}
