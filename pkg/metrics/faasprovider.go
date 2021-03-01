package metrics

import (
	"github.com/smvfal/faas-monitor/pkg/metrics/apiserver"
	"github.com/smvfal/faas-monitor/pkg/metrics/gateway"
	"github.com/smvfal/faas-monitor/pkg/metrics/metricsserver"
	"github.com/smvfal/faas-monitor/pkg/metrics/prometheus"
)

type FaasProvider struct{}

func (*FaasProvider) Functions() ([]string, error) {
	return gateway.Functions()
}

func (*FaasProvider) FunctionReplicas(functionName string) (int, error) {
	return prometheus.FunctionReplicas(functionName)
}

func (*FaasProvider) ResponseTime(functionName string, sinceSeconds int64) (float64, error) {
	return prometheus.ResponseTime(functionName, sinceSeconds)
}

func (*FaasProvider) ProcessingTime(functionName string, sinceSeconds int64) (float64, error) {
	return prometheus.ProcessingTime(functionName, sinceSeconds)
}

func (*FaasProvider) Throughput(functionName string, sinceSeconds int64) (float64, error) {
	return prometheus.Throughput(functionName, sinceSeconds)
}

func (*FaasProvider) ColdStart(functionName string, sinceSeconds int64) (float64, error) {
	return apiserver.ColdStart(functionName, sinceSeconds)
}

func (*FaasProvider) Top(functionName string) (map[string]int64, map[string]int64, error) {
	return metricsserver.Top(functionName)
}