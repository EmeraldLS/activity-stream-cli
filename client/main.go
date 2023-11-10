package main

import (
	"log"

	"github.com/EmeraldLS/daily_activity/client/activity"
	pb "github.com/EmeraldLS/daily_activity/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("An error occured while connecting to the gRPC server :: %v", err)
	}

	client := pb.NewActivityServiceClient(conn)

	rootCmd := &cobra.Command{
		Use:   "activity",
		Short: "Get listed of activities for a day",
		Long: `This program helps to get the list of activities for a day e.g
				Monday - It will returns the list of activties on monday
				Sunday - It will returns the list of activties on sunday
				and so on....
		`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			day := args[0]
			if err := activity.GetActivitiesByDay(client, day); err != nil {
				log.Fatal(err)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
