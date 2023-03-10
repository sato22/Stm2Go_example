// This file was generated by a program.
// Please do not edit this file directly.
package mypackage

import (
// package names to be imported
)

type CountdownStm0State uint8

const (
	CountdownStm0WaitOneSec CountdownStm0State = iota
	CountdownStm0CountdownWaitRelease
)

var CountdownStm0Eod Eod

var CountdownStm0CurrentState CountdownStm0State

func init() {
	CountdownStm0Initialize()
}

func CountdownStm0Initialize() {
	CountdownStm0CurrentState = CountdownStm0WaitOneSec
}

func CountdownStm0Task() {
	switch CountdownStm0CurrentState {
	case CountdownStm0WaitOneSec:
		if CountdownStm0Eod == Entry {
			CountdownStm0WaitOneSecEntry()
			CountdownStm0Eod = Do
		}
		if CountdownStm0Eod == Do {
			CountdownStm0WaitOneSecDo()
			if CountdownStm0WaitOneSeconeSecPassCond() {
				CountdownStm0WaitOneSeconeSecPassAction()
				CountdownStm0CurrentState = CountdownStm0WaitOneSec
				CountdownStm0Eod = Exit
			}
			if CountdownStm0WaitOneSecstopButtonPushCond() {
				CountdownStm0WaitOneSecstopButtonPushAction()
				CountdownStm0CurrentState = CountdownStm0CountdownWaitRelease
				CountdownStm0Eod = Exit
			}
		}
		if CountdownStm0Eod == Exit {
			CountdownStm0WaitOneSecExit()
			CountdownStm0Eod = Entry
		}
	case CountdownStm0CountdownWaitRelease:
		if CountdownStm0Eod == Entry {
			CountdownStm0CountdownWaitReleaseEntry()
			CountdownStm0Eod = Do
		}
		if CountdownStm0Eod == Do {
			CountdownStm0CountdownWaitReleaseDo()
		}
		if CountdownStm0Eod == Exit {
			CountdownStm0CountdownWaitReleaseExit()
			CountdownStm0Eod = Entry
		}
	}
}
