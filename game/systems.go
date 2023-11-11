// systems.go
package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// System interface defines the generic update method.
type System interface {
	Update(*Entity)
}

// OrbitalSystem updates the position of the components based on the orbit data.
type OrbitalSystem struct{}

// Update implements the logic to update the orbital system.
func (os *OrbitalSystem) Update(entity *Entity) {
	orbitComponent, ok := GetComponent[OrbitComponent](entity, Orbit)
	if !ok {
		return
	}

	positionComponent, ok := GetComponent[PositionComponent](entity, Position)
	if !ok {
		return
	}

	// Update the angle
	orbitComponent.Angle += orbitComponent.Speed

	// Get the parent position
	parentPosition, _ := GetComponent[PositionComponent](orbitComponent.Parent, Position)

	// Calculate the new position
	positionComponent.X = parentPosition.X + orbitComponent.Radius*math.Cos(orbitComponent.Angle)
	positionComponent.Y = parentPosition.Y + orbitComponent.Radius*math.Sin(orbitComponent.Angle)

}

// RenderSystem renders the screen.
type RenderSystem struct{}

// Update implements the logic to update the render system.
func (rs *RenderSystem) Update(entity *Entity) {}

func DrawCircleOutline(screen *ebiten.Image, centerX, centerY, radius float64, clr color.Color) {
	const numSteps = 360
	const angleStep = 2 * math.Pi / numSteps

	for i := 0; i < numSteps; i++ {
		angle := float64(i) * angleStep
		nextAngle := float64(i+1) * angleStep

		x1 := centerX + radius*math.Cos(angle)
		y1 := centerY + radius*math.Sin(angle)
		x2 := centerX + radius*math.Cos(nextAngle)
		y2 := centerY + radius*math.Sin(nextAngle)

		ebitenutil.DrawLine(screen, x1, y1, x2, y2, clr)
	}
}

func (rs *RenderSystem) Draw(screen *ebiten.Image, entity *Entity) {
	orbitComponent, ok := GetComponent[OrbitComponent](entity, Orbit)
	if ok {
		// Draw the orbit path
		parentPosition, _ := GetComponent[PositionComponent](orbitComponent.Parent, Position)
		DrawCircleOutline(screen, parentPosition.X, parentPosition.Y, orbitComponent.Radius, color.White)
	}

	positionComponent, ok := GetComponent[PositionComponent](entity, Position)
	if !ok {
		return
	}

	graphicsComponent, ok := GetComponent[GraphicsComponent](entity, Graphics)
	if !ok {
		return
	}

	// Draw planets
	ebitenutil.DrawCircle(screen, positionComponent.X, positionComponent.Y, 5, graphicsComponent.Color)
}
