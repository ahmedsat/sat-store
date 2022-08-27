package resources

type Entity interface {
	Create() (Entity, error)
	// save() (entity, error)
	// update() (entity, error)
	// patch() (entity, error)
	// delete() (entity, error)
	// findOne(map[string]string) (entity, error)
	Find(map[string]string) ([]Entity, error)
	New() Entity
}
