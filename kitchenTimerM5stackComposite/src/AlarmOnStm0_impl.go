// This file was generated by a program.
package mypackage

// package names to be imported

///////////////////////////////////////////////
// functions for State AlarmOnStm0Beep
///////////////////////////////////////////////

func AlarmOnStm0BeepEntry() {
	if debug {
		logger.Println("Entering State AlarmOnStm0Beep")
	}
	alarm.Beep()
}

func AlarmOnStm0BeepDo() {
	// Please write a do process for State AlarmOnStm0Beep
}

func AlarmOnStm0BeepExit() {
	if debug {
		logger.Println("Leaving State AlarmOnStm0Beep")
	}
	alarm.Mute()
}

///////////////////////////////////////////////
// functions for State AlarmOnStm0AlarmOnWaitRelease
///////////////////////////////////////////////

func AlarmOnStm0AlarmOnWaitReleaseEntry() {
	if debug {
		logger.Println("Entering State AlarmOnStm0AlarmOnWaitRelease")
	}
	// Please write an enter process for State AlarmOnStm0AlarmOnWaitRelease
}

func AlarmOnStm0AlarmOnWaitReleaseDo() {
	// Please write a do process for State AlarmOnStm0AlarmOnWaitRelease
}

func AlarmOnStm0AlarmOnWaitReleaseExit() {
	if debug {
		logger.Println("Leaving State AlarmOnStm0AlarmOnWaitRelease")
	}
	// Please write an exit process for State AlarmOnStm0AlarmOnWaitRelease
}

///////////////////////////////////////////////
// Conditions
///////////////////////////////////////////////

func AlarmOnStm0BeepstopButtonPushCond() bool {
	return !rRelease
}

///////////////////////////////////////////////
// Actions
///////////////////////////////////////////////

func AlarmOnStm0BeepstopButtonPushAction() {
	// Please edit the action when stopButtonPush occurs at State Beep
}
