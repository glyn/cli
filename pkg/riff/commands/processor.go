/*
 * Copyright 2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"strings"

	"github.com/projectriff/riff/pkg/cli"
	"github.com/spf13/cobra"
)

func NewProcessorCommand(c *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "processor",
		Short: "processors apply functions to messages on streams",
		Long: strings.TrimSpace(`
<todo>
`),
		Args:    cli.Args(),
		Aliases: []string{"processors"},
	}

	cmd.AddCommand(NewProcessorListCommand(c))
	cmd.AddCommand(NewProcessorCreateCommand(c))
	cmd.AddCommand(NewProcessorDeleteCommand(c))
	cmd.AddCommand(NewProcessorTailCommand(c))

	return cmd
}
