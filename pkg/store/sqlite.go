package store

// SqliteStore is a store that uses bbolt as a backend.
// it implements the Store interface.
// This is more suitable for lightweight deployments where you need a single
// instance of KAJO Server.
type SqliteStore struct {
}
