package template

var DigTemplate = ParseTemplate(`
package dig

import (
	"go.uber.org/dig"
	"lark/apps/apis/{{.PackageName}}/internal/config"
	"lark/apps/apis/{{.PackageName}}/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
	"log/slog"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(repo.New{{.UpperServiceName}}Repository)
	Provide(cache.New{{.UpperServiceName}}Cache)
	Provide(service.New{{.UpperServiceName}}Service)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		slog.Warn(err.Error())
	}
}
`)