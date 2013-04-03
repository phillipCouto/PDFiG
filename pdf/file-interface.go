package pdf

type ObjectNumber struct {
	number     uint32
	generation uint16
}

type File interface {
	// AddObjectAt() adds the object to the File using the
	// ObjectNumber obtained by an earlier call to
	// ReserveObjectNumber()
	AddObjectAt(ObjectNumber, Object)

	// AddObject() adds the passed object to the File.  The
	// returned indirect reference may be used for backward
	// references to the object.
	AddObject(object Object) (reference* Indirect)

	// ReserveObjectNumber() reserves a position (ObjectNumber)
	// for the passed object in the File.
	ReserveObjectNumber(Object) ObjectNumber

	// DeleteObject() deletes the specified object from the file.
	DeleteObject(ObjectNumber)

	// Close() writes the xref, trailer, etc., and closes the underlying file.
	Close()
}
