package utils

import (
	"github.com/RossJ24/ZendeskCo-opChallenge/models"
)

// Searches the tickets that are currently held in memory for the target ticket
func GetLocalTicket(currentTickets, cachedTickets *[]*models.Ticket, target int) *models.Ticket {
	// Create a concatenated slice of the two ticket lists of a CLI struct
	var localTickets []*models.Ticket
	if cachedTickets == nil {
		localTickets = *currentTickets
	} else {
		localTickets = append(*cachedTickets, *currentTickets...)
	}
	// If there are any nil tickets in the concatenated List return nil
	for _, element := range localTickets {
		if element == nil {
			return nil
		}
	}
	// Return the results of a binary search in the sorted ticket list
	return binarySearch(mergeSort(localTickets), target)
}

// Sorts the localTickets using the MergeSort divide-and-conquer algorithm
func mergeSort(localTickets []*models.Ticket) []*models.Ticket {
	length := len(localTickets)
	// Base Case
	if length == 1 {
		return localTickets
	}
	// Construct left and right slices
	mid := int(length / 2)
	left := make([]*models.Ticket, mid)
	right := make([]*models.Ticket, length-mid)

	for i := 0; i < length; i++ {
		if i < mid {
			left[i] = localTickets[i]
		} else {
			right[i-mid] = localTickets[i]
		}
	}
	// Recurse on the left and right
	return merge(mergeSort(left), mergeSort(right))
}

// Merge sort helper function that merges two subslices (arrays)
func merge(left, right []*models.Ticket) (result []*models.Ticket) {
	// Construct a slice to merge into
	result = make([]*models.Ticket, len(left)+len(right))

	// Merge the left and right slice in sorted order
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].ID < right[0].ID {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// Fill in any remainders
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

// Binary Search of the tickets held in memory for a target ticket
func binarySearch(localTickets []*models.Ticket, target int) *models.Ticket {

	// Initialize the left and right parameters
	l := 0
	r := len(localTickets) - 1
	// While the left bound is not greater than the right bound, search and return the ticket if found
	for l <= r {
		mid := l + int((r-l)/2)
		if localTickets[mid].ID == target {
			return localTickets[mid]
		} else if localTickets[mid].ID > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// Return nil if the target ticket is not present
	return nil
}
