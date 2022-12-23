// This file was generated by a program.
// Please do not edit this file directly.
package mypackage

// package names to be imported

type rootStm0State uint8

const (
	rootStm0AlarmOn rootStm0State = iota
	rootStm0Countdown
	rootStm0TimerSet
	rootStm0Idle
)

var rootStm0Eod Eod

var rootStm0CurrentState rootStm0State

func init() {
	rootStm0Initialize()
}

func rootStm0Initialize() {
	rootStm0CurrentState = rootStm0Idle
}

func EntryrootStm0Task() {
	rootStm0Task()
}

func rootStm0Task() {
	switch rootStm0CurrentState {
	case rootStm0AlarmOn:
		if rootStm0Eod == Entry {
			rootStm0AlarmOnEntry()
			rootStm0Eod = Do
		}
		if rootStm0Eod == Do {
			rootStm0AlarmOnDo()
			if rootStm0AlarmOnAlarmOffCond() {
				rootStm0AlarmOnAlarmOffAction()
				rootStm0CurrentState = rootStm0Idle
				rootStm0Eod = Exit
			}
		}
		if rootStm0Eod == Exit {
			rootStm0AlarmOnExit()
			rootStm0Eod = Entry
		}
	case rootStm0Countdown:
		if rootStm0Eod == Entry {
			rootStm0CountdownEntry()
			rootStm0Eod = Do
		}
		if rootStm0Eod == Do {
			rootStm0CountdownDo()
			if rootStm0CountdownEndCond() {
				rootStm0CountdownEndAction()
				rootStm0CurrentState = rootStm0AlarmOn
				rootStm0Eod = Exit
			}
			if rootStm0CountdownStopCond() {
				rootStm0CountdownStopAction()
				rootStm0CurrentState = rootStm0TimerSet
				rootStm0Eod = Exit
			}
		}
		if rootStm0Eod == Exit {
			rootStm0CountdownExit()
			rootStm0Eod = Entry
		}
	case rootStm0TimerSet:
		if rootStm0Eod == Entry {
			rootStm0TimerSetEntry()
			rootStm0Eod = Do
		}
		if rootStm0Eod == Do {
			rootStm0TimerSetDo()
			if rootStm0TimerSetStartCond() {
				rootStm0TimerSetStartAction()
				rootStm0CurrentState = rootStm0Countdown
				rootStm0Eod = Exit
			}
		}
		if rootStm0Eod == Exit {
			rootStm0TimerSetExit()
			rootStm0Eod = Entry
		}
	case rootStm0Idle:
		if rootStm0Eod == Entry {
			rootStm0IdleEntry()
			rootStm0Eod = Do
		}
		if rootStm0Eod == Do {
			rootStm0IdleDo()
			if rootStm0IdleSetTimeCond() {
				rootStm0IdleSetTimeAction()
				rootStm0CurrentState = rootStm0TimerSet
				rootStm0Eod = Exit
			}
		}
		if rootStm0Eod == Exit {
			rootStm0IdleExit()
			rootStm0Eod = Entry
		}
	}
}
