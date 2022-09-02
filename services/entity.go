package services

type Entity interface {
	Create() (Entity, error)

	Update() (Entity, error)
	Delete(map[string]string) (bool, error)
	FindOne(map[string]string) (Entity, error)
	Find(map[string]string) ([]Entity, error)
	New() Entity
}
