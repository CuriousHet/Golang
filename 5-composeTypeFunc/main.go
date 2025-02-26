package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type TransformFunc func(string) string

type Server struct {
	transformFuncs []TransformFunc
}

// Apply all registered transformations sequentially
func (s *Server) handleRequest(filename string) error {
	newFilename := filename
	fmt.Println("Original filename:", filename)
	for _, transform := range s.transformFuncs {
		newFilename = transform(newFilename)
		fmt.Println("Transformed filename:", newFilename)
	}
	return nil
}

// Hash the filename using SHA-256
func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	return hex.EncodeToString(hash[:])
}

// Prefix the filename with a given string
func prefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

// Suffix the filename with a given string
func suffixFilename(suffix string) TransformFunc {
	return func(filename string) string {
		return filename + suffix
	}
}

// Trim spaces from filename
func trimSpaces(filename string) string {
	return strings.TrimSpace(filename)
}

// Append timestamp to filename
func appendTimestamp(filename string) string {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	return fmt.Sprintf("%s_%d%s", name, time.Now().Unix(), ext)
}

// Replace spaces with underscores
func replaceSpaces(filename string) string {
	return strings.ReplaceAll(filename, " ", "_")
}

func main() {
	s := &Server{
		transformFuncs: []TransformFunc{
			trimSpaces,
			replaceSpaces,
			prefixFilename("BOSS_"),
			suffixFilename("_FINAL"),
			appendTimestamp,
			hashFilename,
		},
	}

	s.handleRequest("Cool_pict.jpg")
}
