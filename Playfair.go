package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func prepareMatrix(keyword string) [5][5]rune {
	alphabet := []rune("ABCDEFGHIJKLMNOPRSTUVWXYZ")
	seen := make(map[rune]bool)
	matrix := [5][5]rune{}

	keyword = strings.ToUpper(keyword)
	cleanKey := ""
	for _, ch := range keyword {
		if !seen[ch] && strings.ContainsRune(string(alphabet), ch) {
			cleanKey += string(ch)
			seen[ch] = true
		}
	}
	for _, ch := range alphabet {
		if !seen[ch] {
			cleanKey += string(ch)
			seen[ch] = true
		}
	}

	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = rune(cleanKey[k])
			k++
		}
	}
	return matrix
}

func findPos(matrix [5][5]rune, ch rune) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == ch {
				return i, j
			}
		}
	}
	return -1, -1
}

func makePairs(text string) []string {
	text = strings.ToUpper(strings.ReplaceAll(text, " ", ""))
	pairs := []string{}
	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		a := runes[i]
		var b rune
		if i+1 < len(runes) {
			b = runes[i+1]
			if a == b {
				b = 'X'
			} else {
				i++
			}
		} else {
			b = 'X'
		}
		pairs = append(pairs, string([]rune{a, b}))
	}
	return pairs
}

func encode(keyword, text string) string {
	matrix := prepareMatrix(keyword)
	pairs := makePairs(text)
	result := ""

	for _, pair := range pairs {
		a, b := rune(pair[0]), rune(pair[1])
		r1, c1 := findPos(matrix, a)
		r2, c2 := findPos(matrix, b)

		if r1 == r2 {
			result += string(matrix[r1][(c1+1)%5])
			result += string(matrix[r2][(c2+1)%5])
		} else if c1 == c2 {
			result += string(matrix[(r1+1)%5][c1])
			result += string(matrix[(r2+1)%5][c2])
		} else {
			result += string(matrix[r1][c2])
			result += string(matrix[r2][c1])
		}
	}
	return result
}

func decode(keyword, text string) string {
	matrix := prepareMatrix(keyword)
	pairs := makePairs(text)
	result := ""

	for _, pair := range pairs {
		a, b := rune(pair[0]), rune(pair[1])
		r1, c1 := findPos(matrix, a)
		r2, c2 := findPos(matrix, b)

		if r1 == r2 {
			result += string(matrix[r1][(c1+4)%5])
			result += string(matrix[r2][(c2+4)%5])
		} else if c1 == c2 {
			result += string(matrix[(r1+4)%5][c1])
			result += string(matrix[(r2+4)%5][c2])
		} else {
			result += string(matrix[r1][c2])
			result += string(matrix[r2][c1])
		}
	}
	return result
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Шифр Плейфера")

	keyEntry := widget.NewEntry()
	keyEntry.SetPlaceHolder("Ключевое слово")

	textEntry := widget.NewMultiLineEntry()
	textEntry.SetPlaceHolder("Введите текст")

	resultEntry := widget.NewMultiLineEntry()
	resultEntry.Disable()

	encryptButton := widget.NewButton("Зашифровать", func() {
		key := keyEntry.Text
		txt := textEntry.Text
		res := encode(key, txt)
		resultEntry.SetText(res)
	})

	decryptButton := widget.NewButton("Расшифровать", func() {
		key := keyEntry.Text
		txt := textEntry.Text
		res := decode(key, txt)
		resultEntry.SetText(res)
	})

	buttons := container.NewHBox(encryptButton, decryptButton)

	content := container.NewVBox(
		keyEntry,
		textEntry,
		buttons,
		resultEntry,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
