package registry

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Register maintains a registry discovered golems
type Register interface {
	Persist(entry RegEntry) (err error)
	Retrieve(id uuid.UUID) (entry RegEntry, err error)
	RetrieveAll() (entries []RegEntry, err error)
}

// FileReg is a file-based registry of discovered golems
type FileReg struct {
	FilePath string
}

// An entry in the registry representing a discovered golem
type RegEntry struct {
	id         uuid.UUID
	uri        string
	discovered int64
}

// Persist a new RegEntry to the file system in a file named after the entry id.
// If an entry with the id already exists, the method returns without further action.
func (fr *FileReg) Persist(entry RegEntry) (err error) {
	path := filepath.Join(fr.FilePath, entry.id.String())
	if _, err := os.Stat(path); err == nil {
		// file exists
		return nil
	}

	return ioutil.WriteFile(path, entry.Bytes(), 0664)
}

// Retrieve a single RegEntry object from the registry by its unique ID
func (fr *FileReg) Retrieve(id uuid.UUID) (entry RegEntry, err error) {
	path := filepath.Join(fr.FilePath, id.String())
	if _, err := os.Stat(path); err == nil {
		return readFromFile(path)
	} else {
		return RegEntry{}, errors.New("no such file: " + path)
	}
}

// RetrieveAll gets all entries from the registry
func (fr *FileReg) RetrieveAll() (entries []RegEntry, err error) {
	files, err := ioutil.ReadDir(fr.FilePath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		path := filepath.Join(fr.FilePath, file.Name())
		entry, err := readFromFile(path)
		if err != nil {
			return entries, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// Bytes serializes a RegEntry object into a byte format
func (re *RegEntry) Bytes() []byte {
	return []byte(re.id.String() + "|" + re.uri + "|" + strconv.FormatInt(re.discovered, 10))
}

// FromBytes parses bytes to a RegEntry object.
// It assumes a specific format as created by the RegEntry's Bytes() function.
func FromBytes(bytes []byte) RegEntry {
	fullStr := string(bytes)
	s := strings.Split(fullStr, "|")
	if len(s) != 3 {
		panic("unreadable registry entry: " + fullStr)
	}

	discovered, err := strconv.ParseInt(s[2], 10, 64)
	if err != nil {
		panic("unreadable registry entry: " + fullStr)
	}

	return RegEntry{
		id:         uuid.MustParse(s[0]),
		uri:        s[1],
		discovered: discovered,
	}
}

func readFromFile(path string) (entry RegEntry, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return RegEntry{}, err
	}

	return FromBytes(bytes), nil
}
