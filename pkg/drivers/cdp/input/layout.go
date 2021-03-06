package input

var usKeyboardLayout = map[string]KeyboardKey{
	"0": {
		KeyCode: 48,
		Key:     "0",
		Code:    "Digit0",
	},
	"1": {
		KeyCode: 49,
		Key:     "1",
		Code:    "Digit1",
	},
	"2": {
		KeyCode: 50,
		Key:     "2",
		Code:    "Digit2",
	},
	"3": {
		KeyCode: 51,
		Key:     "3",
		Code:    "Digit3",
	},
	"4": {
		KeyCode: 52,
		Key:     "4",
		Code:    "Digit4",
	},
	"5": {
		KeyCode: 53,
		Key:     "5",
		Code:    "Digit5",
	},
	"6": {
		KeyCode: 54,
		Key:     "6",
		Code:    "Digit6",
	},
	"7": {
		KeyCode: 55,
		Key:     "7",
		Code:    "Digit7",
	},
	"8": {
		KeyCode: 56,
		Key:     "8",
		Code:    "Digit8",
	},
	"9": {
		KeyCode: 57,
		Key:     "9",
		Code:    "Digit9",
	},
	"Power": {
		Key:  "Power",
		Code: "Power",
	},
	"Eject": {
		Key:  "Eject",
		Code: "Eject",
	},
	"Abort": {
		KeyCode: 3,
		Code:    "Abort",
		Key:     "Cancel",
	},
	"Help": {
		KeyCode: 6,
		Code:    "Help",
		Key:     "Help",
	},
	"Backspace": {
		KeyCode: 8,
		Code:    "Backspace",
		Key:     "Backspace",
	},
	"Tab": {
		KeyCode: 9,
		Code:    "Tab",
		Key:     "Tab",
	},
	"Numpad5": {
		KeyCode:  12,
		Key:      "Clear",
		Code:     "Numpad5",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"NumpadEnter": {
		KeyCode:  13,
		Code:     "NumpadEnter",
		Key:      "Enter",
		Location: 3,
	},
	"Enter": {
		KeyCode: 13,
		Code:    "Enter",
		Key:     "Enter",
	},
	`\r`: {
		KeyCode: 13,
		Code:    "Enter",
		Key:     "Enter",
	},
	`\n`: {
		KeyCode: 13,
		Code:    "Enter",
		Key:     "Enter",
	},
	"ShiftLeft": {
		KeyCode:  16,
		Code:     "ShiftLeft",
		Key:      "Shift",
		Location: 1,
	},
	"ShiftRight": {
		KeyCode:  16,
		Code:     "ShiftRight",
		Key:      "Shift",
		Location: 2,
	},
	"ControlLeft": {
		KeyCode:  17,
		Code:     "ControlLeft",
		Key:      "Control",
		Location: 1,
	},
	"ControlRight": {
		KeyCode:  17,
		Code:     "ControlRight",
		Key:      "Control",
		Location: 2,
	},
	"AltLeft": {
		KeyCode:  18,
		Code:     "AltLeft",
		Key:      "Alt",
		Location: 1,
	},
	"AltRight": {
		KeyCode:  18,
		Code:     "AltRight",
		Key:      "Alt",
		Location: 2,
	},
	"Pause": {
		KeyCode: 19,
		Code:    "Pause",
		Key:     "Pause",
	},
	"CapsLock": {
		KeyCode: 20,
		Code:    "CapsLock",
		Key:     "CapsLock",
	},
	"Escape": {
		KeyCode: 27,
		Code:    "Escape",
		Key:     "Escape",
	},
	"Convert": {
		KeyCode: 28,
		Code:    "Convert",
		Key:     "Convert",
	},
	"NonConvert": {
		KeyCode: 29,
		Code:    "NonConvert",
		Key:     "NonConvert",
	},
	"Space": {
		KeyCode: 32,
		Code:    "Space",
		Key:     " ",
	},
	"Numpad9": {
		KeyCode:  33,
		Key:      "PageUp",
		Code:     "Numpad9",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"PageUp": {
		KeyCode: 33,
		Code:    "PageUp",
		Key:     "PageUp",
	},
	"Numpad3": {
		KeyCode:  34,
		Key:      "PageDown",
		Code:     "Numpad3",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"PageDown": {
		KeyCode: 34,
		Code:    "PageDown",
		Key:     "PageDown",
	},
	"End": {
		KeyCode: 35,
		Code:    "End",
		Key:     "End",
	},
	"Numpad1": {
		KeyCode:  35,
		Key:      "End",
		Code:     "Numpad1",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"Home": {
		KeyCode: 36,
		Code:    "Home",
		Key:     "Home",
	},
	"Numpad7": {
		KeyCode:  36,
		Key:      "Home",
		Code:     "Numpad7",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"ArrowLeft": {
		KeyCode: 37,
		Code:    "ArrowLeft",
		Key:     "ArrowLeft",
	},
	"Numpad4": {
		KeyCode:  37,
		Key:      "ArrowLeft",
		Code:     "Numpad4",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"Numpad8": {
		KeyCode:  38,
		Key:      "ArrowUp",
		Code:     "Numpad8",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"ArrowUp": {
		KeyCode: 38,
		Code:    "ArrowUp",
		Key:     "ArrowUp",
	},
	"ArrowRight": {
		KeyCode: 39,
		Code:    "ArrowRight",
		Key:     "ArrowRight",
	},
	"Numpad6": {
		KeyCode:  39,
		Key:      "ArrowRight",
		Code:     "Numpad6",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"Numpad2": {
		KeyCode:  40,
		Key:      "ArrowDown",
		Code:     "Numpad2",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"ArrowDown": {
		KeyCode: 40,
		Code:    "ArrowDown",
		Key:     "ArrowDown",
	},
	"Select": {
		KeyCode: 41,
		Code:    "Select",
		Key:     "Select",
	},
	"Open": {
		KeyCode: 43,
		Code:    "Open",
		Key:     "Execute",
	},
	"PrintScreen": {
		KeyCode: 44,
		Code:    "PrintScreen",
		Key:     "PrintScreen",
	},
	"Insert": {
		KeyCode: 45,
		Code:    "Insert",
		Key:     "Insert",
	},
	"Numpad0": {
		KeyCode: 45,
		// "shiftKeyCode": 96,
		Key:      "Insert",
		Code:     "Numpad0",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"Delete": {
		KeyCode: 46,
		Code:    "Delete",
		Key:     "Delete",
	},
	"NumpadDecimal": {
		KeyCode: 46,
		// "shiftKeyCode": 110,
		Code:     "NumpadDecimal",
		Key:      "\u0000",
		Modifier: KeyboardModifierShift,
		Location: 3,
	},
	"Digit0": {
		KeyCode:  48,
		Code:     "Digit0",
		Modifier: KeyboardModifierShift,
		Key:      "0",
	},
	"Digit1": {
		KeyCode:  49,
		Code:     "Digit1",
		Modifier: KeyboardModifierShift,
		Key:      "1",
	},
	"Digit2": {
		KeyCode:  50,
		Code:     "Digit2",
		Modifier: KeyboardModifierShift,
		Key:      "2",
	},
	"Digit3": {
		KeyCode:  51,
		Code:     "Digit3",
		Modifier: KeyboardModifierShift,
		Key:      "3",
	},
	"Digit4": {
		KeyCode:  52,
		Code:     "Digit4",
		Modifier: KeyboardModifierShift,
		Key:      "4",
	},
	"Digit5": {
		KeyCode:  53,
		Code:     "Digit5",
		Modifier: KeyboardModifierShift,
		Key:      "5",
	},
	"Digit6": {
		KeyCode:  54,
		Code:     "Digit6",
		Modifier: KeyboardModifierShift,
		Key:      "6",
	},
	"Digit7": {
		KeyCode:  55,
		Code:     "Digit7",
		Modifier: KeyboardModifierShift,
		Key:      "7",
	},
	"Digit8": {
		KeyCode:  56,
		Code:     "Digit8",
		Modifier: KeyboardModifierShift,
		Key:      "8",
	},
	"Digit9": {
		KeyCode:  57,
		Code:     "Digit9",
		Modifier: KeyboardModifierShift,
		Key:      "9",
	},
	"KeyA": {
		KeyCode:  65,
		Code:     "KeyA",
		Modifier: KeyboardModifierShift,
		Key:      "a",
	},
	"KeyB": {
		KeyCode:  66,
		Code:     "KeyB",
		Modifier: KeyboardModifierShift,
		Key:      "b",
	},
	"KeyC": {
		KeyCode:  67,
		Code:     "KeyC",
		Modifier: KeyboardModifierShift,
		Key:      "c",
	},
	"KeyD": {
		KeyCode:  68,
		Code:     "KeyD",
		Modifier: KeyboardModifierShift,
		Key:      "d",
	},
	"KeyE": {
		KeyCode:  69,
		Code:     "KeyE",
		Modifier: KeyboardModifierShift,
		Key:      "e",
	},
	"KeyF": {
		KeyCode:  70,
		Code:     "KeyF",
		Modifier: KeyboardModifierShift,
		Key:      "f",
	},
	"KeyG": {
		KeyCode:  71,
		Code:     "KeyG",
		Modifier: KeyboardModifierShift,
		Key:      "g",
	},
	"KeyH": {
		KeyCode:  72,
		Code:     "KeyH",
		Modifier: KeyboardModifierShift,
		Key:      "h",
	},
	"KeyI": {
		KeyCode:  73,
		Code:     "KeyI",
		Modifier: KeyboardModifierShift,
		Key:      "i",
	},
	"KeyJ": {
		KeyCode:  74,
		Code:     "KeyJ",
		Modifier: KeyboardModifierShift,
		Key:      "j",
	},
	"KeyK": {
		KeyCode:  75,
		Code:     "KeyK",
		Modifier: KeyboardModifierShift,
		Key:      "k",
	},
	"KeyL": {
		KeyCode:  76,
		Code:     "KeyL",
		Modifier: KeyboardModifierShift,
		Key:      "l",
	},
	"KeyM": {
		KeyCode:  77,
		Code:     "KeyM",
		Modifier: KeyboardModifierShift,
		Key:      "m",
	},
	"KeyN": {
		KeyCode:  78,
		Code:     "KeyN",
		Modifier: KeyboardModifierShift,
		Key:      "n",
	},
	"KeyO": {
		KeyCode:  79,
		Code:     "KeyO",
		Modifier: KeyboardModifierShift,
		Key:      "o",
	},
	"KeyP": {
		KeyCode:  80,
		Code:     "KeyP",
		Modifier: KeyboardModifierShift,
		Key:      "p",
	},
	"KeyQ": {
		KeyCode:  81,
		Code:     "KeyQ",
		Modifier: KeyboardModifierShift,
		Key:      "q",
	},
	"KeyR": {
		KeyCode:  82,
		Code:     "KeyR",
		Modifier: KeyboardModifierShift,
		Key:      "r",
	},
	"KeyS": {
		KeyCode:  83,
		Code:     "KeyS",
		Modifier: KeyboardModifierShift,
		Key:      "s",
	},
	"KeyT": {
		KeyCode:  84,
		Code:     "KeyT",
		Modifier: KeyboardModifierShift,
		Key:      "t",
	},
	"KeyU": {
		KeyCode:  85,
		Code:     "KeyU",
		Modifier: KeyboardModifierShift,
		Key:      "u",
	},
	"KeyV": {
		KeyCode:  86,
		Code:     "KeyV",
		Modifier: KeyboardModifierShift,
		Key:      "v",
	},
	"KeyW": {
		KeyCode:  87,
		Code:     "KeyW",
		Modifier: KeyboardModifierShift,
		Key:      "w",
	},
	"KeyX": {
		KeyCode:  88,
		Code:     "KeyX",
		Modifier: KeyboardModifierShift,
		Key:      "x",
	},
	"KeyY": {
		KeyCode:  89,
		Code:     "KeyY",
		Modifier: KeyboardModifierShift,
		Key:      "y",
	},
	"KeyZ": {
		KeyCode:  90,
		Code:     "KeyZ",
		Modifier: KeyboardModifierShift,
		Key:      "z",
	},
	"MetaLeft": {
		KeyCode:  91,
		Code:     "MetaLeft",
		Key:      "Meta",
		Location: 1,
	},
	"MetaRight": {
		KeyCode:  92,
		Code:     "MetaRight",
		Key:      "Meta",
		Location: 2,
	},
	"ContextMenu": {
		KeyCode: 93,
		Code:    "ContextMenu",
		Key:     "ContextMenu",
	},
	"NumpadMultiply": {
		KeyCode:  106,
		Code:     "NumpadMultiply",
		Key:      "*",
		Location: 3,
	},
	"NumpadAdd": {
		KeyCode:  107,
		Code:     "NumpadAdd",
		Key:      "+",
		Location: 3,
	},
	"NumpadSubtract": {
		KeyCode:  109,
		Code:     "NumpadSubtract",
		Key:      "-",
		Location: 3,
	},
	"NumpadDivide": {
		KeyCode:  111,
		Code:     "NumpadDivide",
		Key:      "/",
		Location: 3,
	},
	"F1": {
		KeyCode: 112,
		Code:    "F1",
		Key:     "F1",
	},
	"F2": {
		KeyCode: 113,
		Code:    "F2",
		Key:     "F2",
	},
	"F3": {
		KeyCode: 114,
		Code:    "F3",
		Key:     "F3",
	},
	"F4": {
		KeyCode: 115,
		Code:    "F4",
		Key:     "F4",
	},
	"F5": {
		KeyCode: 116,
		Code:    "F5",
		Key:     "F5",
	},
	"F6": {
		KeyCode: 117,
		Code:    "F6",
		Key:     "F6",
	},
	"F7": {
		KeyCode: 118,
		Code:    "F7",
		Key:     "F7",
	},
	"F8": {
		KeyCode: 119,
		Code:    "F8",
		Key:     "F8",
	},
	"F9": {
		KeyCode: 120,
		Code:    "F9",
		Key:     "F9",
	},
	"F10": {
		KeyCode: 121,
		Code:    "F10",
		Key:     "F10",
	},
	"F11": {
		KeyCode: 122,
		Code:    "F11",
		Key:     "F11",
	},
	"F12": {
		KeyCode: 123,
		Code:    "F12",
		Key:     "F12",
	},
	"F13": {
		KeyCode: 124,
		Code:    "F13",
		Key:     "F13",
	},
	"F14": {
		KeyCode: 125,
		Code:    "F14",
		Key:     "F14",
	},
	"F15": {
		KeyCode: 126,
		Code:    "F15",
		Key:     "F15",
	},
	"F16": {
		KeyCode: 127,
		Code:    "F16",
		Key:     "F16",
	},
	"F17": {
		KeyCode: 128,
		Code:    "F17",
		Key:     "F17",
	},
	"F18": {
		KeyCode: 129,
		Code:    "F18",
		Key:     "F18",
	},
	"F19": {
		KeyCode: 130,
		Code:    "F19",
		Key:     "F19",
	},
	"F20": {
		KeyCode: 131,
		Code:    "F20",
		Key:     "F20",
	},
	"F21": {
		KeyCode: 132,
		Code:    "F21",
		Key:     "F21",
	},
	"F22": {
		KeyCode: 133,
		Code:    "F22",
		Key:     "F22",
	},
	"F23": {
		KeyCode: 134,
		Code:    "F23",
		Key:     "F23",
	},
	"F24": {
		KeyCode: 135,
		Code:    "F24",
		Key:     "F24",
	},
	"NumLock": {
		KeyCode: 144,
		Code:    "NumLock",
		Key:     "NumLock",
	},
	"ScrollLock": {
		KeyCode: 145,
		Code:    "ScrollLock",
		Key:     "ScrollLock",
	},
	"AudioVolumeMute": {
		KeyCode: 173,
		Code:    "AudioVolumeMute",
		Key:     "AudioVolumeMute",
	},
	"AudioVolumeDown": {
		KeyCode: 174,
		Code:    "AudioVolumeDown",
		Key:     "AudioVolumeDown",
	},
	"AudioVolumeUp": {
		KeyCode: 175,
		Code:    "AudioVolumeUp",
		Key:     "AudioVolumeUp",
	},
	"MediaTrackNext": {
		KeyCode: 176,
		Code:    "MediaTrackNext",
		Key:     "MediaTrackNext",
	},
	"MediaTrackPrevious": {
		KeyCode: 177,
		Code:    "MediaTrackPrevious",
		Key:     "MediaTrackPrevious",
	},
	"MediaStop": {
		KeyCode: 178,
		Code:    "MediaStop",
		Key:     "MediaStop",
	},
	"MediaPlayPause": {
		KeyCode: 179,
		Code:    "MediaPlayPause",
		Key:     "MediaPlayPause",
	},
	"Semicolon": {
		KeyCode: 186,
		Code:    "Semicolon",
		Key:     ";",
	},
	"Equal": {
		KeyCode: 187,
		Code:    "Equal",
		Key:     "=",
	},
	"NumpadEqual": {
		KeyCode:  187,
		Code:     "NumpadEqual",
		Key:      "=",
		Location: 3,
	},
	"Comma": {
		KeyCode: 188,
		Code:    "Comma",
		Key:     ",",
	},
	"Minus": {
		KeyCode: 189,
		Code:    "Minus",
		Key:     "-",
	},
	"Period": {
		KeyCode: 190,
		Code:    "Period",
		Key:     ".",
	},
	"Slash": {
		KeyCode: 191,
		Code:    "Slash",
		Key:     "/",
	},
	"Backquote": {
		KeyCode: 192,
		Code:    "Backquote",
		Key:     "`",
	},
	"BracketLeft": {
		KeyCode: 219,
		Code:    "BracketLeft",
		Key:     "[",
	},
	"Backslash": {
		KeyCode: 220,
		Code:    "Backslash",
		Key:     "\\",
	},
	"BracketRight": {
		KeyCode: 221,
		Code:    "BracketRight",
		Key:     "]",
	},
	"Quote": {
		KeyCode: 222,
		Code:    "Quote",
		Key:     "\"",
	},
	"AltGraph": {
		KeyCode: 225,
		Code:    "AltGraph",
		Key:     "AltGraph",
	},
	"Props": {
		KeyCode: 247,
		Code:    "Props",
		Key:     "CrSel",
	},
	"Cancel": {
		KeyCode: 3,
		Key:     "Cancel",
		Code:    "Abort",
	},
	"Clear": {
		KeyCode:  12,
		Key:      "Clear",
		Code:     "Numpad5",
		Location: 3,
	},
	"Shift": {
		KeyCode:  16,
		Key:      "Shift",
		Code:     "ShiftLeft",
		Location: 1,
	},
	"Control": {
		KeyCode:  17,
		Key:      "Control",
		Code:     "ControlLeft",
		Location: 1,
	},
	"Alt": {
		KeyCode:  18,
		Key:      "Alt",
		Code:     "AltLeft",
		Location: 1,
	},
	"Accept": {
		KeyCode: 30,
		Key:     "Accept",
	},
	"ModeChange": {
		KeyCode: 31,
		Key:     "ModeChange",
	},
	" ": {
		KeyCode: 32,
		Key:     " ",
		Code:    "Space",
	},
	"Print": {
		KeyCode: 42,
		Key:     "Print",
	},
	"Execute": {
		KeyCode: 43,
		Key:     "Execute",
		Code:    "Open",
	},
	"": {
		KeyCode:  46,
		Key:      "\u0000",
		Code:     "NumpadDecimal",
		Location: 3,
	},
	"a": {
		KeyCode: 65,
		Key:     "a",
		Code:    "KeyA",
	},
	"b": {
		KeyCode: 66,
		Key:     "b",
		Code:    "KeyB",
	},
	"c": {
		KeyCode: 67,
		Key:     "c",
		Code:    "KeyC",
	},
	"d": {
		KeyCode: 68,
		Key:     "d",
		Code:    "KeyD",
	},
	"e": {
		KeyCode: 69,
		Key:     "e",
		Code:    "KeyE",
	},
	"f": {
		KeyCode: 70,
		Key:     "f",
		Code:    "KeyF",
	},
	"g": {
		KeyCode: 71,
		Key:     "g",
		Code:    "KeyG",
	},
	"h": {
		KeyCode: 72,
		Key:     "h",
		Code:    "KeyH",
	},
	"i": {
		KeyCode: 73,
		Key:     "i",
		Code:    "KeyI",
	},
	"j": {
		KeyCode: 74,
		Key:     "j",
		Code:    "KeyJ",
	},
	"k": {
		KeyCode: 75,
		Key:     "k",
		Code:    "KeyK",
	},
	"l": {
		KeyCode: 76,
		Key:     "l",
		Code:    "KeyL",
	},
	"m": {
		KeyCode: 77,
		Key:     "m",
		Code:    "KeyM",
	},
	"n": {
		KeyCode: 78,
		Key:     "n",
		Code:    "KeyN",
	},
	"o": {
		KeyCode: 79,
		Key:     "o",
		Code:    "KeyO",
	},
	"p": {
		KeyCode: 80,
		Key:     "p",
		Code:    "KeyP",
	},
	"q": {
		KeyCode: 81,
		Key:     "q",
		Code:    "KeyQ",
	},
	"r": {
		KeyCode: 82,
		Key:     "r",
		Code:    "KeyR",
	},
	"s": {
		KeyCode: 83,
		Key:     "s",
		Code:    "KeyS",
	},
	"t": {
		KeyCode: 84,
		Key:     "t",
		Code:    "KeyT",
	},
	"u": {
		KeyCode: 85,
		Key:     "u",
		Code:    "KeyU",
	},
	"v": {
		KeyCode: 86,
		Key:     "v",
		Code:    "KeyV",
	},
	"w": {
		KeyCode: 87,
		Key:     "w",
		Code:    "KeyW",
	},
	"x": {
		KeyCode: 88,
		Key:     "x",
		Code:    "KeyX",
	},
	"y": {
		KeyCode: 89,
		Key:     "y",
		Code:    "KeyY",
	},
	"z": {
		KeyCode: 90,
		Key:     "z",
		Code:    "KeyZ",
	},
	"Meta": {
		KeyCode:  91,
		Key:      "Meta",
		Code:     "MetaLeft",
		Location: 1,
	},
	"*": {
		KeyCode:  106,
		Key:      "*",
		Code:     "NumpadMultiply",
		Location: 3,
	},
	"+": {
		KeyCode:  107,
		Key:      "+",
		Code:     "NumpadAdd",
		Location: 3,
	},
	"-": {
		KeyCode:  109,
		Key:      "-",
		Code:     "NumpadSubtract",
		Location: 3,
	},
	"/": {
		KeyCode:  111,
		Key:      "/",
		Code:     "NumpadDivide",
		Location: 3,
	},
	";": {
		KeyCode: 186,
		Key:     ";",
		Code:    "Semicolon",
	},
	"=": {
		KeyCode: 187,
		Key:     "=",
		Code:    "Equal",
	},
	",": {
		KeyCode: 188,
		Key:     ",",
		Code:    "Comma",
	},
	".": {
		KeyCode: 190,
		Key:     ".",
		Code:    "Period",
	},
	"`": {
		KeyCode: 192,
		Key:     "`",
		Code:    "Backquote",
	},
	"[": {
		KeyCode: 219,
		Key:     "[",
		Code:    "BracketLeft",
	},
	`\`: {
		KeyCode: 220,
		Key:     "\\",
		Code:    "Backslash",
	},
	"]": {
		KeyCode: 221,
		Key:     "]",
		Code:    "BracketRight",
	},
	`"`: {
		KeyCode: 222,
		Key:     "\"",
		Code:    "Quote",
	},
	"Attn": {
		KeyCode: 246,
		Key:     "Attn",
	},
	"CrSel": {
		KeyCode: 247,
		Key:     "CrSel",
		Code:    "Props",
	},
	"ExSel": {
		KeyCode: 248,
		Key:     "ExSel",
	},
	"EraseEof": {
		KeyCode: 249,
		Key:     "EraseEof",
	},
	"Play": {
		KeyCode: 250,
		Key:     "Play",
	},
	"ZoomOut": {
		KeyCode: 251,
		Key:     "ZoomOut",
	},
	")": {
		KeyCode: 48,
		Key:     ")",
		Code:    "Digit0",
	},
	"!": {
		KeyCode: 49,
		Key:     "!",
		Code:    "Digit1",
	},
	"@": {
		KeyCode: 50,
		Key:     "@",
		Code:    "Digit2",
	},
	"#": {
		KeyCode: 51,
		Key:     "#",
		Code:    "Digit3",
	},
	"$": {
		KeyCode: 52,
		Key:     "$",
		Code:    "Digit4",
	},
	"%": {
		KeyCode: 53,
		Key:     "%",
		Code:    "Digit5",
	},
	"^": {
		KeyCode: 54,
		Key:     "^",
		Code:    "Digit6",
	},
	"&": {
		KeyCode: 55,
		Key:     "&",
		Code:    "Digit7",
	},
	"(": {
		KeyCode: 57,
		Key:     "\\(",
		Code:    "Digit9",
	},
	"A": {
		KeyCode: 65,
		Key:     "A",
		Code:    "KeyA",
	},
	"B": {
		KeyCode: 66,
		Key:     "B",
		Code:    "KeyB",
	},
	"C": {
		KeyCode: 67,
		Key:     "C",
		Code:    "KeyC",
	},
	"D": {
		KeyCode: 68,
		Key:     "D",
		Code:    "KeyD",
	},
	"E": {
		KeyCode: 69,
		Key:     "E",
		Code:    "KeyE",
	},
	"F": {
		KeyCode: 70,
		Key:     "F",
		Code:    "KeyF",
	},
	"G": {
		KeyCode: 71,
		Key:     "G",
		Code:    "KeyG",
	},
	"H": {
		KeyCode: 72,
		Key:     "H",
		Code:    "KeyH",
	},
	"I": {
		KeyCode: 73,
		Key:     "I",
		Code:    "KeyI",
	},
	"J": {
		KeyCode: 74,
		Key:     "J",
		Code:    "KeyJ",
	},
	"K": {
		KeyCode: 75,
		Key:     "K",
		Code:    "KeyK",
	},
	"L": {
		KeyCode: 76,
		Key:     "L",
		Code:    "KeyL",
	},
	"M": {
		KeyCode: 77,
		Key:     "M",
		Code:    "KeyM",
	},
	"N": {
		KeyCode: 78,
		Key:     "N",
		Code:    "KeyN",
	},
	"O": {
		KeyCode: 79,
		Key:     "O",
		Code:    "KeyO",
	},
	"P": {
		KeyCode: 80,
		Key:     "P",
		Code:    "KeyP",
	},
	"Q": {
		KeyCode: 81,
		Key:     "Q",
		Code:    "KeyQ",
	},
	"R": {
		KeyCode: 82,
		Key:     "R",
		Code:    "KeyR",
	},
	"S": {
		KeyCode: 83,
		Key:     "S",
		Code:    "KeyS",
	},
	"T": {
		KeyCode: 84,
		Key:     "T",
		Code:    "KeyT",
	},
	"U": {
		KeyCode: 85,
		Key:     "U",
		Code:    "KeyU",
	},
	"V": {
		KeyCode: 86,
		Key:     "V",
		Code:    "KeyV",
	},
	"W": {
		KeyCode: 87,
		Key:     "W",
		Code:    "KeyW",
	},
	"X": {
		KeyCode: 88,
		Key:     "X",
		Code:    "KeyX",
	},
	"Y": {
		KeyCode: 89,
		Key:     "Y",
		Code:    "KeyY",
	},
	"Z": {
		KeyCode: 90,
		Key:     "Z",
		Code:    "KeyZ",
	},
	":": {
		KeyCode: 186,
		Key:     ":",
		Code:    "Semicolon",
	},
	"<": {
		KeyCode: 188,
		Key:     "\\<",
		Code:    "Comma",
	},
	"_": {
		KeyCode: 189,
		Key:     "_",
		Code:    "Minus",
	},
	">": {
		KeyCode: 190,
		Key:     ">",
		Code:    "Period",
	},
	"?": {
		KeyCode: 191,
		Key:     "?",
		Code:    "Slash",
	},
	"~": {
		KeyCode: 192,
		Key:     "~",
		Code:    "Backquote",
	},
	"{": {
		KeyCode: 219,
		Key:     "{",
		Code:    "BracketLeft",
	},
	"|": {
		KeyCode: 220,
		Key:     "|",
		Code:    "Backslash",
	},
	"}": {
		KeyCode: 221,
		Key:     "}",
		Code:    "BracketRight",
	},
	"SoftLeft": {
		Key:      "SoftLeft",
		Code:     "SoftLeft",
		Location: 4,
	},
	"SoftRight": {
		Key:      "SoftRight",
		Code:     "SoftRight",
		Location: 4,
	},
	"Camera": {
		KeyCode:  44,
		Key:      "Camera",
		Code:     "Camera",
		Location: 4,
	},
	"Call": {
		Key:      "Call",
		Code:     "Call",
		Location: 4,
	},
	"EndCall": {
		KeyCode:  95,
		Key:      "EndCall",
		Code:     "EndCall",
		Location: 4,
	},
	"VolumeDown": {
		KeyCode:  182,
		Key:      "VolumeDown",
		Code:     "VolumeDown",
		Location: 4,
	},
	"VolumeUp": {
		KeyCode:  183,
		Key:      "VolumeUp",
		Code:     "VolumeUp",
		Location: 4,
	},
}
