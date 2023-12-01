// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package scripts

import (
	"bytes"
	"os"
	"path"
	"text/template"

	"github.com/openGemini/gemix/embed"
	"github.com/pkg/errors"
)

// TSStoreScript represent the data to generate ts-store config
type TSStoreScript struct {
	DeployDir string
	LogDir    string
}

// ConfigToFile write config content to specific path
func (c *TSStoreScript) ConfigToFile(file string) error {
	fp := path.Join("templates", "scripts", "run_ts_store.sh.tpl")
	tpl, err := embed.ReadTemplate(fp)
	if err != nil {
		return errors.WithStack(err)
	}

	tmpl, err := template.New("TSStore").Parse(string(tpl))
	if err != nil {
		return errors.WithStack(err)
	}

	content := bytes.NewBufferString("")
	if err := tmpl.Execute(content, c); err != nil {
		return errors.WithStack(err)
	}

	return os.WriteFile(file, content.Bytes(), 0750)
}
