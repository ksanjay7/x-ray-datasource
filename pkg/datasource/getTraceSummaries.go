package datasource

import (
  "context"
  "fmt"
  "github.com/grafana/grafana-plugin-sdk-go/backend"
)

func (ds *Datasource) getTraceSummaries(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
  return nil, fmt.Errorf("not implemented")
}