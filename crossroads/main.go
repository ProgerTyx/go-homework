package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

type TrafficLight struct {
	color string
	timer int
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		time.Sleep(2 * time.Second)
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func (receiver *TrafficLight) changeLight() {
	switch receiver.color {
	case "green":
		receiver.color = "yellow"
	case "yellow":
		receiver.color = "red"
	case "red":
		receiver.color = "green"
	default:
		fmt.Println(errors.New("unknown color"))
	}
}

func (receiver *TrafficLight) start(delay int) {
	receiver.timer = delay
	ticker := time.NewTicker(time.Second * 1)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go func(waitGroup *sync.WaitGroup) {
		for range ticker.C {
			fmt.Println(receiver.timer)
			CallClear()
			receiver.timer--
			if receiver.timer == 0 {
				receiver.changeLight()
				fmt.Println(receiver.color)
				os.Stdout.Fd()
				waitGroup.Done()
			}
		}

	}(&waitGroup)
	waitGroup.Wait()
}

func main() {
	slice := [4]TrafficLight{
		TrafficLight{"red", 0},
		TrafficLight{"red", 0},
		TrafficLight{"green", 0},
		TrafficLight{"green", 0},
	}

	for _, v := range slice {
		v.start(30)
	}
}
