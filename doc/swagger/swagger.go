package swagger

import (
	"bytes"
	"github.com/zeromicro/go-zero/rest"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterSwagger(env string, server *rest.Server) {

	swaggerFile, err := os.Open("doc/swagger/swagger.json")
	if err != nil {
		log.Println(err)
	}
	defer swaggerFile.Close()
	SwaggerByte, err := ioutil.ReadAll(swaggerFile)
	if err != nil {
		log.Println(err)
	}

	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/swagger",
			Handler: Doc("/swagger", env),
		},
		{
			Method: http.MethodGet,
			Path:   "/swagger-json",
			Handler: func(writer http.ResponseWriter, request *http.Request) {
				writer.Header().Set("Content-Type", "application/json; charset=utf-8")
				_, err := writer.Write(SwaggerByte)
				if err != nil {
					httpx.Error(writer, err)
				}
			},
		},
	})

}

type Opts func(*swaggerConfig)

// SwaggerOpts configures the Doc gmiddlewares.
type swaggerConfig struct {
	// SpecURL the url to find the spec for
	SpecURL string
	// SwaggerHost for the js that generates the swagger ui site, defaults to: http://petstore.swagger.io/
	SwaggerHost string
}

func Doc(basePath, env string, opts ...Opts) http.HandlerFunc {
	config := &swaggerConfig{
		SpecURL:     basePath + "-json",
		SwaggerHost: "https://petstore.swagger.io",
	}
	for _, opt := range opts {
		opt(config)
	}

	// swagger html
	tmpl := template.Must(template.New("swaggerdoc").Parse(swaggerTemplateV2))
	buf := bytes.NewBuffer(nil)
	err := tmpl.Execute(buf, config)
	uiHTML := buf.Bytes()

	// permission
	needPermission := false
	if env == "prod" {
		needPermission = true
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		if err != nil {
			httpx.Error(rw, err)
			return
		}
		if r.URL.Path == basePath {
			if needPermission {
				rw.WriteHeader(http.StatusOK)
				rw.Header().Set("Content-Type", "text/plain")
				_, err = rw.Write([]byte("Swagger not open on prod"))
				if err != nil {
					httpx.Error(rw, err)
				}
				return
			}

			rw.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, err = rw.Write(uiHTML)
			if err != nil {
				httpx.Error(rw, err)
				return
			}

			rw.WriteHeader(http.StatusOK)
			return
		}
	}
}

const swaggerTemplateV2 = `
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>API documentation</title>
    <link rel="stylesheet" type="text/css" href="{{ .SwaggerHost }}/swagger-ui.css" >
    <link rel="icon" type="image/png" href="{{ .SwaggerHost }}/favicon-32x32.png" sizes="32x32" />
    <link rel="icon" type="image/png" href="{{ .SwaggerHost }}/favicon-16x16.png" sizes="16x16" />
    <style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }

      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
    </style>
  </head>

  <body>
    <div id="swagger-ui"></div>

    <script src="{{ .SwaggerHost }}/swagger-ui-bundle.js"> </script>
    <script src="{{ .SwaggerHost }}/swagger-ui-standalone-preset.js"> </script>
    <script>
    window.onload = function() {
      // Begin Swagger UI call region
      const ui = SwaggerUIBundle({
        "dom_id": "#swagger-ui",
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout",
		validatorUrl: null,
        url: "{{ .SpecURL }}",
      })

      // End Swagger UI call region
      window.ui = ui
    }
  </script>
  </body>
</html>`
