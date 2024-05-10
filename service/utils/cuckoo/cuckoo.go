package cuckoo

import (
	"fmt"
	"os"
	"refugio/repository"

	cuckoo "github.com/panmari/cuckoofilter"
)

var DEFAULT_CUCKOO_CAPACITY uint = 250000

func createCuckooFilter(capacity uint) *cuckoo.Filter {
	filter := cuckoo.NewFilter(capacity)

	return filter
}

func GetCuckooFilter(key string) (*cuckoo.Filter, error) {
	filterBytes, err := repository.FetchFilterFromFirestore(key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching filter from Firestore: %v. Creating from scratch\n", err)
		return createCuckooFilter(DEFAULT_CUCKOO_CAPACITY), nil
	}

	filter, err := cuckoo.Decode(filterBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding filter: %v\n", err)
		return nil, err
	}

	return filter, nil
}
