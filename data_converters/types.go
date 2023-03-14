package main

import (
	"encoding/json"
	"fmt"
)

type (
	IRepository interface {
		Get() string
	}

	Repository struct {
		Repository IRepository `json:"repository"`
	}

	GitRepository struct {
		Url string `json:"url"`
	}

	FsRepository struct {
		Path string `json:"path"`
	}
)

func (g *GitRepository) Get() string {
	return g.Url
}

func (f *FsRepository) Get() string {
	return f.Path
}

// MarshalJSON serializes the helm chart to JSON
// It's required due to the interface type in the Repository field
func (h *Repository) MarshalJSON() ([]byte, error) {
	type Alias Repository
	j, err := json.Marshal(
		struct {
			*Alias
			RType string `json:"rtype"`
		}{
			Alias: (*Alias)(h),
			RType: fmt.Sprintf("%T", h.Repository),
		})

	fmt.Printf("MarshalJSON: %s\n", j)

	return j, err
}

// UnmarshalJSON deserializes the helm chart from JSON
func (h *Repository) UnmarshalJSON(data []byte) error {
	fmt.Printf("UnmarshalJSON %v\n", string(data))

	type Alias Repository
	aux := &struct {
		*Alias
		Repository json.RawMessage `json:"repository"`
		RType      string          `json:"rtype"`
	}{
		Alias: (*Alias)(h),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch aux.RType {
	case "*main.GitRepository":
		h.Repository = &GitRepository{}
	case "*main.FsRepository":
		h.Repository = &FsRepository{}
	default:
		return fmt.Errorf("unknown rtype: %s", aux.RType)
	}
	return json.Unmarshal(aux.Repository, &h.Repository)
}
