package constants

const (
	// BuildingHouse represents houses.
	BuildingHouse = iota

	// BuildingPark represents parks.
	BuildingPark

	// BuildingRoad represents roads.
	BuildingRoad

	// BuildingPower represents power buildings.
	BuildingPower

	// BuildingBusiness represents a business building that generates revenue.
	BuildingBusiness
)

const (
	// BuildingCostBusiness is the price of a business building.
	BuildingCostBusiness = 3 * EconomyMultiplier

	// BuildingCostPower is the price of power buildings.
	BuildingCostPower = 1 * EconomyMultiplier

	// BuildingCostRoad is the price of roads.
	BuildingCostRoad = int(0.5 * EconomyMultiplier)

	// BuildingCostPark is the price of parks.
	BuildingCostPark = int(0.25 * EconomyMultiplier)

	// BuildingCostHouse is the price of houses.
	BuildingCostHouse = int(0.75 * EconomyMultiplier)
)

const (
	// BuildingRevenueBusiness is the price of a business building.
	BuildingRevenueBusiness = int(0.3 * EconomyMultiplier)

	// BuildingRevenuePower is the price of power buildings.
	BuildingRevenuePower = int(-0.10 * EconomyMultiplier)

	// BuildingRevenueRoad is the price of roads.
	BuildingRevenueRoad = 0

	// BuildingRevenuePark is the price of parks.
	BuildingRevenuePark = 0

	// BuildingRevenueHouse is the price of houses.
	BuildingRevenueHouse = int(0.15 * EconomyMultiplier)
)

const (
	// NeighborNorth represents the direction use on buildings that are above another.
	NeighborNorth = "NORTH"
	// NeighborEast represents the direction use on buildings that are located to the right of another.
	NeighborEast = "EAST"
	// NeighborSouth represents the direction use on buildings that are located below another.
	NeighborSouth = "SOUTH"
	// NeighborWest represents the direction use on buildings that are located to the left another.
	NeighborWest = "WEST"
	// NeighborPower is a special neighbor relationship. Represents the building providing power to another.
	NeighborPower = "POWER"
)
