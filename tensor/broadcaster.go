package tensor

import "math"

type Broadcaster func(complex128) complex128

type SecondTypeBroadcaster func(complex128, complex128) complex128

func Sin(a complex128) complex128 {
	return complex(math.Sin(real(a))*math.Cosh(imag(a)), math.Cos(real(a))*math.Sinh(imag(a)))
}

func Cos(a complex128) complex128 {
	return complex(math.Cos(real(a))*math.Cosh(imag(a)), -math.Sin(real(a))*math.Sinh(imag(a)))
}

func Tan(a complex128) complex128 {
	return Sin(a) / Cos(a)
}

func Sqrt(a complex128) complex128 {
	return complex(r(a)*math.Cos(arg(a)/2), r(a)*math.Sin(arg(a)/2))
}

func Pow(a, b complex128) complex128 {
	return 0
}

func r(a complex128) float64 {
	return math.Sqrt((real(a) * real(a)) + (imag(a) * imag(a)))
}

func arg(a complex128) float64 {
	return math.Tan(imag(a) / real(a))
}
