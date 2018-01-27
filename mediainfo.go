package mediainfo

// #cgo CFLAGS: -DUNICODE
// #cgo LDFLAGS: -lz -lzen -lpthread -lstdc++ -lmediainfo -ldl
// #include "go_mediainfo.h"
import "C"

import (
	"fmt"
	"strconv"
	"unsafe"
)

const (
	General = C.MediaInfo_Stream_General
	Video   = C.MediaInfo_Stream_Video
	Audio   = C.MediaInfo_Stream_Audio
	Image   = C.MediaInfo_Stream_Image
	Text    = C.MediaInfo_Stream_Text
)

// MediaInfo - represents MediaInfo class, all interaction with libmediainfo through it
type MediaInfo struct {
	handle unsafe.Pointer
}

func init() {
	C.setlocale(C.LC_CTYPE, C.CString(""))
	C.MediaInfoDLL_Load()

	if C.MediaInfoDLL_IsLoaded() == 0 {
		panic("Cannot load mediainfo")
	}
}

// NewMediaInfo - constructs new MediaInfo
func NewMediaInfo() *MediaInfo {
	return &MediaInfo{handle: C.GoMediaInfo_New()}
}

// OpenFile - opens file
func (mi *MediaInfo) OpenFile(path string) error {
	p := C.CString(path)
	s := C.GoMediaInfo_OpenFile(mi.handle, p)
	if s == 0 {
		return fmt.Errorf("mediainfo can't open file: %s", path)
	}
	C.free(unsafe.Pointer(p))
	return nil
}

// OpenMemory - opens memory buffer
func (mi *MediaInfo) OpenMemory(bytes []byte) error {
	if len(bytes) == 0 {
		return fmt.Errorf("buffer is empty")
	}
	s := C.GoMediaInfo_OpenMemory(mi.handle, (*C.char)(unsafe.Pointer(&bytes[0])), C.size_t(len(bytes)))
	if s == 0 {
		return fmt.Errorf("mediainfo can't open memory buffer")
	}
	return nil
}

// Close - closes file
func (mi *MediaInfo) Close() {
	C.GoMediaInfo_Close(mi.handle)
}

// GetStream - allow to read stream info from file
func (mi *MediaInfo) Get(typ uint32, param string) (result string) {
	p := C.CString(param)
	r := C.GoMediaInfoGet(mi.handle, typ, p)
	result = C.GoString(r)
	C.free(unsafe.Pointer(p))
	C.free(unsafe.Pointer(r))
	return
}

// Inform returns string with summary file information, like mediainfo util
func (mi *MediaInfo) Inform() (result string) {
	r := C.GoMediaInfoInform(mi.handle)
	result = C.GoString(r)
	C.free(unsafe.Pointer(r))
	return
}

// Option configure or get information about MediaInfoLib
func (mi *MediaInfo) Option(option string, value string) (result string) {
	o := C.CString(option)
	v := C.CString(value)
	r := C.GoMediaInfoOption(mi.handle, o, v)
	C.free(unsafe.Pointer(o))
	C.free(unsafe.Pointer(v))
	result = C.GoString(r)
	C.free(unsafe.Pointer(r))
	return
}

// AvailableParameters returns string with all available Get params and it's descriptions
func (mi *MediaInfo) AvailableParameters() string {
	return mi.Option("Info_Parameters", "")
}

// StreamCount returns count of streams
func (mi *MediaInfo) StreamCount(typ string) int {
	val := mi.Get(General, typ)
	cnt, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return cnt
}

// VideoCount returns count of video streams
func (mi *MediaInfo) VideoCount() int {
	return mi.StreamCount("VideoCount")
}

// AudioCount returns count of audio streams
func (mi *MediaInfo) AudioCount() int {
	return mi.StreamCount("AudioCount")
}

// TextCount returns count of texts
func (mi *MediaInfo) TextCount() int {
	return mi.StreamCount("TextCount")
}

// ImageCount returns count of images
func (mi *MediaInfo) ImageCount() int {
	return mi.StreamCount("ImageCount")
}
