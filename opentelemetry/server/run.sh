export OTEL_EXPORTER_OTLP_ENDPOINT=":4317"
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
export OTEL_EXPORTER_OTLP_HEADERS=authorization="Bearer oidc_token",foo=bar
export OTEL_EXPORTER_OTLP_COMPRESSION=gzip

go run ./main.go