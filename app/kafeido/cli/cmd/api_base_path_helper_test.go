package cmd

import "os"

// osUnsetenv exists so the tests can express "this machine has no override"
// distinctly from "the override is empty", which resolveApiBasePath treats as
// two different answers.
func osUnsetenv(k string) error { return os.Unsetenv(k) }
