package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ReadToolArgs struct {
	FilePath string `json:"file_path"`
}

func ParseReadToolArgs(unprocessedArgs string) (*ReadToolArgs, error) {
	var args ReadToolArgs
	if err := json.Unmarshal([]byte(unprocessedArgs), &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func ReadTool(args *ReadToolArgs) error {
	b, err := os.ReadFile(args.FilePath)
	if err != nil {
		return err
	}
	fmt.Println(b)
	return nil
}
