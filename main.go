package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/logrusorgru/aurora"
)

func main() {
	ebiten.SetWindowSize(450, 450)
	ebiten.SetWindowTitle("Magic ball")

	game := &Game{}
	game.Init()
	ebiten.SetMaxTPS(30)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
func cons() {
	var resp = [...]string{
		"Пока не ясно, попробуй ещё раз",
		"Спроси попозже",
		"Лучше я не буду тебе об этом говорить",
		"Не могу сейчас предсказать",
		"Сконцентрируйся и спроси опять",

		"Даже не думай",
		"Мой ответ — «нет»",
		"По моим данным — «нет»",
		"Выглядит не очень хорошо",
		"Весьма сомнительно",
	}
	// rand.Seed(int64(time.Now().Nanosecond()))
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	ans := r.Intn(len(resp))
	//ans = 16
	color := aurora.Green
	switch {
	case ans >= 10:
		color = aurora.Yellow
	case ans > 14:
		color = aurora.BgRed
	}
	fmt.Println(color(resp[ans]))
}
