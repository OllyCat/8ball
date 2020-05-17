package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	// текстуры
	textures []*ebiten.Image
	// цвет фона
	bg color.RGBA
	// альфа канал всплывающей текстуры
	alpha float64
	// состояние автомата:
	// 0 - начало, ничего не делаем
	// 1 - текстура всплывает
	// 2 - текстура показывается
	// 3 - текстура тонет
	state int
}

func (g *Game) Init() {
	// инициализация
	// грузим текстуры
	img, _, err := ebitenutil.NewImageFromFile("pics/ball.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal("Could not load resurse: ", err)
	}

	// собираем их в слайс
	g.textures = append(g.textures, img)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("pics/ans_%02d.png", i)
		img, _, err = ebitenutil.NewImageFromFile(name, ebiten.FilterDefault)
		if err != nil {
			log.Fatal("Could not load resurse: ", err)
		}

		// собираем их в слайс
		g.textures = append(g.textures, img)
	}

	// устанавливаем цвет фона
	g.bg = color.RGBA{R: 128, G: 128, B: 128, A: 255}
	g.alpha = 0
	g.state = 0
}

func (g *Game) Update(screen *ebiten.Image) error {
	log.Printf("State: %v\n", g.state)

	switch g.state {
	case 0:
		return nil
	case 1:
		if g.alpha > 1 {
			g.alpha = 1
			g.state++
			return nil
		}
		g.alpha += 0.015
		log.Printf("Alpha: %v\n", g.alpha)
	case 2:
		return nil
	case 3:
		if g.alpha < 0.02 {
			g.alpha = 0
			g.state = 0
			return nil
		}
		g.alpha -= 0.015
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
	w, h := g.textures[0].Size()
	// смещаем текстуру в её центр
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	// берём размер экрана
	sw, sh := screen.Size()
	// смещаем текстуру туда
	op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	screen.DrawImage(g.textures[0], op)

	// если состояние автомата больше 0 - отрисовываем ответ
	if g.state > 0 {
		op = &ebiten.DrawImageOptions{}
		w, h = g.textures[1].Size()
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
		op.ColorM.Scale(1, 1, 1, g.alpha)
		screen.DrawImage(g.textures[1], op)
	}
}

func (g *Game) Layout(outW, outH int) (screenWidth, screenHeight int) {
	return 450, 450
}
