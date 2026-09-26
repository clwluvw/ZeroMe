# ZeroMe

> [!WARNING]
> **This project is retired and no longer maintained.**
> Prometheus now injects the missing zero samples itself with the [Start (Created) Timestamps Zero Injection](https://prometheus.io/docs/prometheus/latest/feature_flags/#start-created-timestamps-zero-injection) feature flag.
> Use it instead of ZeroMe; see [Migrating to Prometheus](#migrating-to-prometheus).

ZeroMe provides a workaround for the issue described in [Prometheus issue #3886](https://github.com/prometheus/prometheus/issues/3886).
It fetches dynamic (appearing and disappearing) metrics at their scrape interval and checks for missing zero values before the metrics appear.

## Migrating to Prometheus

1. Start Prometheus with `--enable-feature=created-timestamp-zero-ingestion`.
1. Make sure the scraped applications expose start (created) timestamps. Unlike ZeroMe, Prometheus cannot inject the zero sample without them.
   Prometheus reads them from the Prometheus protobuf format (recommended) or from OpenMetrics 1.0 `_created` series, and negotiates protobuf first when the flag is enabled, unless `scrape_protocols` is set explicitly.
1. Stop ZeroMe. If the remote write receiver or the out-of-order ingestion window was enabled only for ZeroMe, disable it as well.

The rest of this document describes how ZeroMe works and is kept for reference.

## Prometheus Query

To detect this situation, it uses the following Prometheus query:

```promql
count_over_time(metric_family[1m]) == 1
    and
metric_family
    and on (job, instance)
(count_over_time(up[1m]) == 2)
```

This query performs the following steps:

1. Queries all series with one data point in the last minute (1m represents twice the scrape interval).
1. Uses `and` to compare with all present series of the `metric_family` to ensure the missing data point is in the past, not present.
1. Verifies successful scrapes within the time interval based on job scrape identifiers (e.g., `job` and `instance` labels).

## Prometheus Writer

After fetching the query results, ZeroMe creates new time series with the same label set, setting the value to zero at an interval of `-$scrape_interval`.
These new series are then written to the [remote write API](https://prometheus.io/docs/prometheus/latest/querying/api/#remote-write-receiver).

## Configuration

```yaml
queriers:  # List of queriers for metrics
  querier1:  # Name of the querier
    address: http://localhost:9090/  # Prometheus address to query
    headers:  # Optional headers
      X-Scope-OrgID: tenant  # Example: tenant for querying Mimir
    type: prometheus  # Type of the querier (currently only Prometheus is supported)
writers:  # List of writers for metrics
  writer1:  # Name of the writer
    address: http://localhost:9090/api/v1/write  # Remote write API address
    headers:  # Optional additional headers
      X-Scope-OrgID: tenant  # Example: tenant for Mimir
    timeout: 30s  # Timeout for write requests
    type: prometheus  # Type of the writer (currently only Prometheus is supported)
metrics:  # List of metrics to apply ZeroMe
  - name: metric  # Metric family name
    interval: 30s  # Scrape interval
    querier: querier1  # Querier from the queriers list
    writer: writer1  # Writer from the writers list
    up_labels:  # Label names that identify the `up` metric
      - job
      - instance
```
