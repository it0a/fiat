package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"os/signal"
	"sort"
)

type Event struct {
	Code  byte
	Value byte
}

var BufWriter bufio.Writer

var FreqStatMap = map[int]uint32{
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
}

var FreqMap = map[byte]byte{
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
	87:  "F11",
	88:  "F12",
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
	117: "KPEQUAL",
	118: "KPPLUSMINUS",
	119: "PAUSE",
	121: "KPCOMMA",
	125: "LEFTMETA",
	126: "RIGHTMETA",
}

var ValueMap = map[byte]string{
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

func flushFrequencyMap(freqMap map[byte]byte) []byte {
	buf := []byte{}
	for key := range freqMap {
		if freqMap[key] > 0 {
			fmt.Printf("Writing key %s %v times\n", CodeMap[key], freqMap[key])
			buf = append(buf, key)
			buf = append(buf, freqMap[key])
			freqMap[key] = 0
		}
	}
	return buf
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
	BufWriter := bufio.NewWriterSize(tmp, 64)
	defer func() {
		if err := BufWriter.Flush(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	attachInterrupt()
	for {
		ev := ReadEvent(kbd)
		if ev.Code >= 1 && ev.Code <= 126 && ev.Value == 1 {
			FreqMap[ev.Code] = FreqMap[ev.Code] + 1
			if FreqMap[ev.Code] == 8 {
				BufWriter.Write(flushFrequencyMap(FreqMap))
			}
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

type sortedFreqStatMap struct {
	m map[int]uint32
	k []int
}

func (sm *sortedFreqStatMap) Len() int {
	return len(sm.m)
}

func (sm *sortedFreqStatMap) Less(i, j int) bool {
	return sm.m[sm.k[i]] > sm.m[sm.k[j]]
}

func (sm *sortedFreqStatMap) Swap(i, j int) {
	sm.k[i], sm.k[j] = sm.k[j], sm.k[i]
}

func viewStatistics() {
	statFile, err := os.Open("/tmp/fiat")
	if err != nil {
		log.Fatal("Couldn't open /tmp/fiat! ", err)
	}
	BufReader := bufio.NewReader(statFile)
	for {
		code, err := BufReader.ReadByte()
		freq, err := BufReader.ReadByte()
		if err != nil {
			fmt.Println("EOF")
			break
		}
		FreqStatMap[int(code)] = FreqStatMap[int(code)] + binary.LittleEndian.Uint32([]byte{freq, 0, 0, 0})
	}
	sm := new(sortedFreqStatMap)
	sm.m = FreqStatMap
	sm.k = make([]int, len(sm.m))
	i := 0
	for key, _ := range FreqStatMap {
		sm.k[i] = key
		i++
	}
	sort.Sort(sm)
	for i := 0; i < len(sm.m); i++ {
		if sm.m[sm.k[i]] > 0 {
			fmt.Printf("Key/Freq: %s => %v\n", CodeMap[byte(sm.k[i])], sm.m[sm.k[i]])
		}
	}
	displayKeyboard()
}

func displayKeyboard() {
	fmt.Printf("ESC F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12\n")
	fmt.Printf("~  1  2  3  4  5  6  7  8  9  0  -  +   BS\n")
	fmt.Printf("TAB  Q  W  E  R  T  Y  U  I  O  P  [  ]  \\\n")
	fmt.Printf("CAPS  A  S  D  F  G  H  J  K  L  ;  \\'\n")
	fmt.Printf("LSHIFT Z  X  C  V  B  N  M  ,  .  / RSHIFT\n")
	fmt.Printf("LCTRL LSUPER LALT SPACE         RALT RCTRL\n")
}

func main() {
	app := cli.NewApp()
	app.Name = "fiat"
	app.Usage = "free input analysis tool"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "stat",
			ShortName: "st",
			Usage:     "view current statistics",
			Action: func(c *cli.Context) {
				viewStatistics()
			},
		},
	}
	app.Action = func(c *cli.Context) {
		listen()
	}
	app.Run(os.Args)
}
