// Copyright 2016 Palm Stone Games, Inc. All rights reserved.

package elements

import (
	"encoding/json"
	"fmt"
	"strconv"

	"code.palmstonegames.com/polymer"
	"code.psg.io/polymer-template/json"
)

func init() {
	polymer.Register("changeme-adder", &Adder{})
}

type Adder struct {
	*polymer.Proto

	Addends []string `polymer:"bind"`
	Sum     int      `polymer:"bind"`
	Error   string   `polymer:"bind"`

	Create chan *polymer.Event `polymer:"handler"`
	Delete chan *polymer.Event `polymer:"handler"`
	Submit chan *polymer.Event `polymer:"handler"`
}

func (a *Adder) Created() {
	// Add your own to be triggered on creation
}

func (a *Adder) Ready() {
	a.ListenEvents()
}

func (a *Adder) ListenEvents() {
	go func() {
		for {
			select {
			case <-a.Create:
				a.Addends = append(a.Addends, "")
				a.Notify("addends")
			case e := <-a.Delete:
				idx, err := strconv.ParseInt(e.RootTarget.GetAttribute("data-index"), 10, 32)
				if err != nil {
					a.Error = fmt.Sprintf("Error while parsing index: %v\n", err)
					a.Notify("error")
					continue
				}
				copy(a.Addends[idx:], a.Addends[idx+1:])
				a.Addends = a.Addends[:len(a.Addends)-1]

				a.Notify("addends")
			case <-a.Submit:
				a.submit()
			}
		}
	}()
}

func (a *Adder) submit() {
	defer a.Notify("error", "sum")
	a.Sum = 0

	var addReq jsonrpc.AddRequest
	addReq.Addends = make([]int, 0, len(a.Addends))
	for _, addend := range a.Addends {
		val, err := strconv.Atoi(addend)
		if err != nil {
			a.Error = fmt.Sprintf("Error converting string(%q) to integer: %v", addend, err)
			return
		}
		addReq.Addends = append(addReq.Addends, val)
	}

	resp, err := DoHTTPJSON("POST", "/add", &addReq)
	if err != nil {
		a.Error = fmt.Sprintf("Error during RPC: %v", err)
		return
	}
	defer resp.Body.Close()

	var addResp jsonrpc.AddResponse
	if err := json.NewDecoder(resp.Body).Decode(&addResp); err != nil {
		a.Error = fmt.Sprintf("Error while decoding json: %v", err)
		return
	}

	a.Sum = addResp.Sum
	a.Error = ""
}

func (a *Adder) ComputeSum() int {
	return a.Sum
}
