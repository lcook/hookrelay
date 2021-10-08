package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/lcook/hookrelay"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	ExamplePort  string = "8080"
	ExampleYaml  string = "discord.yaml"
	ExampleToken string = "BOT_TOKEN"
)

type ExamplePayload struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Handler []hookrelay.Hook

type Example struct {
	Config struct {
		Endpoint string
		Channel  string
	} `yaml:"example"`
	Option byte
}

func (e *Example) Endpoint() string { return e.Config.Endpoint }
func (e *Example) Options() byte    { return e.Option }

func (e *Example) LoadConfig(config string) error {
	file, err := os.Open(config)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	cfg := Example{}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	e.Config.Endpoint = cfg.Config.Endpoint
	e.Config.Channel = cfg.Config.Channel

	return nil
}

func (e *Example) Response(i interface{}) func(w http.ResponseWriter, r *http.Request) {
	dg := i.(*discordgo.Session)

	return func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error(err)

			return
		}

		var payload ExamplePayload

		err = json.Unmarshal(buf, &payload)
		if err != nil {
			log.Error(err)

			return
		}

		_, err = dg.ChannelMessageSend(e.Config.Channel, fmt.Sprintf("%s said: %s", payload.Username, payload.Message))
		if err != nil {
			log.Error(err)
		}
	}
}

func main() {
	dg, err := discordgo.New("Bot " + ExampleToken)
	if err != nil {
		log.Error(err)
	}

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	e := &Example{
		Option: (hookrelay.DefaultOptions),
	}

	srv, err := hookrelay.InitMux(dg, Handler{e}, ExampleYaml, ExamplePort)
	if err != nil {
		log.Error(err)
	}

	log.Println("Listening on port", ExamplePort, "with endpoint", e.Config.Endpoint)

	if err := srv.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Error(err)
	}

	dg.Close()
}
