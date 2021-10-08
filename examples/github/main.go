package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/lcook/hookrelay"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	ExamplePort string = "8080"
	ExampleYaml string = "github.yaml"
)

type Commit struct {
	ID      string
	Message string
	Author  struct {
		Email    string
		Name     string
		Username string
	} `json:"author,omitempty"`
}

type Payload struct {
	Commits []Commit `json:"commits,omitempty"`
}

type Handler []hookrelay.Hook

type GitHubCommit struct {
	Config struct {
		Endpoint string
	} `yaml:"github"`
	Option byte
}

func (gh *GitHubCommit) Endpoint() string { return gh.Config.Endpoint }
func (gh *GitHubCommit) Options() byte    { return gh.Option }

func (gh *GitHubCommit) LoadConfig(config string) error {
	file, err := os.Open(config)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	cfg := GitHubCommit{}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	gh.Config.Endpoint = cfg.Config.Endpoint

	return nil
}

func (gh *GitHubCommit) Response(i interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request from client", r.RemoteAddr)

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error(err)

			return
		}

		var payload Payload

		err = json.Unmarshal(buf, &payload)
		if err != nil {
			log.Error(err)

			return
		}

		log.Println("Parsed", len(payload.Commits), "commits")

		for _, commit := range payload.Commits {
			log.Printf("[%s] %s: %s\n", commit.ID[0:7], commit.Author.Username, commit.Message)
		}
	}
}

func main() {
	gh := &GitHubCommit{
		Option: (hookrelay.DefaultOptions),
	}

	srv, err := hookrelay.InitMux(nil, Handler{gh}, ExampleYaml, ExamplePort)
	if err != nil {
		log.Error(err)
	}

	log.Println("Listening on port", ExamplePort, "with endpoint", gh.Config.Endpoint)

	if err := srv.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Error(err)
	}
}
