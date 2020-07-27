package rnnoise

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L${SRCDIR}/res/lib -lrnnoise
#include "./res/libs/include/rnnoise.h"
*/
import "C"
import (
	"unsafe"
)

const FRAME_SIZE = 480

func Denoise(in *[]byte) *[]int16 {

	floatArr := intToFloat(*upto16bpps(*in))

	var DS *C.struct_DenoiseState
	DS = C.rnnoise_create(nil)
	//defer C.free(DS)
	out := make([]float32, FRAME_SIZE)
	counter := 0
	currPeace := make([]float32, FRAME_SIZE)
	all := make([]int16, 1)

	for index, _ := range floatArr {

		currPeace[counter] = floatArr[index]
		counter += 1

		if counter == FRAME_SIZE {
			value := C.rnnoise_process_frame(DS, (*C.float)(unsafe.Pointer(&out[0])), (*C.float)(unsafe.Pointer(&currPeace[0])))
			if (value * 100) >= 0.5 {
				for _, val := range out {
					all = append(all, int16(val))
				}
			}
			counter = 0
		}
	}
	return &all
}

func intToFloat(in []int16) []float32 {

	out := make([]float32, 1)

	for _, element := range in {
		tmp := float32(element)
		out = append(out, tmp)
	}
	return out
}

func upto16bpps(inbuffer []byte) *[]int16 {

	data := make([]int16, 1)

	for _, element := range inbuffer {
		tmpelem := int16(uint16(element-128) * 256)
		data = append(data, tmpelem)
	}

	return &data

}
