package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"unsafe"
)

var exactlyString = map[bool]string{false: "inexactly", true: "exactly"}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arguments: NUMBER")
		os.Exit(1)
	}
	r, ok := new(big.Rat).SetString(os.Args[1])
	if !ok {
		f, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Printf("Failed to parse NUMBER %v: %v\n", os.Args[1], err)
			os.Exit(1)
		}
		if r.SetFloat64(f) == nil {
			fmt.Println(os.Args[1], "is not finite")
			os.Exit(1)
		}
	}
	fmt.Println("input:", r.RatString())

	f64, e64 := r.Float64()
	fmt.Println("float64:", new(big.Rat).SetFloat64(f64))
	fmt.Println("float64: ", f64, exactlyString[e64])
	f64a, f64b := aroundFloat64(f64)
	fmt.Println("float64-:", f64a)
	fmt.Println("float64+:", f64b)

	f32, e32 := r.Float32()
	fmt.Println("float32:", new(big.Rat).SetFloat64(float64(f32)))
	fmt.Println("float32: ", f32, exactlyString[e32])
	f32a, f32b := aroundFloat32(f32)
	fmt.Println("float32-:", f32a)
	fmt.Println("float32+:", f32b)

	fmt.Println("32 as 64:", float64(f32))
}

func float32uint(x float32) uint32 { return *(*uint32)(unsafe.Pointer(&x)) }
func float64uint(x float64) uint64 { return *(*uint64)(unsafe.Pointer(&x)) }
func uint32float(x uint32) float32 { return *(*float32)(unsafe.Pointer(&x)) }
func uint64float(x uint64) float64 { return *(*float64)(unsafe.Pointer(&x)) }

func aroundFloat32(x float32) (float32, float32) {
	xi := float32uint(x)
	return uint32float(xi - 1), uint32float(xi + 1)
}

func aroundFloat64(x float64) (float64, float64) {
	xi := float64uint(x)
	return uint64float(xi - 1), uint64float(xi + 1)
}
