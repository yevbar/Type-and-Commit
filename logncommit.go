package main

import (
	"fmt"
	"github.com/MarinX/keylogger"
	"os"
	"os/exec"
	"strings"
	"time"
)

func FormatTime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func Commit(Input string, date time.Time) {
	Formatted := FormatTime(date)

	os.Setenv("GIT_AUTHOR_DATE", Formatted)
	os.Setenv("GIT_COMMIT_DATE", Formatted)

	msg := fmt.Sprintf("Typed " + Input)
	exec.Command("git", "commit", "--allow-empty", "-m "+msg).Start()
	fmt.Println("You typed " + Input)
}

func main() {
	devs, err := keylogger.NewDevices()
	if err != nil {
		panic(err)
	}

	for _, val := range devs {
		fmt.Println("Id->", val.Id, "Device->", val.Name)
	}

	rd := keylogger.NewKeyLogger(devs[3])

	in, err := rd.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	var LastPressed string = ""
	TheTime := time.Now()
	TheTime = TheTime.AddDate(-1, 0, -5)

	for i := range in {
		if i.Type == keylogger.EV_KEY {
			var CurrentlyPressed = i.KeyString()
			if CurrentlyPressed != "" {
				if CurrentlyPressed != LastPressed {
					if IsDesired(CurrentlyPressed) {
						Commit(CurrentlyPressed, TheTime)
						TheTime = TheTime.AddDate(0, 0, 1)
						if time.Now().Before(TheTime) {
							TheTime = time.Now()
							TheTime = TheTime.AddDate(-1, 0, -5)
						}
						LastPressed = CurrentlyPressed
					} else if CurrentlyPressed == "SPACE" {
						Commit("SPACE", TheTime)
						TheTime = TheTime.AddDate(0, 0, 1)
						if time.Now().Before(TheTime) {
							TheTime = time.Now()
							TheTime = TheTime.AddDate(-1, 0, -5)
						}
						LastPressed = CurrentlyPressed
					} else if CurrentlyPressed == "BS" {
						Commit("BACKSPACE", TheTime)
						TheTime = TheTime.AddDate(0, 0, 1)
						if time.Now().Before(TheTime) {
							TheTime = time.Now()
							TheTime = TheTime.AddDate(-1, 0, -5)
						}
						LastPressed = CurrentlyPressed
					}
				} else {
					LastPressed = ""
				}
			}
		}

	}
}
