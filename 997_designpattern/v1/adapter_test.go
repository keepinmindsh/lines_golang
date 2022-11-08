package designpattern

import (
	"fmt"
	"testing"
)

type Client struct {
}

type Computer interface {
	InsertIntoLightningPort()
}

// InsertLightningConnectorIntoComputer
// Adapter 를 위한 Interface를 통해 추상화 작업을 수행함.
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

type Mac struct {
}

// InsertIntoLightningPort - 구현체
func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

// insertIntoUSBPort - Adaptee
func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	windowMachine *Windows
}

// InsertIntoLightningPort - 구현체
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

// Test_Adapter - Client
func Test_Adapter(t *testing.T) {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{} // 실제 호출될 값 - 해당 객체는 요청에 따라 다른 버전, 다른 객체로 구현이 가능함.

	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
