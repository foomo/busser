package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/foomo/busser/table"
	"github.com/foomo/busser/table/validation"
)

type fs struct {
	root string
}

func NewFS(root string) (Store, error) {
	return &fs{root: root}, nil
}

const indexSuffix = "-index.json"

func (fs *fs) getNames(id table.ID, version table.Version) (tableName, validationTableName, indexName string) {
	basename := filepath.Join(fs.root, string(id)+"-"+string(version))
	return basename + "-table.json", basename + "-validation-table.json", basename + indexSuffix
}

func (fs *fs) store(filename string, v interface{}) error {
	jsonBytes, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, jsonBytes, 0644)
}

func (fs *fs) load(filename string, v interface{}) error {
	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, &v)
}

func (fs *fs) Add(t *table.Table, vt *validation.Table) error {
	tName, vtName, indexName := fs.getNames(t.ID, t.Version)
	err := fs.store(tName, t)
	if err != nil {
		return err
	}
	err = fs.store(vtName, vt)
	if err != nil {
		// cleanup
		os.Remove(tName)
		return err
	}
	err = fs.store(indexName, &table.TableSummary{
		ID:        t.ID,
		Version:   t.Version,
		Timestamp: t.Timestamp,
		Valid:     vt.Valid,
	})
	if err != nil {
		// cleanup
		os.Remove(tName)
		os.Remove(vtName)
		return err
	}
	return nil
}

func (fs *fs) getCommitFileName(id table.ID) string {
	return filepath.Join(fs.root, "commit-"+string(id))
}

func (fs *fs) getCommit(id table.ID) (table.Version, error) {
	commitBytes, err := ioutil.ReadFile(fs.getCommitFileName(id))
	if err == os.ErrNotExist {
		return "", nil
	}
	return table.Version(commitBytes), nil
}

func (fs *fs) GetVersion(id table.ID, version table.Version) (t *table.Table, vt *validation.Table, err error) {
	tName, vtName, _ := fs.getNames(id, version)
	err = fs.load(tName, &t)
	if err != nil {
		return nil, nil, err
	}
	return t, vt, fs.load(vtName, &vt)
}

func (fs *fs) Commit(id table.ID, version table.Version) error {
	return ioutil.WriteFile(fs.getCommitFileName(id), []byte(version), 0644)
}

func (fs *fs) GetCommitted(id table.ID) (t *table.Table, vt *validation.Table, err error) {
	version, err := fs.getCommit(id)
	if err != nil {
		return nil, nil, err
	}
	return fs.GetVersion(id, version)
}

func (fs *fs) List() (list table.List, err error) {
	indexFiles := []string{}
	err = filepath.Walk(fs.root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, indexSuffix) {
			indexFiles = append(indexFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	commits := map[table.ID]table.Version{}
	for _, file := range indexFiles {
		index := &table.TableSummary{}
		err := fs.load(file, &index)
		if err != nil {
			return nil, err
		}
		commits[index.ID] = ""
		list = append(list, *index)
	}
	for tableID := range commits {
		commited, err := fs.getCommit(tableID)
		if err != nil {
			return nil, err
		}
		if commited != "" {
			for i, index := range list {
				list[i].Committed = index.ID == tableID && index.Version == commited
			}
		}
	}
	sort.Sort(list)
	return list, nil
}

func (fs *fs) Delete(id table.ID, version table.Version) error {
	committedVersion, err := fs.getCommit(id)
	if err != nil {
		return err
	}
	if committedVersion != "" {
		err := os.Remove(fs.getCommitFileName(id))
		if err != nil {
			return err
		}
	}
	tName, vtName, indexName := fs.getNames(id, version)
	for _, n := range []string{tName, vtName, indexName} {
		err := os.Remove(n)
		if err != nil {
			return err
		}
	}

	return nil
}
