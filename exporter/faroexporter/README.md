# Faro Exporter

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [alpha]: logs, traces   |
| Distributions | [contrib] |
| Issues        | [![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Aexporter%2Ffaro%20&label=open&color=orange&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aopen+is%3Aissue+label%3Aexporter%2Ffaro) [![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Aexporter%2Ffaro%20&label=closed&color=blue&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aclosed+is%3Aissue+label%3Aexporter%2Ffaro) |
| Code coverage | [![codecov](https://codecov.io/github/open-telemetry/opentelemetry-collector-contrib/graph/main/badge.svg?component=exporter_faro)](https://app.codecov.io/gh/open-telemetry/opentelemetry-collector-contrib/tree/main/?components%5B0%5D=exporter_faro&displayType=list) |
| [Code Owners](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/CONTRIBUTING.md#becoming-a-code-owner)    | [@dehaansa](https://www.github.com/dehaansa), [@rlankfo](https://www.github.com/rlankfo), [@mar4uk](https://www.github.com/mar4uk) |

[alpha]: https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/component-stability.md#alpha
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
<!-- end autogenerated section -->

The Faro exporter sends telemetry data to a [Faro](https://grafana.com/oss/faro/) endpoint.

## Configuration

The following settings are required:

- `endpoint` (no default): The URL to send telemetry data to (e.g., https://faro.example.com/collect).

The following settings can be optionally configured:

- `sending_queue`
  - `enabled` (default = true)
  - `num_consumers` (default = 10)
  - `queue_size` (default = 1000)
- `retry_on_failure`
  - `enabled` (default = true)
  - `initial_interval` (default = 5s): Time to wait after the first failure before retrying.
  - `max_interval` (default = 30s): Upper bound on backoff.
  - `max_elapsed_time` (default = 300s): Maximum amount of time spent trying to send a batch.
- `timeout` (default = 5s): HTTP request timeout when sending data.
- `read_buffer_size` (default = 0): Size of the buffer used to read the response body.
- `write_buffer_size` (default = 512 KiB): Size of the buffer used to write the request body.
- `headers` (default = `{}`): Additional headers to send with the request.
- `compression` (default = none): Compression method to use for the request body. Supported values: `none`, `gzip`.

Example:

```yaml
exporters:
  faro:
    endpoint: https://faro.example.com/collect
    timeout: 10s
    headers:
      X-API-Key: "my-api-key"
```

The full list of settings exposed for this exporter are documented [here](./config.go) with detailed sample configurations [here](./testdata/config.yaml).

## Getting Started

The following settings are required:

- `endpoint` (no default): URL to which the exporter is going to send Faro telemetry data. For example: `https://faro.example.com/collect`.

To use TLS, specify `https://` as the protocol scheme in the URL passed to the `endpoint` property.
See [Advanced Configuration](#advanced-configuration) for more TLS options.

Example:

```yaml
exporters:
  faro:
    endpoint: "https://faro.example.com/collect"

  faro/tlsnoverify:
    endpoint: "https://faro.example.com/collect"
    tls:
      insecure_skip_verify: true
```

## Advanced Configuration

Several helper files are leveraged to provide additional capabilities automatically:

- [HTTP client settings](https://github.com/open-telemetry/opentelemetry-collector/blob/main/config/confighttp/README.md#client-configuration)
- [TLS and mTLS settings](https://github.com/open-telemetry/opentelemetry-collector/blob/main/config/configtls/README.md)
- [Queuing, retry and timeout settings](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md) 