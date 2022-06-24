package main

import (
	"drinkingtime/pkg/img"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/robfig/cron/v3"
)

var myApp fyne.App

func main() {
	myApp = app.New()
	myApp.SetIcon(img.ResourceIconPng)
	c := cron.New(cron.WithSeconds())
	c.AddFunc("01 30 7,9,11,13,15,17,19,21 * * *", func() {
		drinkImage()
	})
	c.Start()
	myApp.Run()
	select {}
}

func drinkImage() {
	w := myApp.NewWindow("Drinking time")
	w.CenterOnScreen()
	w.SetTitle("Drinking time")
	image := canvas.NewImageFromResource(img.ResourceDrinkJpeg)
	w.SetContent(image)
	w.Resize(fyne.NewSize(1920, 1080))
	w.CenterOnScreen()
	w.Show()
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	st := time.Now()
	go func() {
		for range time.Tick(time.Second) {
			if time.Now().Sub(st) >= 3*time.Second {
				w.Hide()
				break
			}
		}
	}()
}
