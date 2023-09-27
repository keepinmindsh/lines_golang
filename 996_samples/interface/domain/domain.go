package domain

type Interface interface {
	Print(string)
	Clone() Interface
}
