package arraylist

type List struct {
	elements []interface{}
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink where size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value at the end of the list.
func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array is not empty, otherwise false
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

// Indexof returns index of provided element
func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}

	for index, element := range list.elements {
		if element == value {
			return index
		}
	}
	return -1
}

// Remove removes the element at the given index from the list
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

// Empty returns true if list does not contain any elements.
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *List) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// Sort sorts values (in-place) using.
func (list *List) Sort() {

}

// Swap swaps the two values at the specified positions.
func (list *List) Swap() {

}

func (list *List) Insert() {

}

func (list *List) Set() {

}

func (list *List) String() {

}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *List) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(shrinkFactor*float32(currentCapacity)) {
		list.resize(list.size)
	}
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Check that the index is within bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
