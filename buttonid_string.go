// Code generated by "stringer -type=ButtonID"; DO NOT EDIT.

package joystick

import "strconv"

const _ButtonID_name = "A_ButtonB_ButtonX_ButtonY_ButtonLeftBumperRightBumperBackButtonStartButtonLeftStickRightStickDPadUpDPadDownDPadLeftDPadRight"

var _ButtonID_index = [...]uint8{0, 8, 16, 24, 32, 42, 53, 63, 74, 83, 93, 99, 107, 115, 124}

func (i ButtonID) String() string {
	if i >= ButtonID(len(_ButtonID_index)-1) {
		return "ButtonID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ButtonID_name[_ButtonID_index[i]:_ButtonID_index[i+1]]
}
