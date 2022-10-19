//go:build windows
// +build windows

package envoy

import (
	"fmt"
	"os"
	"os/exec"

	"path/filepath"
	"strings"

	"time"
)

// testSelfExecOverride is a way for the tests to no fork-bomb themselves by
// self-executing the whole test suite for each case recursively. It's gross but
// the least gross option I could think of.
var testSelfExecOverride string

func isHotRestartOption(s string) bool {
	restartOpts := []string{
		"--restart-epoch",
		"--hot-restart-version",
		"--drain-time-s",
		"--parent-shutdown-time-s",
	}
	for _, opt := range restartOpts {
		if s == opt {
			return true
		}
		if strings.HasPrefix(s, opt+"=") {
			return true
		}
	}
	return false
}

func hasHotRestartOption(argSets ...[]string) bool {
	for _, args := range argSets {
		for _, opt := range args {
			if isHotRestartOption(opt) {
				return true
			}
		}
	}
	return false
}

func makeBootstrapTemp(bootstrapJSON []byte) (string, error) {
	tempFile := filepath.Join(os.TempDir(),
		fmt.Sprintf("envoy-%x-bootstrap.json", time.Now().UnixNano()+int64(os.Getpid())))

	f, err := os.Create(tempFile)
	if err != nil {
		return tempFile, err
	}

	defer f.Close()
	f.Write(bootstrapJSON)
	f.Sync()

	// We can't wait for the process since we need to exec into Envoy before it
	// will be able to complete so it will be remain as a zombie until Envoy is
	// killed then will be reaped by the init process (pid 0). This is all a bit
	// gross but the cleanest workaround I can think of for Envoy 1.10 not
	// supporting /dev/fd/<fd> config paths any more. So we are done and leaving
	// the child to run it's course without reaping it.
	return tempFile, nil
}

func startProc(binary string, args []string) (p *os.Process, err error) {
	if binary, err = exec.LookPath(binary); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(binary, args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

func execEnvoy(binary string, prefixArgs, suffixArgs []string, bootstrapJSON []byte) error {
	tempFile, err := makeBootstrapTemp(bootstrapJSON)
	if err != nil {
		os.RemoveAll(tempFile)
		return err
	}
	// We don't defer a cleanup since we are about to Exec into Envoy which means
	// defer will never fire. The child process cleans up for us in the happy
	// path.

	// We default to disabling hot restart because it makes it easier to run
	// multiple envoys locally for testing without them trying to share memory and
	// unix sockets and complain about being different IDs. But if user is
	// actually configuring hot-restart explicitly with the --restart-epoch option
	// then don't disable it!
	disableHotRestart := !hasHotRestartOption(prefixArgs, suffixArgs)

	// First argument needs to be the executable name.
	envoyArgs := []string{}
	envoyArgs = append(envoyArgs, prefixArgs...)
	if disableHotRestart {
		envoyArgs = append(envoyArgs, "--disable-hot-restart")
	}
	envoyArgs = append(envoyArgs, suffixArgs...)
	envoyArgs = append(envoyArgs, "--config-path", tempFile)

	// Exec
	if proc, err := startProc(binary, envoyArgs); err == nil {
		proc.Wait()
	}

	return nil
}
