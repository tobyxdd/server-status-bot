package main

import (
	"log"
	"os"
	"gopkg.in/tucnak/telebot.v2"
	"time"
	"github.com/shirou/gopsutil/process"
	"fmt"
)

const (
	appTitle             = "Simple Server Status Bot"
	appVersion           = "0.1.0"
	configFilename       = "config.json"
	sampleConfigFilename = "config.sample.json"
)

func main() {
	log.Printf("%s v%s\n", appTitle, appVersion)
	config, err := loadBotConfig(configFilename)
	if err != nil {
		if os.IsNotExist(err) {
			generateSampleBotConfig(sampleConfigFilename)
			log.Fatalf("Could not find %s. A sample config file %s has been generated."+
				" Please create a configuration file and restart the program.",
				configFilename, sampleConfigFilename)
		} else {
			log.Fatalf("Unable to load configuration: %s\n", err.Error())
		}
	}
	startBot(config)
}

func startBot(config *botConfig) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  config.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("Initialization failed: %s\n", err.Error())
	}
	bot.Handle("/status", func(m *telebot.Message) {
		if m.FromChannel() {
			//No channel support for now
			return
		}
		if m.FromGroup() {
			log.Printf("%s in group %s by @%s\n", m.Text, m.Chat.Title, m.Sender.Username)
		} else {
			log.Printf("%s by @%s\n", m.Text, m.Sender.Username)
		}
		bot.Send(m.Chat, getStatusText(config))
	})
	log.Println("Up and running...")
	bot.Start()
}

func getStatusText(config *botConfig) string {
	statusMap := make(map[string]bool)
	procs, err := process.Processes()
	if err == nil {
		for _, proc := range procs {
			//log.Println(proc.Name())
			for _, service := range config.Services {
				procName, err := proc.Name()
				if err == nil && service.Process == procName {
					statusMap[service.Name] = true
				}
			}
		}
	}
	var msg string
	for _, service := range config.Services {
		msg += fmt.Sprintf("%s %s\n", getStatusIcon(statusMap[service.Name]), service.Name)
	}
	return msg
}

func getStatusIcon(ok bool) string {
	if ok {
		return "✅"
	} else {
		return "❌"
	}
}
