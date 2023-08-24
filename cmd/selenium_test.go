package main

import (
	"github.com/tebeka/selenium"
	"testing"
	"time"
)

func TestSelenoid(t *testing.T) {
	caps := selenium.Capabilities{
		"name":             "myCoolTestName",
		"browserName":      "chrome",
		"browserVersion":   "chrome_104",
		"enableVNC":        true,
		"screenResolution": "1920x1080x24",
		"labels": map[string]interface{}{
			"manual": "true",
		},
		"sessionTimeout": "60m",
		"alwaysMatch": map[string]interface{}{
			"browserName":    "chrome",
			"browserVersion": "chrome_104",
			"selenoid:options": map[string]interface{}{
				"enableVNC":      true,
				"sessionTimeout": "60m",
				"labels": map[string]interface{}{
					"manual": "true",
				},
			},
		},
	}
	driver, err := selenium.NewRemote(caps, "http://175.178.161.110:8201/wd/hub")
	if err != nil {
		panic(err)
	}
	defer func() {
		driver.Close()
		driver.Quit()
	}()
	selenium.SetDebug(true)
	err = driver.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	windows, err := driver.WindowHandles()
	for _, w := range windows {
		err = driver.MaximizeWindow(w)
		if err != nil {
			panic(err)
		}
	}

	time.Sleep(10 * time.Second)
}
