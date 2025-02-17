package robber

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"path/filepath"
)

type Flags struct {
	Org       *string
	User      *string
	Repo      *string
	Rules     *os.File
	Context   *int
	Entropy   *bool
	Both      *bool
	NoContext *bool
	Debug     *bool
}

func ParseFlags() *Flags {
	parser := argparse.NewParser("yar", "Sail ye seas of git for booty is to be found")
	flags := &Flags{
		Org: parser.String("o", "org", &argparse.Options{
			Required: false,
			Help:     "Organization to plunder",
			Default:  "",
		}),

		User: parser.String("u", "user", &argparse.Options{
			Required: false,
			Help:     "User to plunder",
			Default:  "",
		}),

		Repo: parser.String("r", "repo", &argparse.Options{
			Required: false,
			Help:     "Repository to plunder",
			Default:  "",
		}),

		Rules: parser.File("", "rules", os.O_RDONLY, 0600, &argparse.Options{
			Required: false,
			Help:     "JSON file containing regex rulesets",
			Default:  filepath.Join(GetGoPath(), "src", "github.com", "Furduhlutur", "yar", "rules.json"),
		}),

		Context: parser.Int("c", "context", &argparse.Options{
			Required: false,
			Help:     "Show N number of lines for context",
			Default:  2,
		}),

		Entropy: parser.Flag("e", "entropy", &argparse.Options{
			Required: false,
			Help:     "Search for secrets using entropy analysis",
			Default:  false,
		}),

		// Overrides entropy flag
		Both: parser.Flag("b", "both", &argparse.Options{
			Required: false,
			Help:     "Search by using both regex and entropy analysis. Overrides entropy flag",
			Default:  false,
		}),

		// Overrides context flag
		NoContext: parser.Flag("n", "no-context", &argparse.Options{
			Required: false,
			Help:     "Only show the secret itself, similar to trufflehog's regex output. Overrides context flag",
			Default:  false,
		}),

		Debug: parser.Flag("d", "debug", &argparse.Options{
			Required: false,
			Default:  false,
		}),
	}

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	validateFlags(flags, parser)
	return flags
}

func validateFlags(flags *Flags, parser *argparse.Parser) {
	if *flags.User == "" && *flags.Repo == "" && *flags.Org == "" {
		fmt.Print(parser.Usage("Must give atleast one of org/user/repo"))
		os.Exit(1)
	}
}
