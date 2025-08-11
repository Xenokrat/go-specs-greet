package interactions

import (
	"testing"

	"github.com/Xenokrat/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(Greet))
}

