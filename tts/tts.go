package tts

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ghostiam/binstruct"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"server/wave"
	"time"
)

const SynthesConnError = "syn_srv_conn_error.wav"

type WriteCloser struct {
	*bufio.Writer
}

func (wc *WriteCloser) Close() error {
	return nil
}

func TTS(host string, port string, text string, lang string, gender string) (*wave.WaveStruct, error) {

	wav, err := getwav(host, port, text, lang, gender)
	if err != nil {
		return nil, err
	}

	if wav.Bitspersample == 16 {
		reader := binstruct.NewReaderFromBytes(wav.Data, binary.LittleEndian, false)

		byteslen := len(wav.Data) / 2
		counter := 0
		newarr := make([]byte, 1)
		for true {
			if counter <= byteslen {
				val, err := reader.ReadInt16()
				if err == io.EOF {
					oldRate := wav.Bitspersample
					wav.Bitspersample = 8
					wav.Byterate = wav.Byterate / (oldRate / wav.Bitspersample)
					break
				}
				if err != nil {
					return nil, err
				}
				tmp := byte((int32(val) + 32767) >> 8)
				newarr = append(newarr, tmp)
				counter++
			}
		}
		wav.Data = newarr
		wav.Duration = time.Duration(1 + len(wav.Data)/int(wav.Bitspersample/8)/int(wav.Numchannels)/int(wav.Samplerate))
		return wav, nil
	}
	if wav.Bitspersample == 8 {
		wav.Duration = time.Duration(1 + len(wav.Data)/int(wav.Bitspersample/8)/int(wav.Numchannels)/int(wav.Samplerate))
		return wav, nil
	}

	return nil, errors.New("Possible bits per sample 8 and 16")

}

func Save(host string, port string, text string, path string, lang string, gender string) error {

	wav, err := getwav(host, port, text, lang, gender)
	if err != nil {
		return err
	}
	err = wave.WriteWaveByte(path, int(wav.Numchannels), int(wav.Samplerate), int(wav.Bitspersample), false, &wav.Data)
	if err != nil {
		return err
	}
	return nil
}

func getwav(host string, port string, text string, lang string, gender string) (*wave.WaveStruct, error) {

	var bb bytes.Buffer
	bw := bufio.NewWriter(&bb)
	outBuff := &WriteCloser{bw}

	Url := "http://" + host + ":" + port + "/api/v1/say?text=" + url.QueryEscape(text) + "&rate=30"

	resp, err := http.Get(Url)

	if err != nil {
		fmt.Println(err)
		wavbuf, err := wave.ReadFileToBuf("./res/wavs/languages/" + lang + "/" + gender + "/" + SynthesConnError)
		if err != nil {
			return nil, err
		}
		return wavbuf, err

	} else {

		if resp.StatusCode != 200 {
			fmt.Println(err)
			wavbuf, err := wave.ReadFileToBuf("./res/wavs/languages/" + lang + "/" + gender + "/" + SynthesConnError)
			if err != nil {
				return nil, err
			}
			return wavbuf, err

		}
	}

	defer resp.Body.Close()

	_, err = io.Copy(outBuff, resp.Body)
	if err != nil {
		return nil, err
	}

	wav := wave.ParseWaveStruct(bb.Bytes())

	return wav, nil
}
