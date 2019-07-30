package engine

type renderer interface {
	generate(gameState) []renderable
}
