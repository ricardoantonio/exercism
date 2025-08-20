// Package twofer implements a simple function that returns a sharing phrase.
// Given a name, it returns a string in the form "One for <name>, one for me."
package twofer

// ShareWith returns a sharing phrase. If name is empty, it defaults to "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
