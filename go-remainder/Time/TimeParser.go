package TimeParser

import (
	"fmt"
	"github.com/jangirvipin/go-reminder/Notification"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"
	"log"
	"sync"
	"time"
)

func TimeParse(Time string, NotificationMessage string, wg *sync.WaitGroup) {

	w := when.New(nil)
	w.Add(en.All...)

	result, err := w.Parse(Time, time.Now())

	if err != nil {
		log.Fatal("Error parsing time:", err)
	}

	if result == nil {
		log.Fatal("No time found in the input string")
	}

	fmt.Println("Scheduled time:", result.Time)

	duration := time.Until(result.Time)
	fmt.Println("Duration:", duration)

	if duration < 0 {
		log.Fatal("The scheduled time is in the past")
	}

	time.AfterFunc(duration, func() {
		Notification.SendNotification(NotificationMessage)
		wg.Done()
	})

}
