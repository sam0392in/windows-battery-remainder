package battery

import (
	"fmt"

	"github.com/distatus/battery"
	"tawesoft.co.uk/go/dialog"
)

func batteryState(bat []*battery.Battery) string {
	var state string
	for _, battery := range bat {
		state = battery.State.String()
	}
	return state
}

func currentPercentage(bat []*battery.Battery) int {
	var percentBattery int
	for _, battery := range bat {
		current := battery.Current
		total := battery.Full
		percentBattery = int((current / total) * 100)
	}
	return percentBattery
}

func alertFullCharted() {
	dialog.Alert("100% Charged, unplug the charger")
}

func alertlowBattery(current int) {
	dialog.Alert("Plugin the charger\nRemaining battery percentage: %d", current)
}

func CheckBattery() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return
	}
	status := batteryState(batteries)
	currentBat := currentPercentage(batteries)
	if currentBat == 100 && (status == "Full" || status == "Charging") {
		alertFullCharted()
	} else if currentBat < 20 && status == "Discharging" {
		alertlowBattery(currentBat)
	}
}
