# nogo-buildssa

Running `bazel build -- //...` returns

```
INFO: Analyzed 3 targets (0 packages loaded, 0 targets configured).
INFO: Found 3 targets...
ERROR: <redacted>/nogo-buildssa/BUILD.bazel:10:11: GoCompilePkg types.a failed: (Exit 1): builder failed: error executing command bazel-out/darwin_arm64-opt-exec-2B5CBBC6/bin/external/go_sdk/builder compilepkg -sdk external/go_sdk -installsuffix darwin_arm64 -src types.go -embedroot '' -embedroot ... (remaining 21 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
compilepkg: panic: unexpected expression-could not match index list to instantiation

goroutine 18 [running]:
golang.org/x/tools/go/ssa.(*builder).expr0(0x140001339f0, 0x140001ddba0, {0x10072cb48?, 0x1400012bb00?}, {0x7, {0x10072c328, 0x140001cce40}, {0x0, 0x0}})
	external/org_golang_x_tools/go/ssa/builder.go:845 +0x2250
golang.org/x/tools/go/ssa.(*builder).expr(0x14000132978?, 0x140001ddba0, {0x10072cb48?, 0x1400012bb00?})
	external/org_golang_x_tools/go/ssa/builder.go:610 +0x134
golang.org/x/tools/go/ssa.(*builder).setCallFunc(0x100a60a68?, 0x80?, 0x100873c00?, 0x1400014a8c0)
	external/org_golang_x_tools/go/ssa/builder.go:1006 +0x2f4
golang.org/x/tools/go/ssa.(*builder).setCall(0x1006ed6c0?, 0x14000113620?, 0x1400012bbc0, 0x1400014a8c0)
	external/org_golang_x_tools/go/ssa/builder.go:1085 +0x30
golang.org/x/tools/go/ssa.(*builder).expr0(0x140001339f0, 0x140001ddba0, {0x10072c758?, 0x1400012bbc0?}, {0x7, {0x10072c210, 0x10086a500}, {0x0, 0x0}})
	external/org_golang_x_tools/go/ssa/builder.go:686 +0x2354
golang.org/x/tools/go/ssa.(*builder).expr(0x14000132e38?, 0x140001ddba0, {0x10072c758?, 0x1400012bbc0?})
	external/org_golang_x_tools/go/ssa/builder.go:610 +0x134
golang.org/x/tools/go/ssa.(*builder).emitCallArgs(0x140001ddba0?, 0x10072ccf8?, 0x140001cc1c0, 0x1400012bc00, {0x0?, 0x0, 0x14000132f78?})
	external/org_golang_x_tools/go/ssa/builder.go:1030 +0x734
golang.org/x/tools/go/ssa.(*builder).setCall(0x1006ed6c0?, 0x14000113620?, 0x1400012bc00, 0x1400014a840)
	external/org_golang_x_tools/go/ssa/builder.go:1092 +0x8c
golang.org/x/tools/go/ssa.(*builder).expr0(0x140001339f0, 0x140001ddba0, {0x10072c758?, 0x1400012bc00?}, {0x7, {0x10072c3a0, 0x1400011ea20}, {0x0, 0x0}})
	external/org_golang_x_tools/go/ssa/builder.go:686 +0x2354
golang.org/x/tools/go/ssa.(*builder).expr(0x100873c00?, 0x140001ddba0, {0x10072c758?, 0x1400012bc00?})
	external/org_golang_x_tools/go/ssa/builder.go:610 +0x134
golang.org/x/tools/go/ssa.(*builder).stmt(0x14000133598?, 0x140001ddba0, {0x10072c908?, 0x140001115c0?})
	external/org_golang_x_tools/go/ssa/builder.go:2101 +0x8e8
golang.org/x/tools/go/ssa.(*builder).stmtList(0x0?, 0x0?, {0x140001115d0?, 0x1, 0x14000133608?})
	external/org_golang_x_tools/go/ssa/builder.go:911 +0x6c
golang.org/x/tools/go/ssa.(*builder).stmt(0x140001ddba0?, 0x140001ddba0, {0x10072c6f8?, 0x140001134d0?})
	external/org_golang_x_tools/go/ssa/builder.go:2218 +0x8b8
golang.org/x/tools/go/ssa.(*builder).buildFunctionBody(0x140001dd860?, 0x140001ddba0)
	external/org_golang_x_tools/go/ssa/builder.go:2327 +0x3a4
golang.org/x/tools/go/ssa.(*builder).buildFunction(0x100651570?, 0x140001ddba0)
	external/org_golang_x_tools/go/ssa/builder.go:2267 +0x34
golang.org/x/tools/go/ssa.(*builder).buildCreated(0x140001859f0)
	external/org_golang_x_tools/go/ssa/builder.go:2349 +0x2c
golang.org/x/tools/go/ssa.(*Package).build(0x1400014a680)
	external/org_golang_x_tools/go/ssa/builder.go:2529 +0xab8
sync.(*Once).doSlow(0x140001c00c0?, 0x1400010d680?)
	GOROOT/src/sync/once.go:68 +0x10c
sync.(*Once).Do(...)
	GOROOT/src/sync/once.go:59
golang.org/x/tools/go/ssa.(*Package).Build(...)
	external/org_golang_x_tools/go/ssa/builder.go:2413
golang.org/x/tools/go/analysis/passes/buildssa.run(0x140001816c0)
	external/org_golang_x_tools/go/analysis/passes/buildssa/buildssa.go:73 +0x13c
main.(*action).execOnce(0x14000156510)
	external/io_bazel_rules_go/go/tools/builders/nogo_main.go:301 +0x764
sync.(*Once).doSlow(0x0?, 0x0?)
	GOROOT/src/sync/once.go:68 +0x10c
sync.(*Once).Do(...)
	GOROOT/src/sync/once.go:59
main.(*action).exec(0x0?)
	external/io_bazel_rules_go/go/tools/builders/nogo_main.go:245 +0x44
main.execAll.func1(0x0?)
	external/io_bazel_rules_go/go/tools/builders/nogo_main.go:239 +0x50
created by main.execAll
	external/io_bazel_rules_go/go/tools/builders/nogo_main.go:237 +0x48
INFO: Elapsed time: 0.273s, Critical Path: 0.11s
INFO: 2 processes: 2 internal.
FAILED: Build did NOT complete successfully
```

when it should succeed.
