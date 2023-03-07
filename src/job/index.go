package job

import (
	"log"

	"github.com/robfig/cron"

	config "github.com/lokesh-go/youtube-data-golang/src/config"
	dal "github.com/lokesh-go/youtube-data-golang/src/dal"
	ytPkg "github.com/lokesh-go/youtube-data-golang/src/pkg/youtube"
)

const searchText = "bhola teaser"

// Methods ...
type Methods interface {
	Start()
}

type client struct {
	config      *config.Config
	ytServices  ytPkg.Methods
	dalServices dal.Methods
}

// New ...
func New(config *config.Config, ytServices ytPkg.Methods, dalServices dal.Methods) Methods {
	return &client{config, ytServices, dalServices}
}

func (c *client) Start() {
	// Checks
	if !c.config.Job.Enabled {
		return
	}

	// Gets cron job
	cron := cron.New()

	// Adds function
	cron.AddFunc(c.config.Job.Interval, c.fetchLatestVideoAndPushData)

	// Start job
	cron.Start()
}

func (c *client) fetchLatestVideoAndPushData() {
	// Search on youtube
	res, err := c.ytServices.Search(searchText)
	if err != nil {
		log.Println("youtube search error: ", err.Error())
		return
	}

	
}
