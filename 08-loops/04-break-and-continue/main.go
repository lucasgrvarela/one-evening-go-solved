package main

import "strings"

func CountCreatedEvents(events []string) (i int) {
	for _, s := range events {
		if strings.HasSuffix(s, "_deleted") {
			break
		}

		if strings.HasSuffix(s, "_created") {
			i++
		} else {
			continue
		}
	}
	return i
}

func main() {
	events := []string{
		"product_created",
		"product_updated",
		"product_assigned",
		"order_created",
		"order_updated",
		"client_created",
		"client_updated",
		"client_refreshed",
		"client_deleted",
		"order_updated",
	}

	CountCreatedEvents(events)
}
