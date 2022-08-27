package resources

type Entity interface {
	Create() (Entity, error)
	// save() (entity, error)
	// update() (entity, error)
	// patch() (entity, error)
	// delete() (entity, error)
	// findOne(map[string]any) (entity, error)
	Find(map[string]any) ([]Entity, error)
	New() Entity
}
