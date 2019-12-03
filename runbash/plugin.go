package runbash

import (
	"errors"
	"os/exec"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "plugins.flutter.io/runbash"

// RunBashPlugin implements flutter.Plugin and handles method.
type RunBashPlugin struct{}

var _ flutter.Plugin = &RunBashPlugin{}

// InitPlugin ...
func (p *RunBashPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("run_bash", p.run)
	return nil
}

func (p *RunBashPlugin) run(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	runPath := argsMap["run_path"].(string)
	runCmd := argsMap["run_cmd"].(string)

	if runPath == "" {
		return nil, errors.New("run path is empty")
	}

	if runCmd == "" {
		return nil, errors.New("run cmd is empty")
	}

	cmd := exec.Command("bash", "-c", runCmd)
	cmd.Dir = runPath

	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return string(out), nil
}
