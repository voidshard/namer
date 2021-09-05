package name

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/rakyll/statik/fs"
	_ "github.com/voidshard/namer/internal/names/statik"
)

const (
	statikConfig = "/config.json"
)

type emConf struct {
	Name  string `json:"name"`
	Town  string `json:"town"`
	River string `json:"river"`
	Place string `json:"place"`
}

type nameList struct {
	Male    []string `json:"male"`
	Female  []string `json:"female"`
	Neutral []string `json:"neutral"`
	Surname []string `json:"surname"`
}

type townList struct {
	Prefix []string `json:"prefix"`
	Middle []string `json:"middle"`
	Suffix []string `json:"suffix"`
}

type riverList struct {
	Type   []string `json:"type"`
	Prefix []string `json:"prefix"`
	Suffix []string `json:"suffix"`
}

type placeList struct {
	Name []string `json:"name"`
}

// NewEmbeddedNamer reads data from an embedded filesystem that holds naming information.
func NewEmbeddedNamer() (Namer, error) {
	sfs, err := fs.New()
	if err != nil {
		return nil, err
	}

	enamer := &EmbeddedNamer{
		rng:       rand.New(rand.NewSource(time.Now().UnixNano())),
		filesys:   sfs,
		config:    map[string]*emConf{},
		tags:      []string{},
		towndata:  map[string]*townList{},
		namedata:  map[string]*nameList{},
		placedata: map[string][]string{},
		riverdata: map[string]*riverList{},
	}

	err = enamer.read(statikConfig, &enamer.config)
	if err != nil {
		return nil, err
	}

	for tag, paths := range enamer.config {
		pdata := &placeList{}
		err = enamer.read(paths.Place, pdata)
		if err != nil {
			return nil, fmt.Errorf("unable to read place data (%w): %s", err, tag)
		}

		rdata := &riverList{}
		err = enamer.read(paths.River, rdata)
		if err != nil {
			return nil, fmt.Errorf("unable to read river data (%w): %s", err, tag)
		}

		tdata := &townList{}
		err = enamer.read(paths.Town, tdata)
		if err != nil {
			return nil, fmt.Errorf("unable to read town data (%w): %s", err, tag)
		}

		ndata := &nameList{}
		err = enamer.read(paths.Name, ndata)
		if err != nil {
			return nil, fmt.Errorf("unable to read name data (%w): %s", err, tag)
		}

		enamer.towndata[tag] = tdata
		enamer.namedata[tag] = ndata
		enamer.riverdata[tag] = rdata
		enamer.placedata[tag] = pdata.Name
		enamer.tags = append(enamer.tags, tag)
	}

	return enamer, nil
}

// EmbeddedNamer is a Namer that used embedded name data
type EmbeddedNamer struct {
	rng     *rand.Rand
	filesys http.FileSystem
	config  map[string]*emConf

	tags      []string
	towndata  map[string]*townList
	namedata  map[string]*nameList
	placedata map[string][]string
	riverdata map[string]*riverList
}

// read & unmarshal object at the given path
func (e *EmbeddedNamer) read(path string, obj interface{}) error {
	f, err := e.filesys.Open(path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}

// Tags returns all knowns tags (categories of names)
func (e *EmbeddedNamer) Tags() []string {
	return e.tags
}

// randomChoice picks one of a list of things
func (e *EmbeddedNamer) randomChoice(in []string) string {
	length := len(in)
	if length == 0 {
		return ""
	}
	return in[e.rng.Intn(length-1)]
}

// Town randomly chooses a town name
func (e *EmbeddedNamer) Town() string {
	return e.town("")
}

// Male randomly chooses a male character name
func (e *EmbeddedNamer) Male() (string, string) {
	return e.male("")
}

// Female randomly chooses a female character name
func (e *EmbeddedNamer) Female() (string, string) {
	return e.female("")
}

// Place randomly chooses a place name
func (e *EmbeddedNamer) Place() string {
	return e.place("")
}

// River randomly chooses a river name / type
func (e *EmbeddedNamer) River() (string, string) {
	return e.river("")
}

// town randomly chooses / generates a town name
func (e *EmbeddedNamer) town(tag string) string {
	names, ok := e.towndata[tag]
	if tag == "" || !ok {
		return e.town(e.randomChoice(e.Tags()))
	}

	pre := e.randomChoice(names.Prefix)
	mid := e.randomChoice(names.Middle)
	suf := e.randomChoice(names.Suffix)

	if pre == mid || suf == mid {
		mid = ""
	}
	if pre == suf {
		suf = ""
	}

	return fmt.Sprintf("%s%s%s", pre, mid, suf)
}

// river randomly chooses / generates a river name
func (e *EmbeddedNamer) river(tag string) (string, string) {
	rivers, ok := e.riverdata[tag]
	if tag == "" || !ok {
		return e.river(e.randomChoice(e.Tags()))
	}

	pre := e.randomChoice(rivers.Prefix)
	suf := e.randomChoice(rivers.Suffix)
	typ := e.randomChoice(rivers.Type)

	return pre + suf, typ
}

// place randomly chooses / generates a place name
func (e *EmbeddedNamer) place(tag string) string {
	places, ok := e.placedata[tag]
	if tag == "" || !ok {
		return e.place(e.randomChoice(e.Tags()))
	}
	return e.randomChoice(places)
}

// male randomly chooses a male name & surname
func (e *EmbeddedNamer) male(tag string) (string, string) {
	return e.character(tag, true)
}

// female randomly chooses a female name & surname
func (e *EmbeddedNamer) female(tag string) (string, string) {
	return e.character(tag, false)
}

// character randomly chooses a character name
func (e *EmbeddedNamer) character(tag string, isMale bool) (string, string) {
	names, ok := e.namedata[tag]
	if tag == "" || !ok {
		return e.character(e.randomChoice(e.Tags()), isMale)
	}

	surname := e.randomChoice(names.Surname)
	if e.rng.Intn(100) > 95 {
		return e.randomChoice(names.Neutral), surname
	}

	if isMale {
		return e.randomChoice(names.Male), surname
	} else {
		return e.randomChoice(names.Female), surname
	}
}

// embeddedChooser
type embeddedChooser struct {
	parent *EmbeddedNamer
	tag    string
}

// Tag selects a name by the given tag
func (e *EmbeddedNamer) Tag(tag string) nameChooser {
	return &embeddedChooser{parent: e, tag: tag}
}

// Town randomly chooses a town name
func (e *embeddedChooser) Town() string {
	return e.parent.town(e.tag)
}

// Male randomly chooses a male character name
func (e *embeddedChooser) Male() (string, string) {
	return e.parent.male(e.tag)
}

// Female randomly chooses a female character name
func (e *embeddedChooser) Female() (string, string) {
	return e.parent.female(e.tag)
}

// Place randomly chooses a place name
func (e *embeddedChooser) Place() string {
	return e.parent.place(e.tag)
}

// River randomly chooses a river name / type
func (e *embeddedChooser) River() (string, string) {
	return e.parent.river(e.tag)
}
