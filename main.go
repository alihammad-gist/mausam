package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	externalip "github.com/glendc/go-external-ip"
)

// Please insert your token here
//
// used with ipinfo urls in the query string
// https://ipinfo.io/....?token=xxxxx
var IpinfoToken = ""

type geoLoc struct {
	City string
	Loc  string // "lat,long"
}

func main() {
	// flag to get only loc
	// used to open wttr.io/:lat,long
	// in terminal using curl / in the browser
	locOnly := flag.Bool("loc", false, "Output coords of host-ip's location")
	flag.Parse()
	g := ipToGeo(getIp())

	if *locOnly {
		// outputting only location and exiting
		// without errors
		fmt.Println(g.Loc)
		return
	}

	w := getFormattedWeather(g)
	withLabel := fmt.Sprintf(
		"%s: %s", g.City, w,
	)
	fmt.Printf(withLabel)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ipToGeo provides location details of provided ip
// using ipinfo.io/ip to get `geoIp` location

func ipToGeo(ip net.IP) geoLoc {
	url := fmt.Sprintf(
		"https://ipinfo.io/%s?token=%s",
		ip.String(),
		IpinfoToken,
	)
	res, err := http.Get(url)
	fatal(err)

	jsonStr, err := io.ReadAll(res.Body)
	fatal(err)

	g := geoLoc{}
	err = json.Unmarshal(jsonStr, &g)
	fatal(err)

	return g
}

// getip uses the externalip library which in turn uses
// multiple services sequentially, most probably until one
// successfully provides the public ip of the host.
func getIp() net.IP {
	consensus := externalip.DefaultConsensus(nil, nil)

	ip, err := consensus.ExternalIP()
	if err != nil {
		fatal(err)
	}

	return ip
}

// getFormattedWeather returns the response from wttr.in and
// appends the city name
func getFormattedWeather(g geoLoc) string {
	url := fmt.Sprintf(
		"https://wttr.in/%s?format=2",
		g.Loc,
	)

	res, err := http.Get(url)
	fatal(err)

	b, err := io.ReadAll(res.Body)
	fatal(err)

	return string(b)
}
