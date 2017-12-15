package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func send(message, chatid string) {
	mybuddy := fmt.Sprintf("set mybuddy to a reference to text chat id \"%s\"", chatid)
	send := fmt.Sprintf("send \"%s\" to mybuddy", strings.Replace(message, "\"", "\\\"", -1))
	exec.Command("/usr/bin/osascript", "-e", "tell application \"Messages\"", "-e", mybuddy, "-e", send, "-e", "end tell").Run()
}
func testsend(message, chatid string) {
	fmt.Println(message)
}

func readandparsesettings(location string) Results {

	file, err := ioutil.ReadFile(location)

	if err != nil {
		panic(err)
	}
	Data := Results{}
	err = json.Unmarshal(file, &Data)
	if err != nil {
		newlocation := fmt.Sprintf("%sbackup.json", location[:len(location)-5])
		//try backup
		newfile, err := ioutil.ReadFile(newlocation)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(newfile, &Data)
		if err != nil {
			panic(err)
		}
	}
	return Data
}
func writesettings(location string, Data Results) error {
	jsondata, err := json.MarshalIndent(Data, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(location, jsondata, 0644)
	if err != nil {
		return err
	}
	return nil
}

func checkandwriteallowed(from, chatid string) bool {
	allowedtorun := true
	if from == Data.Chat.Lastperson {
		Data.Chat.Lastamount += 1
		if Data.Chat.Lastamount > 5 {
			allowedtorun = false
		}
		if Data.Chat.Lastamount == 5 {
			send(Data.Maxmessage, chatid)
		}

	} else {
		Data.Chat.Lastperson = from
		Data.Chat.Lastamount = 1
	}
	return allowedtorun
}

var Data Results

func main() {

}
