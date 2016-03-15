package core

type IEquatable interface {
	//
	// Summary:
	//     Indicates whether the current object is equal to another object of the same type.
	//
	// Parameters:
	//   other:
	//     An object to compare with this object.
	//
	// Returns:
	//     true if the current object is equal to the other parameter; otherwise, false.
	Equals(other interface{}) bool
}
