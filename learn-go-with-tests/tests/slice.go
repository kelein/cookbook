package tests

import "github.com/pkg/errors"

func sum(number []int) int {
	count := 0
	for _, n := range number {
		count += n
	}
	return count
}

func sumAll(nums ...[]int) []int {
	sumarr := make([]int, len(nums))
	for i, num := range nums {
		sumarr[i] = sum(num)
	}
	return sumarr
}

func sumAllTails(nums ...[]int) []int {
	sums := make([]int, len(nums))
	for i, num := range nums {
		if len(num) > 0 {
			sums[i] = sum(num[1:])
		} else {
			sums[i] = 0
		}
	}
	return sums
}

func search(dict map[string]string, word string) string {
	return dict[word]
}

const (
	// ErrNotFound stands for word not found error
	ErrNotFound = DictErr("word not found in dict")

	// ErrWordExists stands for word exists error
	ErrWordExists = DictErr("word already exists")
)

// DictErr stands dict error type
type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

// Dict alias for map[string]string
type Dict map[string]string

// Search looks up a word in the dictionary
func (d Dict) Search(word string) (string, error) {
	result, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

// Add put word and value into dictionary
func (d Dict) Add(word, value string) error {
	d[word] = value
	return nil
}

// Update modify a exists word value
func (d Dict) Update(word, value string) error {
	_, err := d.Search(word)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}
	d[word] = value
	return nil
}

// Delete drop word from dictionary
func (d Dict) Delete(word string) {
	delete(d, word)
}
