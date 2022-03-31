package server

import (
	"github.com/gofiber/fiber/v2"
)

type SearchKeywordQueryParams struct {
	Keyword string `validate:"required"`
}

func SearchKeyword(server *Server) Handler {
	return func(ctx *fiber.Ctx) error {
		params := new(SearchKeywordQueryParams)

		if err := ctx.QueryParser(params); err != nil {
			return BadRequest(ctx, err)
		}

		if err := server.validate.Struct(params); err != nil {
			return BadRequest(ctx, err)
		}

		response, err := server.kakao.LocalKeyword(params.Keyword)

		if err != nil {
			return InternalServerError(ctx, err)
		}

		return OK(ctx, response.Documents, response.Meta)
	}
}
