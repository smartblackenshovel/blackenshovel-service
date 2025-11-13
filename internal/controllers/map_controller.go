package controllers

import (
	"blackenshovel-service/config"
	"bytes"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
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
		zoom    = 18.5
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
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Mapbox returned status %d: %s", resp.StatusCode, string(body)), resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	// Decode image (supports PNG or JPEG)
	srcImg, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode image: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert to RGBA for consistent RGB access
	bounds := srcImg.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	draw.Draw(rgbaImg, bounds, srcImg, bounds.Min, draw.Src)

	// Convert to RGB565
	rgb565 := make([]byte, bounds.Dx()*bounds.Dy()*2)
	i := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := rgbaImg.RGBAAt(x, y)
			r := uint16(c.R) >> 3
			g := uint16(c.G) >> 2
			b := uint16(c.B) >> 3
			val := (r << 11) | (g << 5) | b
			rgb565[i] = byte(val >> 8)
			rgb565[i+1] = byte(val & 0xFF)
			i += 2
		}
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("X-Image-Width", fmt.Sprint(bounds.Dx()))
	w.Header().Set("X-Image-Height", fmt.Sprint(bounds.Dy()))
	w.Write(rgb565)
}
