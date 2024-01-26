package app

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strings"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/initsystem"
	"github.com/essentialkaos/ek/v12/options"
	"github.com/essentialkaos/ek/v12/sortutil"
	"github.com/essentialkaos/ek/v12/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "RDS Migrator"
	VER  = "1.1.0"
	DESC = "Utility for migrating Redis-Split metadata to RDS format"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Command line argument list
const (
	OPT_DRY_RUN  = "D:dry"
	OPT_CONVERT  = "C:convert"
	OPT_ROLE     = "r:role"
	OPT_NO_COLOR = "nc:no-color"
	OPT_HELP     = "h:help"
	OPT_VER      = "v:version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// optMap is map with options data
var optMap = options.Map{
	OPT_DRY_RUN:  {Type: options.BOOL},
	OPT_CONVERT:  {Type: options.BOOL},
	OPT_ROLE:     {},
	OPT_NO_COLOR: {Type: options.BOOL},
	OPT_HELP:     {Type: options.BOOL},
	OPT_VER:      {Type: options.BOOL},
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Run(gitRev string) {
	preConfigureUI()

	args, errList := options.Parse(optMap)

	if len(errList) != 0 {
		for _, err := range errList {
			printError(err.Error())
		}

		os.Exit(1)
	}

	configureUI()

	switch {
	case options.GetB(OPT_VER):
		genAbout(gitRev).Print()
		os.Exit(0)
	case options.GetB(OPT_HELP) || len(args) == 0:
		genUsage().Print()
		os.Exit(0)
	}

	if !options.GetB(OPT_DRY_RUN) && !options.GetB(OPT_CONVERT) {
		checkSyncDaemonStatus()
	}

	process(args.Get(0).Clean().String())
}

// preConfigureUI configure user interface
func preConfigureUI() {
	term := os.Getenv("TERM")

	fmtc.DisableColors = true

	if term != "" {
		switch {
		case strings.Contains(term, "xterm"),
			strings.Contains(term, "color"),
			term == "screen":
			fmtc.DisableColors = false
		}
	}

	// Check for output redirect using pipes
	if fsutil.IsCharacterDevice("/dev/stdin") &&
		!fsutil.IsCharacterDevice("/dev/stdout") &&
		os.Getenv("FAKETTY") == "" {
		fmtc.DisableColors = true
	}

	if os.Getenv("NO_COLOR") != "" {
		fmtc.DisableColors = true
	}
}

// configureUI configure user interface
func configureUI() {
	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}
}

// checkSyncDaemonStatus checks sync daemon status
func checkSyncDaemonStatus() {
	works, err := initsystem.IsWorks("rds-sync")

	if err != nil {
		printErrorAndExit("Can't check RDS Sync status: %v", err)
	}

	if works {
		printErrorAndExit("You must stop RDS Sync daemon before metadata conversion")
	}
}

// process starts data processing
func process(dir string) {
	dir = strings.TrimRight(dir, "/")

	checkMetaDir(dir)

	files := getMetaFiles(dir)

	backupMeta(dir, files)

	convertFiles(files)
}

// checkMetaDir checks meta data directory
func checkMetaDir(dir string) {
	err := fsutil.ValidatePerms("DRWX", dir)

	if err != nil {
		printErrorAndExit(err.Error())
	}

	if fsutil.IsEmptyDir(dir) {
		printErrorAndExit("Directory %s is empty", dir)
	}
}

// getMetaFiles returns list of meta files
func getMetaFiles(dir string) []string {
	files := fsutil.List(
		dir, true,
		fsutil.ListingFilter{
			NotMatchPatterns: []string{"*.*"},
		},
	)

	fsutil.ListToAbsolute(dir, files)
	sortutil.StringsNatural(files)

	return files
}

// backupMeta backups meta before migration
func backupMeta(dir string, files []string) {
	if options.GetB(OPT_DRY_RUN) {
		return
	}

	backupName := getBackupName()
	err := backup(backupName, dir, files)

	if err != nil {
		printErrorAndExit(err.Error())
	} else {
		fmtc.Printf("\n{s-}Backup archive created as %s{!}\n", backupName)
	}

	fmtc.NewLine()
}

// convertFiles convert all meta files
func convertFiles(files []string) {
	for _, file := range files {
		err := convert(file, options.GetB(OPT_DRY_RUN))
		printFileActionStatus(file, err)
	}
}

// printFileActionStatus prints file processing status
func printFileActionStatus(file string, err error) {
	if err == nil {
		fmtc.Printf("{g}✔ {!} %s\n", file)
	} else {
		fmtc.Printf("{r}✖ {!} %s\n", file, err)
	}
}

// printError prints error message to console
func printError(f string, a ...any) {
	fmtc.Fprintf(os.Stderr, "{r}"+f+"{!}\n", a...)
}

// printErrorAndExit print error message and exit with exit code 1
func printErrorAndExit(f string, a ...any) {
	printError(f, a...)
	os.Exit(1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genUsage generates usage info
func genUsage() *usage.Info {
	info := usage.NewInfo("", "dir")

	info.AddOption(OPT_DRY_RUN, "Dry run {s-}(do not convert anything){!}")
	info.AddOption(OPT_CONVERT, "Just convert meta {s-}(do not check anything){!}")
	info.AddOption(OPT_ROLE, "Overwrite instances role", "role")
	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show information about version")

	info.AddExample(
		"/opt/redis-split/meta",
		"Convert all metadata in /opt/redis-split/meta to the latest version",
	)

	return info
}

// genAbout generates basic info about app
func genAbout(gitRev string) *usage.About {
	about := &usage.About{
		App:     APP,
		Version: VER,
		Desc:    DESC,
		Year:    2009,
		License: "Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>",
		Owner:   "ESSENTIAL KAOS",
	}

	if gitRev != "" {
		about.Build = "git:" + gitRev
	}

	return about
}
