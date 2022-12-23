// This file was generated by a program.
// Please do not edit this file directly.
package mypackage

import (
// package names to be imported
)

type TimerSetStm1State uint8

const (
	TimerSetStm1Idle TimerSetStm1State = iota
	TimerSetStm1AddOneMin
	TimerSetStm1AddOneSec
	TimerSetStm1TimerWaitRelease
	TimerSetStm1Set
	TimerSetStm1AddMinHold
	TimerSetStm1AddSecHold
)

var TimerSetStm1Eod Eod

var TimerSetStm1CurrentState TimerSetStm1State

func init() {
	TimerSetStm1Initialize()
}

func TimerSetStm1Initialize() {
	TimerSetStm1CurrentState = TimerSetStm1Idle
}

func TimerSetStm1Task() {
	switch TimerSetStm1CurrentState {
	case TimerSetStm1Idle:
		if TimerSetStm1Eod == Entry {
			TimerSetStm1IdleEntry()
			TimerSetStm1Eod = Do
		}
		if TimerSetStm1Eod == Do {
			TimerSetStm1IdleDo()
			if TimerSetStm1IdleSecButtonPushCond() {
				TimerSetStm1IdleSecButtonPushAction()
				TimerSetStm1CurrentState = TimerSetStm1AddOneSec
				TimerSetStm1Eod = Exit
			}
			if TimerSetStm1IdleMinButtonPushCond() {
				TimerSetStm1IdleMinButtonPushAction()
				TimerSetStm1CurrentState = TimerSetStm1AddOneMin
				TimerSetStm1Eod = Exit
			}
		}
		if TimerSetStm1Eod == Exit {
			TimerSetStm1IdleExit()
			TimerSetStm1Eod = Entry
		}
	case TimerSetStm1AddOneMin:
		if TimerSetStm1Eod == Entry {
			TimerSetStm1AddOneMinEntry()
			TimerSetStm1Eod = Do
		}
		if TimerSetStm1Eod == Do {
			TimerSetStm1AddOneMinDo()
			if TimerSetStm1AddOneMinMinButtonReleaseCond() {
				TimerSetStm1AddOneMinMinButtonReleaseAction()
				TimerSetStm1CurrentState = TimerSetStm1Set
				TimerSetStm1Eod = Exit
			}
			if TimerSetStm1AddOneMinMinButtonHold3sCond() {
				TimerSetStm1AddOneMinMinButtonHold3sAction()
				TimerSetStm1CurrentState = TimerSetStm1AddMinHold
				TimerSetStm1Eod = Exit
			}
		}
		if TimerSetStm1Eod == Exit {
			TimerSetStm1AddOneMinExit()
			TimerSetStm1Eod = Entry
		}
	case TimerSetStm1AddOneSec:
		if TimerSetStm1Eod == Entry {
			TimerSetStm1AddOneSecEntry()
			TimerSetStm1Eod = Do
		}
		if TimerSetStm1Eod == Do {
			TimerSetStm1AddOneSecDo()
			if TimerSetStm1AddOneSecSecButtonReleaseCond() {
				TimerSetStm1AddOneSecSecButtonReleaseAction()
				TimerSetStm1CurrentState = TimerSetStm1Set
				TimerSetStm1Eod = Exit
			}
			if TimerSetStm1AddOneSecSecButtonHold3sCond() {
				TimerSetStm1AddOneSecSecButtonHold3sAction()
				TimerSetStm1CurrentState = TimerSetStm1AddSecHold
				TimerSetStm1Eod = Exit
			}
		}
		if TimerSetStm1Eod == Exit {
			TimerSetStm1AddOneSecExit()
			TimerSetStm1Eod = Entry
		}
	case TimerSetStm1TimerWaitRelease:
		if TimerSetStm1Eod == Entry {
			TimerSetStm1TimerWaitReleaseEntry()
			TimerSetStm1Eod = Do
		}
		if TimerSetStm1Eod == Do {
			TimerSetStm1TimerWaitReleaseDo()
		}
		if TimerSetStm1Eod == Exit {
			TimerSetStm1TimerWaitReleaseExit()
			TimerSetStm1Eod = Entry
		}
	case TimerSetStm1Set:
		if TimerSetStm1Eod == Entry {
			TimerSetStm1SetEntry()
			TimerSetStm1Eod = Do
		}
		if TimerSetStm1Eod == Do {
			TimerSetStm1SetDo()
			if TimerSetStm1SetstartButtonPushCond() {
				TimerSetStm1SetstartButtonPushAction()
				TimerSetStm1CurrentState = TimerSetStm1TimerWaitRelease
				TimerSetStm1Eod = Exit
			}
			if TimerSetStm1SetSecButtonPushCond() {
				TimerSetStm1SetSecButtonPushAction()
				TimerSetStm1CurrentState = TimerSetStm1AddOneSec
				TimerSetStm1Eod = Exit
			}
			if TimerSetStm1SetMinButtonPushCond() {
				TimerSetStm1SetMinButtonPushAction()
				TimerSetStm1CurrentState = TimerSetStm1AddOneMin
				TimerSetStm1Eod = Exit
			}
		}
		if TimerSetStm1Eod == Exit {
			TimerSetStm1SetExit()
			TimerSetStm1Eod = Entry
		}
	case TimerSetStm1AddMinHold:
		if TimerSetStm1Eod == Entry {
			TimerSetStm1AddMinHoldEntry()
			TimerSetStm1Eod = Do
		}
		if TimerSetStm1Eod == Do {
			TimerSetStm1AddMinHoldDo()
			if TimerSetStm1AddMinHoldMinButtonReleaseCond() {
				TimerSetStm1AddMinHoldMinButtonReleaseAction()
				TimerSetStm1CurrentState = TimerSetStm1Set
				TimerSetStm1Eod = Exit
			}
			if TimerSetStm1AddMinHoldPass250msCond() {
				TimerSetStm1AddMinHoldPass250msAction()
				TimerSetStm1CurrentState = TimerSetStm1AddMinHold
				TimerSetStm1Eod = Exit
			}
		}
		if TimerSetStm1Eod == Exit {
			TimerSetStm1AddMinHoldExit()
			TimerSetStm1Eod = Entry
		}
	case TimerSetStm1AddSecHold:
		if TimerSetStm1Eod == Entry {
			TimerSetStm1AddSecHoldEntry()
			TimerSetStm1Eod = Do
		}
		if TimerSetStm1Eod == Do {
			TimerSetStm1AddSecHoldDo()
			if TimerSetStm1AddSecHoldSecButtonReleaseCond() {
				TimerSetStm1AddSecHoldSecButtonReleaseAction()
				TimerSetStm1CurrentState = TimerSetStm1Set
				TimerSetStm1Eod = Exit
			}
			if TimerSetStm1AddSecHoldPass250msCond() {
				TimerSetStm1AddSecHoldPass250msAction()
				TimerSetStm1CurrentState = TimerSetStm1AddSecHold
				TimerSetStm1Eod = Exit
			}
		}
		if TimerSetStm1Eod == Exit {
			TimerSetStm1AddSecHoldExit()
			TimerSetStm1Eod = Entry
		}
	}
}