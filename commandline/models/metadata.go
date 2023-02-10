package models

import (
	"fmt"

	"github.com/kc-workspace/go-lib/mapper"
)

type Metadata struct {
	// Short name
	Short string
	// Full name, always existed
	Name string
	// Version
	Version string
	// Commit SHA
	Commit string
	// Build date
	Date string
	// Built by
	BuiltBy string
	// Help message
	Usage string
}

func (m *Metadata) String() string {
	return fmt.Sprintf("%s: %s (%s)", m.Name, m.Version, m.Commit)
}

func (m *Metadata) ToMapper() mapper.Mapper {
	return mapper.Mapper{
		"short":   m.Short,
		"name":    m.Name,
		"version": m.Version,
		"commit":  m.Commit,
		"date":    m.Date,
		"buildby": m.BuiltBy,
		"usage":   m.Usage,
	}
}

func EmptyMetadata() *Metadata {
	return &Metadata{
		Short:   "",
		Name:    "",
		Version: "",
		Commit:  "",
		Date:    "",
		BuiltBy: "",
		Usage:   "",
	}
}
