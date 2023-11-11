// components.go
package game

import "image/color"

// ComponentType represents the type of a component.
type ComponentType int

const (
	Graphics ComponentType = iota
	Orbit
	Position
)

// Component represents a generic component.
type Component interface {
	ComponentType() ComponentType
}

// GraphicsComponent holds data related to rendering.
type GraphicsComponent struct {
	Color color.Color
}

func (c *GraphicsComponent) ComponentType() ComponentType { return Graphics }

// OrbitComponent contains data on the orbit of the planet.
type OrbitComponent struct {
	Radius float64
	Speed  float64
	Angle  float64
	Parent *Entity
}

func (c *OrbitComponent) ComponentType() ComponentType { return Orbit }

// PositionComponent holds data on the current position of the planet.
type PositionComponent struct {
	X, Y float64
}

func (c *PositionComponent) ComponentType() ComponentType { return Position }
