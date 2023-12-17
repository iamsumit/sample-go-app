package router

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/api/apigokit"
	redoc "github.com/mvrilo/go-redoc"
	swgui "github.com/swaggest/swgui/v5cdn"
)

// ConfigureDocsRoutes configures the docs routes.
func ConfigureDocsRoutes(a *apigokit.Handler, cfg Config) {
	ServeSWGUIDocsRoutes(a.RoutingHandler(), "v1")
	ServeReDocsRoutes(a.RoutingHandler(), "v1")
}

// ServeSWGUIDocsRoutes serves the swagger documentation routes usign swgui.
func ServeSWGUIDocsRoutes(a *api.API, version string) {
	// -------------------------------------------------------------------
	// Swagger JSON
	// -------------------------------------------------------------------
	vSwagPath := fmt.Sprintf("/%s/swagger.json", version)
	a.Handle(http.MethodGet, vSwagPath, func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		// Construct the full path to the JSON file
		swagJSONPath := filepath.Join("./docs", "swagger", version, "/swagger.json")
		http.ServeFile(
			w,
			r,
			swagJSONPath,
		)

		return nil
	})

	// -------------------------------------------------------------------
	// Swagger UI
	// -------------------------------------------------------------------
	docPath := fmt.Sprintf("/%s/swgui-doc", version)
	a.Handle(http.MethodGet, docPath, func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		swgui.New(
			"Activity API",
			vSwagPath,
			docPath,
		).ServeHTTP(w, r)

		return nil
	})
}

// ServeReDocsRoutes serves the swagger documentation using redoc library.
func ServeReDocsRoutes(a *api.API, version string) {
	swagJSONPath := filepath.Join("./docs", "swagger", version, "/swagger.json")
	doc := redoc.Redoc{
		Title:       "Activity API",
		Description: "Activity V1 API",
		SpecFile:    swagJSONPath,
		SpecPath:    fmt.Sprintf("/%s/swagger.json", version),
		DocsPath:    fmt.Sprintf("/%s/redoc", version),
	}

	a.Handle(http.MethodGet, doc.DocsPath, func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		doc.Handler().ServeHTTP(w, r)
		return nil
	})
}
