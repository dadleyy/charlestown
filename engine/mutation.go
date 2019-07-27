package engine

type mutation = func(state gameState) gameState

func dup(state gameState) gameState {
	return gameState{
		dimensions: dimensions{
			width:  state.dimensions.width,
			height: state.dimensions.height,
		},
		cursor: cursor{
			x:    state.cursor.x,
			y:    state.cursor.y,
			kind: state.cursor.kind,
		},
	}
}

func move(x int, y int) mutation {
	return func(state gameState) gameState {
		loc := cursor{x: state.cursor.x + x, y: state.cursor.y + y}
		next := dup(state)
		next.cursor = loc
		return next
	}
}

func cursorChange(kind int) mutation {
	return func(state gameState) gameState {
		next := dup(state)

		if next.cursor.kind == kind {
			kind = cursorDefault
		}

		next.cursor.kind = kind
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
