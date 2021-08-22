package handler

import (
	"testing/services"

	"github.com/gofiber/fiber/v2"
)

// handler tempatnya validasi

func HandlerStudentDetail(c *fiber.Ctx) error {

	// ada endpoint login, login ngirim username and password

	// di check/validasi username and passwordnya itu kosong apa nggak

	// endpoint getDetail
	// /siswa/get-detail/:siswaID or
	// /siswa/get-detail?siswaID=10

	// validasi si ID-nya
	// - kosong apa nggak
	// - kalau dia int dia nilai 0 apa nggak
	// return error

	services.ServicesGetStudent()

	return c.SendString("GetDetail Studen")
}
