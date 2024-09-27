/**
# Copyright 2024 NVIDIA CORPORATION
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package runtime

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/NVIDIA/nvidia-container-toolkit/tools/container"
)

const (
	defaultSetAsDefault = true
	// defaultRuntimeName specifies the NVIDIA runtime to be use as the default runtime if setting the default runtime is enabled
	defaultRuntimeName   = "nvidia"
	defaultHostRootMount = "/host"

	runtimeSpecificDefault = "RUNTIME_SPECIFIC_DEFAULT"
)

type Options struct {
	container.Options
}

func Flags(opts *Options) []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Usage:       "Path to docker config file",
			Value:       runtimeSpecificDefault,
			Destination: &opts.Config,
			EnvVars:     []string{"RUNTIME_CONFIG"},
		},
		&cli.StringFlag{
			Name:        "socket",
			Usage:       "Path to the docker socket file",
			Value:       runtimeSpecificDefault,
			Destination: &opts.Socket,
			EnvVars:     []string{"RUNTIME_SOCKET"},
		},
		&cli.StringFlag{
			Name:        "restart-mode",
			Usage:       "Specify how docker should be restarted; If 'none' is selected it will not be restarted [signal | systemd | none ]",
			Value:       runtimeSpecificDefault,
			Destination: &opts.RestartMode,
			EnvVars:     []string{"RUNTIME_RESTART_MODE"},
		},
		&cli.StringFlag{
			Name:        "host-root",
			Usage:       "Specify the path to the host root to be used when restarting docker using systemd",
			Value:       defaultHostRootMount,
			Destination: &opts.HostRootMount,
			EnvVars:     []string{"HOST_ROOT_MOUNT"},
		},
		&cli.StringFlag{
			Name:        "runtime-name",
			Aliases:     []string{"nvidia-runtime-name", "runtime-class"},
			Usage:       "Specify the name of the `nvidia` runtime. If set-as-default is selected, the runtime is used as the default runtime.",
			Value:       defaultRuntimeName,
			Destination: &opts.RuntimeName,
			EnvVars:     []string{"NVIDIA_RUNTIME_NAME"},
		},
		&cli.BoolFlag{
			Name:        "set-as-default",
			Usage:       "Set the `nvidia` runtime as the default runtime.",
			Value:       defaultSetAsDefault,
			Destination: &opts.SetAsDefault,
			EnvVars:     []string{"NVIDIA_RUNTIME_SET_AS_DEFAULT"},
			Hidden:      true,
		},
	}

	return flags
}

// ValidateOptions checks whether the specified options are valid
func ValidateOptions(opts *Options, toolkitRoot string) error {
	// We set this option here to ensure that it is avalable in future calls.
	opts.RuntimeDir = toolkitRoot
	return nil
}

func Setup(c *cli.Context, opts *Options, runtime string) error {
	switch runtime {
	}
	return fmt.Errorf("undefined runtime %v", runtime)
}

func Cleanup(c *cli.Context, opts *Options, runtime string) error {
	switch runtime {
	}
	return fmt.Errorf("undefined runtime %v", runtime)
}
