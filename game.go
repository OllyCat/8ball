package main

import (
	"errors"
	"image/color"
	"log"
	"math/rand"
	"path/filepath"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	// шар
	ball *ebiten.Image
	// стекло
	glass *ebiten.Image
	// текстуры ответов
	answers []*ebiten.Image
	// цвет фона
	bg color.RGBA
	// параметры текущей текстуры
	alpha float64
	angle float64
	rand  int
	// состояние автомата:
	// 0 - начало: генерим случайности
	// 1 - текстура всплывает
	// 2 - текстура показывается
	// 3 - текстура тонет
	state int
}

func (g *Game) Init() {
	// инициализация
	// грузим текстуры

	// шар:
	img, _, err := ebitenutil.NewImageFromFile("pics/ball.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal("Could not load resurse: ", err)
	}
	g.ball = img

	// стекло:
	img, _, err = ebitenutil.NewImageFromFile("pics/glass.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal("Could not load resurse: ", err)
	}
	g.glass = img

	// ответы:
	fl, err := filepath.Glob("./pics/ans_*.png")
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%#v\n", fl)

	for _, name := range fl {
		img, _, err = ebitenutil.NewImageFromFile(name, ebiten.FilterDefault)
		if err != nil {
			log.Fatal("Could not load resurse: ", err)
		}

		// собираем их в слайс
		g.answers = append(g.answers, img)
	}

	// устанавливаем цвет фона
	g.bg = color.RGBA{R: 0x51, G: 0x9c, B: 0x52, A: 0xff}
	// началный альфа канал
	g.alpha = 0
	// начальное состояние конечного автомата
	g.state = 0
}

func (g *Game) Update(screen *ebiten.Image) error {
	// проверяем нажатие клавишь
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyEscape):
		// если нажата Esc - выходим из игры возвращая сообщение об ошибке
		// не знаю на сколько правильно так выходить, но пока не понял как сделать лучше в ebiten
		return errors.New("\nGame finished\n")
	}

	switch g.state {
	case 0:
		// если клацнули мышкой - меняем состояние автомата
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.state++
			// случайные небольшой угол поворота
			g.angle = 0.45/2 - rand.Float64()*0.45
			// случайный ответ
			g.rand = rand.Intn(len(g.answers))
		}
	case 1:
		// если всплытие - увеличиваем альфа канал для плавного появления
		if g.alpha > 1 {
			// достигли 1-ы в альфа канале - меняем состояние автомата
			g.alpha = 1
			g.state++
		} else {
			// добавляем приращение в альфа канал для плавного появления
			g.alpha += 0.015
		}
	case 2:
		// если клацнули мышкой - меняем состояние автомата
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.state++
		}
	case 3:
		// если альфа меньше 0.02, значит изображение почти исчезло и можно его убирать
		// и менять состояние автомата на нулевое
		if g.alpha < 0.02 {
			g.alpha = 0
			g.state = 0
		} else {
			// инаяе - вычитаем из альфа канала приращение для плавного затухания
			g.alpha -= 0.015
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// отрисовка экрана

	// заполняем цветом фона
	screen.Fill(g.bg)

	// отрисовываем шар
	op := &ebiten.DrawImageOptions{}
	// берём размеры текстуры
	w, h := g.ball.Size()
	// смещаем текстуру в её центр
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	// берём размер экрана
	sw, sh := screen.Size()
	// смещаем текстуру туда
	op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	screen.DrawImage(g.ball, op)

	// если состояние автомата больше 0 - отрисовываем ответ
	if g.state > 0 {
		op = &ebiten.DrawImageOptions{}
		w, h = g.answers[g.rand].Size()

		op.Filter = ebiten.FilterLinear
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(g.angle)
		op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
		op.ColorM.Scale(1, 1, 1, g.alpha)
		screen.DrawImage(g.answers[g.rand], op)
	}

	// отрисовываем стекло
	op = &ebiten.DrawImageOptions{}
	// берём размеры текстуры
	w, h = g.glass.Size()
	// смещаем текстуру в её центр
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	// берём размер экрана
	sw, sh = screen.Size()
	// смещаем текстуру туда
	op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	screen.DrawImage(g.glass, op)

}

func (g *Game) Layout(outW, outH int) (screenWidth, screenHeight int) {
	return 450, 450
}
