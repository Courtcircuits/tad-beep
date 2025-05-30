package graphql

import (
	"context"
	nethttp "net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/application/graphql/graph"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"

	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type GraphQL interface {
	Register(app *fiber.App)
	Query() (*fiber.Ctx, error)
	Playground() (*fiber.Ctx, error)
}

type graphqlController struct {
	message_service core.MessageService
}

func wrapHandler(f func(nethttp.ResponseWriter, *nethttp.Request)) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(nethttp.HandlerFunc(f))(ctx.Context())
	}
}

func NewGraphQL(message_service core.MessageService) *graphqlController {
	return &graphqlController{
		message_service: message_service,
	}
}

func (g *graphqlController) Query() fiber.Handler {
	h := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			MessageService: g.message_service,
		},
	}))
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *nethttp.Request) bool {
				// TODO: check origin
				return true
			},
		},
	})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.Options{})

	// Add the introspection middleware.
	h.Use(extension.Introspection{})

	h.AroundFields(func(ctx context.Context, next graphql.Resolver) (res any, err error) {
		res, err = next(ctx)
		return res, err
	})
	return func(ctx *fiber.Ctx) error {
		wrapHandler(h.ServeHTTP)(ctx)
		return nil
	}

}

func (g *graphqlController) Playground() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		wrapHandler(playground.Handler("GraphQL playground", "/query"))(ctx)
		return nil
	}
}

func (g *graphqlController) Register(app *fiber.App) {
	app.Post("/graphql", g.Query())
	app.Get("/graphql", g.Query())
	app.Options("/graphql", g.Query())
	app.Post("/query", g.Query())
	app.Get("/query", g.Query())
	app.Options("/query", g.Query())

	app.Get("/playground", g.Playground())
}
