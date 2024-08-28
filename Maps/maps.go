package maps

type Dictionary map[string] string

func (d Dictionary) Search (key string) (string, error) {
	return d[key], nil
}

func (d Dictionary) Add (key, value string) {
	d[key] = value
}