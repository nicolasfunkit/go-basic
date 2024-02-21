package main

import "strings"

func CountCreatedEvents(events []string) int {
	var count = 0
	for _, status := range events {
		if strings.HasSuffix(status, "_deleted") {
			break
		}
		if !strings.HasSuffix(status, "_created") {
			continue
		}
		count++
	}

	return count
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
