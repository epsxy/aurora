package global

var (
	verbose = false
	dryRun  = false
)

func SetVerbose(v bool) {
	verbose = v
}

func GetVerbose() bool {
	return verbose
}

func SetDryRun(d bool) {
	dryRun = d
}

func GetDryRun() bool {
	return dryRun
}
