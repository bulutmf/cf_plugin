package main


import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
)

type PluginTest struct{}

func (c *PluginTest) Run(cliConnection plugin.CliConnection, args []string) {
  // Ensure that we called the command basic-plugin-command
  if args[0] == "basic-plugin-command" {
    fmt.Println("Running the basic-plugin-command")
  }
}

func (c *PluginTest) GetMetadata() plugin.PluginMetadata {
  return plugin.PluginMetadata{
    Name: "PluginTest",
    Version: plugin.VersionType{
      Major: 1,
      Minor: 0,
      Build: 0,
    },
    Commands: []plugin.Command{
      plugin.Command{
        Name:     "basic-plugin-command",
        HelpText: "Basic plugin command's help text",

        // UsageDetails is optional
        // It is used to show help of usage of each command
        UsageDetails: plugin.Usage{
          Usage: "basic-plugin-command\n   cf basic-plugin-command",
        },
      },
    },
  }
}



func main() {
	fmt.Printf("Hello world, registering my plugin\n")
	plugin.Start(new(PluginTest))

}


