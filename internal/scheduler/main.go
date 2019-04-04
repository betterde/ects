package scheduler

var (
	service *Cron
)

func Initialize() {
	service = New()
	service.Start()
}