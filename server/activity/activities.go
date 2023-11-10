package activity

import (
	pb "github.com/EmeraldLS/daily_activity/proto"
	"github.com/golang-module/carbon"
)

type Activity struct {
	activityName       string
	activityTime       string
	activitySupervisor ActivitySupervisor
}

type ActivitySupervisor struct {
	supervisorName   string
	supervisorAge    int32
	supervisorGender pb.ActivitySupervisor_Gender
}

var Activities = map[string][]Activity{
	"Monday": {
		{
			activityName: "Wash Plate",
			activityTime: carbon.Now().AddSeconds(5).ToDateTimeString(),
			activitySupervisor: ActivitySupervisor{
				supervisorName:   "Lawrence Segun",
				supervisorAge:    20,
				supervisorGender: pb.ActivitySupervisor_Male,
			},
		},
		{
			activityName: "Wash Toilet",
			activityTime: carbon.Now().AddHours(2).ToDateTimeString(),
			activitySupervisor: ActivitySupervisor{
				supervisorName:   "Sanni Abdullah",
				supervisorAge:    20,
				supervisorGender: pb.ActivitySupervisor_Male,
			},
		},
	},
	"Tuesday": {
		{
			activityName: "Fetch Water",
			activityTime: carbon.Now().AddSeconds(30).ToDateTimeString(),
			activitySupervisor: ActivitySupervisor{
				supervisorName:   "Sanni Motunrayo Eniola",
				supervisorAge:    19,
				supervisorGender: pb.ActivitySupervisor_Female,
			},
		},
	},
	"Wednesday": {
		{
			activityName: "Break Up",
			activityTime: carbon.Now().ToDateTimeString(),
			activitySupervisor: ActivitySupervisor{
				supervisorName:   "Adeboye Boluwatife Dorcas",
				supervisorAge:    21,
				supervisorGender: pb.ActivitySupervisor_Female,
			},
		},
	},
	"Thursday": {},
	"Friday":   {},
	"Saturday": {},
	"Sunday":   {},
}
