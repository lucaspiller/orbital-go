package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 480
	ScreenHeight = 640
)

type Game struct {
	EntityManager *EntityManager
}

func (g *Game) Update() error {
	g.EntityManager.UpdateSystems()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, system := range g.EntityManager.systems {
		if renderSystem, ok := system.(*RenderSystem); ok {
			for _, entity := range g.EntityManager.entities {
				if entity.Enabled {
					renderSystem.Draw(screen, entity)
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	game := &Game{}

	// Create an entity manager
	em := &EntityManager{}

	// Create entity representing sun
	sun := NewEntity("Sun")
	sun.AddComponent(&GraphicsComponent{Color: color.RGBA{255, 255, 0, 255}})
	sun.AddComponent(&PositionComponent{X: ScreenWidth / 2, Y: ScreenHeight / 2})
	em.AddEntity(sun)

	// Create entities representing planets
	planet1 := NewEntity("Planet1")
	planet1.AddComponent(&GraphicsComponent{Color: color.RGBA{0, 0, 255, 255}})
	planet1.AddComponent(&OrbitComponent{Parent: sun, Radius: 50, Speed: 0.02})
	planet1.AddComponent(&PositionComponent{})

	planet2 := NewEntity("Planet2")
	planet2.AddComponent(&GraphicsComponent{Color: color.RGBA{255, 0, 0, 255}})
	planet2.AddComponent(&OrbitComponent{Parent: sun, Radius: 150, Speed: 0.01})
	planet2.AddComponent(&PositionComponent{})

	planet2Moon1 := NewEntity("Planet2Moon1")
	planet2Moon1.AddComponent(&GraphicsComponent{Color: color.RGBA{255, 255, 0, 255}})
	planet2Moon1.AddComponent(&OrbitComponent{Parent: planet2, Radius: 20, Speed: 0.1})
	planet2Moon1.AddComponent(&PositionComponent{})

	// Add planets to the entity manager
	em.AddEntity(planet1)
	em.AddEntity(planet2)
	em.AddEntity(planet2Moon1)

	// Create systems
	em.AddSystem(&OrbitalSystem{})
	em.AddSystem(&RenderSystem{})

	game.EntityManager = em

	return game
}
