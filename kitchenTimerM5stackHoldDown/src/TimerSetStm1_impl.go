// This file was generated by a program.
package mypackage

///////////////////////////////////////////////
// functions for State TimerSetStm1Idle
///////////////////////////////////////////////

func TimerSetStm1IdleEntry() {
	if debug {
		logger.Println("Entering State TimerSetStm1Idle")
	}
	// Please write an enter process for State TimerSetStm1Idle
}

func TimerSetStm1IdleDo() {
	// Please write a do process for State TimerSetStm1Idle
}

func TimerSetStm1IdleExit() {
	if debug {
		logger.Println("Leaving State TimerSetStm1Idle")
	}
	// Please write an exit process for State TimerSetStm1Idle
}

///////////////////////////////////////////////
// functions for State TimerSetStm1AddOneMin
///////////////////////////////////////////////

func TimerSetStm1AddOneMinEntry() {
	if debug {
		logger.Println("Entering State TimerSetStm1AddOneMin")
	}
	addMinute()
	countHold = 0
	// Please write an enter process for State TimerSetStm1AddOneMin
}

func TimerSetStm1AddOneMinDo() {
	countHold++
}

func TimerSetStm1AddOneMinExit() {
	if debug {
		logger.Println("Leaving State TimerSetStm1AddOneMin")
	}
	// Please write an exit process for State TimerSetStm1AddOneMin
}

///////////////////////////////////////////////
// functions for State TimerSetStm1AddOneSec
///////////////////////////////////////////////

func TimerSetStm1AddOneSecEntry() {
	if debug {
		logger.Println("Entering State TimerSetStm1AddOneSec")
	}
	addSecond()
	countHold = 0
}

func TimerSetStm1AddOneSecDo() {
	countHold++
}

func TimerSetStm1AddOneSecExit() {
	if debug {
		logger.Println("Leaving State TimerSetStm1AddOneSec")
	}
	// Please write an exit process for State TimerSetStm1AddOneSec
}

///////////////////////////////////////////////
// functions for State TimerSetStm1TimerWaitRelease
///////////////////////////////////////////////

func TimerSetStm1TimerWaitReleaseEntry() {
	if debug {
		logger.Println("Entering State TimerSetStm1TimerWaitRelease")
	}
	// Please write an enter process for State TimerSetStm1TimerWaitRelease
}

func TimerSetStm1TimerWaitReleaseDo() {
	// Please write a do process for State TimerSetStm1TimerWaitRelease
}

func TimerSetStm1TimerWaitReleaseExit() {
	if debug {
		logger.Println("Leaving State TimerSetStm1TimerWaitRelease")
	}
	// Please write an exit process for State TimerSetStm1TimerWaitRelease
}

///////////////////////////////////////////////
// functions for State TimerSetStm1Set
///////////////////////////////////////////////

func TimerSetStm1SetEntry() {
	if debug {
		logger.Println("Entering State TimerSetStm1Set")
	}
	// Please write an enter process for State TimerSetStm1Set
}

func TimerSetStm1SetDo() {
	// Please write a do process for State TimerSetStm1Set
}

func TimerSetStm1SetExit() {
	if debug {
		logger.Println("Leaving State TimerSetStm1Set")
	}
	// Please write an exit process for State TimerSetStm1Set
}

///////////////////////////////////////////////
// functions for State TimerSetStm1AddMinHold
///////////////////////////////////////////////

func TimerSetStm1AddMinHoldEntry() {
	if debug {
		logger.Println("Entering State TimerSetStm1AddMinHold")
	}
	count = 0
}

func TimerSetStm1AddMinHoldDo() {
	count++
}

func TimerSetStm1AddMinHoldExit() {
	if debug {
		logger.Println("Leaving State TimerSetStm1AddMinHold")
	}
	// Please write an exit process for State TimerSetStm1AddMinHold
}

///////////////////////////////////////////////
// functions for State TimerSetStm1AddSecHold
///////////////////////////////////////////////

func TimerSetStm1AddSecHoldEntry() {
	if debug {
		logger.Println("Entering State TimerSetStm1AddSecHold")
	}
	count = 0
}

func TimerSetStm1AddSecHoldDo() {
	count++
}

func TimerSetStm1AddSecHoldExit() {
	if debug {
		logger.Println("Leaving State TimerSetStm1AddSecHold")
	}
	// Please write an exit process for State TimerSetStm1AddSecHold
}

///////////////////////////////////////////////
// Conditions
///////////////////////////////////////////////

func TimerSetStm1IdleSecButtonPushCond() bool {
	return !mRelease
}

func TimerSetStm1IdleMinButtonPushCond() bool {
	return !lRelease
}

func TimerSetStm1AddOneMinMinButtonReleaseCond() bool {
	return lRelease
}

func TimerSetStm1AddOneMinMinButtonHold3sCond() bool {
	return countHold == countSecond*3
}

func TimerSetStm1AddOneSecSecButtonReleaseCond() bool {
	return mRelease
}

func TimerSetStm1AddOneSecSecButtonHold3sCond() bool {
	return countHold == countSecond*3
}

func TimerSetStm1SetstartButtonPushCond() bool {
	return !rRelease
}

func TimerSetStm1SetSecButtonPushCond() bool {
	return !mRelease
}

func TimerSetStm1SetMinButtonPushCond() bool {
	return !lRelease
}

func TimerSetStm1AddMinHoldMinButtonReleaseCond() bool {
	return lRelease
}

func TimerSetStm1AddMinHoldPass250msCond() bool {
	return count == countSecond/4
}

func TimerSetStm1AddSecHoldSecButtonReleaseCond() bool {
	return mRelease
}

func TimerSetStm1AddSecHoldPass250msCond() bool {
	return count == countSecond/4
}

///////////////////////////////////////////////
// Actions
///////////////////////////////////////////////

func TimerSetStm1IdleSecButtonPushAction() {
	// Please edit the action when SecButtonPush occurs at State Idle
}

func TimerSetStm1IdleMinButtonPushAction() {
	// Please edit the action when MinButtonPush occurs at State Idle
}

func TimerSetStm1AddOneMinMinButtonReleaseAction() {
	// Please edit the action when MinButtonRelease occurs at State AddOneMin
}

func TimerSetStm1AddOneMinMinButtonHold3sAction() {
	// Please edit the action when MinButtonHold3s occurs at State AddOneMin
}

func TimerSetStm1AddOneSecSecButtonReleaseAction() {
	// Please edit the action when SecButtonRelease occurs at State AddOneSec
}

func TimerSetStm1AddOneSecSecButtonHold3sAction() {
	// Please edit the action when SecButtonHold3s occurs at State AddOneSec
}

func TimerSetStm1SetstartButtonPushAction() {
	// Please edit the action when startButtonPush occurs at State Set
}

func TimerSetStm1SetSecButtonPushAction() {
	// Please edit the action when SecButtonPush occurs at State Set
}

func TimerSetStm1SetMinButtonPushAction() {
	// Please edit the action when MinButtonPush occurs at State Set
}

func TimerSetStm1AddMinHoldMinButtonReleaseAction() {
	// Please edit the action when MinButtonRelease occurs at State AddMinHold
}

func TimerSetStm1AddMinHoldPass250msAction() {
	addMinute()
}

func TimerSetStm1AddSecHoldSecButtonReleaseAction() {
	// Please edit the action when SecButtonRelease occurs at State AddSecHold
}

func TimerSetStm1AddSecHoldPass250msAction() {
	addSecond()
}
