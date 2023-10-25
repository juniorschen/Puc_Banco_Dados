package controllers

import (
	"db.sampes.puc/entities"
	"db.sampes.puc/services"
	"github.com/gofiber/fiber/v2"
	"os"
)

func BuscarTodosRobozinhos(ctx *fiber.Ctx) error {
	result, err := services.GetAllRobozinhos(getRawQuery(ctx))

	if err != nil {
		/*if errors.Is(err, context.DeadlineExceeded) || os.IsTimeout(err) {
			return fiber.NewError(fiber.StatusTooManyRequests, err.Error())
		}*/
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(result)
}

func BuscarUmRobozinho(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	result, err := services.GetRobozinho(uint(id), getRawQuery(ctx))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || os.IsTimeout(err) {
			return fiber.NewError(fiber.StatusTooManyRequests, "Atenção não.....")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(result)
}

func IncluirUmRobozinho(ctx *fiber.Ctx) error {
	var body entities.Robozinho
	ctx.BodyParser(&body)

	err := services.CreateRobozinho(body, getRawQuery(ctx))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(nil)
}

func AtualizarUmRobozinho(ctx *fiber.Ctx) error {
	var body entities.Robozinho
	ctx.BodyParser(&body)

	err := services.UpdateRobozinho(body, getRawQuery(ctx))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(nil)
}

func RemoverUmRobozinho(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)
	err := services.DeleteRobozinho(uint(id), getRawQuery(ctx))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(nil)
}

func getRawQuery(ctx *fiber.Ctx) bool {
	return ctx.QueryBool("Raw", false)
}
