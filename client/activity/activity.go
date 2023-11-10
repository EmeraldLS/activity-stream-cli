package activity

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/EmeraldLS/daily_activity/proto"
)

func GetActivitiesByDay(client pb.ActivityServiceClient, day string) error {
	actReq := &pb.ActivityRequest{
		Day: day,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	stream, err := client.GetActivities(ctx, actReq)
	if err != nil {
		return fmt.Errorf("an error occured getting the activities from the server :: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("Streaming Completed")
			return nil
		}

		if err != nil {
			return err
		}

		act := res.GetActivity()
		supervisor := act.GetActivitySupervisor()

		log.Printf("Activity Name :: %v\n", act.GetActivityName())
		log.Printf("Activity Time :: %v\n", act.GetActivityTime())
		log.Printf("SuperVisor Name :: %v\n", supervisor.GetSupervisorName())
		log.Printf("SuperVisor Gender :: %v\n", supervisor.GetSupervisorGender())
		log.Printf("SuperVisor Age :: %v\n", supervisor.GetSupervisorAge())
		log.Println("----------------------")
		time.Sleep(time.Duration(time.Second * 2))
	}
}
