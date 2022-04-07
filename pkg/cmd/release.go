package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/epsxy/aurora/pkg/git"
	"github.com/epsxy/aurora/pkg/global"
	"github.com/spf13/cobra"
)

func releaseNewVersion(isMajor bool, isMinor bool, isFix bool) {
	// Steps
	// 1. get current version from VERSION file located at the root of the repo
	versionFilePath := filepath.Join(git.ShowTopLevel(), "VERSION")
	data, err := os.ReadFile(versionFilePath)
	if err != nil {
		log.Fatalf("your repo must have a VERSION file located at its root")
	}
	version := strings.Replace(string(data), "v", "", 1)
	// 2.1. Verify it's valid
	r := "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)$"
	isMatching, _ := regexp.MatchString(r, version)
	if !isMatching {
		log.Fatalf("version number in VERSION isn't valid; it doesn't match v1.2.3 format")
	}
	// 2.2 Increment it accordingly
	splitVersion := strings.Split(version, ".")
	major, _ := strconv.Atoi(splitVersion[0])
	minor, _ := strconv.Atoi(splitVersion[1])
	fix, _ := strconv.Atoi(splitVersion[2])
	if isMajor {
		major += 1
		minor = 0
		fix = 0
	} else if isMinor {
		minor += 1
		fix = 0
	} else if isFix {
		fix += 1
	}
	// 2.3 Save it
	newSemVer := fmt.Sprintf("v%d.%d.%d", major, minor, fix)
	if global.GetDryRun() {
		fmt.Printf("[dry-run] would increment version to: %s\n", newSemVer)
		return
	}
	err = os.WriteFile(versionFilePath, []byte(newSemVer), os.ModeExclusive)
	if err != nil {
		log.Fatalf("unable to save new version in file")
	}
	// 3. Create release commit
	// TODO: Add template in conf file for this type of commit?
	git.AddFile(versionFilePath)
	git.CreateCommit(fmt.Sprintf("chore: release %s", newSemVer))
}

var Major = &cobra.Command{
	Use:   "major",
	Short: "Major release",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			v = false
		}
		global.SetVerbose(v)

		d, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			d = false
		}
		global.SetDryRun(d)

		releaseNewVersion(true, false, false)
		return
	},
}

var Minor = &cobra.Command{
	Use:   "minor",
	Short: "Minor release",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			v = false
		}
		global.SetVerbose(v)

		d, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			d = false
		}
		global.SetDryRun(d)

		releaseNewVersion(false, true, false)
		return
	},
}

var Fix = &cobra.Command{
	Use:   "fix",
	Short: "Fix release",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			v = false
		}
		global.SetVerbose(v)

		d, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			d = false
		}
		global.SetDryRun(d)

		releaseNewVersion(false, false, true)
		return
	},
}

//Release is the release command entrypoint
var Release = &cobra.Command{
	Use:   "release",
	Short: "Handle releases",
	Long:  "Create a new release following (partially) the SemVer spec (vX.X.X format)",
}
