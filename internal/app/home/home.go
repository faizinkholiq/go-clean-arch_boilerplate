package home

import (
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
)

func getMacAddr() ([]string, error) {
    ifas, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    var as []string
    for _, ifa := range ifas {
        a := ifa.HardwareAddr.String()
        if a != "" {
            as = append(as, a)
        }
    }
    return as, nil
}

func Index(c *fiber.Ctx) error {
	result := make(map[string]interface{})
    // Mac Address
    as, err := getMacAddr()
    if err != nil {
        log.Fatal(err)
    }
    for _, a := range as {
        result["mac_address"] = a
	}
    result["user_agent"] = c.Get("User-Agent")
    result["ip_address"] = c.IP()
    result["message"] = "Hello, Home 👋!"

	return c.JSON(result)
}