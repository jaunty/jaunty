package cli

// CLI is the entrypoint for Jaunty's app.
type CLI struct {
	Debug bool `short:"d" help:"Run in debug mode?" default:"false" required:"false"`

	Web Web `cmd:"" help:"Run the web server."`
}
