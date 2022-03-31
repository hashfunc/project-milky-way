package server

import "github.com/gofiber/fiber/v2"

type DefaultResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(ctx *fiber.Ctx, data, meta interface{}) error {
	return ctx.JSON(&DefaultResponse{
		Code:    "S200",
		Message: "OK",
		Meta:    meta,
		Data:    data,
	})
}

func BadRequest(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).
		JSON(&DefaultResponse{
			Code:    "F400",
			Message: "올바르지 않은 요청입니다.",
			Data:    err.Error(),
		})
}

func InternalServerError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).
		JSON(&DefaultResponse{
			Code:    "F500",
			Message: "서버 오류가 발생하였습니다.",
			Data:    err.Error(),
		})
}
