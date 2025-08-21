//go:build generate
// +build generate

// Пакет, не участвующий в сборке приложения.
package generate

//go:generate go tool oapi-codegen --config=api/oapi-server.cfg.yaml api/openapi.yaml
