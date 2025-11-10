package controllers

import (
	"blackenshovel-service/config"
	"fmt"
	"io"
	"net/http"
)

func GetStaticMapHandler(w http.ResponseWriter, r *http.Request) {
	lon := r.URL.Query().Get("lon")
	lat := r.URL.Query().Get("lat")

	C := config.AppConfig

	if lon == "" || lat == "" {
		http.Error(w, "Missing 'lon' or 'lat' parameter", http.StatusBadRequest)
		return
	}

	const (
		zoom    = 17.5
		bearing = 0
		width   = 320
		height  = 240
	)

	url := fmt.Sprintf(
		"https://api.mapbox.com/styles/v1/mapbox/satellite-v9/static/%s,%s,%.1f,%d/%dx%d?access_token=%s",
		lon, lat, zoom, bearing, width, height, C.MapboxToken,
	)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch map", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Mapbox returned status %d", resp.StatusCode), resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, resp.Body)
}
