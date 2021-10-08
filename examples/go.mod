module examples

go 1.16

require (
	github.com/bwmarrin/discordgo v0.23.2
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/lcook/hookrelay v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/lcook/hookrelay => ../
