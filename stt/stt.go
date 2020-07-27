package stt

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io"
	"net/url"
)

const Host = "assistant"
const Port = "2700"
const buffsize = 8000

type Message struct {
	Result []struct {
		Conf  float64
		End   float64
		Start float64
		Word  string
	}
	Text string
}

var m Message

func STT(data *bytes.Buffer) (string, error) {

	u := url.URL{Scheme: "ws", Host: Host + ":" + Port, Path: ""}
	log.Info("connecting to ", u.String())

	// Opening websocket connection
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return "", err
	}

	defer c.Close()

	for {
		buf := make([]byte, buffsize)
		dat, err := data.Read(buf)
		//buf := make([]byte, buffsize)
		//dat, err := f.Read(data)

		if dat == 0 && err == io.EOF {
			err = c.WriteMessage(websocket.TextMessage, []byte("{\"eof\" : 1}"))
			if err != nil {
				return "", err
			}
			break
		}

		err = c.WriteMessage(websocket.BinaryMessage, buf)
		if err != nil {
			return "", nil
		}

		// Read message from server
		_, _, err = c.ReadMessage()
		if err != nil {
			return "", nil
		}
	}

	// Read final message from server
	_, msg, err := c.ReadMessage()
	if err != nil {
		return "", nil
	}

	// Closing websocket connection
	c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

	// Unmarshalling received message
	err = json.Unmarshal(msg, &m)
	if err != nil {
		return "", nil
	}
	log.Info(m.Text)
	return m.Text, nil
}
