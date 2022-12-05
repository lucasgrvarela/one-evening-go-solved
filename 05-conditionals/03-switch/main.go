package main

func Direction(d string) string {
	switch d {
	case "N":
		return "north"
	case "E":
		return "east"
	case "S":
		return "south"
	case "W":
		return "west"
	default:
		return "unknown"
	}
}
