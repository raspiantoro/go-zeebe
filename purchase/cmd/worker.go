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

	"github.com/raspiantoro/go-zeebe/commons/bootstrap/database"
	"github.com/raspiantoro/go-zeebe/commons/bootstrap/zeebe"
	"github.com/raspiantoro/go-zeebe/purchase/internal/commons"
	"github.com/raspiantoro/go-zeebe/purchase/internal/handler"
	"github.com/raspiantoro/go-zeebe/purchase/internal/handler/workflow"
	"github.com/raspiantoro/go-zeebe/purchase/internal/repository"
	"github.com/raspiantoro/go-zeebe/purchase/internal/service"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use: "worker",
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

		option := commons.Option{
			DB: db,
		}

		purchaseRepository := repository.NewPurchaseRepository(repository.Option{option})

		purchaseService := service.NewPurchaseService(service.Option{
			Option: option,
			Repository: repository.Repository{
				Purchase: purchaseRepository,
			},
		})

		purchaseHandler := workflow.NewPurchaseHandler(handler.Option{
			Option: option,
			Service: service.Service{
				Purchase: purchaseService,
			},
		})

		jobWorker := zeebe.NewJobWorker("prepare-purchase", purchaseHandler.Prepare)

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		done := make(chan struct{})

		go func() {
			<-sigs
			close(done)
		}()

		<-done
		fmt.Println("exiting worker")
		jobWorker.Close()
		fmt.Println("bye")
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
