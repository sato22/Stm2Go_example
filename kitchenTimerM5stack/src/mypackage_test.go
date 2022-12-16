// This file was generated by a program.
package mypackage

import (
	"fmt"
	"log"
	"testing"
	"time"

	stm2go "github.com/JuliaMBD/stm2go/testing"
)

type DebugStruct struct{}

func (l DebugStruct) Println(s string) {
	log.Println(s)
}

// Button
type Button struct {
	name    string
	release bool // true → release状態，false → push状態
}

func (b *Button) Push() {
	b.release = false
	log.Println(b.name, "Push")
}

func (b *Button) Release() {
	b.release = true
	log.Println(b.name, "Release")
}

func (b Button) Get() bool {
	return b.release
}

// Screen
type Screen struct{}

func (s Screen) PrintVal(str string) {
	logger.Println(fmt.Sprintf("screen display: %s", str))
}

// Alarm
type Buzzer struct{}

func (b Buzzer) Beep() {
	logger.Println("alarm beep!")
}

func (b Buzzer) Mute() {
	logger.Println("alarm mute")
}

// define struct
var leftButton = &Button{"leftButton", true}
var middleButton = &Button{"middleButton", true}
var rightButton = &Button{"rightButton", true}

var screen = &Screen{}

var buzzer = &Buzzer{}

func TestExample(t *testing.T) {
	env := stm2go.NewTestEnv() // TestEnv構造体

	ConfigureLeftButton(leftButton)
	ConfigureMiddleButton(middleButton)
	ConfigureRightButton(rightButton)

	ConfigureMonitor(screen)

	ConfigureAlarm(buzzer)

	ConfigureLog(DebugStruct{})

	// goroutine(base.go Task())
	env.Add(stm2go.Continue, func() {
		for {
			time.Sleep(10 * time.Millisecond)
			stm0Task()
		}
	},
	)

	// goroutine(user operation)
	env.Add(stm2go.Done, func() {
		logger.Println("----------------leftButtonPush-------------")
		env.Sleep(1 * time.Second)
		leftButton.Push()
		env.Sleep(50 * time.Millisecond)
		leftButton.Release()

		logger.Println("----------------middleButtonPush-------------")
		env.Sleep(1 * time.Second)
		middleButton.Push()
		env.Sleep(50 * time.Millisecond)
		middleButton.Release()

		logger.Println("----------------rightButtonPush-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(50 * time.Millisecond)
		rightButton.Release()
		env.Sleep(15 * time.Second)
	},
	)

	env.Set(1)
	env.Go()
}
