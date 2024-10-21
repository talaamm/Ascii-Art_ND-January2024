package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\u001b[38;5;196m"
	Green  = "\033[32m"
	Yellow = "\u001b[38;5;11m"
	orange = "\u001b[38;5;208m"
	Blue   = "\033[34m"
	Purple = "\u001b[38;5;13m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Black  = "\u001b[38;5;232m"
	pink   = "\033[31m"
)

var thecolor string

func main() {
	myflag := flag.String("color", "white", "red , green , yellow, orange, blue, purple, cyan, white, black , pink")
	flag.Parse()

	if *myflag == "red" {
		thecolor = Red
	} else if *myflag == "green" {
		thecolor = Green
	} else if *myflag == "yellow" {
		thecolor = Yellow
	} else if *myflag == "orange" {
		thecolor = orange
	} else if *myflag == "blue" {
		thecolor = Blue
	} else if *myflag == "purple" {
		thecolor = Purple
	} else if *myflag == "cyan" {
		thecolor = Cyan
	} else if *myflag == "white" {
		thecolor = White
	} else if *myflag == "black" {
		thecolor = Black
	} else if *myflag == "pink" {
		thecolor = pink
	} else {
		fmt.Println("This color is no available! colors: red , green , yellow, orange, blue, purple, cyan, white, black , pink ")
	}

	if os.Args[1] == "--color" {
		fmt.Print(`Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <letters to be colored> "something"`)
		return
	}

	if len(flag.Args()) == 1 {
		firstcase(flag.Args()[0])
	} else if len(flag.Args()[0]) == 1 {
		secondcase(flag.Args()[0], flag.Args()[1])
	} else if strings.Contains(flag.Args()[1], flag.Args()[0]) {
		thirdcase(flag.Args()[0], flag.Args()[1])
	} else {
		forthcase(flag.Args()[0], flag.Args()[1])
	}
}

// colored everything
func firstcase(s string) {
	for i := 0; i < 8; i++ {
		for x := 0; x < len(s); x++ {
			fmt.Print(thecolor + std(rune(s[x]))[i] + Reset)
		}
		fmt.Println()
	}
}

// color 1 letter
func secondcase(letter, sen string) {
	for i := 0; i < 8; i++ {
		for x := 0; x < len(sen); x++ {
			if string(sen[x]) == letter {
				fmt.Print(thecolor + std(rune(sen[x]))[i] + Reset)
			} else {
				fmt.Print(std(rune(sen[x]))[i])
			}
		}
		fmt.Println()
	}
}

// colored 1 word from sen
func thirdcase(word, sen string) {
	index := strings.Index(sen, word) // 4

	for i := 0; i < 8; i++ {
		for x := 0; x < len(sen); x++ {
			if x >= index && x < len(word)+index { // 7
				fmt.Print(thecolor + std(rune(sen[x]))[i] + Reset)
			} else {
				fmt.Print(std(rune(sen[x]))[i])
			}
		}
		fmt.Println()
	}
}

var RuThere = false

func forthcase(let, sen string) {
	for i := 0; i < 8; i++ {
		for _, x := range sen {
			for _, q := range let {
				if q == x {
					RuThere = true
					break
				}
			}
			if RuThere == true {
				fmt.Print(thecolor + std((x))[i] + Reset)
			} else {
				fmt.Print(std(x)[i])
			}
			RuThere = false

		}
		fmt.Println()
	}
}

func std(r rune) []string {
	allLetters := make(map[rune][]string)
	allLetters[' '] = []string{
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	}
	allLetters['!'] = []string{
		" _  ",
		"| | ",
		"| | ",
		"| | ",
		"|_| ",
		"(_) ",
		"    ",
		"    ",
	}
	allLetters['"'] = []string{
		" _ _  ",
		"( | ) ",
		" V V  ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	allLetters['#'] = []string{
		"     _  _    ",
		"   _| || |_  ",
		"  |_  __  _| ",
		"   _| || |_  ",
		"  |_  __  _| ",
		"    |_||_|   ",
		"             ",
		"             ",
	}
	allLetters['$'] = []string{
		"    _   ",
		"   | |  ",
		"  / __) ",
		"  \\__ \\ ",
		"  (   / ",
		"   |_|  ",
		"        ",
		"        ",
	}
	allLetters['%'] = []string{
		" _   __ ",
		"(_) / / ",
		"   / /  ",
		"  / /   ",
		" / / _  ",
		"/_/ (_) ",
		"        ",
		"        ",
	}
	allLetters['&'] = []string{
		"           ",
		"    ___    ",
		"   ( _ )   ",
		"   / _ \\/\\ ",
		"  | (_>  < ",
		"   \\___/\\/ ",
		"           ",
		"           ",
	}
	allLetters['\''] = []string{
		" _  ",
		"( ) ",
		"|/  ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
	}
	allLetters['('] = []string{
		"    __ ",
		"   / / ",
		"  | |  ",
		"  | |  ",
		"  | |  ",
		"  | |  ",
		"   \\_\\ ",
		"       ",
	}
	allLetters[')'] = []string{
		"__   ",
		"\\ \\  ",
		" | | ",
		" | | ",
		" | | ",
		" | | ",
		"/_/  ",
		"     ",
	}
	allLetters['*'] = []string{
		"    _     ",
		" /\\| |/\\  ",
		" \\ ` ' /  ",
		"|_     _| ",
		" / , . \\  ",
		" \\/|_|\\/  ",
		"          ",
		"          ",
	}
	allLetters['+'] = []string{
		"        ",
		"   _    ",
		" _| |_  ",
		"|_   _| ",
		"  |_|   ",
		"        ",
		"        ",
		"        ",
	}
	allLetters['-'] = []string{
		"         ",
		"         ",
		" ______  ",
		"|______| ",
		"         ",
		"         ",
		"         ",
		"         ",
	}
	allLetters['.'] = []string{
		"    ",
		"    ",
		"    ",
		"    ",
		" _  ",
		"(_) ",
		"    ",
		"    ",
	}
	allLetters['/'] = []string{
		"     __ ",
		"    / / ",
		"   / /  ",
		"  / /   ",
		" / /    ",
		"/_/     ",
		"        ",
		"        ",
	}
	allLetters['0'] = []string{
		"        ",
		"  ___   ",
		" / _ \\  ",
		"| | | | ",
		"| |_| | ",
		" \\___/  ",
		"        ",
		"        ",
	}
	allLetters['1'] = []string{
		"    ",
		" _  ",
		"/ | ",
		"| | ",
		"| | ",
		"|_| ",
		"    ",
		"    ",
	}
	allLetters['2'] = []string{
		"        ",
		" ____   ",
		"|___ \\  ",
		"  __) | ",
		" / __/  ",
		"|_____| ",
		"        ",
		"        ",
	}
	allLetters['3'] = []string{
		"        ",
		" _____  ",
		"|___ /  ",
		"  |_ \\  ",
		" ___) | ",
		"|____/  ",
		"        ",
		"        ",
	}
	allLetters['4'] = []string{
		"         ",
		" _  _    ",
		"| || |   ",
		"| || |_  ",
		"|__   _| ",
		"   |_|   ",
		"         ",
		"         ",
	}
	allLetters['5'] = []string{
		"        ",
		" ____   ",
		"| ___|  ",
		"|___ \\  ",
		"  __) | ",
		"|____/  ",
		"        ",
		"        ",
	}
	allLetters['6'] = []string{
		"        ",
		"  __    ",
		" / /    ",
		"| '_ \\  ",
		"| (_) | ",
		" \\___/  ",
		"        ",
		"        ",
	}
	allLetters['7'] = []string{
		"        ",
		" _____  ",
		"|___  | ",
		"   / /  ",
		"  / /   ",
		" /_/    ",
		"        ",
		"        ",
	}
	allLetters['8'] = []string{
		"          ",
		"    ___   ",
		"   ( _ )  ",
		"   / _ \\  ",
		"  | (_) | ",
		"   \\___/  ",
		"          ",
		"          ",
	}
	allLetters['9'] = []string{
		"    ___   ",
		"   / _ \\  ",
		"  | (_) | ",
		"   \\__, | ",
		"     / /  ",
		"    /_/   ",
		"          ",
		"          ",
	}
	allLetters[':'] = []string{
		" _  ",
		"(_) ",
		"    ",
		" _  ",
		"(_) ",
		"    ",
		"    ",
		"    ",
	}
	allLetters[';'] = []string{
		" _  ",
		"(_) ",
		"    ",
		" _  ",
		"( ) ",
		"|/  ",
		"    ",
		"    ",
	}
	allLetters['<'] = []string{
		"   __ ",
		"  / / ",
		" / /  ",
		"< <   ",
		" \\ \\  ",
		"  \\_\\ ",
		"      ",
		"      ",
	}
	allLetters['='] = []string{
		"         ",
		" ______  ",
		"|______| ",
		" ______  ",
		"|______| ",
		"         ",
		"         ",
		"         ",
	}
	allLetters['>'] = []string{
		"__    ",
		"\\ \\   ",
		" \\ \\  ",
		"  > > ",
		" / /  ",
		"/_/   ",
		"      ",
		"      ",
	}
	allLetters['?'] = []string{
		" ___   ",
		"|__ \\  ",
		"   ) | ",
		"  / /  ",
		" |_|   ",
		" (_)   ",
		"       ",
		"       ",
	}
	allLetters['@'] = []string{
		"          ",
		"   ____   ",
		"  / __ \\  ",
		" / / _` | ",
		"| | (_| | ",
		" \\ \\__,_| ",
		"  \\____/  ",
		"          ",
	}
	allLetters['A'] = []string{
		"           ",
		"    /\\     ",
		"   /  \\    ",
		"  / /\\ \\   ",
		" / ____ \\  ",
		"/_/    \\_\\ ",
		"           ",
		"           ",
	}
	allLetters['B'] = []string{
		" ____   ",
		"|  _ \\  ",
		"| |_) | ",
		"|  _ <  ",
		"| |_) | ",
		"|____/  ",
		"        ",
		"        ",
	}
	allLetters['C'] = []string{
		"  _____  ",
		" / ____| ",
		"| |      ",
		"| |      ",
		"| |____  ",
		" \\_____| ",
		"         ",
		"         ",
	}
	allLetters['D'] = []string{
		" _____   ",
		"|  __ \\  ",
		"| |  | | ",
		"| |  | | ",
		"| |__| | ",
		"|_____/  ",
		"         ",
		"         ",
	}
	allLetters['E'] = []string{
		" ______  ",
		"|  ____| ",
		"| |__    ",
		"|  __|   ",
		"| |____  ",
		"|______| ",
		"         ",
		"         ",
	}
	allLetters['F'] = []string{
		" ______  ",
		"|  ____| ",
		"| |__    ",
		"|  __|   ",
		"| |      ",
		"|_|      ",
		"         ",
		"         ",
	}
	allLetters['G'] = []string{
		"  _____  ",
		" / ____| ",
		"| |  __  ",
		"| | |_ | ",
		"| |__| | ",
		" \\_____| ",
		"         ",
		"         ",
	}
	allLetters['H'] = []string{
		" _    _  ",
		"| |  | | ",
		"| |__| | ",
		"|  __  | ",
		"| |  | | ",
		"|_|  |_| ",
		"         ",
		"         ",
	}
	allLetters['I'] = []string{
		" _____  ",
		"|_   _| ",
		"  | |   ",
		"  | |   ",
		" _| |_  ",
		"|_____| ",
		"        ",
		"        ",
	}
	allLetters['J'] = []string{
		"      _  ",
		"     | | ",
		"     | | ",
		" _   | | ",
		"| |__| | ",
		" \\____/  ",
		"         ",
		"         ",
	}
	allLetters['K'] = []string{
		" _  __ ",
		"| |/ / ",
		"| ' /  ",
		"|  <   ",
		"| . \\  ",
		"|_|\\_\\ ",
		"       ",
		"       ",
	}
	allLetters['L'] = []string{
		" _       ",
		"| |      ",
		"| |      ",
		"| |      ",
		"| |____  ",
		"|______| ",
		"         ",
		"         ",
	}
	allLetters['M'] = []string{
		" __  __  ",
		"|  \\/  | ",
		"| \\  / | ",
		"| |\\/| | ",
		"| |  | | ",
		"|_|  |_| ",
		"         ",
		"         ",
	}
	allLetters['N'] = []string{
		" _   _  ",
		"| \\ | | ",
		"|  \\| | ",
		"| . ` | ",
		"| |\\  | ",
		"|_| \\_| ",
		"        ",
		"        ",
	}
	allLetters['O'] = []string{
		"  ____   ",
		" / __ \\  ",
		"| |  | | ",
		"| |  | | ",
		"| |__| | ",
		" \\____/  ",
		"         ",
		"         ",
	}
	allLetters['P'] = []string{
		" _____   ",
		"|  __ \\  ",
		"| |__) | ",
		"|  ___/  ",
		"| |      ",
		"|_|      ",
		"         ",
		"         ",
	}
	allLetters['Q'] = []string{
		"  ____   ",
		" / __ \\  ",
		"| |  | | ",
		"| |  | | ",
		"| |__| | ",
		" \\___\\_\\ ",
		"         ",
		"         ",
	}
	allLetters['R'] = []string{
		" _____   ",
		"|  __ \\  ",
		"| |__) | ",
		"|  _  /  ",
		"| | \\ \\  ",
		"|_|  \\_\\ ",
		"         ",
		"         ",
	}
	allLetters['S'] = []string{
		"  _____  ",
		" / ____| ",
		"| (___   ",
		" \\___ \\  ",
		" ____) | ",
		"|_____/  ",
		"         ",
		"         ",
	}
	allLetters['T'] = []string{
		" _______  ",
		"|__   __| ",
		"   | |    ",
		"   | |    ",
		"   | |    ",
		"   |_|    ",
		"          ",
		"          ",
	}
	allLetters['U'] = []string{
		" _    _  ",
		"| |  | | ",
		"| |  | | ",
		"| |  | | ",
		"| |__| | ",
		" \\____/  ",
		"         ",
		"         ",
	}
	allLetters['V'] = []string{
		"__      __ ",
		"\\ \\    / / ",
		" \\ \\  / /  ",
		"  \\ \\/ /   ",
		"   \\  /    ",
		"    \\/     ",
		"           ",
		"           ",
	}
	allLetters['W'] = []string{
		"__          __ ",
		"\\ \\        / / ",
		" \\ \\  /\\  / /  ",
		"  \\ \\/  \\/ /   ",
		"   \\  /\\  /    ",
		"    \\/  \\/     ",
		"               ",
		"               ",
	}
	allLetters['X'] = []string{
		"__   __ ",
		"\\ \\ / / ",
		" \\ V /  ",
		"  > <   ",
		" / . \\  ",
		"/_/ \\_\\ ",
		"        ",
		"        ",
	}
	allLetters['Y'] = []string{
		"__     __ ",
		"\\ \\   / / ",
		" \\ \\_/ /  ",
		"  \\   /   ",
		"   | |    ",
		"   |_|    ",
		"          ",
		"          ",
	}
	allLetters['Z'] = []string{
		" ______ ",
		"|___  / ",
		"   / /  ",
		"  / /   ",
		" / /__  ",
		"/_____| ",
		"        ",
		"        ",
	}
	allLetters['['] = []string{
		" ___  ",
		"|  _| ",
		"| |   ",
		"| |   ",
		"| |   ",
		"| |_  ",
		"|___| ",
		"      ",
	}
	allLetters[']'] = []string{
		" ___  ",
		"|_  | ",
		"  | | ",
		"  | | ",
		"  | | ",
		" _| | ",
		"|___| ",
		"      ",
	}
	allLetters['\\'] = []string{
		"__      ",
		"\\ \\     ",
		" \\ \\    ",
		"  \\ \\   ",
		"   \\ \\  ",
		"    \\_\\ ",
		"        ",
		"        ",
	}
	allLetters['^'] = []string{
		" /\\   ",
		"|/\\|  ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	allLetters['_'] = []string{
		"         ",
		"         ",
		"         ",
		"         ",
		"         ",
		"         ",
		" ______  ",
		"|______| ",
	}
	allLetters[','] = []string{
		"    ",
		"    ",
		"    ",
		"    ",
		" _  ",
		"( ) ",
		"|/  ",
		"    ",
	}
	allLetters['`'] = []string{
		" _  ",
		"( ) ",
		" \\| ",
		"    ",
		"    ",
		"    ",
		"    ",

		"    ",
	}
	allLetters['a'] = []string{
		"        ",
		"        ",
		"  __ _  ",
		" / _` | ",
		"| (_| | ",
		" \\__,_| ",
		"        ",
		"        ",
	}
	allLetters['b'] = []string{
		" _      ",
		"| |     ",
		"| |__   ",
		"| '_ \\  ",
		"| |_) | ",
		"|_.__/  ",
		"        ",
		"        ",
	}
	allLetters['c'] = []string{
		"       ",
		"       ",
		"  ___  ",
		" / __| ",
		"| (__  ",
		" \\___| ",
		"       ",
		"       ",
	}
	allLetters['e'] = []string{
		"       ",
		"       ",
		"  ___  ",
		" / _ \\ ",
		"|  __/ ",
		" \\___| ",
		"       ",
		"       ",
	}
	allLetters['f'] = []string{
		"  __  ",
		" / _| ",
		"| |_  ",
		"|  _| ",
		"| |   ",
		"|_|   ",
		"      ",
		"      ",
	}
	allLetters['g'] = []string{
		"        ",
		"        ",
		"  __ _  ",
		" / _` | ",
		"| (_| | ",
		" \\__, | ",
		"  __/ | ",
		" |___/  ",
	}
	allLetters['d'] = []string{
		"     _  ",
		"    | | ",
		"  __| | ",
		" / _` | ",
		"| (_| | ",
		" \\__,_| ",
		"        ",
		"        ",
	}
	allLetters['h'] = []string{
		" _      ",
		"| |     ",
		"| |__   ",
		"|  _ \\  ",
		"| | | | ",
		"|_| |_| ",
		"        ",
		"        ",
	}
	allLetters['i'] = []string{
		" _  ",
		"(_) ",
		" _  ",
		"| | ",
		"| | ",
		"|_| ",
		"    ",
		"    ",
	}
	allLetters['j'] = []string{
		"   _  ",
		"  (_) ",
		"   _  ",
		"  | | ",
		"  | | ",
		"  | | ",
		" _/ | ",
		"|__/  ",
	}
	allLetters['k'] = []string{
		"       ",
		" _     ",
		"| | _  ",
		"| |/ / ",
		"|   <  ",
		"|_|\\_\\ ",
		"       ",
		"       ",
	}
	allLetters['l'] = []string{
		" _  ",
		"| | ",
		"| | ",
		"| | ",
		"| | ",
		"|_| ",
		"    ",
		"    ",
	}
	allLetters['m'] = []string{
		"            ",
		"            ",
		" _ __ ___   ",
		"| '_ ` _ \\  ",
		"| | | | | | ",
		"|_| |_| |_| ",
		"            ",
		"            ",
	}
	allLetters['n'] = []string{
		"        ",
		"        ",
		" _ __   ",
		"| '_ \\  ",
		"| | | | ",
		"|_| |_| ",
		"        ",
		"        ",
	}
	allLetters['o'] = []string{
		"        ",
		"        ",
		"  ___   ",
		" / _ \\  ",
		"| (_) | ",
		" \\___/  ",
		"        ",
		"        ",
	}
	allLetters['p'] = []string{
		"        ",
		"        ",
		" _ __   ",
		"| '_ \\  ",
		"| |_) | ",
		"| .__/  ",
		"| |     ",
		"|_|     ",
	}
	allLetters['q'] = []string{
		"        ",
		"        ",
		"  __ _  ",
		" / _` | ",
		"| (_| | ",
		" \\__, | ",
		"    | | ",
		"    |_| ",
	}
	allLetters['r'] = []string{
		"       ",
		"       ",
		" _ __  ",
		"| '__| ",
		"| |    ",
		"|_|    ",
		"       ",
		"       ",
	}
	allLetters['s'] = []string{
		"      ",
		"      ",
		" ___  ",
		"/ __| ",
		"\\__ \\ ",
		"|___/ ",
		"      ",
		"      ",
	}
	allLetters['t'] = []string{
		" _    ",
		"| |   ",
		"| |_  ",
		"| __| ",
		"\\ |_  ",
		" \\__| ",
		"      ",
		"      ",
	}
	allLetters['u'] = []string{
		"        ",
		"        ",
		" _   _  ",
		"| | | | ",
		"| |_| | ",
		" \\__,_| ",
		"        ",
		"        ",
	}
	allLetters['v'] = []string{
		"        ",
		"        ",
		"__   __ ",
		"\\ \\ / / ",
		" \\ V /  ",
		"  \\_/   ",
		"        ",
		"        ",
	}
	allLetters['w'] = []string{
		"           ",
		"           ",
		"__      __ ",
		"\\ \\ /\\ / / ",
		" \\ V  V /  ",
		"  \\_/\\_/   ",
		"           ",
		"           ",
	}
	allLetters['x'] = []string{
		"       ",
		"       ",
		"__  __ ",
		"\\ \\/ / ",
		" >  <  ",
		"/_/\\_\\ ",
		"       ",
		"       ",
	}
	allLetters['y'] = []string{
		"        ",
		"        ",
		" _   _  ",
		"| | | | ",
		"| |_| | ",
		" \\__, | ",
		" __/ /  ",
		"|___/   ",
	}
	allLetters['z'] = []string{
		"      ",
		"      ",
		" ____ ",
		"|_  / ",
		" / /  ",
		"/___| ",
		"      ",
		"      ",
	}
	allLetters['{'] = []string{
		"   __ ",
		"  / / ",
		" | |  ",
		"/ /   ",
		"\\ \\   ",
		" | |  ",
		"  \\_\\ ",

		"      ",
	}
	allLetters['}'] = []string{
		"__    ",
		"\\ \\   ",
		" | |  ",
		"  \\ \\ ",
		"  / / ",
		" | |  ",
		"/_/   ",

		"      ",
	}
	allLetters['|'] = []string{
		" _  ",
		"| | ",
		"| | ",
		"| | ",
		"| | ",
		"| | ",
		"| | ",
		"|_| ",
	}
	allLetters['~'] = []string{
		" /\\/| ",
		"|/\\/  ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	for harf, drawn := range allLetters {
		if harf == r {
			return drawn
		}
	}

	return allLetters[' '] // return nil
}
