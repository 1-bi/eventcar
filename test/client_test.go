package test

import (
	"github.com/1-bi/eventcar"
	"github.com/1-bi/eventcar/test/fixture"
	"github.com/smartystreets/gunit"
	"testing"
)

// ---- setup method ---
var client *eventcar.ClientEndpoint

func TestClient(t *testing.T) {
	// define method ---
	gunit.Run(new(fixture.ClientFixture), t)
}
