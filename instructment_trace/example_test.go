package trace_test

import trace "github.com/Howie27Chen/instrument_trace"

func a() {
	defer trace.Trace()()
	b()
}

func b() {
	defer trace.Trace()()
	c()
}

func c() {
	defer trace.Trace()()
	d()
}

func d() {
	defer trace.Trace()()
}

func ExampleTrace() {
	a()
	// Output:
	// g[000001]: ->github.com/Howie27Chen/instrument_trace_test.a
}
