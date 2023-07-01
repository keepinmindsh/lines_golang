package pattern

import (
	"fmt"
	"testing"
)

type ComputerBridge interface {
	Print()
	SetPrinter(Printer)
}

type MacBridge struct {
	printer Printer
}

func (m *MacBridge) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *MacBridge) SetPrinter(p Printer) {
	m.printer = p
}

type WindowsBridge struct {
	printer Printer
}

func (w *WindowsBridge) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *WindowsBridge) SetPrinter(p Printer) {
	w.printer = p
}

type Printer interface {
	PrintFile()
}

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func Test_Bridge(t *testing.T) {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &MacBridge{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &WindowsBridge{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
