package beater

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/qingyunha/consolebeat/config"
)

type Consolebeat struct {
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Consolebeat{
		config: config,
	}
	return bt, nil
}

func (bt *Consolebeat) Run(b *beat.Beat) error {
	var msg string
	logp.Info("consolebeat is running! Hit CTRL-D to stop it.")
	bt.client = b.Publisher.Connect()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg = scanner.Text()
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"message":    msg,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
	}
	return nil
}

func (bt *Consolebeat) Stop() {
	bt.client.Close()
}
