package main

import (
	"fmt"
	"log"
)

type Route struct{}

// the first version return err and log( handle) the error
//
// when the log output it will output two lines
// if many process running concurrently, it will be hard to make sense of the log

func GetRoute1(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates1(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates")
		return Route{}, err
	}

	err = validateCoordinates1(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates")
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates1(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		log.Printf("invalid latitude: %f", lat)
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		log.Printf("invalid longitude: %f", lng)
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

// we should only handle the error OR return the error, not doing it multiple time

func GetRoute2(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{}, err
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates2(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

// rather than just return the err in GetRoute, we should wrap it
// to know where the error come from( src or dest)

func GetRoute3(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate source coordinates: %w", err)
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate target coordinates: %w", err)
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func getRoute(lat, lng, lat2, lng2 float32) (Route, error) {
	return Route{}, nil
}
