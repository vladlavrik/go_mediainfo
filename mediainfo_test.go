package mediainfo_test

import (
	"github.com/zhulik/go_mediainfo"
	"io/ioutil"
	"os"
	"testing"
)

const (
	ogg        = "testdata/test.ogg"
	mp3        = "testdata/test.mp3"
	non_exists = "testdata/non_exists.ogg"
)

func TestOpenWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(ogg)
	if error != nil {
		t.Fail()
	}
}

func TestOpenWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(mp3)
	if error != nil {
		t.Fail()
	}
}

func TestOpenWithUnexistsFile(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenFile(non_exists)
	if error == nil {
		t.Fail()
	}
}

func TestOpenMemoryWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	f, _ := os.Open(ogg)
	bytes, _ := ioutil.ReadAll(f)

	error := mi.OpenMemory(bytes)
	if error != nil {
		t.Fail()
	}
}

func TestOpenMemoryWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	f, _ := os.Open(mp3)
	bytes, _ := ioutil.ReadAll(f)

	error := mi.OpenMemory(bytes)
	if error != nil {
		t.Fail()
	}
}

func TestOpenMemoryWithEmptyArray(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	error := mi.OpenMemory([]byte{})
	if error == nil {
		t.Fail()
	}
}


func TestInformWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if len(mi.Inform()) < 10 {
		t.Fail()
	}
}

func TestInformWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if len(mi.Inform()) < 10 {
		t.Fail()
	}
}

func TestAvailableParametersWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if len(mi.AvailableParameters()) < 10 {
		t.Fail()
	}
}

func TestAvailableParametersWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if len(mi.AvailableParameters()) < 10 {
		t.Fail()
	}
}

func TestDurationWithOgg(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(ogg)

	if mi.Duration() != 3494 {
		t.Fail()
	}
}

func TestDurationWithMp3(t *testing.T) {
	mi := mediainfo.NewMediaInfo()
	mi.OpenFile(mp3)

	if mi.Duration() != 87771 {
		t.Fail()
	}
}