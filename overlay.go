package pe

import (
	"errors"
	"io"
)

// error
var (
	ErrNoOverlayFound = errors.New("pe: not have overlay data")
)

// NewOverlayReader returns a new ReadSeeker reading the PE overlay data.
func (pe *File) NewOverlayReader() (*io.SectionReader, error) {
	if pe.data == nil {
		return nil, errors.New("pe: file reader is nil")
	}
	return io.NewSectionReader(pe.f, pe.OverlayOffset, 1<<63-1), nil
}

// Overlay returns the overlay of the PE file.
func (pe *File) Overlay() ([]byte, error) {
	sr, err := pe.NewOverlayReader()
	if err != nil {
		return nil, err
	}

	overlay := make([]byte, int64(pe.size)-pe.OverlayOffset)
	n, err := sr.ReadAt(overlay, 0)
	if n == len(overlay) {
		err = nil
	}
	return overlay, err
}

func (pe *File) OverlayLength() int64 {
	return int64(pe.size) - pe.OverlayOffset
}
