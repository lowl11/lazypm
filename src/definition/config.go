package definition

var Printer *ConsolePrinter

func Init() {
	Printer = newPrinter()
}
