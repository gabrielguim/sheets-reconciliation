package commons

type Entry struct {
	Name  string
	Value uint32
}

type Area struct {
	Receipts []Entry
	Payments []Entry
	Name     string
}
