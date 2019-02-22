package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Medium/medium-sdk-go"
	"github.com/cappyzawa/medium-resource/out"
)

type Out struct {
	InStream  io.Reader
	OutStream io.Writer
	ErrStream io.Writer
}

func (o *Out) Execute(args []string) int {
	var req out.Request
	if err := o.inputRequest(&req); err != nil {
		fmt.Fprintf(o.ErrStream, "invalid payload: %s\n", err)
		return 1
	}

	if len(args) < 2 {
		fmt.Fprintf(o.ErrStream, "destination path not specified")
		return 1
	}
	source := args[1]

	// https://godoc.org/github.com/Medium/medium-sdk-go#NewClientWithAccessToken
	mc := medium.NewClientWithAccessToken(req.Source.AccessToken)
	command := out.Command{
		MediumClient: mc,
	}
	res, err := command.Run(source, req)
	if err != nil {
		fmt.Fprintf(o.ErrStream, "failed to run o command: %s\n", err)
		return 1
	}

	if err := o.outputResponse(res); err != nil {
		fmt.Fprintf(o.ErrStream, "failed to decode response: %s\n", err)
		return 1
	}
	return 0
}

func (o *Out) inputRequest(req *out.Request) error {
	if err := json.NewDecoder(o.InStream).Decode(req); err != nil {
		return err
	}
	return nil
}

func (o *Out) outputResponse(res *out.Response) error {
	if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
		return err
	}
	return nil
}

func main() {
	c := &Out{
		InStream:  os.Stdin,
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
	}
	os.Exit(c.Execute(os.Args))
}