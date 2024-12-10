package tensor

import "fmt"

func AddTwoTensors(A *Tensor, B *Tensor) *Tensor {
	if !checkSameShape(A.Shapes, B.Shapes) {
		panic("tensors must have the same shape")
	}

	m := NewTensor(make([]complex128, len(A.Values)), A.Shapes)

	for i := range A.Values {
		m.Values[i] = A.Values[i] + B.Values[i]
	}

	return m
}

func ElementwiseMultiplication(A *Tensor, B *Tensor) *Tensor {
	if !checkSameShape(A.Shapes, B.Shapes) {
		panic("tensors must have the same shape")
	}

	m := NewTensor(make([]complex128, len(A.Values)), A.Shapes)

	for i := range A.Values {
		m.Values[i] = A.Values[i] * B.Values[i]
	}

	return m
}

func SubtractTensors(A *Tensor, B *Tensor) *Tensor {
	if !checkSameShape(A.Shapes, B.Shapes) {
		panic("tensors must have the same shape")
	}

	m := NewTensor(make([]complex128, len(A.Values)), A.Shapes)

	for i := range A.Values {
		m.Values[i] = A.Values[i] - B.Values[i]
	}

	return m
}

func ElementwiseDivision(A *Tensor, B *Tensor) *Tensor {
	if !checkSameShape(A.Shapes, B.Shapes) {
		panic("tensors must have the same shape")
	}

	m := NewTensor(make([]complex128, len(A.Values)), A.Shapes)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC recovered: %v", r)
		}
	}()

	for i := range A.Values {
		m.Values[i] = A.Values[i] / B.Values[i]
	}

	return m
}

func BroadcastFunction[T Broadcaster | SecondTypeBroadcaster](A *Tensor, f T, secondParameter *complex128) *Tensor {
	m := make([]complex128, len(A.Values))
	for i := range A.Values {
		switch f := any(f).(type) {
		case Broadcaster:
			m[i] = f(m[i])
		case SecondTypeBroadcaster:
			if secondParameter == nil {
				panic("second parameter must exist when having a func(complex128)complex128 broadcasting function")
			}
			m[i] = f(m[i], *secondParameter)
		}
	}

	return NewTensor(m, A.Shapes)
}

func checkSameShape(aShape []int, bShape []int) bool {
	if len(aShape) != len(bShape) {
		return false
	}

	for i := range aShape {
		if aShape[i] != bShape[i] {
			return false
		}
	}

	return true
}
