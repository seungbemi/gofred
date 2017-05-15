package gofred

import (
	"encoding/json"
	"fmt"
	"strings"
)

// IconInfo includes item's icon type and path
type IconInfo struct {
	Type string `json:"type,omitempty"` // Optional
	Path string `json:"path"`
}

// Item that will be shown as a result
type Item struct {
	Title        string   `json:"title"`    // Essential
	Subtitle     string   `json:"subtitle"` // Optional
	Icon         IconInfo `json:"icon"`
	Arg          string   `json:"arg"`            // Recommended
	Autocomplete string   `json:"autocomplete"`   // Recommended
	UID          string   `json:"uid,omitempty"`  // Optional
	Type         string   `json:"type,omitempty"` // Default = "default"
	Valid        bool     `json:"valid"`          // Default = true

}

type alfredWorkflow struct {
	VarMap map[string]string `json:"variables"`
}

// Response has all the items for showing on alfred
type Response struct {
	alfredWorkflow `json:"alfredworkflow,omitempty"`
	Items          []Item `json:"items,omitempty"`
}

// NewResponse returns a instance Response
func NewResponse() *Response {
	resp := &Response{}
	resp.VarMap = make(map[string]string)
	return resp
}

// AddVariable add a alfred environment value to pass
func (r *Response) AddVariable(key, value string) {
	r.VarMap[key] = value
}

// AddItem add a item to response
func (r *Response) AddItem(title, subtitle, iconType, iconPath, arg, autocomplete, uid, itemType string, valid bool) {
	r.Items = append(r.Items, Item{
		Title:    title,
		Subtitle: subtitle,
		Icon: IconInfo{
			Type: iconType,
			Path: iconPath,
		},
		Arg:          arg,
		Autocomplete: autocomplete,
		UID:          uid,
		Type:         itemType,
		Valid:        valid,
	})
}

// MatchCommand add a item if the title matches with given command
func (r *Response) MatchCommand(cmd, title, subtitle, iconType, iconPath, arg, autocomplete, uid, itemType string, valid bool) {
	if len(cmd) == 0 || strings.Contains(title, cmd) {
		r.AddItem(title, subtitle, iconPath, iconPath, arg, autocomplete, uid, itemType, valid)
	}
}

// AddMatchedListItem add an not runnable item which is matched with given command
func (r *Response) AddMatchedListItem(cmd, title, subtitle, iconPath, autoComplete string) {
	r.MatchCommand(cmd, title, subtitle, "", iconPath, "", autoComplete, title+"."+subtitle, "default", false)
}

// AddMatchedExecutableItem is simplified function of MatchCommand
func (r *Response) AddMatchedExecutableItem(cmd, title, subtitle, iconPath, autoComplete, arg string) {
	r.MatchCommand(cmd, title, subtitle, "", iconPath, arg, autoComplete, title+"."+subtitle, "default", true)
}

func (r *Response) String() string {
	bytes, err := json.Marshal(r)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return string(bytes)
}
