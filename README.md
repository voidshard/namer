## Namer

Random names for various things.

### What

The `namer` lib randomly chooses and/or generates names for various things including characters, towns, rivers, places ..

A name can be chosen by tag (a type of name eg. Human, Goblin, French, Spanish etc) or entirely at random.

### How

```golang
import (
	"fmt"

	"github.com/voidshard/namer"
)

func main() {
	n, _ := namer.New()

        firstname, lastname := n.Female()
        fmt.Println(firstname, lastname)
        // ellena black

        fmt.Println(n.Tags())
        // ["fantasy-human01", "fantasy-elf01"]

        firstname, lastname = n.Tag("fantasy-elf01").Female()
        fmt.Println(firstname, lastname)
        // lusha leoceran
}
```

The current implementation reads data from an embedded data set to choose / generate names. Other implementations might be added in future!

### Contribution

Any contributions welcome. I'll be adding more name sets, types of things & whatnot as I need them myself. If you add any be sure to make a PR to share with the community.

### Credit

These names / rules for making names I've collected from various sources. I don't recall the sources exactly, but thought I'd mention.
