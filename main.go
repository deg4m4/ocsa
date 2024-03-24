package main

import "deg4m4/ocsa/core"

func main() {
	ocsaServer := core.Ocsa{}

	ocsaServer.SetHost("localhost")
	ocsaServer.SetPort(8052)
	ocsaServer.SetVerbose(true)
	ocsaServer.SetRootDir("./")

	ocsaServer.RunServer()

}
