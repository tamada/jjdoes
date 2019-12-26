package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/tjdoe"
)

/*
VERSION represents the version of ttt.
*/
const VERSION = "1.0.0"

type options struct {
	dest        string
	mapping     string
	seed        string
	helpFlag    bool
	versionFlag bool
}

func getVersionMessage(prog string) string {
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

func getHelpMessage(prog string) string {
	return fmt.Sprintf(`%s [OPTIONS] <SCORES.CSV> <ROOT_DIR>
OPTIONS
    -d, --dest <DIR>       specifies destination of anonymized programs.
                           if this option was not specified, output to 'dest' directory.
    -s, --score <SCORE>    specifies destination of score file. default is 'anonymized_score.csv'
    -h, --help             print this message and exit.
    -v, --version          print version and exit.
SCORES.CSV
    shows score file, the first row is the header, and following rows
    represent each student, and must be formatted as follows.  The
    first column is id, the second column shows the name, and the
    following columns represent the scores of assignments.
ROOT_DIR
    the directory contains the programs.  The layout of the directory is arbitrary.
    The user arbitrary defines the names of sub-directories and files.`, prog)
}

func buildFlagSet() (*flag.FlagSet, *options) {
	opts := new(options)
	flags := flag.NewFlagSet("tjdoe", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(getHelpMessage("ttt")) }
	flags.StringVarP(&opts.dest, "dest", "d", "dest", "specifies destination of anonymized programs")
	flags.StringVarP(&opts.mapping, "score", "s", "anonymized_score.csv", "specifies the destination of anonymized score file.")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this message")
	flags.BoolVarP(&opts.versionFlag, "version", "v", false, "print version")
	return flags, opts
}

func (opts *options) generateSeed() int64 {
	if opts.seed == "" {
		seed, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		return seed.Int64()
	}
	seed := int64(0)
	for _, rune := range opts.seed {
		seed = seed + int64(rune)
	}
	return seed
}

func parseArgs(args []string) (*options, []string, error) {
	flags, opts := buildFlagSet()
	if err := flags.Parse(args); err != nil {
		return nil, nil, err
	}
	newArgs := flags.Args()[1:]
	if len(newArgs) != 2 {
		return opts, newArgs, fmt.Errorf("scores.csv or rootdir does not specified")
	}
	return opts, newArgs, nil
}

func printVersionAndOrHelp(prog string, opts *options) int {
	if opts.versionFlag {
		fmt.Println(getVersionMessage(prog))
	}
	if opts.helpFlag {
		fmt.Println(getHelpMessage(prog))
	}
	return 0
}

func outputAnonymizedScores(tjdoe *tjdoe.TJDoe, students []*tjdoe.Student, dest string) int {
	file, err := os.OpenFile(dest, os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return 3
	}
	defer file.Close()
	err = tjdoe.OutputAnonymizedScores(students, file)
	if err != nil {
		fmt.Println(err.Error())
		return 4
	}
	return 0
}

func perform(opts *options, args []string) int {
	tjdoe := tjdoe.New(opts.generateSeed())
	students, err := tjdoe.BuildScores(args[1:])
	if err != nil {
		return 2
	}
	mapping := tjdoe.BuildMappings(students)
	tjdoe.AnonymizeDirectory(args[0], opts.dest, mapping)
	return outputAnonymizedScores(tjdoe, students, opts.mapping)
}

func goMain(args []string) int {
	opts, newArgs, err := parseArgs(args)
	if err != nil {
		fmt.Println(getHelpMessage(args[0]))
		return 1
	}
	if opts.helpFlag || opts.versionFlag {
		return printVersionAndOrHelp(args[0], opts)
	}
	return perform(opts, newArgs)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
