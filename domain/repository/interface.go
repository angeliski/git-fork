package repository

// Service provide a way to interact with the repository
type Service interface {
	Fetch(repository Model) error
	Pull(repository Model) error
	Push(repository Model) error
	AddRemote(repository Model) error
	HasRemote(remoteName string) bool
}
