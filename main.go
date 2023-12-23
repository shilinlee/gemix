// Copyright 2023 Huawei Cloud Computing Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"time"

	"github.com/openGemini/gemix/cmd"
)

var (
	Version   string
	Commit    string
	Branch    string
	BuildTime string
)

func init() {
	// If commit, branch, or build time are not set, make that clear.
	if Version == "" {
		Version = "unknown"
	}
	if Commit == "" {
		Commit = "unknown"
	}
	if Branch == "" {
		Branch = "unknown"
	}

	if BuildTime == "" {
		BuildTime = time.Now().UTC().String()
	}
}

func main() {
	cmd.Version = Version
	cmd.Commit = Commit
	cmd.Branch = Branch
	cmd.BuildTime = BuildTime
	cmd.Execute()
}
