# Logging (Beyond slog)

This directory compares three popular structured logging libraries for Go and offers quick guidance:

- `zerolog` — extremely low-allocation, fast, JSON-first logger. Good for high-throughput services where allocations and GC pressure matter. By default it encodes JSON; use `ConsoleWriter` for human-friendly output during development.

- `zap` — high-performance, well-tested, and widely used (Uber). Provides both a fast, structured production logger (`zap.NewProduction`) and a developer-friendly sugared logger (`zap.S().`). Strong type-safety and good ecosystem; pairs well with tracing and metrics.

- `logrus` — feature-rich, easy-to-use, and familiar API (similar to other languages). Uses reflection and has higher allocation overhead; good for small services, CLIs, and when you prefer a familiar ergonomics over maximum speed.

Quick comparison

- Performance: `zerolog` ≈ `zap` (generated code path) >> `logrus`.
- Ergonomics: `logrus` (easy) ≈ `zap` (sugared) > `zerolog` (fluent, minimal API).
- Output: All support JSON; `zerolog`/`zap` target JSON-first production. Use text/console writers for readability in dev.
- Type-safety: `zap`'s typed fields reduce mistakes; `zerolog` is fluent and compact; `logrus` uses interface{} for field values.

When to choose

- Choose `zerolog` or `zap` when you need low-overhead, production-grade structured logs and care about allocations/latency.
- Choose `logrus` for quick prototypes, CLIs, or when you want lots of integrations/plugins and easy migration from other ecosystems.

Usage guidance

- Configure log level via environment variables or config files and avoid hardcoding verbose levels in production.
- Prefer JSON logs in production for log ingestion systems (Datadog, ELK, Loki, etc.).
- Use `ConsoleWriter` or human formatters for local development and CI logs if needed.
- Do not log sensitive data; redact or avoid logging secrets/PII.
- For file rotation, use external rotation tools or libraries (e.g., `lumberjack`) rather than reimplementing rotation logic in your app.
- Integrate logs with tracing (W3C traceparent or OpenTelemetry) by injecting trace IDs as log fields.

Examples in this folder

- `01_zerolog/` — minimal `ConsoleWriter` example.
- `02_zap/` — `zap.NewProduction` example.
- `03_logrus/` — simple `logrus` usage with structured fields.

Next steps

- Add a short section showing how to attach trace IDs to logs (OpenTelemetry example).
- Add CI snippets demonstrating how to run examples and verify output format.
