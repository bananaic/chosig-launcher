package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 애플리케이션 생성
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello Fyne")

	// 위젯 생성
	label := widget.NewLabel("Hello, World!")
	button := widget.NewButton("Click Me", func() {
		label.SetText("Button Clicked!")
	})

	// 컨테이너 설정
	myWindow.SetContent(container.NewVBox(
		label,
		button,
	))

	// 윈도우 크기와 실행
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
