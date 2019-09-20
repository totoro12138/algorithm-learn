package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	mark, err := os.Open("mark.png")
	if err != nil {
		panic(err)
	}
	defer mark.Close()

	img, err := os.Open("image.jpeg")
	if err != nil {
		panic(err)
	}
	defer img.Close()

	markJpg, err := png.Decode(mark)
	if err != nil {
		panic(err)
	}

	imgJpg, err := jpeg.Decode(img)
	if err != nil {
		panic(err)
	}

	offset := image.Pt(imgJpg.Bounds().Dx()-markJpg.Bounds().Dx()-10, imgJpg.Bounds().Dy()-markJpg.Bounds().Dy()-10)
	b := imgJpg.Bounds()
	m := image.NewRGBA(b)

	draw.Draw(m, b, imgJpg, image.ZP, draw.Src)
	draw.Draw(m, markJpg.Bounds().Add(offset), markJpg, image.ZP, draw.Over)

	imgDraw, err := os.Create("water_image.jpg")
	if err != nil {
		panic(err)
	}
	defer imgDraw.Close()

	if err := jpeg.Encode(imgDraw, m, &jpeg.Options{1000}); err != nil {
		panic(err)
	}
}
