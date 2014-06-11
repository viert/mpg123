package mpg123

/*
#cgo LDFLAGS: -lmpg123
#include "mpg123.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

const (
	MPG123_DONE              = C.MPG123_DONE
	MPG123_NEW_FORMAT        = C.MPG123_NEW_FORMAT
	MPG123_NEED_MORE         = C.MPG123_NEED_MORE
	MPG123_ERR               = C.MPG123_ERR
	MPG123_OK                = C.MPG123_OK
	MPG123_BAD_OUTFORMAT     = C.MPG123_BAD_OUTFORMAT
	MPG123_BAD_CHANNEL       = C.MPG123_BAD_CHANNEL
	MPG123_BAD_RATE          = C.MPG123_BAD_RATE
	MPG123_ERR_16TO8TABLE    = C.MPG123_ERR_16TO8TABLE
	MPG123_BAD_PARAM         = C.MPG123_BAD_PARAM
	MPG123_BAD_BUFFER        = C.MPG123_BAD_BUFFER
	MPG123_OUT_OF_MEM        = C.MPG123_OUT_OF_MEM
	MPG123_NOT_INITIALIZED   = C.MPG123_NOT_INITIALIZED
	MPG123_BAD_DECODER       = C.MPG123_BAD_DECODER
	MPG123_BAD_HANDLE        = C.MPG123_BAD_HANDLE
	MPG123_NO_BUFFERS        = C.MPG123_NO_BUFFERS
	MPG123_BAD_RVA           = C.MPG123_BAD_RVA
	MPG123_NO_GAPLESS        = C.MPG123_NO_GAPLESS
	MPG123_NO_SPACE          = C.MPG123_NO_SPACE
	MPG123_BAD_TYPES         = C.MPG123_BAD_TYPES
	MPG123_BAD_BAND          = C.MPG123_BAD_BAND
	MPG123_ERR_NULL          = C.MPG123_ERR_NULL
	MPG123_ERR_READER        = C.MPG123_ERR_READER
	MPG123_NO_SEEK_FROM_END  = C.MPG123_NO_SEEK_FROM_END
	MPG123_BAD_WHENCE        = C.MPG123_BAD_WHENCE
	MPG123_NO_TIMEOUT        = C.MPG123_NO_TIMEOUT
	MPG123_BAD_FILE          = C.MPG123_BAD_FILE
	MPG123_NO_SEEK           = C.MPG123_NO_SEEK
	MPG123_NO_READER         = C.MPG123_NO_READER
	MPG123_BAD_PARS          = C.MPG123_BAD_PARS
	MPG123_BAD_INDEX_PAR     = C.MPG123_BAD_INDEX_PAR
	MPG123_OUT_OF_SYNC       = C.MPG123_OUT_OF_SYNC
	MPG123_RESYNC_FAIL       = C.MPG123_RESYNC_FAIL
	MPG123_NO_8BIT           = C.MPG123_NO_8BIT
	MPG123_BAD_ALIGN         = C.MPG123_BAD_ALIGN
	MPG123_NULL_BUFFER       = C.MPG123_NULL_BUFFER
	MPG123_NO_RELSEEK        = C.MPG123_NO_RELSEEK
	MPG123_NULL_POINTER      = C.MPG123_NULL_POINTER
	MPG123_BAD_KEY           = C.MPG123_BAD_KEY
	MPG123_NO_INDEX          = C.MPG123_NO_INDEX
	MPG123_INDEX_FAIL        = C.MPG123_INDEX_FAIL
	MPG123_BAD_DECODER_SETUP = C.MPG123_BAD_DECODER_SETUP
	MPG123_MISSING_FEATURE   = C.MPG123_MISSING_FEATURE
	MPG123_BAD_VALUE         = C.MPG123_BAD_VALUE
	MPG123_LSEEK_FAILED      = C.MPG123_LSEEK_FAILED
	MPG123_BAD_CUSTOM_IO     = C.MPG123_BAD_CUSTOM_IO
	MPG123_LFS_OVERFLOW      = C.MPG123_LFS_OVERFLOW
	MPG123_INT_OVERFLOW      = C.MPG123_INT_OVERFLOW

	MPG123_VERBOSE                   = C.MPG123_VERBOSE
	MPG123_FLAGS                     = C.MPG123_FLAGS
	MPG123_ADD_FLAGS                 = C.MPG123_ADD_FLAGS
	MPG123_FORCE_RATE                = C.MPG123_FORCE_RATE
	MPG123_DOWN_SAMPLE               = C.MPG123_DOWN_SAMPLE
	MPG123_RVA                       = C.MPG123_RVA
	MPG123_DOWNSPEED                 = C.MPG123_DOWNSPEED
	MPG123_UPSPEED                   = C.MPG123_UPSPEED
	MPG123_START_FRAME               = C.MPG123_START_FRAME
	MPG123_DECODE_FRAMES             = C.MPG123_DECODE_FRAMES
	MPG123_ICY_INTERVAL              = C.MPG123_ICY_INTERVAL
	MPG123_OUTSCALE                  = C.MPG123_OUTSCALE
	MPG123_TIMEOUT                   = C.MPG123_TIMEOUT
	MPG123_REMOVE_FLAGS              = C.MPG123_REMOVE_FLAGS
	MPG123_RESYNC_LIMIT              = C.MPG123_RESYNC_LIMIT
	MPG123_INDEX_SIZE                = C.MPG123_INDEX_SIZE
	MPG123_PREFRAMES                 = C.MPG123_PREFRAMES
	MPG123_FEEDPOOL                  = C.MPG123_FEEDPOOL
	MPG123_FEEDBUFFER                = C.MPG123_FEEDBUFFER
	MPG123_MONO_LEFT                 = C.MPG123_MONO_LEFT
	MPG123_MONO_RIGHT                = C.MPG123_MONO_RIGHT
	MPG123_MONO_MIX                  = C.MPG123_MONO_MIX
	MPG123_FORCE_STEREO              = C.MPG123_FORCE_STEREO
	MPG123_FORCE_8BIT                = C.MPG123_FORCE_8BIT
	MPG123_QUIET                     = C.MPG123_QUIET
	MPG123_GAPLESS                   = C.MPG123_GAPLESS
	MPG123_NO_RESYNC                 = C.MPG123_NO_RESYNC
	MPG123_SEEKBUFFER                = C.MPG123_SEEKBUFFER
	MPG123_FUZZY                     = C.MPG123_FUZZY
	MPG123_FORCE_FLOAT               = C.MPG123_FORCE_FLOAT
	MPG123_PLAIN_ID3TEXT             = C.MPG123_PLAIN_ID3TEXT
	MPG123_IGNORE_STREAMLENGTH       = C.MPG123_IGNORE_STREAMLENGTH
	MPG123_SKIP_ID3V2                = C.MPG123_SKIP_ID3V2
	MPG123_IGNORE_INFOFRAME          = C.MPG123_IGNORE_INFOFRAME
	MPG123_AUTO_RESAMPLE             = C.MPG123_AUTO_RESAMPLE
	MPG123_RVA_MIX                   = C.MPG123_RVA_MIX
	MPG123_RVA_ALBUM                 = C.MPG123_RVA_ALBUM
	MPG123_RVA_MAX                   = C.MPG123_RVA_MAX
	MPG123_FEATURE_OUTPUT_8BIT       = C.MPG123_FEATURE_OUTPUT_8BIT
	MPG123_FEATURE_OUTPUT_16BIT      = C.MPG123_FEATURE_OUTPUT_16BIT
	MPG123_FEATURE_OUTPUT_32BIT      = C.MPG123_FEATURE_OUTPUT_32BIT
	MPG123_FEATURE_INDEX             = C.MPG123_FEATURE_INDEX
	MPG123_FEATURE_PARSE_ID3V2       = C.MPG123_FEATURE_PARSE_ID3V2
	MPG123_FEATURE_DECODE_LAYER1     = C.MPG123_FEATURE_DECODE_LAYER1
	MPG123_FEATURE_DECODE_LAYER2     = C.MPG123_FEATURE_DECODE_LAYER2
	MPG123_FEATURE_DECODE_LAYER3     = C.MPG123_FEATURE_DECODE_LAYER3
	MPG123_FEATURE_DECODE_ACCURATE   = C.MPG123_FEATURE_DECODE_ACCURATE
	MPG123_FEATURE_DECODE_DOWNSAMPLE = C.MPG123_FEATURE_DECODE_DOWNSAMPLE
	MPG123_FEATURE_DECODE_NTOM       = C.MPG123_FEATURE_DECODE_NTOM
	MPG123_FEATURE_PARSE_ICY         = C.MPG123_FEATURE_PARSE_ICY
	MPG123_FEATURE_TIMEOUT_READ      = C.MPG123_FEATURE_TIMEOUT_READ
)

var initialized = false

type Handle *[0]byte

type CodeError struct {
	ret C.int
}

type HandleError struct {
	handle Handle
}

func (he *HandleError) Error() string {
	err := C.mpg123_strerror(he.handle)
	return C.GoString(err)
}

func (mce *CodeError) Error() string {
	err := C.mpg123_plain_strerror(mce.ret)
	return C.GoString(err)
}

func notInitializedError() error {
	return errors.New("mpg123 is not initialized, call mpg123.Init() first")
}

func Init() error {
	initialized = true
	ret := C.mpg123_init()
	if ret != MPG123_OK {
		return &CodeError{ret}
	}
	return nil
}

func New(decoder string) (Handle, error) {
	if !initialized {
		return nil, notInitializedError()
	}

	var ret C.int
	var handle *[0]byte

	if decoder == "" {
		handle = C.mpg123_new(nil, &ret)
	} else {
		handle = C.mpg123_new(C.CString(decoder), &ret)
	}

	if ret == MPG123_OK {
		return handle, nil
	}

	return nil, &CodeError{ret}
}

func OpenFeed(handle Handle) error {
	ret := C.mpg123_open_feed(handle)
	if ret == MPG123_OK {
		return nil
	}
	return &CodeError{ret}
}

func Param(handle Handle, param C.int, value int, fvalue float64) error {
	ret := C.mpg123_param(handle, uint32(param), C.long(value), C.double(fvalue))
	if ret == MPG123_OK {
		return nil
	}
	return &CodeError{ret}
}

func Decode(handle Handle, input *[]byte, maxOutputSize int) (output []byte, retcode C.int, err error) {
	output = make([]byte, maxOutputSize)
	err = nil

	var cInput *C.uchar
	var cInputLen C.size_t

	if input == nil {
		cInput = nil
		cInputLen = C.size_t(0)
	} else {
		cInput = (*C.uchar)(unsafe.Pointer(&(*input)[0]))
		cInputLen = C.size_t(len(*input))
	}

	cOutput := (*C.uchar)(unsafe.Pointer(&output[0]))
	cOutputLen := C.size_t(maxOutputSize)

	var cSize C.size_t

	retcode = C.mpg123_decode(handle, cInput, cInputLen, cOutput, cOutputLen, &cSize)
	output = output[0:int(cSize)]
	if retcode == MPG123_ERR {
		err = &HandleError{handle}
	}
	return
}

func Delete(handle Handle) {
	C.mpg123_delete(handle)
}

func Exit() {
	initialized = false
	C.mpg123_exit()
}
