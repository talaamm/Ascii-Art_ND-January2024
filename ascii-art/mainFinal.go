package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 { // for empty string -->  go run main.go ""
		fmt.Print("")
		return
	}

	str := os.Args[1]

	if str == "\\n" {
		fmt.Print("\n")
		return
	}

	split := strings.Split(str, "\\n") //["hello" "\n"] --> ["hello" ""]

	for _, word := range split {
		if word == "" {
			fmt.Print("\n")
		} else {
			for k := 0; k < 8; k++ {
				for i := 0; i < len(word); i++ {
					fmt.Print(letters3(rune(word[i]))[k])
				}
				fmt.Print("\n")
			}
		}
	}
}

func letters3(r rune) []string {
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

	return allLetters[' ']
}
