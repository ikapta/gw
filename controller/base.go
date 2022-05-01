package controller

// import "github.com/gofiber/fiber/v2"

type ResultType[T any] struct {
  data T
  err error
  msg string
  meta map[string]interface{}
}

// func JSONResult[T any](c fiber.Ctx) func() error {
//   return func (data T) error {
//     return c.JSON(fiber.Map{...data})
//   }
// }