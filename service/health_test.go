package service

import (
	"log"
	"testing"
)

func TestHealthCheckDump(t *testing.T) {
	var hc HealthCheck
	hc.SetDefault()
	passtest := true
	if hc.CheckName != "defaultcheck" {
		passtest = false
	}
	if hc.CheckID != "defaultid" {
		passtest = false
	}
	if hc.TTL != 10 {
		passtest = false
	}
	if hc.Interval != 10 {
		passtest = false
	}
	if hc.Notes != "Health check Notes not given." {
		passtest = false
	}
	if passtest != true {
		t.Fatalf("SetDefault error")
	} else {
		log.Println("Healthcheck Setdefault success")
	}
}

func TestHealthCheckParseJSON(t *testing.T) {
	var hc HealthCheck
	hc.SetDefault()
	if res, err := hc.ParseJSON(); err == nil {
		t.Log(string(res))
		log.Println("test healthCheck parseJSON success")
	} else {
		t.Fatalf("test healthCheck parseJSON fail")
	}
}

func TestHealthCheck(t *testing.T) {
	var hc HealthCheck
	hc.SetDefault()
	if res, err := hc.Check(); err != nil {
		t.Fatalf(err.Error())
	} else {
		if res != PASS {
			t.Fatalf("default health check fail")
		} else {
			log.Println("Script health check pass")
		}
	}

	hc.TTL = 0
	hc.Script = "cd"
	hc.Interval = 10
	if res, err := hc.Check(); err != nil {
		t.Fatalf(err.Error())
	} else {
		if res != PASS {
			t.Fatalf("Script health check fail")
		} else {
			log.Println("Script health check pass")
		}
	}

	hc.TTL = 0
	hc.Script = ""
	hc.HTTP = "http://baidu.com"
	hc.Interval = 10
	if res, err := hc.Check(); err != nil {
		t.Fatalf(err.Error())
	} else {
		if res != PASS {
			t.Fatalf("HTTP health check fail")
		} else {
			log.Println("HTTP health check pass")
		}
	}

}
