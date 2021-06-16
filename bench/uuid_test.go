package bench

// START OMIT

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ { // HL
		_ = uuid.Must(uuid.NewUUID())
	}
}

// END OMIT
