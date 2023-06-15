package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

const (
	frequency  = 100
	duration   = 4 * time.Second
	targetUrl  = "http://google.com"
	httpMethod = "GET"
)

type AttackPlan struct {
	attacker *vegeta.Attacker
	targeter vegeta.Targeter
	rate     vegeta.ConstantPacer
}

func main() {
	attackPlan := planAttack()
	metrics := execute(attackPlan)

	displayRequest(metrics)
	displayError(metrics)
	displayLatency(metrics)
	displayPayload(metrics)
}

func execute(plan AttackPlan) vegeta.Metrics {
	var metrics vegeta.Metrics
	for res := range plan.attacker.Attack(plan.targeter, plan.rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	return metrics
}

func planAttack() AttackPlan {
	var attackPlan AttackPlan
	attackPlan.rate = vegeta.Rate{Freq: frequency, Per: time.Second}
	attackPlan.targeter = vegeta.NewStaticTargeter(vegeta.Target{
		Method: httpMethod,
		URL:    targetUrl,
	})
	attackPlan.attacker = vegeta.NewAttacker()
	fmt.Print("\n----------------------\n")
	fmt.Printf(
		"Vegeta will attack \"%s\", %d times per second for %d seconds:\n",
		targetUrl,
		frequency,
		duration/time.Second,
	)
	return attackPlan
}

func displayRequest(metrics vegeta.Metrics) {
	fmt.Print("\n----------------------REQUEST\n")
	fmt.Printf("Total: %v requests\n", metrics.Requests)
	fmt.Printf("Rate: %v\n", metrics.Rate)
	fmt.Printf("Throughput: %v\n", metrics.Throughput)
	fmt.Printf("\nSuccess rate: %v%%\n", metrics.Success*100)
	for k, v := range metrics.StatusCodes {
		fmt.Printf("\tstatus %s: %d requests\n", k, v)
	}
}

func displayError(metrics vegeta.Metrics) {
	if len(metrics.Errors) == 0 {
		return
	}
	fmt.Print("\n----------------------ERRORS\n")
	fmt.Print("Messages:\n")
	for _, errorMessage := range metrics.Errors {
		fmt.Printf("\t- \"%s\"\n", errorMessage)
	}
}

func displayLatency(metrics vegeta.Metrics) {
	fmt.Print("\n----------------------LATENCY\n")
	fmt.Printf("Total: %s\n", metrics.Latencies.Total)
	fmt.Printf("Mean: %s\n", metrics.Latencies.Mean)
	fmt.Printf("50th percentile: %s\n", metrics.Latencies.P50)
	fmt.Printf("90th percentile: %s\n", metrics.Latencies.P90)
	fmt.Printf("95th percentile: %s\n", metrics.Latencies.P95)
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("Max: %s\n", metrics.Latencies.Max)
	fmt.Printf("Min: %s\n", metrics.Latencies.Min)
}

func displayPayload(metrics vegeta.Metrics) {
	fmt.Print("\n----------------------PAYLOAD\n")
	fmt.Print("Bytes IN:\n")
	fmt.Printf("\tTotal: %d\n", metrics.BytesIn.Total)
	fmt.Printf("\tMean: %v\n", metrics.BytesIn.Mean)
	fmt.Print("Bytes OUT:\n")
	fmt.Printf("\tTotal: %d\n", metrics.BytesOut.Total)
	fmt.Printf("\tMean: %v\n", metrics.BytesOut.Mean)
}
