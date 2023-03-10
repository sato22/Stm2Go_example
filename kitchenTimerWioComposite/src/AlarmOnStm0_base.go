// This file was generated by a program.
// Please do not edit this file directly.
package mypackage

// package names to be imported

type AlarmOnStm0State uint8

const (
	AlarmOnStm0Beep AlarmOnStm0State = iota
	AlarmOnStm0AlarmOnWaitRelease
)

var AlarmOnStm0Eod Eod

var AlarmOnStm0CurrentState AlarmOnStm0State

func init() {
	AlarmOnStm0Initialize()
}

func AlarmOnStm0Initialize() {
	AlarmOnStm0CurrentState = AlarmOnStm0Beep
}

func AlarmOnStm0Task() {
	switch AlarmOnStm0CurrentState {
	case AlarmOnStm0Beep:
		if AlarmOnStm0Eod == Entry {
			AlarmOnStm0BeepEntry()
			AlarmOnStm0Eod = Do
		}
		if AlarmOnStm0Eod == Do {
			AlarmOnStm0BeepDo()
			if AlarmOnStm0BeepstopButtonPushCond() {
				AlarmOnStm0BeepstopButtonPushAction()
				AlarmOnStm0CurrentState = AlarmOnStm0AlarmOnWaitRelease
				AlarmOnStm0Eod = Exit
			}
		}
		if AlarmOnStm0Eod == Exit {
			AlarmOnStm0BeepExit()
			AlarmOnStm0Eod = Entry
		}
	case AlarmOnStm0AlarmOnWaitRelease:
		if AlarmOnStm0Eod == Entry {
			AlarmOnStm0AlarmOnWaitReleaseEntry()
			AlarmOnStm0Eod = Do
		}
		if AlarmOnStm0Eod == Do {
			AlarmOnStm0AlarmOnWaitReleaseDo()
		}
		if AlarmOnStm0Eod == Exit {
			AlarmOnStm0AlarmOnWaitReleaseExit()
			AlarmOnStm0Eod = Entry
		}
	}
}
