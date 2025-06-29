package errors

import "errors"

var (
	// ErrNotFound adalah error yang kita gunakan ketika sebuah resource tidak ditemukan.
	// Ini adalah "sinyal" bisnis, bukan error teknis.
	ErrNotFound = errors.New("requested resource not found")

	// ErrInternal adalah error untuk semua kesalahan tak terduga dari sistem,
	// seperti koneksi database putus atau query yang salah.
	ErrInternal = errors.New("an unexpected internal error occurred")

)
