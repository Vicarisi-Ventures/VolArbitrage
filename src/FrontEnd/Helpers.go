package FrontEnd

import (
	"io/ioutil"
	"path/filepath"
)

type Resource interface {
	Name() string
	Content() []byte
}

type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

func (r *StaticResource) Name() string {
	return r.StaticName
}

func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

func NewStaticResource(name string, content []byte) *StaticResource {

	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}

}

func LoadResourceFromPath(path string) (Resource, error) {

	bytes, err := ioutil.ReadFile(filepath.Clean(path))

	if err != nil {
		return nil, err
	}

	name := filepath.Base(path)
	return NewStaticResource(name, bytes), nil

}

func findTopFive(arr []float64) ([5]int, [5]float64) {

	var index [5]int
	var values [5]float64

	for i := 0; i < len(arr); i++ {

		for j := 0; j < 5; j++ {

			var isBreak bool

			// Begin Swap & Shift
			if arr[i] > values[j] {

				// Store Old Values
				old_index := index[j]
				old_value := values[j]

				// Update New Values
				index[j] = i
				values[j] = arr[i]

				// Shift Values
				for k := (j + 1); k < 4; k++ {

					index[k] = old_index
					old_index = index[k+1]

					values[k] = old_value
					old_value = values[k+1]

				}

				isBreak = true

			}

			if isBreak {
				break
			}

		}

	}

	return index, values

}

func findBottomFive(arr []float64) ([5]int, [5]float64) {

	var index [5]int
	var values [5]float64

	for i := 0; i < len(arr); i++ {

		for j := 0; j < 5; j++ {

			var isBreak bool

			// Begin Swap & Shift
			if arr[i] < values[j] {

				// Store Old Values
				old_index := index[j]
				old_value := values[j]

				// Update New Values
				index[j] = i
				values[j] = arr[i]

				// Shift Values
				for k := (j + 1); k < 4; k++ {

					index[k] = old_index
					old_index = index[k+1]

					values[k] = old_value
					old_value = values[k+1]

				}

				isBreak = true

			}

			if isBreak {
				break
			}

		}

	}

	return index, values

}
