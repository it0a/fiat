package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"os/signal"
)

type Event struct {
	Code  byte
	Value byte
}

var BufWriter bufio.Writer

var FreqMap = map[byte]uint32{
	1:   0,
	2:   0,
	3:   0,
	4:   0,
	5:   0,
	6:   0,
	7:   0,
	8:   0,
	9:   0,
	10:  0,
	11:  0,
	12:  0,
	13:  0,
	14:  0,
	15:  0,
	16:  0,
	17:  0,
	18:  0,
	19:  0,
	20:  0,
	21:  0,
	22:  0,
	23:  0,
	24:  0,
	25:  0,
	26:  0,
	27:  0,
	28:  0,
	29:  0,
	30:  0,
	31:  0,
	32:  0,
	33:  0,
	34:  0,
	35:  0,
	36:  0,
	37:  0,
	38:  0,
	39:  0,
	40:  0,
	41:  0,
	42:  0,
	43:  0,
	44:  0,
	45:  0,
	46:  0,
	47:  0,
	48:  0,
	49:  0,
	50:  0,
	51:  0,
	52:  0,
	53:  0,
	54:  0,
	55:  0,
	56:  0,
	57:  0,
	58:  0,
	59:  0,
	60:  0,
	61:  0,
	62:  0,
	63:  0,
	64:  0,
	65:  0,
	66:  0,
	67:  0,
	68:  0,
	69:  0,
	70:  0,
	71:  0,
	72:  0,
	73:  0,
	74:  0,
	75:  0,
	76:  0,
	77:  0,
	78:  0,
	79:  0,
	80:  0,
	81:  0,
	82:  0,
	83:  0,
	85:  0,
	86:  0,
	87:  0,
	88:  0,
	89:  0,
	90:  0,
	91:  0,
	92:  0,
	93:  0,
	94:  0,
	95:  0,
	96:  0,
	97:  0,
	98:  0,
	99:  0,
	100: 0,
	101: 0,
	102: 0,
	103: 0,
	104: 0,
	105: 0,
	106: 0,
	107: 0,
	108: 0,
	109: 0,
	110: 0,
	111: 0,
	112: 0,
	113: 0,
	114: 0,
	115: 0,
	116: 0,
	117: 0,
	118: 0,
	119: 0,
	121: 0,
	122: 0,
	123: 0,
	124: 0,
	125: 0,
	126: 0,
	127: 0,
	128: 0,
	129: 0,
	130: 0,
	131: 0,
	132: 0,
	133: 0,
	134: 0,
	135: 0,
	136: 0,
	137: 0,
	138: 0,
	139: 0,
	140: 0,
	141: 0,
	142: 0,
	143: 0,
	144: 0,
	145: 0,
	146: 0,
	147: 0,
	148: 0,
	149: 0,
	150: 0,
	151: 0,
	152: 0,
	153: 0,
	154: 0,
	155: 0,
	156: 0,
	157: 0,
	158: 0,
	159: 0,
	160: 0,
	161: 0,
	162: 0,
	163: 0,
	164: 0,
	165: 0,
	166: 0,
	167: 0,
	168: 0,
	169: 0,
	170: 0,
	171: 0,
	172: 0,
	173: 0,
	174: 0,
	175: 0,
	176: 0,
	177: 0,
	178: 0,
	179: 0,
	180: 0,
	183: 0,
	184: 0,
	185: 0,
	186: 0,
	187: 0,
	188: 0,
	189: 0,
	190: 0,
	191: 0,
	192: 0,
	193: 0,
	194: 0,
	200: 0,
	201: 0,
	202: 0,
	203: 0,
	205: 0,
	206: 0,
	207: 0,
	208: 0,
	209: 0,
	210: 0,
	211: 0,
	212: 0,
	213: 0,
	214: 0,
	215: 0,
	216: 0,
	217: 0,
	218: 0,
	219: 0,
	220: 0,
	221: 0,
	222: 0,
	223: 0,
	224: 0,
	225: 0,
	226: 0,
	240: 0,
}

var CodeMap = map[byte]string{
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

func DecodeEvent(buf []byte, ev *Event) {
	ev.Code = buf[18]
	ev.Value = buf[20]
}

func NewEventBuffer() []byte {
	return make([]byte, 24)
}

func ReadEvent(f *os.File) Event {
	var ev Event
	buf := NewEventBuffer()
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(n, err)
	}
	DecodeEvent(buf, &ev)
	return ev
}

func blitFrequencies(freqMap map[byte]uint32) {
	for key := range freqMap {
		if freqMap[key] > 0 {
			fmt.Printf("%s => %v\n", CodeMap[key], freqMap[key])
		}
	}
}

func listen() {
	kbd, err := os.Open("/dev/input/event0")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer kbd.Close()
	tmp, err := os.Create("/tmp/fiat")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err := tmp.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	BufWriter := bufio.NewWriterSize(tmp, 32)
	defer func() {
		if err := BufWriter.Flush(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	attachInterrupt()
	for {
		ev := ReadEvent(kbd)
		if ev.Code >= 1 && ev.Code <= 240 && ev.Value == 1 {
			FreqMap[ev.Code] = FreqMap[ev.Code] + 1
			blitFrequencies(FreqMap)
			BufWriter.WriteByte(ev.Code)
		}
	}
}

func attachInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Captured %v, stopping data capture", sig)
			if err := BufWriter.Flush(); err != nil {
				log.Fatal("FAILED TO FLUSH THE BUFFER! YOUR DATA MAY BE LOST. ", err.Error())
			} else {
				log.Println("FLUSH OK!")
			}
			os.Exit(1)
		}
	}()
}

func main() {
	app := cli.NewApp()
	app.Name = "fiat"
	app.Usage = "free input analysis tool"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		listen()
	}
	app.Run(os.Args)
}
