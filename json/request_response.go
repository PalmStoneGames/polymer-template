// Copyright 2016 Palm Stone Games, Inc. All rights reserved.

package jsonrpc

type AddRequest struct {
	Addends []int `json:"addends"`
}

type AddResponse struct {
	Sum int `json:"sum"`
}
