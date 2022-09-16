package main

import (
	"context"
	"fmt"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	conn, err := dbus.NewWithContext(context.Background())
	if err != nil {
		panic(fmt.Sprintf("create connection with bus failed: %s", err.Error()))
	}

	sc, ec := conn.SubscribeUnits(10 * time.Second)
	err = conn.Subscribe()
	if err != nil {
		panic(fmt.Sprintf("subscribe dbus failed: %s", err.Error()))
	}

	for {
		select {
		case eventsMap := <-sc:
			for k, v := range eventsMap {
				logrus.Info("[%s] %v", k, v)
			}

		case err = <-ec:
			logrus.Errorf("an error occured when subscribing: %s", err.Error())
		}
	}
}
