package flock

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
)

// PluginHandler ...
func PluginHandler(content []byte) (map[string]interface{}, error) {
	tmp := os.TempDir()

	gof := filepath.Join(tmp, "tmp.go")
	if err := ioutil.WriteFile(gof, content, 0666); err != nil {
		return nil, err
	}

	sof := filepath.Join(tmp, "tmp.so")
	// Compile the file as a plugin
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", sof, gof)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	p, err := plugin.Open(sof)
	if err != nil {
		return nil, err
	}

	fmSym, err := p.Lookup("FM")
	if err != nil {
		return nil, err
	}

	fmt.Println(fmSym)

	fmPlain, ok := fmSym.(*map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to assert interface as *map[string]interface{}, type: %T", fmSym)
	}

	fm := FuncMap(*fmPlain)

	return fm, nil
}
