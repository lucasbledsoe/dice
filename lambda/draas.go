package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/travis-g/dice"
)

func handleRequest(ctx context.Context, props dice.RollerProperties) (roller dice.Roller, err error) {
	if props.Count < 2 {
		roller, err = dice.NewRoller(&props)
	} else {
		roller, err = dice.NewRollerGroup(&props)
	}
	if err != nil {
		return
	}
	err = roller.FullRoll(ctx)
	return
}

func main() {
	lambda.Start(handleRequest)
}
