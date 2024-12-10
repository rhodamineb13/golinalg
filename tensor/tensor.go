package tensor

type Tensor struct {
	Values  []complex128
	Shapes  []int
	Strides []int
}

// Initializes a new n-rank Tensor object
func NewTensor[T int | int32 | int64 | complex64 | complex128 | float32 | float64](array []T, shape []int) *Tensor {
	// Create a slice to hold the complex128 values
	arrComp := make([]complex128, 0)

	// Convert each element of the array to complex128
	for _, v := range array {
		switch v := any(v).(type) {
		case int64:
			// Convert int64 to complex128 (real part is the int64 value, imaginary part is 0)
			arrComp = append(arrComp, complex(float64(v), 0))
		case float32:
			// Convert float32 to complex128
			arrComp = append(arrComp, complex(float64(v), 0))
		case float64:
			// Convert float64 to complex128
			arrComp = append(arrComp, complex(v, 0))
		case complex64:
			// Convert complex64 to complex128
			arrComp = append(arrComp, complex128(v))
		case complex128:
			// Already complex128, append as-is
			arrComp = append(arrComp, v)
		default:
			// This should never happen due to the constraints on T
			panic("Unsupported type")
		}
	}

	// Create the tensor object and return it
	return &Tensor{
		Values:  arrComp,
		Shapes:  shape,
		Strides: calculateStrides(shape),
	}
}

func calculateStrides(shape []int) []int {
	strides := make([]int, len(shape))
	stride := 1
	for i := len(shape) - 1; i >= 0; i-- {
		strides[i] = stride
		stride *= shape[i]
	}
	return strides
}
