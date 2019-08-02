package mutations

import "log"
import "fmt"
import "time"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

// Interact mutates the state based on the current mode.
func Interact() Mutation {
	return interact{}
}

type interact struct {
}

func (i interact) replaceNeighbor(list []objects.Neighbor, addition objects.Neighbor) []objects.Neighbor {
	next := make([]objects.Neighbor, 0, len(list))

	for _, n := range list {
		if n.Direction == addition.Direction {
			continue
		}

		next = append(next, n)
	}

	return append(next, addition)
}

func (i interact) build(next objects.Game) objects.Game {
	addditions := make([]objects.Building, 0, len(next.Cursor.Inventory)+len(next.Buildings))
	cache := make(map[int]map[int]objects.Building)

	// Start the new addition by looping over our currently inventory, indexing them by their x and y coordinates.
	for _, item := range next.Cursor.Inventory {
		construct := objects.Building{next.Cursor.Location, item.Kind, make([]objects.Neighbor, 0, 5)}

		if construct.Cost() > next.Funds {
			message := fmt.Sprintf("Not enough funds")
			expiry := time.Now().Add(time.Second * 3)
			next.Flashes = append(next.Flashes, objects.Flash{message, expiry})
			continue
		}

		x, y := construct.Location.Values()
		column, ok := cache[x]

		if !ok {
			cache[x] = map[int]objects.Building{}
			column = cache[x]
		}

		column[y] = construct
		next.Funds -= construct.Cost()
	}

	log.Printf("addition cache: %v", cache)

	// Loop over our current set of buildings, checking the additions to see if there is a location match.
	for _, building := range next.Buildings {
		x, y := building.Location.Values()
		column, ok := cache[x]

		// If we found our column, we should check north and south.
		if ok {
			// If there is a match at our y coordinate it is a duplicate.
			if _, ok := column[y]; ok {
				log.Printf("dupe building detected not adding")
				continue
			}

			// Check for a neighbor to the north.
			if north, hit := column[y-1]; hit {
				building.Neighbors = i.replaceNeighbor(building.Neighbors, objects.Neighbor{north, constants.NeighborNorth})
				north.Neighbors = i.replaceNeighbor(north.Neighbors, objects.Neighbor{building, constants.NeighborSouth})
				column[y-1] = north
			}

			// Check for a neighbor to the south.
			if south, hit := column[y+1]; hit {
				building.Neighbors = i.replaceNeighbor(building.Neighbors, objects.Neighbor{south, constants.NeighborSouth})
				south.Neighbors = i.replaceNeighbor(south.Neighbors, objects.Neighbor{building, constants.NeighborNorth})
				column[y+1] = south
			}
		}

		// Check for a neighbor to the west.
		if west, ok := cache[x-2]; ok {
			if neighbor, ok := west[y]; ok {
				building.Neighbors = i.replaceNeighbor(building.Neighbors, objects.Neighbor{neighbor, constants.NeighborWest})
				neighbor.Neighbors = i.replaceNeighbor(neighbor.Neighbors, objects.Neighbor{building, constants.NeighborEast})
				cache[x-2][y] = neighbor
			}
		}

		// Check for a neighbor to the east.
		if east, ok := cache[x+2]; ok {
			if neighbor, ok := east[y]; ok {
				building.Neighbors = i.replaceNeighbor(building.Neighbors, objects.Neighbor{neighbor, constants.NeighborWest})
				neighbor.Neighbors = i.replaceNeighbor(neighbor.Neighbors, objects.Neighbor{building, constants.NeighborEast})
				cache[x+2][y] = neighbor
			}
		}

		addditions = append(addditions, building)
	}

	// Cache is now ready - insert them
	for _, column := range cache {
		for _, cell := range column {
			addditions = append(addditions, cell)
		}
	}

	next.Buildings = addditions
	return next
}

func (i interact) demolish(game objects.Game) objects.Game {
	buildings := make([]objects.Building, 0, len(game.Buildings))

	for _, b := range game.Buildings {
		match := b.Location.Equals(game.Cursor.Location)

		if match {
			continue
		}

		neighbors := make([]objects.Neighbor, 0, len(b.Neighbors))

		// Remove the neighbor
		for _, c := range b.Neighbors {
			if c.Building.Location.Equals(game.Cursor.Location) {
				continue
			}

			neighbors = append(neighbors, c)
		}

		b.Neighbors = neighbors

		buildings = append(buildings, b)
	}

	game.Buildings = buildings
	return game
}

func (i interact) move(game objects.Game) objects.Game {
	if len(game.Cursor.Inventory) == 1 {
		log.Printf("completing move")
		next := i.build(game)
		next.Cursor.Inventory = []objects.Building{}
		return next
	}

	log.Printf("initiating move")
	buildings := make([]objects.Building, 0, len(game.Buildings))

	for _, building := range game.Buildings {
		hit := building.Location.Equals(game.Cursor.Location)

		if !hit {
			buildings = append(buildings, building)
			continue
		}

		game.Cursor.Inventory = []objects.Building{building}
	}

	game.Buildings = buildings
	return game
}

func (i interact) Apply(game objects.Game) objects.Game {
	next := game

	if next.Turn.Done() {
		next.Flashes = append(next.Flashes, objects.Flash{"No actions remaining", time.Now().Add(time.Second * 5)})
		return next
	}

	switch next.Cursor.Mode {
	case constants.CursorDemolish:
		next.Turn = next.Turn.Inc()
		return i.demolish(next)
	case constants.CursorBuild:
		next.Turn = next.Turn.Inc()
		return i.build(next)
	case constants.CursorMove:
		return i.move(next)
	}

	return next
}
