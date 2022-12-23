// This file was generated by a program.
// Please do not edit this file directly.
package mypackage

import (
// package names to be imported
)

type AlarmOnStm3State uint8

const (
	AlarmOnStm3Beep AlarmOnStm3State = iota
	AlarmOnStm3AlarmOnWaitRelease
)

var AlarmOnStm3Eod Eod

var AlarmOnStm3CurrentState AlarmOnStm3State

func init() {
	AlarmOnStm3Initialize()
}

func AlarmOnStm3Initialize() {
	AlarmOnStm3CurrentState = AlarmOnStm3Beep
}

func AlarmOnStm3Task() {
	switch AlarmOnStm3CurrentState {
	case AlarmOnStm3Beep:
		if AlarmOnStm3Eod == Entry {
			AlarmOnStm3BeepEntry()
			AlarmOnStm3Eod = Do
		}
		if AlarmOnStm3Eod == Do {
			AlarmOnStm3BeepDo()
			if AlarmOnStm3BeepstopButtonPushCond() {
				AlarmOnStm3BeepstopButtonPushAction()
				AlarmOnStm3CurrentState = AlarmOnStm3AlarmOnWaitRelease
				AlarmOnStm3Eod = Exit
			}
		}
		if AlarmOnStm3Eod == Exit {
			AlarmOnStm3BeepExit()
			AlarmOnStm3Eod = Entry
		}
	case AlarmOnStm3AlarmOnWaitRelease:
		if AlarmOnStm3Eod == Entry {
			AlarmOnStm3AlarmOnWaitReleaseEntry()
			AlarmOnStm3Eod = Do
		}
		if AlarmOnStm3Eod == Do {
			AlarmOnStm3AlarmOnWaitReleaseDo()
		}
		if AlarmOnStm3Eod == Exit {
			AlarmOnStm3AlarmOnWaitReleaseExit()
			AlarmOnStm3Eod = Entry
		}
	}
}