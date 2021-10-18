package swagger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"path"
)

type SwaggerModel struct {
	FileName string
	BasePath string
}

func HandleSwagger(fileName string, basePath string) *SwaggerModel{
	return &SwaggerModel{
		FileName: fileName,
		BasePath: basePath,
	}
}

func (s *SwaggerModel) HandleSwaggerUI(op middleware.SwaggerUIOpts) http.Handler{
	return middleware.SwaggerUI(op, nil)
}

func (s *SwaggerModel)HandleSwaggerFile(handler http.Handler) (http.Handler, error){
	if _, err := os.Stat(s.FileName); os.IsNotExist(err) {
		return nil, errors.New(fmt.Sprintf("%s file is not exist", s.FileName))
	}

	doc, err := loads.Spec(s.FileName)
	if err != nil {
		return nil, err
	}

	b, err := json.MarshalIndent(doc.Spec(), ""," ")
	if err != nil {
		return nil, err
	}

	return handlers.CORS()(middleware.Spec(s.BasePath, b, handler)), nil
}

func (s *SwaggerModel) AddSwagger(app *fiber.App){
	swaggerUI := middleware.SwaggerUIOpts{
		BasePath: s.BasePath,
		Path: "docs",
		SpecURL: path.Join(s.BasePath,"swagger.json"),
	}

	swaggerUIHandler := s.HandleSwaggerUI(swaggerUI)
	swaggerFileHandler, err := s.HandleSwaggerFile(swaggerUIHandler)
	if err != nil {
		panic(err)
	}

	app.Use(path.Join(s.BasePath, swaggerUI.Path), adaptor.HTTPHandler(swaggerUIHandler))
	app.Use(path.Join(s.BasePath, "swagger.json"), adaptor.HTTPHandler(swaggerFileHandler))
}
