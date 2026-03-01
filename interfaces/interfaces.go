package interfaces

type Saver interface {
	Save() (string, error)
}
