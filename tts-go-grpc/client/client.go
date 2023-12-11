package client

type Client interface {
	init()
	Cleanup()
}
