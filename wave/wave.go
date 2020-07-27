package wave

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/NerdDoc/go-wave"
	"math"
	"os"
	"time"
)

type WaveStruct struct {
	ChunkSize     int
	Numchannels   int
	Samplerate    int
	Byterate      int
	Bitspersample int
	Duration      time.Duration
	Data          []byte
}

type WriteCloser struct {
	*bufio.Writer
}

func (wc *WriteCloser) Close() error {
	return nil
}

// Writing wave to file
// if rawpcm - true - writing rawpcm without header
func WriteWaveByte(audioFileName string, inputchannel int, samplerate int, bitspersample int, rawpcm bool, data *[]byte) error {

	waveFile, err := os.Create(audioFileName)
	if err != nil {
		return err
	}

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputchannel,
		SampleRate:    samplerate,
		BitsPerSample: bitspersample,
		RawPcm:        rawpcm,
	}
	waveWriter, err := wave.NewWriter(param)
	if err != nil {
		return err
	}

	if bitspersample == 8 {
		_, err = waveWriter.Write(*data)
	}
	if err != nil {
		return err
	}

	if bitspersample == 16 {
		_, err = waveWriter.WriteSample16(*upto16bpps(*data))
	}
	if err != nil {
		return err
	}

	err = waveWriter.Close()
	if err != nil {
		return err
	}

	return nil
}

func WriteWaveInt16(audioFileName string, inputchannel int, samplerate int, bitspersample int, rawpcm bool, data *[]int16) error {
	waveFile, err := os.Create(audioFileName)
	if err != nil {
		return err
	}

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputchannel,
		SampleRate:    samplerate,
		BitsPerSample: bitspersample,
		RawPcm:        rawpcm,
	}
	waveWriter, err := wave.NewWriter(param)
	if err != nil {
		return err
	}

	if bitspersample == 16 {
		_, err = waveWriter.WriteSample16(*data)
	}
	if err != nil {
		return err
	}

	err = waveWriter.Close()
	if err != nil {
		return err
	}

	return nil
}

// Creating wave file in buffer
func WriteWaveBuffer(inputchannel int, samplerate int, bitspersample int, rawpcm bool, data *[]int16) (*bytes.Buffer, error) {

	var bb bytes.Buffer
	bw := bufio.NewWriter(&bb)
	outBuff := &WriteCloser{bw}

	param := wave.WriterParam{
		Out:           outBuff,
		Channel:       inputchannel,
		SampleRate:    samplerate,
		BitsPerSample: bitspersample,
		RawPcm:        rawpcm,
	}
	waveWriter, err := wave.NewWriter(param)
	if err != nil {
		return nil, err
	}

	if bitspersample == 16 {
		_, err = waveWriter.WriteSample16(*data)
	}
	if err != nil {
		return nil, err
	}

	err = waveWriter.Close()
	if err != nil {
		return nil, err
	}

	return &bb, nil
}

// Dirty downsampling function
func Downsampling(in []int16, InputSamplerate int, OutputSamplerate int) (*[]int16, error) {

	if InputSamplerate < OutputSamplerate {
		err := errors.New("Input sample rate cannot be less then output")
		return nil, err
	}

	inputLenght := len(in)
	offsetResult := 0
	offsetBuffer := 0
	sampleRateRatio := InputSamplerate / OutputSamplerate
	newLenght := math.Round(float64(inputLenght / sampleRateRatio))
	result := make([]int16, int(newLenght))

	for {
		if offsetResult < int(newLenght) {
			nextOffsetBuffer := math.Round(float64((offsetResult + 1) * sampleRateRatio))
			var accum int16 = 0
			count := 0
			for i := offsetBuffer; i < int(nextOffsetBuffer) && i < inputLenght; i++ {
				accum += in[i]
				count++
			}
			result[offsetResult] = accum / int16(count)
			offsetResult++
			fmt.Println(offsetResult)
			offsetBuffer = int(nextOffsetBuffer)
		} else {
			break
		}
	}
	return &result, nil
}

func upto16bpps(inbuffer []byte) *[]int16 {

	data := make([]int16, 1)

	for _, element := range inbuffer {
		tmpelem := int16(uint16(element-128) * 256)
		data = append(data, tmpelem)
	}

	return &data

}

func Downtobyte(in []int16) *[]byte {

	data := make([]byte, 1)

	for _, element := range in {
		tmpelem := byte((int32(element) + 32767) >> 8)
		data = append(data, tmpelem)
	}

	return &data

}

func ParseWaveStruct(wave []byte) *WaveStruct {

	var tmp WaveStruct
	tmp.Numchannels = int(binary.LittleEndian.Uint16(wave[22:24]))
	tmp.Bitspersample = int(binary.LittleEndian.Uint16(wave[34:36]))
	tmp.Samplerate = int(binary.LittleEndian.Uint32(wave[24:28]))
	tmp.Byterate = int(binary.LittleEndian.Uint32(wave[28:32]))
	index := bytes.IndexAny(wave, "data")
	tmp.Data = wave[index+4:]
	tmp.Duration = time.Duration(1 + len(tmp.Data)/(tmp.Bitspersample/8)/tmp.Numchannels/tmp.Samplerate)
	return &tmp
}

func ReadFileToBuf(filename string) (*WaveStruct, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	wavstruct := ParseWaveStruct(buffer)

	return wavstruct, nil
}
