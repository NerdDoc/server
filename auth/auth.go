package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// MsgId - 15 - Get/set use neuralnetwork return {neuralnet} (211)
// MsgId - 14 - Get/set samplerate return {samplerate}(210)
// MsgId - 13 - Get/set threshold return {threshold}- (209)
// MsgId - 12 - Get list of users  - return []{userid - username} (202)
// MsgId - 11 - User authenticated - return {userid - username} (200)
// MsgId - 10 - User added/deleted - return {userid - username}(201)
// MsgId - 9 - Can not identify user (400)
// MsgId - 8 - Can not add user (400)
// MsgId - 7 - User exists (400)
// MsgId - 6 - Empty database. Add users first (400)
// MsgId - 5 - Empty file (400)
// MsgId - 4 - Value is not type of bool (400)
// MsgId - 3 - Value is not type of integer (400)
// MsgId - 2 - Value is not type of float (400)
// MsgId - 1 - Specify user id (400)

type User struct {
	Name   string `json:"name"`
	UserId string `json:"userid"`
}

type ErrorMessage struct {
	MessageTxt string `json:"msgtxt"`
	MessageId  string `json:"msgid"`
}

type ListOfUsers struct {
	MessageId string `json:"msgid"`
	Users     []User
}

type Threshold struct {
	MessageId string  `json:"msgid"`
	Threshold float32 `json:"threshold"`
}

type Samplerate struct {
	MessageId  string `json:"msgid"`
	Samplerate string `json:"samplerate"`
}

type Neuralnet struct {
	MessageId string `json:"msgid"`
	Neurlanet bool   `json:"neurlanet"`
}

type Auth struct {
	Name      string `json:"name"`
	UserId    string `json:"userid"`
	MessageId string `json:"msgid"`
}

var errMessage ErrorMessage
var listOfUsers ListOfUsers
var threshold Threshold
var samplerate Samplerate
var neuralnet Neuralnet
var auth Auth

func GetListOfUsers(host string, port string) (string, error) {

	listOfUsers = ListOfUsers{}
	resp, err := http.Get("http://" + host + ":" + port + "/auth/api/v1.0/list")
	if err != nil {
		return "", err
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 400 {
		err = json.Unmarshal(bytes, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New(errMessage.MessageId)
	}
	if resp.StatusCode == 202 {
		err = json.Unmarshal(bytes, &listOfUsers)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(listOfUsers.MessageId)
		fmt.Println(listOfUsers.Users)
		return "", errors.New("OK")
	}
	return "", nil
}

func UserAdd(host string, port string, username string, wave *bytes.Buffer) (string, error) {

	listOfUsers = ListOfUsers{}
	uri := "http://" + host + ":" + port + "/auth/api/v1.0/useradd"
	filename := username + ".wav"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	_, err = io.Copy(part, wave)
	err = writer.Close()

	resp, _ := http.DefaultClient.Post(uri, writer.FormDataContentType(), body)
	defer resp.Body.Close()
	response_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(response_body, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 201 {
		err = json.Unmarshal(response_body, &listOfUsers)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(listOfUsers.MessageId)
		fmt.Println(listOfUsers.Users)
		return "", errors.New("OK")
	}
	return "", nil
}

func UserDel(host string, port string, id string) (string, error) {

	req, err := http.NewRequest("DELETE", "http://"+host+":"+port+"/auth/api/v1.0/userdel/"+id, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 400 {
		err = json.Unmarshal(bytes, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 201 {
		err = json.Unmarshal(bytes, &listOfUsers)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(listOfUsers.MessageId)
		fmt.Println(listOfUsers.Users)
		return "", errors.New("OK")
	}
	return "", nil
}

func AuthUser(host string, port string, username string, wave *bytes.Buffer) (string, error) {

	listOfUsers = ListOfUsers{}
	uri := "http://" + host + ":" + port + "/auth/api/v1.0/auth"
	filename := username + ".wav"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	_, err = io.Copy(part, wave)
	err = writer.Close()

	resp, _ := http.DefaultClient.Post(uri, writer.FormDataContentType(), body)
	defer resp.Body.Close()
	response_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(response_body, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 200 {
		err = json.Unmarshal(response_body, &auth)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(auth.MessageId)
		fmt.Println(auth.Name)
		return "", errors.New("OK")
	}
	return "", nil
}

func GetThreshold(host string, port string) (string, error) {

	resp, err := http.Get("http://" + host + ":" + port + "/auth/api/v1.0/threshold")
	if err != nil {
		return "", err
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 400 {
		err = json.Unmarshal(bytes, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 209 {
		err = json.Unmarshal(bytes, &threshold)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(threshold.MessageId)
		fmt.Println(threshold.Threshold)
		return "", errors.New("OK")
	}
	return "", nil

}

func GetSamplerate(host string, port string) (string, error) {

	resp, err := http.Get("http://" + host + ":" + port + "/auth/api/v1.0/samplerate")
	if err != nil {
		return "", err
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 400 {
		err = json.Unmarshal(bytes, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 210 {
		err = json.Unmarshal(bytes, &samplerate)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(samplerate.MessageId)
		fmt.Println(samplerate.Samplerate)
		return "", errors.New("OK")
	}
	return "", nil

}

func GetNeuralnetwork(host string, port string) (string, error) {

	resp, err := http.Get("http://" + host + ":" + port + "/auth/api/v1.0/neuralnet")
	if err != nil {
		return "", err
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 400 {
		err = json.Unmarshal(bytes, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 211 {
		err = json.Unmarshal(bytes, &neuralnet)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(neuralnet.MessageId)
		fmt.Println(neuralnet.Neurlanet)
		return "", errors.New("OK")
	}
	return "", nil
}

func SetNeuralnetwork(host string, port string, useNN bool) (string, error) {

	NN := "True"
	if useNN == false {
		NN = "False"
	}

	type Payload struct {
		Neural string `json:"neuralnet"`
	}
	data := Payload{Neural: NN}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	fmt.Println(string(payloadBytes))
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://"+host+":"+port+"/auth/api/v1.0/neuralnet", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		err = json.Unmarshal(bytes, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 211 {
		err = json.Unmarshal(bytes, &neuralnet)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(neuralnet.MessageId)
		fmt.Println(neuralnet.Neurlanet)
		return "", errors.New("OK")
	}
	return "", nil
}

func SetSamplerate(host string, port string, sr int) (string, error) {

	type Payload struct {
		Samplerate int `json:"samplerate"`
	}
	data := Payload{Samplerate: sr}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	fmt.Println(string(payloadBytes))
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://"+host+":"+port+"/auth/api/v1.0/samplerate", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		err = json.Unmarshal(bytes, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 210 {
		err = json.Unmarshal(bytes, &samplerate)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(samplerate.MessageId)
		fmt.Println(samplerate.Samplerate)
		return "", errors.New("OK")
	}
	return "", nil
}

func SetThreshold(host string, port string, tshold float32) (string, error) {

	type Payload struct {
		Threshold float32 `json:"threshold"`
	}
	data := Payload{Threshold: tshold}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	fmt.Println(string(payloadBytes))
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://"+host+":"+port+"/auth/api/v1.0/threshold", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		err = json.Unmarshal(bytes, &errMessage)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(errMessage.MessageId)
		fmt.Println(errMessage.MessageTxt)
		return "", errors.New("empty")
	}
	if resp.StatusCode == 209 {
		err = json.Unmarshal(bytes, &threshold)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(threshold.MessageId)
		fmt.Println(threshold.Threshold)
		return "", errors.New("OK")
	}
	return "", nil
}
