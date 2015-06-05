package gogitlab

import (
	"encoding/json"
)

type File struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Mode string `json:"mode,omitempty"`

	Children []*File
}

func (g *Gitlab) RepoTree(id, ref, path string) ([]*File, error) {

	url, opaque := g.ResourceUrlRaw(repo_url_tree, map[string]string{":id": id, ":ref": ref, ":path": path})

	var files []*File

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &files)
	}

	return files, err
}
