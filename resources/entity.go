package resources

type Entity interface {
	Create() (Entity, error)
	// Save() (entity, error)
	// Update() (entity, error)
	// Patch() (entity, error)
	// Delete() (entity, error)
	FindOne(map[string]string) (Entity, error)
	Find(map[string]string) ([]Entity, error)
	New() Entity
}
