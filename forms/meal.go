package forms

import (
)

type MeatForm struct {
    Id    int         `form:"-"`
    Foods  []string
    LactosePresence  bool
}