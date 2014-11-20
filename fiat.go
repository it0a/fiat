package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var CodeMap = map[uint32]string{
	1:   "ESC",
	2:   "1",
	3:   "2",
	4:   "3",
	5:   "4",
	6:   "5",
	7:   "6",
	8:   "7",
	9:   "8",
	10:  "9",
	11:  "0",
	12:  "MINUS",
	13:  "EQUAL",
	14:  "BACKSPACE",
	15:  "TAB",
	16:  "Q",
	17:  "W",
	18:  "E",
	19:  "R",
	20:  "T",
	21:  "Y",
	22:  "U",
	23:  "I",
	24:  "O",
	25:  "P",
	26:  "LEFTBRACE",
	27:  "RIGHTBRACE",
	28:  "ENTER",
	29:  "LEFTCTRL",
	30:  "A",
	31:  "S",
	32:  "D",
	33:  "F",
	34:  "G",
	35:  "H",
	36:  "J",
	37:  "K",
	38:  "L",
	39:  "SEMICOLON",
	40:  "APOSTROPHE",
	41:  "GRAVE",
	42:  "LEFTSHIFT",
	43:  "BACKSLASH",
	44:  "Z",
	45:  "X",
	46:  "C",
	47:  "V",
	48:  "B",
	49:  "N",
	50:  "M",
	51:  "COMMA",
	52:  "DOT",
	53:  "SLASH",
	54:  "RIGHTSHIFT",
	55:  "KPASTERISK",
	56:  "LEFTALT",
	57:  "SPACE",
	58:  "CAPSLOCK",
	59:  "F1",
	60:  "F2",
	61:  "F3",
	62:  "F4",
	63:  "F5",
	64:  "F6",
	65:  "F7",
	66:  "F8",
	67:  "F9",
	68:  "F10",
	69:  "NUMLOCK",
	70:  "SCROLLLOCK",
	71:  "KP7",
	72:  "KP8",
	73:  "KP9",
	74:  "KPMINUS",
	75:  "KP4",
	76:  "KP5",
	77:  "KP6",
	78:  "KPPLUS",
	79:  "KP1",
	80:  "KP2",
	81:  "KP3",
	82:  "KP0",
	83:  "KPDOT",
	85:  "ZENKAKUHANKAKU",
	86:  "102ND",
	87:  "F11",
	88:  "F12",
	89:  "RO",
	90:  "KATAKANA",
	91:  "HIRAGANA",
	92:  "HENKAN",
	93:  "KATAKANAHIRAGANA",
	94:  "MUHENKAN",
	95:  "KPJPCOMMA",
	96:  "KPENTER",
	97:  "RIGHTCTRL",
	98:  "KPSLASH",
	99:  "SYSRQ",
	100: "RIGHTALT",
	101: "LINEFEED",
	102: "HOME",
	103: "UP",
	104: "PAGEUP",
	105: "LEFT",
	106: "RIGHT",
	107: "END",
	108: "DOWN",
	109: "PAGEDOWN",
	110: "INSERT",
	111: "DELETE",
	112: "MACRO",
	113: "MUTE",
	114: "VOLUMEDOWN",
	115: "VOLUMEUP",
	116: "POWER",
	117: "KPEQUAL",
	118: "KPPLUSMINUS",
	119: "PAUSE",
	121: "KPCOMMA",
	122: "HANGUEL",
	123: "HANJA",
	124: "YEN",
	125: "LEFTMETA",
	126: "RIGHTMETA",
	127: "COMPOSE",
	128: "STOP",
	129: "AGAIN",
	130: "PROPS",
	131: "UNDO",
	132: "FRONT",
	133: "COPY",
	134: "OPEN",
	135: "PASTE",
	136: "FIND",
	137: "CUT",
	138: "HELP",
	139: "MENU",
	140: "CALC",
	141: "SETUP",
	142: "SLEEP",
	143: "WAKEUP",
	144: "FILE",
	145: "SENDFILE",
	146: "DELETEFILE",
	147: "XFER",
	148: "PROG1",
	149: "PROG2",
	150: "WWW",
	151: "MSDOS",
	152: "COFFEE",
	153: "DIRECTION",
	154: "CYCLEWINDOWS",
	155: "MAIL",
	156: "BOOKMARKS",
	157: "COMPUTER",
	158: "BACK",
	159: "FORWARD",
	160: "CLOSECD",
	161: "EJECTCD",
	162: "EJECTCLOSECD",
	163: "NEXTSONG",
	164: "PLAYPAUSE",
	165: "PREVIOUSSONG",
	166: "STOPCD",
	167: "RECORD",
	168: "REWIND",
	169: "PHONE",
	170: "ISO",
	171: "CONFIG",
	172: "HOMEPAGE",
	173: "REFRESH",
	174: "EXIT",
	175: "MOVE",
	176: "EDIT",
	177: "SCROLLUP",
	178: "SCROLLDOWN",
	179: "KPLEFTPAREN",
	180: "KPRIGHTPAREN",
	183: "F13",
	184: "F14",
	185: "F15",
	186: "F16",
	187: "F17",
	188: "F18",
	189: "F19",
	190: "F20",
	191: "F21",
	192: "F22",
	193: "F23",
	194: "F24",
	200: "PLAYCD",
	201: "PAUSECD",
	202: "PROG3",
	203: "PROG4",
	205: "SUSPEND",
	206: "CLOSE",
	207: "PLAY",
	208: "FASTFORWARD",
	209: "BASSBOOST",
	210: "PRINT",
	211: "HP",
	212: "CAMERA",
	213: "SOUND",
	214: "QUESTION",
	215: "EMAIL",
	216: "CHAT",
	217: "SEARCH",
	218: "CONNECT",
	219: "FINANCE",
	220: "SPORT",
	221: "SHOP",
	222: "ALTERASE",
	223: "CANCEL",
	224: "BRIGHTNESSDOWN",
	225: "BRIGHTNESSUP",
	226: "MEDIA",
	240: "UNKNOWN",
}

var ValueMap = map[int32]string{
	0: "RELEASED",
	1: "PRESSED",
	2: "HOLDING",
}

type Input_Event struct {
	Code  uint32
	Value uint32
}

func main() {
	app := cli.NewApp()
	app.Name = "fiat"
	app.Usage = "free input analysis tool"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "capture",
			ShortName: "c",
			Usage:     "capture input",
			Action: func(c *cli.Context) {
				fmt.Println("TODO")
			},
		},
	}
	app.Action = func(c *cli.Context) {
		testHook()
	}
	app.Run(os.Args)
}

func DecodeEvent(buf []byte, event *Input_Event) {
	event.Code = uint32(buf[18]) + uint32(buf[19])<<8
	event.Value = uint32(buf[20]) + uint32(buf[21])<<8
}

func ReadEvent(f *os.File) Input_Event {
	var event Input_Event
	buf := make([]byte, 24)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(n, err)
	}
	DecodeEvent(buf, &event)
	return event
}

func testHook() {
	kbd, err := os.Open("/dev/input/event0")
	if err != nil {
		panic(err)
	}
	defer kbd.Close()
	for {
		event := ReadEvent(kbd)
		if event.Code >= 1 && event.Code <= 240 && event.Value == 1 {
			fmt.Print(CodeMap[event.Code])
		}
	}
}
