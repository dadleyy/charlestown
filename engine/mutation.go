package engine

type mutation = func(state gameState) gameState

func dup(state gameState) gameState {
	return gameState{
		buildings: state.buildings[0:],
		world: dimensions{
			width:  state.world.width,
			height: state.world.height,
		},
		dimensions: dimensions{
			width:  state.dimensions.width,
			height: state.dimensions.height,
		},
		cursor: cursor{
			location: point{x: state.cursor.location.x, y: state.cursor.location.y},
			mode:     state.cursor.mode,
		},
	}
}

func move(x int, y int) mutation {
	return func(state gameState) gameState {
		loc := point{x: state.cursor.location.x + x, y: state.cursor.location.y + y}
		next := dup(state)

		if loc.x < 1 {
			loc.x = 1
		}

		if loc.y < 1 {
			loc.y = 1
		}

		if loc.y > state.world.height-1 {
			loc.y = state.world.height - 1
		}

		if loc.x > state.world.width-1 {
			loc.x = state.world.width - 1
		}

		if loc.y < 0 {
			loc.y = 0
		}

		if loc.x < 0 {
			loc.x = 0
		}

		next.cursor.location = loc
		return next
	}
}

func mode() mutation {
	return func(state gameState) gameState {
		next := dup(state)
		switch state.cursor.mode {
		case cursorNormal:
			next.cursor.mode = cursorBuild
		default:
			next.cursor.mode = cursorNormal
		}
		return next
	}
}

func resize(width, height int) mutation {
	return func(state gameState) gameState {
		next := dup(state)
		next.dimensions = dimensions{width, height}
		return next
	}
}

func build(kind string) mutation {
	return func(state gameState) gameState {
		next := dup(state)
		next.buildings = append(next.buildings, state.cursor.location)
		return next
	}
}
