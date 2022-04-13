package rust

import (
	"opensca/internal/enum/language"
	"opensca/internal/filter"
	"opensca/internal/srt"
)

type Analyzer struct{}

func New() Analyzer {
	return Analyzer{}
}

// GetLanguage Get language of Analyzer
func (a Analyzer) GetLanguage() language.Type {
	return language.Rust
}

// CheckFile Check if it is a parsable file
func (a Analyzer) CheckFile(filename string) bool {
	return filter.RustCargoLock(filename)
}

// FilterFile filters the files that the current parser needs to parse
func (a Analyzer) FilterFile(dirRoot *srt.DirTree, depRoot *srt.DepTree) []*srt.FileData {
	files := []*srt.FileData{}
	for _, f := range dirRoot.Files {
		if a.CheckFile(f.Name) {
			files = append(files, f)
		}
	}
	return files
}

// ParseFile Parse the file
func (a Analyzer) ParseFile(dirRoot *srt.DirTree, depRoot *srt.DepTree, file *srt.FileData) []*srt.DepTree {
	if filter.RustCargoLock(file.Name) {
		return parseCargoLock(dirRoot, depRoot, file)
	}
	return []*srt.DepTree{}
}