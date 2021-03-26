package ui

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
)

// NewInfoPage displays the information about a task.
func NewInfoPage(p *BoardPage, listIdx, taskIdx int) tview.Primitive {
	task, err := p.data.GetTask(listIdx, taskIdx)
	if err != nil {
		app.Stop()
		log.Fatal(err)
	}
	info := tview.NewModal().
		SetText(fmt.Sprintf("Task: %v\n Task Description: %v", task[0], task[1])).
		SetBackgroundColor(theme.PrimitiveBackgroundColor).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				pages.RemovePage("info")
				pages.SwitchToPage("board")
				app.SetFocus(p.lists[p.activeListIdx])
			}
		})
	width, height := GetSize()
	return GetCenteredModal(info, width/2, height/2)
}
