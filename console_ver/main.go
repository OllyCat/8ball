package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/logrusorgru/aurora"
)

func main() {
	var resp = [...]string{
		"Это бесспорно",
		"Это предрешено",
		"Никаких сомнений",
		"Определённо да",
		"Можешь быть уверен в этом",
		"Мне кажется — да",
		"Вероятнее всего",
		"Выглядит хорошо",
		"Знаки говорят — «да»",
		"Да",
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
