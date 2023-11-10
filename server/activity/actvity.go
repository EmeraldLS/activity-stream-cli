package activity

import (
	"fmt"
	"log"

	pb "github.com/EmeraldLS/daily_activity/proto"
)

type ActivityServer struct {
	pb.ActivityServiceServer
}

func (a *ActivityServer) GetActivities(req *pb.ActivityRequest, stream pb.ActivityService_GetActivitiesServer) error {
	day := req.GetDay()
	activities := Activities[day]

	fmt.Printf("Incoming activity request for :: %v\n", day)

	for _, activity := range activities {
		actResp := pb.ActivityResponse{
			Activity: &pb.Activity{
				ActivityName: activity.activityName,
				ActivityTime: activity.activityTime,
				ActivitySupervisor: &pb.ActivitySupervisor{
					SupervisorName:   activity.activitySupervisor.supervisorName,
					SupervisorAge:    activity.activitySupervisor.supervisorAge,
					SupervisorGender: activity.activitySupervisor.supervisorGender,
				},
			},
		}
		if err := stream.Send(&actResp); err != nil {
			log.Fatalf("An error occured while sending activities stream :: %v", err)
		}
	}

	return nil
}
