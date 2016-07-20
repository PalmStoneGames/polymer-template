// Copyright 2016 Palm Stone Games, Inc. All rights reserved.

package main

import (
	"code.psg.io/polymer-template/client/elements" // import elements so their initialization code can run

	"code.palmstonegames.com/polymer"
)

func main() {
	<-polymer.OnReady()

	elements.InitRouter(polymer.GetWindow().Location().Pathname)
}
