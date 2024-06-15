package container

type Publisher interface {
	Publish(message string)
}
