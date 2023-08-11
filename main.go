package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

//go:embed "charset/frak.64c"
var data []byte

/*
//go:embed "other/message.txt"
var message []byte
*/
func byteToByteSliceByBits(b byte, p byte) []byte {
	bitByByteSlice := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		bit := (b >> uint(i)) & 1
		bitByByteSlice[7-i] = byte(bit * p)
	}
	return bitByByteSlice
}

// renderChar func input Ascii capital letter byte code and returns 8x8 font consist of 0 and 1 - 8 strings by 8x Zeros or Ones
func renderChar(b byte, padding byte) *[][]byte {

	var items = []rune{
		'@', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
		'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '[', '~', ']', '|', '\\', ' ', '!', '"', '#', '$', '%', '&',
		'\'', '(', ')', '*', '+', ',', '-', '.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		':', ';', '<', '=', '>', '?',
	}

	//this func maps code of letters with indices of pixels to read data from, regarding shape of certain semigraphics image of the letter
	translator := make(map[byte]int)
	for i, x := range items {
		translator[byte(x)] = i * 8

	}
	//charset data starts from the 3rd byte
	charset := data[2:]
	//var rendered = make([]string, 0, 8) //create space for semigraphics image consist of Zeros and Ones!
	var rendered = make([][]byte, 8)
	for i := range rendered {
		rendered[i] = make([]byte, 8)
	}

	for y := 0; y < 8; y++ {

		t := translator[b]
		z := t + y
		x := charset[z]
		q := byteToByteSliceByBits(x, padding)
		copy(rendered[y], q)

		//struna := fmt.Sprintf("%08b", x)
		//rendered = append(rendered, struna)
	}
	return &rendered //return address of semigraphics "Big" image 8x8 cursors size.
}

type gTerm struct {
	xmax int
	ymax int
}

type colPair struct {
	fg termbox.Attribute
	bg termbox.Attribute
}

func wrapAndCenterText(text string, maxWidth int) string {
	var result string

	paragraphs := strings.Split(text, "\n")
	for _, paragraph := range paragraphs {
		words := strings.Fields(paragraph)
		line := ""
		for _, word := range words {
			if len(line)+len(word)+1 > maxWidth {
				padding := (maxWidth - len(line)) / 2
				result += strings.Repeat(" ", padding) + line + "\n"
				line = ""
			}
			if line != "" {
				line += " "
			}
			line += word
		}

		padding := (maxWidth - len(line)) / 2
		result += strings.Repeat(" ", padding) + line + "\n"
	}

	return result
}

type shadowScroll struct {
	info          []byte
	currenBitmap  [8][8]byte
	messageShadow [][8]byte
}

