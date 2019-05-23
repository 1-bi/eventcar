package test

import (
	"github.com/1-bi/eventcar"
	"github.com/1-bi/eventcar/test/fixture"
	"github.com/smartystreets/gunit"
	"testing"
)

// ---- setup method ---
var agent *eventcar.Agent

func TestAgent(t *testing.T) {
	// define method ---
	gunit.Run(new(fixture.AgentFixture), t)
}
