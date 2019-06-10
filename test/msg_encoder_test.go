package test

import (
	"github.com/1-bi/eventcar/test/fixture"
	"github.com/smartystreets/gunit"
	"testing"
)

func TestProtobufEncoderFixture(t *testing.T) {
	// define method ---
	gunit.Run(new(fixture.ProtobufEncoderFixture), t)
}
