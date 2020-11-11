package definitions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var collectionRegex = regexp.MustCompile("[a-z]+ | [a-z][a-z_]*[a-z]")

func IsValidId(id Id) error {
	if id <= 0 {
		return fmt.Errorf("id must be positive")

	}
	return nil
}

func IsValidFqid(fqid Fqid) error {
	parts := strings.Split(fqid, keyseparator)
	if len(parts) != 2 {
		return fmt.Errorf("'%s' is not a valid fqid", fqid)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("The id in '%s' is not an int", fqid)
	}

	if err = IsValidId(id); err != nil {
		return fmt.Errorf("invalid fqid: %w", err)
	}
	if err = IsValidCollection(parts[0]); err != nil {
		return fmt.Errorf("invalid collection: %w", err)
	}

	return nil
}

func IsValidCollection(collection Collection) error {
	if !collectionRegex.MatchString(collection) {
		return fmt.Errorf("The collection '%s' is invalid", collection)
	}
	return nil
}
