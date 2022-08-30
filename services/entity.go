package services

type Entity interface {
	Create() (Entity, error)
	// Save() (entity, error)
	// Update() (entity, error)
	// Patch() (entity, error)
	Delete(map[string]string) (bool, error)
	FindOne(map[string]string) (Entity, error)
	Find(map[string]string) ([]Entity, error)
	New() Entity
}

// DELETE FROM products WHERE id='10'
// DELETE FROM products WHERE WHERE id = '10'
