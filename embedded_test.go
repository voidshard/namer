package name

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var (
	testTagElf01   = "fantasy-elf01"
	testTagHuman01 = "fantasy-human01"
	testAllTags    = []string{testTagElf01, testTagHuman01}
)

func TestNewEmbeddedNamer(t *testing.T) {
	result, err := NewEmbeddedNamer()

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestTags(t *testing.T) {
	nmr, _ := NewEmbeddedNamer()
	sort.Strings(testAllTags)

	tags := nmr.Tags()
	sort.Strings(tags)

	assert.Equal(t, testAllTags, tags)
}

func TestCharacterNames(t *testing.T) {
	nmr, err := NewEmbeddedNamer()

	assert.Nil(t, err)

	for i, tag := range testAllTags {
		t.Run(fmt.Sprintf("%d %s", i, tag), func(t *testing.T) {
			mName, mSurname := nmr.Tag(tag).Male()
			assert.NotEqual(t, "", mName)
			assert.NotEqual(t, "", mSurname)

			fName, fSurname := nmr.Tag(tag).Female()
			assert.NotEqual(t, "", fName)
			assert.NotEqual(t, "", fSurname)

			mName, mSurname = nmr.Male()
			assert.NotEqual(t, "", mName)
			assert.NotEqual(t, "", mSurname)

			fName, fSurname = nmr.Female()
			assert.NotEqual(t, "", fName)
			assert.NotEqual(t, "", fSurname)
		})
	}

}
