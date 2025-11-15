package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func prepareMatrix(kw string) [5][5]rune {
	ltrs := []rune("ABCDEFGHIJKLMNOPRSTUVWXYZ")
	seen := make(map[rune]bool)
	mtx := [5][5]rune{}

	kw = strings.ToUpper(kw)
	ck := ""
	for _, ch := range kw {
		if !seen[ch] && strings.ContainsRune(string(ltrs), ch) {
			ck += string(ch)
			seen[ch] = true
		}
	}
	for _, ch := range ltrs {
		if !seen[ch] {
			ck += string(ch)
			seen[ch] = true
		}
	}

	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			mtx[i][j] = rune(ck[k])
			k++
		}
	}
	return mtx
}

func findPos(mtx [5][5]rune, ch rune) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if mtx[i][j] == ch {
				return i, j
			}
		}
	}
	return -1, -1
}

func makePairs(text string) []string {
	text = strings.ToUpper(strings.ReplaceAll(text, " ", ""))
	ps := []string{}
	rs := []rune(text)

	for i := 0; i < len(rs); i++ {
		a := rs[i]
		var b rune
		if i+1 < len(rs) {
			b = rs[i+1]
			if a == b {
				b = 'X'
			} else {
				i++
			}
		} else {
			b = 'X'
		}
		ps = append(ps, string([]rune{a, b}))
	}
	return ps
}

func encode(kw, text string) string {
	mtx := prepareMatrix(kw)
	ps := makePairs(text)
	res := ""

	for _, p := range ps {
		a, b := rune(p[0]), rune(p[1])
		r1, c1 := findPos(mtx, a)
		r2, c2 := findPos(mtx, b)

		if r1 == r2 {
			res += string(mtx[r1][(c1+1)%5])
			res += string(mtx[r2][(c2+1)%5])
		} else if c1 == c2 {
			res += string(mtx[(r1+1)%5][c1])
			res += string(mtx[(r2+1)%5][c2])
		} else {
			res += string(mtx[r1][c2])
			res += string(mtx[r2][c1])
		}
	}
	return res
}

func decode(kw, text string) string {
	mtx := prepareMatrix(kw)
	ps := makePairs(text)
	res := ""

	for _, pair := range ps {
		a, b := rune(pair[0]), rune(pair[1])
		r1, c1 := findPos(mtx, a)
		r2, c2 := findPos(mtx, b)

		if r1 == r2 {
			res += string(mtx[r1][(c1+4)%5])
			res += string(mtx[r2][(c2+4)%5])
		} else if c1 == c2 {
			res += string(mtx[(r1+4)%5][c1])
			res += string(mtx[(r2+4)%5][c2])
		} else {
			res += string(mtx[r1][c2])
			res += string(mtx[r2][c1])
		}
	}
	return res
}

func main() {
	a := app.New()
	w := a.NewWindow("Шифр Плейфера")

	key := widget.NewEntry()
	key.SetPlaceHolder("Ключевое слово")

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Введите текст")

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Результат")
	output.Disable()

	encryptButton := widget.NewButton("Зашифровать", func() {
		key := key.Text
		txt := input.Text
		res := encode(key, txt)
		output.SetText(res)
	})

	decryptButton := widget.NewButton("Расшифровать", func() {
		key := key.Text
		txt := input.Text
		res := decode(key, txt)
		output.SetText(res)
	})

	buttons := container.NewHBox(encryptButton, decryptButton)

	content := container.NewVBox(
		key,
		input,
		buttons,
		output,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
