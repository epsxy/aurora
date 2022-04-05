package global

var (
	verbose = false
)

func SetVerbose(v bool) {
	verbose = v
}

func GetVerbose() bool {
	return verbose
}
