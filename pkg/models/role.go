// models/role.go

package models

type Role int

const (
	Default Role = iota
	Standard
	Admin
)
