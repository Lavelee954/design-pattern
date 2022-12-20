package ex_code

import (
	_interface "example.com/design-pattern/Adapter/ex_code/interface"
	"fmt"
)

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com _interface.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
