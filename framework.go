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
	// Executes otto.applescript to send the actual message to the iMessage chat
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
	// only lets same user call Otto 5 times in a row
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
	fulltext := strings.Split(os.Args[1:][0], "|~|")
	message, from, chatid, settingslocation := fulltext[0], fulltext[1], fulltext[2], fulltext[3]
	Data = readandparsesettings(settingslocation)
	ottomessage := false
	if strings.ToLower(message[:4]) != "" {
		allowedtorun := checkandwriteallowed(from, chatid)
		if allowedtorun {
			hasntBeenCalled := true
			for key, value := range ottomap
}
