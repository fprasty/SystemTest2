package mddlwr

import (
	"systemtest2/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticateUser(c *fiber.Ctx) error {
	cookie := c.Cookies("Dyjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateAdmin(c *fiber.Ctx) error {
	cookie := c.Cookies("DyAdmin-jwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated admin",
		})
	}
	return c.Next()

}

func IsAuthenticateFo(c *fiber.Ctx) error {
	cookie := c.Cookies("FOjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateMR(c *fiber.Ctx) error {
	cookie := c.Cookies("MRjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateMRTI(c *fiber.Ctx) error {
	cookie := c.Cookies("MRTIjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateRS(c *fiber.Ctx) error {
	cookie := c.Cookies("RSjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateKKD(c *fiber.Ctx) error {
	cookie := c.Cookies("KKDjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateMLI(c *fiber.Ctx) error {
	cookie := c.Cookies("MLIjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateOC(c *fiber.Ctx) error {
	cookie := c.Cookies("OCjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateDO(c *fiber.Ctx) error {
	cookie := c.Cookies("DOjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateVilla(c *fiber.Ctx) error {
	cookie := c.Cookies("Villajwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateKR(c *fiber.Ctx) error {
	cookie := c.Cookies("KRjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticatePB(c *fiber.Ctx) error {
	cookie := c.Cookies("PBjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateTOS(c *fiber.Ctx) error {
	cookie := c.Cookies("TOSjwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}
