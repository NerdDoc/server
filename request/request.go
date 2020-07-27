package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mod/rnnoise"
	"mod/stt"
	"mod/tts"
	"mod/wave"
	"net/http"
	"server/rnnoise"
	"server/stt"
	"server/tts"
	"server/wave"
	"strings"
	"time"
)

var req Request

type Wave struct {
	BitsPerSample int           `json:"bitspersample"`
	Samplerate    int           `json:"samplerate"` //Samplerate of record.
	Duration      time.Duration `json: "duration"`
	WavData       []byte        `json:"wavdata"`
}

type Request struct {
	Wave        Wave
	BsRecordNrj float64 `json:"bsrecordnrj"` //Energy of record before  denoise by rnnoise - base energy of record
	Gender      string  `json:"gender"`      //Gender of voice assistant.
	Language    string  `json:"language"`    // Language of voice assistant
	MachineId   string  `json:"machineid"`
	QueryId     string  `json:"queryid"`
}

type PlayWaveReq struct {
	Wave Wave
	Beep bool
}

func Req(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get request")
	addr := strings.Split(r.RemoteAddr, ":")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("size of wav data")
	fmt.Println(len(req.Wave.WavData))
	denoisedwav := rnnoise.Denoise(&req.Wave.WavData)
	buf, err := wave.WriteWaveBuffer(1, req.Wave.Samplerate, 16, false, denoisedwav)
	if err != nil {
		fmt.Println(err)
	}

	text, err := stt.STT(buf)
	if err != nil {
		fmt.Println(err)
	}
	err = wave.WriteWaveByte("./received.wav", 1, req.Wave.Samplerate, req.Wave.BitsPerSample, false, &req.Wave.WavData)
	if err != nil {
		panic(err)
	}
	err = wave.WriteWaveInt16("./received_denoised.wav", 1, req.Wave.Samplerate, 16, false, denoisedwav)
	if err != nil {
		panic(err)
	}

	//tmp, err := wave.WriteWaveBuffer(1,req.Wave.Samplerate,16,false, denoisedwav)
	//if err != nil {
	//	fmt.Println(err)
	//}

	wav, err := tts.TTS("assistant", "8080", text, req.Language, req.Gender)
	if err != nil {
		fmt.Println(err)
	}

	var tmp = request.Wave{
		BitsPerSample: int(wav.Bitspersample),
		Samplerate:    int(wav.Samplerate),
		Duration:      wav.Duration,
		WavData:       wav.Data,
	}

	var req = request.PlayWaveReq{
		Wave: tmp,
		Beep: false,
	}

	data, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}

	reqt := bytes.NewReader(data)
	resp, err := http.Post("http://"+addr[0]+":9999/api/v1/play", "application/json", reqt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)

	//res, err := auth.GetListOfUsers("127.0.0.1","5000")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

	//res, err = auth.GetThreshold("127.0.0.1","5000")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//
	//res, err = auth.GetSamplerate("127.0.0.1","5000")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//
	//res, err = auth.GetNeuralnetwork("127.0.0.1","5000")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//
	//res, err = auth.SetNeuralnetwork("127.0.0.1","5000",true)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//
	//res, err = auth.SetSamplerate("127.0.0.1","5000",48000)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//
	//res, err = auth.SetThreshold("127.0.0.1","5000",0.9)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

	//res, err := auth.UserAdd("127.0.0.1", "5000", "Anna", tmp)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

	//res, err := auth.GetListOfUsers("127.0.0.1","5000")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

	//res, err = auth.AuthUser("127.0.0.1", "5000", "Unknown", tmp)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

	//res, err = auth.UserDel("127.0.0.1", "5000", "mMeu_053")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

}

//func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
//	var request Request
//	err := fastjson.NewDecoder(r.Body).Decode(&request)
//	if err != nil || len(request.WavData) == 0 {
//		log.Error("params decode problem %s", err)
//		return
//	}
//
//}
