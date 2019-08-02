package mutations

import "log"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

// Mode toggles the game cursor's current mode.
func Mode() Mutation {
	return mode{}
}

// Intramode interacts with the current state based on the mode.
func Intramode() Mutation {
	return intramode{}
}

type mode struct {
}

func (m mode) Apply(game objects.Game) objects.Game {
	// If we're currently in move mode and we are carrying something, dump it back to where it belongs.
	if len(game.Cursor.Inventory) == 1 && game.Cursor.Mode == constants.CursorMove {
		b := game.Cursor.Inventory[0]
		game.Buildings = append(game.Buildings, b)
	}

	game.Cursor.Inventory = make([]objects.Building, 0)
	game.Cursor.Mode = (game.Cursor.Mode + 1) % (constants.CursorDemolish + 1)

	if game.Cursor.Mode == constants.CursorBuild {
		game.Cursor.Inventory = []objects.Building{objects.Building{Kind: constants.BuildingHouse}}
	}

	return game
}

type intramode struct {
}

func (i intramode) Apply(game objects.Game) objects.Game {
	if game.Cursor.Mode == constants.CursorBuild && len(game.Cursor.Inventory) == 1 {
		target := game.Cursor.Inventory[0]
		target.Kind = (target.Kind + 1) % (constants.BuildingBusiness + 1)
		log.Printf("toggling build")
		game.Cursor.Inventory = []objects.Building{target}
	}

	return game
}
