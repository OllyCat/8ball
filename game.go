package main

import (
	"errors"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"golang.org/x/exp/shiny/materialdesign/colornames"
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
	r := strings.NewReader(ball_png)
	ball_img, _, err := image.Decode(r)
	if err != nil {
		log.Fatalf("Could not load ball resouse: %v\n", err)
	}
	bimg, err := ebiten.NewImageFromImage(ball_img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal("Could not load resurse: ", err)
	}
	g.ball = bimg

	// стекло:
	r = strings.NewReader(glass_png)
	glass_img, _, err := image.Decode(r)
	if err != nil {
		log.Fatalf("Could not load glass resouse: %v\n", err)
	}
	gimg, err := ebiten.NewImageFromImage(glass_img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal("Could not load resurse: ", err)
	}
	g.glass = gimg

	// ответы:
	for _, a := range answers {
		r = strings.NewReader(a)
		i, _, err := image.Decode(r)
		if err != nil {
			log.Fatalf("Could not load glass resouse: %v\n", err)
		}
		img, err := ebiten.NewImageFromImage(i, ebiten.FilterDefault)
		if err != nil {
			log.Fatal("Could not load resurse: ", err)
		}

		// собираем их в слайс
		g.answers = append(g.answers, img)
	}

	// устанавливаем цвет фона
	//g.bg = color.RGBA{R: 0x51, G: 0x9c, B: 0x52, A: 0xff}
	g.bg = colornames.Green400
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
		return errors.New("Game finished")
	}

	t := ebiten.Touches()

	switch g.state {
	case 0:
		// если клацнули мышкой - меняем состояние автомата
		if len(t) > 0 || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
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
		if len(t) > 0 || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
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
	// берём размер экрана
	sw, sh := screen.Size()

	// отрисовка экрана

	// заполняем цветом фона
	screen.Fill(g.bg)

	// отрисовываем шар
	op := &ebiten.DrawImageOptions{}
	// берём размеры текстуры
	w, h := g.ball.Size()

	// рассчитываем фактор увеличения исходя из размеров экрана
	s := math.Min(float64(sw)/float64(w), float64(sh)/float64(h))

	// увеличиваем текстуру до нужных размеров
	op.GeoM.Scale(s, s)
	w, h = int(float64(w)*s), int(float64(h)*s)
	op.Filter = ebiten.FilterLinear
	// смещаем текстуру в её центр
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	// смещаем текстуру туда
	op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	screen.DrawImage(g.ball, op)

	// если состояние автомата больше 0 - отрисовываем ответ
	if g.state > 0 {
		op = &ebiten.DrawImageOptions{}
		w, h = g.answers[g.rand].Size()
		op.GeoM.Scale(s, s)
		w, h = int(float64(w)*s), int(float64(h)*s)

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

	// увеличиваем
	w, h = g.answers[g.rand].Size()
	op.GeoM.Scale(s, s)
	w, h = int(float64(w)*s), int(float64(h)*s)
	// смещаем текстуру в её центр
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	// смещаем текстуру туда
	op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	screen.DrawImage(g.glass, op)

}

func (g *Game) Layout(outW, outH int) (screenWidth, screenHeight int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outW) * s), int(float64(outH) * s)
}
