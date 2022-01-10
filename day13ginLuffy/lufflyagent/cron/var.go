package cron

import "lufflyagent/settings"


func GetTasksListUrl() string {
	return "/cron/jobs/" + settings.IP()
}

func GetTasksCronIDUrl() string {
	return "/cron/cronid/jobs/" + settings.IP() + "/"
}

func GetTasksIDUrl() string {
	return GetTasksListUrl() + "/"
}
