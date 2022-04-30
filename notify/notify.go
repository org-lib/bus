package notify

import "github.com/martinlindhe/notify"

func Notify(appname, notice, text string) {
	// show a notification
	notify.Notify("app name", "notice", "some text", "assets/ico.ico")

	// show a notification and play a alert sound
	notify.Alert("app name", "alert", "some text", "assets/ico.ico")
}
