package main

import (
	"flag"
	"time"
	"errors"
	"fmt"
	"context"

	"github.com/kettsun0123/loadtest/scenarios"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

// option parameters
var (
	rate      uint64
	worker    uint64
	parallel  uint64
	duration  time.Duration
	urlPrefix string
	scenario  string
	output    string
	forceFlag bool
)

func main() {
	// parse options
	if err := parseOptions(); err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	var eg *errgroup.Group
	switch Scenario(scenario) {
	case ScenarioALL:
		var err error
		if eg, err = scenarios.AttacksAll(ctx); err != nil {
			fmt.Println(err)
			return
		}
	case ScenarioBetsOnly:
		eg = scenarios.AttacksBets(ctx)
	case ScenarioLaunch:
		eg = scenarios.AttacksLaunch(ctx)
	default:
		fmt.Printf("unknown scenario: %s\n", scenario)
		return
	}

	fmt.Println("attack...")

	// cancel
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)

	go func() {
		select {
		case <-sigCh:
			fmt.Println("catch cancel signal")
			cancel()
			select {
			case <-ctx.Done():
				switch err := ctx.Err(); err {
				case nil, context.Canceled:
				default:
					fmt.Println(err)
				}
				fmt.Println("cancel done")
			}
		}
	}()

	// wait error group
	if err := eg.Wait(); err != nil {
		switch err := ctx.Err(); err {
		case nil, context.Canceled:
		default:
			fmt.Println(err)
		}
	}
	fmt.Println("attack... done")

}

func parseOptions() error {
	flag.Uint64Var(&rate, "rate", 1, "Requests per second per thread")
	flag.Uint64Var(&worker, "worker", 2, "worker")
	flag.Uint64Var(&parallel, "parallel", 1, "concurrency thread size")
	flag.DurationVar(&duration, "duration", time.Second*10, "duration to load")
	flag.StringVar(&urlPrefix, "url", "http://localhost:8000", "url prefix")
	flag.StringVar(&scenario, "scenario", "all", "all or bets or launch")
	flag.StringVar(&output, "output", "stdout", "stdout or json or text")
	flag.BoolVar(&forceFlag, "force", false, "Options for executing load-scenario ignoring Confirmation")
	flag.Parse()

	if rate == 0 {
		return errors.New("rate must be greater than 0")
	}
	if worker == 0 {
		return errors.New("worker must be greater than 0")
	}
	if parallel == 0 {
		return errors.New("parallel must be greater than 0")
	}
	if duration < time.Second {
		return errors.New("duration must be greater than or equal to 1s")
	}
	if urlPrefix == "" {
		return errors.New("urlPrefix required")
	}
	switch Scenario(scenario) {
	case ScenarioALL, ScenarioBetsOnly, ScenarioLaunch:
	default:
		return errors.New("unknown scenario")
	}
	switch output {
	case OutputStdout, OutputJson, OutputText:
	default:
		return errors.New("unknown output format")
	}

	fmt.Printf("start load test:\n")
	fmt.Printf("scenario=%s\n", scenario)
	fmt.Printf("duration=%.0f\n", duration.Seconds())
	fmt.Printf("rate=%d\n", rate)
	fmt.Printf("worker=%d\n", worker)
	fmt.Printf("parallel=%d\n", parallel)
	fmt.Printf("urlPrefix=%s\n", urlPrefix)
	fmt.Printf("output=%s\n", output)
	fmt.Printf("force=%s\n", forceFlag)

	if forceFlag {
		if err := Confirm(); err != nil {
			return err
		}
	}
	return nil
}

type Scenario string

const (
	ScenarioALL          Scenario = "all"
	ScenarioBetsOnly Scenario = "bets"
	ScenarioLaunch       Scenario = "launch"
)

type attackOptions struct {
	rate     uint64
	worker   uint64
	duration time.Duration
}

const (
	OutputStdout string = "stdout"
	OutputJson   string = "json"
	OutputText   string = "text"
)

func Confirm() error {
	fmt.Println("continue (y/N)")
	var confirm string
	if _, err := fmt.Scanln(&confirm); err != nil {
		return err
	}
	if confirm != "y" {
		return errors.New("exit")
	}
	return nil
}