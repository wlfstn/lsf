package lsfDraw

type FileEntry struct {
	Name   string
	RawLen int
	IsDir  bool
}
type FileEntries []FileEntry
