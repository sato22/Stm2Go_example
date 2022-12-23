// This file was generated by a program.
package mypackage

// package names to be imported

///////////////////////////////////////////////
// functions for State AlarmOnStm3Beep
///////////////////////////////////////////////

func AlarmOnStm3BeepEntry() {
	if debug {
		logger.Println("Entering State AlarmOnStm3Beep")
	}
	alarm.Beep()
}

func AlarmOnStm3BeepDo() {
	// Please write a do process for State AlarmOnStm3Beep
}

func AlarmOnStm3BeepExit() {
	if debug {
		logger.Println("Leaving State AlarmOnStm3Beep")
	}
	alarm.Mute()
}

///////////////////////////////////////////////
// functions for State AlarmOnStm3AlarmOnWaitRelease
///////////////////////////////////////////////

func AlarmOnStm3AlarmOnWaitReleaseEntry() {
	if debug {
		logger.Println("Entering State AlarmOnStm3AlarmOnWaitRelease")
	}
	// Please write an enter process for State AlarmOnStm3AlarmOnWaitRelease
}

func AlarmOnStm3AlarmOnWaitReleaseDo() {
	// Please write a do process for State AlarmOnStm3AlarmOnWaitRelease
}

func AlarmOnStm3AlarmOnWaitReleaseExit() {
	if debug {
		logger.Println("Leaving State AlarmOnStm3AlarmOnWaitRelease")
	}
	// Please write an exit process for State AlarmOnStm3AlarmOnWaitRelease
}

///////////////////////////////////////////////
// Conditions
///////////////////////////////////////////////

func AlarmOnStm3BeepstopButtonPushCond() bool {
	return !rRelease
}

///////////////////////////////////////////////
// Actions
///////////////////////////////////////////////

func AlarmOnStm3BeepstopButtonPushAction() {
	// Please edit the action when stopButtonPush occurs at State Beep
}
