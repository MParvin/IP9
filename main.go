package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/service"
)

const (
	defaultV4DB = "data/ip2region_v4.xdb"
	defaultV6DB = "data/ip2region_v6.xdb"
)

func main() {
	v4DB := envOrDefault("IP2REGION_V4_DB", defaultV4DB)
	v6DB := envOrDefault("IP2REGION_V6_DB", defaultV6DB)

	geo, err := newGeoLookup(v4DB, v6DB)
	if err != nil {
		log.Fatalf("geo lookup init: %v", err)
	}
	defer geo.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		userIP := c.ClientIP()
		country, code := geo.lookup(userIP)
		if code != "" {
			c.String(http.StatusOK, "%s\n%s (%s)\n", userIP, country, code)
			return
		}
		c.String(http.StatusOK, "%s\n", userIP)
	})

	r.Run()
}

type geoLookup struct {
	ip2region *service.Ip2Region
}

func newGeoLookup(v4DB, v6DB string) (*geoLookup, error) {
	var v4Config *service.Config
	if _, err := os.Stat(v4DB); err == nil {
		cfg, err := service.NewV4Config(service.VIndexCache, v4DB, 20)
		if err != nil {
			return nil, fmt.Errorf("v4 config: %w", err)
		}
		v4Config = cfg
	}

	var v6Config *service.Config
	if _, err := os.Stat(v6DB); err == nil {
		cfg, err := service.NewV6Config(service.VIndexCache, v6DB, 20)
		if err != nil {
			return nil, fmt.Errorf("v6 config: %w", err)
		}
		v6Config = cfg
	}

	if v4Config == nil && v6Config == nil {
		return nil, fmt.Errorf("no ip2region database found (checked %q and %q)", v4DB, v6DB)
	}

	ip2region, err := service.NewIp2Region(v4Config, v6Config)
	if err != nil {
		return nil, fmt.Errorf("ip2region service: %w", err)
	}

	return &geoLookup{ip2region: ip2region}, nil
}

func (g *geoLookup) Close() {
	g.ip2region.Close()
}

func (g *geoLookup) lookup(ip string) (country, code string) {
	region, err := g.ip2region.Search(ip)
	if err != nil || region == "" {
		return "", ""
	}

	parts := strings.Split(region, "|")
	if len(parts) < 5 {
		return "", ""
	}

	return parts[0], parts[4]
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
