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
	"context"
	"fmt"
	"strings"

	"github.com/projectriff/cli/pkg/cli"
	"github.com/projectriff/cli/pkg/cli/options"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KafkaProviderDeleteOptions struct {
	options.DeleteOptions
}

var (
	_ cli.Validatable = (*KafkaProviderDeleteOptions)(nil)
	_ cli.Executable  = (*KafkaProviderDeleteOptions)(nil)
)

func (opts *KafkaProviderDeleteOptions) Validate(ctx context.Context) cli.FieldErrors {
	errs := cli.FieldErrors{}

	errs = errs.Also(opts.DeleteOptions.Validate(ctx))

	return errs
}

func (opts *KafkaProviderDeleteOptions) Exec(ctx context.Context, c *cli.Config) error {
	client := c.StreamingRuntime().KafkaProviders(opts.Namespace)

	if opts.All {
		if err := client.DeleteCollection(nil, metav1.ListOptions{}); err != nil {
			return err
		}
		c.Successf("Deleted kafka providers in namespace %q\n", opts.Namespace)
		return nil
	}

	for _, name := range opts.Names {
		if err := client.Delete(name, nil); err != nil {
			return err
		}
		c.Successf("Deleted kafka provider %q\n", name)
	}

	return nil
}

func NewKafkaProviderDeleteCommand(ctx context.Context, c *cli.Config) *cobra.Command {
	opts := &KafkaProviderDeleteOptions{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete kafka provider(s)",
		Long: strings.TrimSpace(`
Delete one or more kafka providers by name or all kafka providers within a
namespace.

Deleting a kafka provider will disrupt all processors consuming streams managed
by the provider. Existing messages in the stream may be preserved by the
underlying kafka broker, depending on the implementation.
`),
		Example: strings.Join([]string{
			fmt.Sprintf("%s streaming kafka-provider delete my-kafka-provider", c.Name),
			fmt.Sprintf("%s streaming kafka-provider delete %s ", c.Name, cli.AllFlagName),
		}, "\n"),
		PreRunE: cli.ValidateOptions(ctx, opts),
		RunE:    cli.ExecOptions(ctx, c, opts),
	}

	cli.Args(cmd,
		cli.NamesArg(&opts.Names),
	)

	cli.NamespaceFlag(cmd, c, &opts.Namespace)
	cmd.Flags().BoolVar(&opts.All, cli.StripDash(cli.AllFlagName), false, "delete all kafka providers within the namespace")

	return cmd
}
