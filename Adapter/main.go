package main

import (
	"example.com/design-pattern/Adapter/ex_code"
	"example.com/design-pattern/Adapter/ex_code/mac"
	"example.com/design-pattern/Adapter/ex_code/window"
)

func main() {

	client := &ex_code.Client{}

	mac := &mac.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &window.Windows{}
	windowsMachineAdapter := &window.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