func (sh *shadowScroll) loadInfo(s string) {
	s = strings.ToUpper(s)
	for _, c := range s {
		sh.info = append(sh.info, byte(c))
	}
}
func main() {

	message := "ala ma kota feliksa który jest bardzo przytulaśnym kotem."
	message = strings.ToUpper(message)

	//	var sc shadowScroll
	//	sc.info := append(sc.info, []byte(message...))

	//os.Exit(1)
	err := termbox.Init()
	if err != nil {
		fmt.Printf("panic: %v\n", err)
	}
	var console gTerm
	console.xmax, console.ymax = termbox.Size()
	defer termbox.Close()

	asciiChars := []rune{
		'\u2588', // Full block
		'\u2593', // Dark shade
		'\u2592', // Medium shade
		'\u2591', // Light shade
	}

	//fields := []rune{'█', '▓', '▒', '░'}
	//colors := []termbox.Attribute{2, 4} //2 is foreground and 4 is background color
	// termometr is a slice of termbox.Cell which represents foreground background and character for certain color gradient
	var termometr []termbox.Cell
	// spectrum is a bunch of foreground and backround colours to feed the termometr slice
	spectrum := []colPair{{1, 2}, {2, 10}, {10, 12}, {12, 11}, {11, 3}, {3, 7}, {7, 13},
		{13, 14}, {14, 6}, {6, 1},
		{1, 2}, {2, 10}, {10, 16}, {16, 10}, {10, 2}, {2, 1},
		{1, 4}, {4, 12}, {12, 16}, {16, 12}, {12, 4}, {4, 1},
		{1, 3}, {3, 11}, {11, 16}, {16, 11}, {11, 3}, {3, 1},
		{1, 7}, {7, 15}, {15, 16}, {16, 15}, {15, 7}, {7, 1},
		{1, 5}, {5, 13}, {13, 16}, {16, 13}, {13, 5}, {5, 1},
		{1, 6}, {6, 14}, {14, 16}, {16, 14}, {14, 6}, {6, 1},
		{1, 6}, {6, 7}, {7, 11}, {11, 12}, {12, 2}, {2, 1},
		{1, 3}, {3, 11}, {11, 11}, {11, 12}, {12, 10}, {10, 2}, {2, 1},
		{1, 3}, {3, 11}, {11, 12}, {12, 4}, {4, 1},
	}
	// elem - an element to append to termometr slice
	var elem termbox.Cell
	for _, colAtrr := range spectrum {
		for _, char := range asciiChars {
			elem.Fg = colAtrr.fg
			elem.Bg = colAtrr.bg
			elem.Ch = char
			termometr = append(termometr, elem)
		}

	}
	////termbox.Close()
	// sinTab is 256 element slice that helps to save time for calculations of plasma effect
	sinTab := make([]float64, 256, 256)
	// cosTab is 256 element slice consist of cosinus(0-2pi) range and is precalculated from same reason as sinTab
	cosTab := make([]float64, 256, 256)

	var i float64
	index := 0
	//termbox.SetCursor(50, 30)
	for i = 0; i < math.Pi*2; i = i + 0.024639942381096416 {
		valueSin := math.Sin(i)*8 - 4
		valueCos := math.Cos(i)*8 - 4
		sinTab[index] = valueSin
		cosTab[index] = valueCos
		////	fmt.Print(index, " ")
		index++
	}

	go func() {
	label:
		ev := termbox.PollEvent()
		if ev.Key != termbox.KeyEsc {
			goto label
		}
		termbox.Close()
		os.Exit(0)
	}()

	defer fmt.Println(len(termometr))
	defer fmt.Println("program finished with exit code:", 0)
	defer termbox.Close()

	//	go func() {
	var x0start uint8 = 0x40
	var x0_step uint8 = 0x02
	var x0speed uint8 = 0x03
	var y0start uint8 = 0x80
	var y0_step uint8 = 0xfa
	var y0speed uint8 = 0xff

	var x1start uint8 = 0x20
	var x1_step uint8 = 0x01
	var x1speed uint8 = 0xff
	var y1start uint8 = 0x40
	var y1_step uint8 = 0x01
	var y1speed uint8 = 0x01

	var x2start uint8 = 0xff
	var x2_step uint8 = 0xfd
	var x2speed uint8 = 0xff
	var y2start uint8 = 0x07
	var y2_step uint8 = 0xff
	var y2speed uint8 = 0xfe

	go func() {
	label:
		ev := termbox.PollEvent()
		if ev.Key != termbox.KeyEsc {
			goto label
		}
		termbox.Close()
		os.Exit(0)
	}()
	indices := make([]uint8, (console.xmax * console.ymax), (console.xmax * console.ymax))
	for {

		start := time.Now()
		var index float64 = 0
		for y := 0; y < console.ymax; y++ {
			for x := 0; x < console.xmax; x++ {

				sinX0index := x0start + uint8(x)*x0_step
				cosY0index := y0start + uint8(y)*y0_step
				index = sinTab[sinX0index] + cosTab[cosY0index]
				sinX1index := x1start + uint8(x)*x1_step
				cosY1index := y1start + uint8(y)*y1_step
				index = index + sinTab[sinX1index] + cosTab[cosY1index]
				sinX2index := x2start + uint8(x)*x2_step
				cosY2index := y2start + uint8(y)*y2_step
				index = index + sinTab[sinX2index] + cosTab[cosY2index]
				uindex := uint8(index)
				indices[y*console.xmax+x] = uindex

			}
		}
		sign := message[0]
		graph := renderChar(byte(sign), 16)
		baseX, baseY := (console.xmax-8)/2, (console.ymax-8)/2
		for i, g := range *graph {
			for x, valueAdd := range g {
				copyToIndices := baseX + console.xmax*baseY + i*console.xmax + x
				value := indices[copyToIndices] + valueAdd
				indices[copyToIndices] = uint8(value)
			}

		}
		for y := 0; y < console.ymax; y++ {
			for x := 0; x < console.xmax; x++ {
				uindex := indices[y*console.xmax+x]
				termbox.SetCell(x, y, termometr[uindex].Ch, termometr[uindex].Fg, termometr[uindex].Bg)
			}
		}

		x0start = x0start + x0speed
		y0start = y0start + y0speed
		x1start = x1start + x1speed
		y1start = y1start + y1speed
		x2start = x2start + x2speed
		y2start = y2start + y2speed

		termbox.Flush()
		duration := time.Since(start)
		ms := duration.Microseconds()
		rs := time.Duration(33333-ms) * time.Microsecond
		time.Sleep(rs)

		x0start++
	}

}
