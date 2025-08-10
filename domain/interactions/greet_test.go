package go_specs_greet_test

import (
	"testing"

	go_spec_greet "github.com/Xenokrat/go-specs-greet/domain/interactions"
	"github.com/Xenokrat/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(go_spec_greet.Greet))
}

