package gonagioschecks

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestNagios(t *testing.T) {

	Convey("When a Nagios is created it looks correct", t, func() {
		n := Nagios{}

		So(n.Status(), ShouldEqual, OK)
		So(n.Message, ShouldBeBlank)

		Convey("... and the EscalateIf logic checks out", func() {
			n.Code = UNKNOWN
			n.EscalateIf(OK)
			So(n.Status(), ShouldEqual, OK)

			n.EscalateIf(WARNING)
			So(n.Status(), ShouldEqual, WARNING)
			n.EscalateIf(OK)
			So(n.Status(), ShouldEqual, WARNING) // regression?
			n.EscalateIf(CRITICAL)
			So(n.Status(), ShouldEqual, CRITICAL)
			n.EscalateIf(WARNING)
			So(n.Status(), ShouldEqual, CRITICAL) // regression?
			n.EscalateIf(OK)
			So(n.Status(), ShouldEqual, CRITICAL) // regression?
			n.EscalateIf(212)
			So(n.Status(), ShouldEqual, CRITICAL) // regression?
		})

		Convey("... and the AddMessage[If/etc] logic checks out", func() {
			So(n.Message, ShouldBeBlank)
			n.AddMessage("Hello World")
			So(n.Message, ShouldEqual, "Hello World")
			n.AddMessageIfBool("\nIt's Me!!\t", OK == 0)
			So(n.Message, ShouldEqual, "Hello World It's Me!! ")
			n.AddMessageIf("Really\nit\tis!", "haha")
			So(n.Message, ShouldEqual, "Hello World It's Me!! Really it is!")
			n.AddMessageIfBool("It's Me!!\t", OK == CRITICAL)
			So(n.Message, ShouldEqual, "Hello World It's Me!! Really it is!") // regression?
			n.PrependMessage("Oi!\n")
			So(n.Message, ShouldEqual, "Oi! Hello World It's Me!! Really it is!")
			So(n.FullMessage(), ShouldEqual, n.Message)
		})

		Convey("... and the AddMetrics stuff works fine too", func() {
			So(n.Metrics, ShouldBeEmpty)
			n.AddMetricNumbers("HR", 67, 100, 140, nil, nil)
			So(n.Metrics, ShouldContain, "'HR'=67;100;140;;")
			n.AddMetrics("'STRESS'=31;50;75;0;100")
			So(n.Metrics, ShouldContain, "'HR'=67;100;140;;")
			So(n.Metrics, ShouldContain, "'STRESS'=31;50;75;0;100")
			So(n.FullMessage(), ShouldEqual, n.Message+"| "+"'HR'=67;100;140;;"+" "+"'STRESS'=31;50;75;0;100")
		})
	})
}
