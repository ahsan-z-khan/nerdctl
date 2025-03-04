/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/containerd/nerdctl/pkg/cmd/namespace"
	"github.com/spf13/cobra"
)

func newNamespaceCreateCommand() *cobra.Command {
	namespaceCreateCommand := &cobra.Command{
		Use:           "create NAMESPACE",
		Short:         "Create a new namespace",
		Args:          cobra.MinimumNArgs(1),
		RunE:          namespaceCreateAction,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	namespaceCreateCommand.Flags().StringArrayP("label", "l", nil, "Set labels for a namespace")
	return namespaceCreateCommand
}

func processNamespaceCreateCommandOption(cmd *cobra.Command) (types.NamespaceCreateOptions, error) {
	globalOptions, err := processRootCmdFlags(cmd)
	if err != nil {
		return types.NamespaceCreateOptions{}, err
	}
	labels, err := cmd.Flags().GetStringArray("label")
	if err != nil {
		return types.NamespaceCreateOptions{}, err
	}
	return types.NamespaceCreateOptions{
		GOptions: globalOptions,
		Labels:   labels,
	}, nil
}

func namespaceCreateAction(cmd *cobra.Command, args []string) error {
	options, err := processNamespaceCreateCommandOption(cmd)
	if err != nil {
		return err
	}
	return namespace.Create(cmd.Context(), args[0], options)
}
