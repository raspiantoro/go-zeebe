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
	"github.com/raspiantoro/go-zeebe/commons/bootstrap/zeebe"
	"github.com/raspiantoro/go-zeebe/purchase/internal/command"
	"github.com/raspiantoro/go-zeebe/purchase/internal/commons"
	"github.com/raspiantoro/go-zeebe/purchase/internal/handler"
	"github.com/raspiantoro/go-zeebe/purchase/internal/handler/api"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		zeebeConfig := zeebe.Config{
			Gateway:  os.Getenv("ZEEBE_ADDRESS"),
			Insecure: true,
		}
		err := zeebe.InitClient(zeebeConfig)
		if err != nil {
			panic(err)
		}

		purchaseCommand, err := zeebe.NewCommand("purchasing-process")
		if err != nil {
			panic(err)
		}

		approvalCommand := zeebe.NewMessageCommand("purchase_approval_message")

		command := command.Command{
			Purchase: purchaseCommand,
			Approval: approvalCommand,
		}

		handlerOption := handler.Option{
			Option: commons.Option{
				Command: command,
			},
		}

		purchaseHandler := api.NewPurchaseHandler(handlerOption)
		approvalHandler := api.NewApprovalHandler(handlerOption)

		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.POST("/purchase", purchaseHandler.Submit)
		e.POST("/approval", approvalHandler.Submit)
		e.Logger.Fatal(e.Start(":3000"))

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
