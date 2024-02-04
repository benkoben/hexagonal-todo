package repository

import (
    "time"
)

type RepositoryOptions struct {
    // Timeout sets the amount of time a query can take before context.WithTimeout times out.
    Timeout time.Duration
}
