# Metrics

this code is wrapping open telemetry metrics to serve metrics that can be collected using prometheus

## How to use
- init with manual registration see example [Use Register](example/use_register)
- init with config [Use Config](example/use_register_config)

## Output
```
...
# HELP test_sent_counter data test sent counter
# TYPE test_sent_counter counter
test_sent_counter{service_name="unknown_service:main",telemetry_sdk_language="go",telemetry_sdk_name="opentelemetry",telemetry_sdk_version="1.8.0"} 16
# HELP test_sent_histogram data test sent histogram
# TYPE test_sent_histogram histogram
test_sent_histogram_bucket{service_name="unknown_service:main",telemetry_sdk_language="go",telemetry_sdk_name="opentelemetry",telemetry_sdk_version="1.8.0",le="+Inf"} 16
test_sent_histogram_sum{service_name="unknown_service:main",telemetry_sdk_language="go",telemetry_sdk_name="opentelemetry",telemetry_sdk_version="1.8.0"} 6.452e+09
test_sent_histogram_count{service_name="unknown_service:main",telemetry_sdk_language="go",telemetry_sdk_name="opentelemetry",telemetry_sdk_version="1.8.0"} 16
```