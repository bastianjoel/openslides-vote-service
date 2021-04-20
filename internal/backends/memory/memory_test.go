package memory_test

import (
	"testing"

	"github.com/OpenSlides/openslides-vote-service/internal/backends/memory"
	"github.com/OpenSlides/openslides-vote-service/internal/backends/test"
)

func TestConfig(t *testing.T) {
	m := memory.New()

	test.Config(t, m)
}

func TestBakend(t *testing.T) {
	m := memory.New()

	test.Backend(t, m)
}
