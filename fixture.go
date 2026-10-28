package fixture

import "github.com/google/uuid"

// NewID gives tests a real use of the intentionally outdated module.
func NewID() string { return uuid.NewString() }
