// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package cli

import (
	"github.com/cockroachdb/cockroach/pkg/cli/systembench"
	"github.com/spf13/cobra"
)

// An seqWriteBench command runs I/O benchmarks on cockroach.
var seqWriteBench = &cobra.Command{
	Use:   "seqwrite",
	Short: "Runs the sequential disk write benchmark.",
	Long: `
Runs the sequential disk write benchmark.
`,
	Args: cobra.NoArgs,
	RunE: MaybeDecorateGRPCError(RunSeqWriteBench),
}

// RunSeqWriteBench runs a sequential write I/O benchmark.
func RunSeqWriteBench(cmd *cobra.Command, args []string) error {
	iOOpts := systembench.DiskOptions{
		Concurrency: systemBenchCtx.concurrency,
		Duration:    systemBenchCtx.duration,
		Dir:         systemBenchCtx.tempDir,

		Type:         systembench.SeqWriteTest,
		WriteSize:    systemBenchCtx.writeSize,
		SyncInterval: systemBenchCtx.syncInterval,
	}
	return systembench.Run(iOOpts)
}

var systemBenchCmds = []*cobra.Command{
	seqWriteBench,
}

var systemBenchCmd = &cobra.Command{
	Use:   "systembench [command]",
	Short: "Run systembench",
	Long: `
Run cockroach hardware benchmarks, for options use --help.`,
	RunE: usageAndErr,
}

func init() {
	systemBenchCmd.AddCommand(systemBenchCmds...)
}
