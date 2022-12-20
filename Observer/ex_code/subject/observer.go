package subject

type Observer interface {
	update(string)
	getID() string
}
