// Copyright 2020 The Atlas Authors. All rights reserved.
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

package cmd

import (
	"net"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	cmd "github.com/spf13/cobra"
	conf "github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/binarly-io/atlas/pkg/controller"
	"github.com/binarly-io/atlas/pkg/app"
	"github.com/binarly-io/atlas/pkg/api/atlas"
	atlas_health "github.com/binarly-io/atlas/pkg/api/health"
	"github.com/binarly-io/atlas/pkg/controller/service"
	"github.com/binarly-io/atlas/pkg/transport/service"

	"github.com/binarly-io/atlas-server/pkg/softwareid"
)

var (
	// readFilename specifies file to read and send to service
	sendFilename string

	// readSTDIN specifies whether to read STDIN
	sendSTDIN bool
)

var serveCmd = &cmd.Command{
	Use:   "serve [OPTION] [FILE]",
	Short: "Serve service",
	Long: heredoc.Docf(`
			Serve
			`,
	),
	Args: func(cmd *cmd.Command, args []string) error {
		//if len(args) < 1 {
		//	return errors.New("requires an filename as argument")
		//}
		return nil
	},
	Run: func(cmd *cmd.Command, args []string) {
		//filename := args[0]

		// Init termination context
		ctx := app.ContextInit()

		log.Infof("Starting service. Version:%s GitSHA:%s BuiltAt:%s\n", softwareid.Version, softwareid.GitSHA, softwareid.BuiltAt)

		serverInit()

		log.Infof("Listening on %s", serviceAddress)
		listener, err := net.Listen("tcp", serviceAddress)
		if err != nil {
			log.Fatalf("failed to Listen() %v", err)
		}

		// Create gRPC server
		grpcServer := grpc.NewServer(service_transport.GetGRPCServerOptions(tls, auth, tlsCertFile, tlsKeyFile, jwtPublicKeyFile)...)
		serverRegister(grpcServer)

		// Serve gRPC requests
		go func() {
			if err := grpcServer.Serve(listener); err != nil {
				log.Fatalf("failed to Serve() %v", err)
			}
		}()

		serverRun()

		app.ContextWait(ctx)
	},
}

func init() {
	// Bind full flag set to the configuration
	if err := conf.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(serveCmd)
}

// serverInit initializes all server internals
func serverInit() {
	controller.Init()
}

// serverRegister registers Atlas servers with gRPC server
func serverRegister(grpcServer *grpc.Server) {
	atlas.RegisterControlPlaneServer(grpcServer, controller_service.NewControlPlaneServer())
	atlas_health.RegisterHealthServer(grpcServer, controller_service.NewHealthServer())
}

// serverRun runs additional - non-rRPC-functions
func serverRun() {
	go controller_service.IncomingCommandsHandler(controller_service.GetIncomingQueue(), controller_service.GetOutgoingQueue())
}
