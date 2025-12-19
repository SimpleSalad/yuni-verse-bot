module github.com/SimpleSalad/yuni-verse-bot

go 1.25.5

require (
	github.com/SimpleSalad/yuni-verse-bot/lamps v0.0.0-20251219052104-7ad98fea7deb
	github.com/bwmarrin/discordgo v0.29.0
)

require (
	github.com/SimpleSalad/yuni-verse-bot/lamps/testlamp v0.0.0-20251219054506-b268373fd273 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
)

replace github.com/SimpleSalad/yuni-verse-bot/lamps => ./lamps
