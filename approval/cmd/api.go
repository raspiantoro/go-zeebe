/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/raspiantoro/go-zeebe/approval/internal/command"
	"github.com/raspiantoro/go-zeebe/approval/internal/commons"
	"github.com/raspiantoro/go-zeebe/approval/internal/handler"
	"github.com/raspiantoro/go-zeebe/approval/internal/handler/api"
	"github.com/raspiantoro/go-zeebe/approval/internal/repository"
	"github.com/raspiantoro/go-zeebe/approval/internal/service"
	"github.com/raspiantoro/go-zeebe/builder/database"
	"github.com/raspiantoro/go-zeebe/builder/zeebe"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		zeebeConfig := zeebe.Config{
			Gateway:  os.Getenv("ZEEBE_ADDRESS"),
			Insecure: true,
		}
		err := zeebe.InitClient(zeebeConfig)
		if err != nil {
			panic(err)
		}

		db, err := database.GetDB(os.Getenv("POSTGRES_URL"))
		if err != nil {
			panic(err)
		}

		approvalCommand := zeebe.NewMessageCommand("purchase_approval_message")

		command := command.Command{
			Approval: approvalCommand,
		}

		option := commons.Option{
			Command: command,
			DB:      db,
		}

		approvalRepository := repository.NewApprovalRepository(repository.Option{Option: option})

		approvalService := service.NewApprovalService(service.Option{
			Option: option,
			Repository: repository.Repository{
				Approval: approvalRepository,
			},
		})

		approvalHandler := api.NewApprovalHandler(handler.Option{
			Option: option,
			Service: service.Service{
				Approval: approvalService,
			},
		})

		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.POST("/approval", approvalHandler.Submit)
		e.Logger.Fatal(e.Start(":3001"))

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		done := make(chan struct{})

		go func() {
			<-sigs
			close(done)
		}()

		<-done
		fmt.Println("bye")
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
