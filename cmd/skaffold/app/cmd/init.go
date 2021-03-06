/*
Copyright 2019 The Skaffold Authors

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

package cmd

import (
	"context"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/initializer"
)

var (
	composeFile            string
	cliArtifacts           []string
	cliKubernetesManifests []string
	skipBuild              bool
	skipDeploy             bool
	force                  bool
	analyze                bool
	enableJibInit          bool
	enableBuildpacksInit   bool
	buildpacksBuilder      string
)

// for testing
var initEntrypoint = initializer.DoInit

// NewCmdInit describes the CLI command to generate a Skaffold configuration.
func NewCmdInit() *cobra.Command {
	return NewCmd("init").
		WithDescription("[alpha] Generate configuration for deploying an application").
		WithFlags(func(f *pflag.FlagSet) {
			f.StringVarP(&opts.ConfigurationFile, "filename", "f", "skaffold.yaml", "Filename or URL to the pipeline file")
			f.BoolVar(&skipBuild, "skip-build", false, "Skip generating build artifacts in Skaffold config")
			f.BoolVar(&skipDeploy, "skip-deploy", false, "Skip generating deploy stanza in Skaffold config")
			f.MarkHidden("skip-deploy")
			f.BoolVar(&force, "force", false, "Force the generation of the Skaffold config")
			f.StringVar(&composeFile, "compose-file", "", "Initialize from a docker-compose file")
			f.StringArrayVarP(&cliArtifacts, "artifact", "a", nil, "'='-delimited Dockerfile/image pair, or JSON string, to generate build artifact\n(example: --artifact='{\"builder\":\"Docker\",\"payload\":{\"path\":\"/web/Dockerfile.web\"},\"image\":\"gcr.io/web-project/image\"}')")
			f.StringArrayVarP(&cliKubernetesManifests, "kubernetes-manifest", "k", nil, "a path or a glob pattern to kubernetes manifests (can be non-existent) to be added to the kubectl deployer (overrides detection of kubernetes manifests). Repeat the flag for multiple entries. E.g.: skaffold init -k pod.yaml -k k8s/*.yml")
			f.BoolVar(&analyze, "analyze", false, "Print all discoverable Dockerfiles and images in JSON format to stdout")
			f.BoolVar(&enableJibInit, "XXenableJibInit", false, "")
			f.MarkHidden("XXenableJibInit")
			f.BoolVar(&enableBuildpacksInit, "XXenableBuildpacksInit", false, "")
			f.MarkHidden("XXenableBuildpacksInit")
			f.StringVar(&buildpacksBuilder, "XXdefaultBuildpacksBuilder", "heroku/buildpacks", "")
			f.MarkHidden("XXdefaultBuildpacksBuilder")
		}).
		NoArgs(cancelWithCtrlC(context.Background(), doInit))
}

func doInit(ctx context.Context, out io.Writer) error {
	return initEntrypoint(ctx, out, initializer.Config{
		ComposeFile:            composeFile,
		CliArtifacts:           cliArtifacts,
		CliKubernetesManifests: cliKubernetesManifests,
		SkipBuild:              skipBuild,
		SkipDeploy:             skipDeploy,
		Force:                  force,
		Analyze:                analyze,
		EnableJibInit:          enableJibInit,
		EnableBuildpacksInit:   enableBuildpacksInit,
		BuildpacksBuilder:      buildpacksBuilder,
		Opts:                   opts,
	})
}
