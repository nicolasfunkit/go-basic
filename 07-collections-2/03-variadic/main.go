package main

func DebugLog(args ...string) []string {
	return updateLog("[DEBUG]", args)
}

func InfoLog(args ...string) []string {
	return updateLog("[INFO]", args)
}

func ErrorLog(args ...string) []string {
	return updateLog("[ERROR]", args)
}

func updateLog(prefix string, args []string) []string {
	return append([]string{prefix}, args...)
}
