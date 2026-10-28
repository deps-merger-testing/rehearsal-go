package fixture

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewID(t *testing.T) {
	_, err := uuid.Parse(NewID())
	require.NoError(t, err)
}
