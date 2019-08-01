package constants

const (
	// CursorNormal is the kind that allows the user to navigate around.
	CursorNormal = iota
	// CursorBuild allows the user to create buildings.
	CursorBuild
	// CursorMove allows the user to move buildings.
	CursorMove
	// CursorDemolish allows the user to delete buildings.
	CursorDemolish
)
