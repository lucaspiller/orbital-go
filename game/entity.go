package game

// Entity represents an entity in the ECS system.
type Entity struct {
	Tag        string
	Enabled    bool
	Components map[ComponentType]interface{}
}

func NewEntity(tag string) *Entity {
	return &Entity{
		Tag:        tag,
		Enabled:    true,
		Components: make(map[ComponentType]interface{}),
	}
}

// AddComponent adds a component to an entity.
func (e *Entity) AddComponent(component Component) {
	e.Components[component.ComponentType()] = component
}

func GetComponent[T any](entity *Entity, componentType ComponentType) (*T, bool) {
	component := entity.Components[componentType]
	if component == nil {
		return nil, false
	}

	specificComponent, ok := component.(*T)
	if !ok {
		return nil, false
	}

	return specificComponent, true
}

// EntityManager manages entities in the ECS system.
type EntityManager struct {
	entities []*Entity
	systems  []System
}

// AddEntity adds a new entity to the ECS system.
func (em *EntityManager) AddEntity(entity *Entity) {
	em.entities = append(em.entities, entity)
}

// AddSystem adds a new system to the ECS.
func (em *EntityManager) AddSystem(system System) {
	em.systems = append(em.systems, system)
}

// UpdateSystems updates all systems in the ECS.
func (em *EntityManager) UpdateSystems() {
	for _, system := range em.systems {
		for _, entity := range em.entities {
			if entity.Enabled {
				system.Update(entity)
			}
		}
	}
}
