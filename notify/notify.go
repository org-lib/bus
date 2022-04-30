package notify

import "github.com/martinlindhe/notify"

func Notify(appname, text string) {
	// show a notification
	notify.Notify(appname, "notice", text, "assets/ico.ico")
}
func Alert(appname, notice, text string) {
	// show a notification and play a alert sound
	notify.Alert(appname, "alert", text, "assets/ico.ico")
}
