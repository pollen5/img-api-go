# Img API Go
This is a [GoLang](https://golang.org) wrapper for my [Image Manipulation API](https://github.com/pollen5/img-api) that I made.

This is an official wrapper and will be updated quickly as the API changes.

## Install
```sh
$ go get github.com/pollen5/img-api
```

## Usage
**Import**
```go
import (
  "github.com/pollen5/img-api-go"
)
```
**Create a client.**
```go
// All options optional.
client := &imgapi.NewClient(&imgapi.Options{
  Port: 3030,
  Host: "localhost",
  Password: "",
  Client: http.DefaultClient,
})

// If the default values (like above) is enough for you just use: 
client := imgapi.DefaultClient // A default client instance, also makes it easier to carry it around.
```
When using `DefaultClient` feel free to modify the options too but be careful that the changes are global and any dependencies using it could conflict but this is mostly for personal use so it's perfectly fine to setup the options in `DefaultClient` so it's easier to carry around the client between files.

**Call endpoints:**
```go
pong, err := client.Ping()
stats, err := client.Stats(true)
image, err := client.Tom(user.AvatarURL("1024"))
```
The image returns are of type `[]byte` to send it via [discordgo](https://github.com/bwmarrin/discordgo) wrap it with a `Buffer` e.g `session.ChannelFileSend(id, "out.png", bytes.NewBuffer(image))`

View the documentation at [GoDoc](https://godoc.org/github.com/pollen5/img-api-go)

## Changelog

#### v1.0.0 (19/5/2020)
- Initial release

## License
Released under the [MIT License](LICENSE)
