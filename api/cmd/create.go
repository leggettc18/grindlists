/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new database migration",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		dir = filepath.Clean(dir)
		fmt.Println(dir)
		ext := ".sql"
		startTime := time.Now()
		format := "unix"
		version, err := timeVersion(startTime, format)
		if err != nil {
			return err
		}
		versionGlob := filepath.Join(dir, version+"_"+ext)
		matches, err := filepath.Glob(versionGlob)
		if err != nil {
			return err
		}
		if len(matches) > 0 {
			return fmt.Errorf("duplicate migration version: %s", version)
		}
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		for _, direction := range []string{"up", "down"} {
			basename := fmt.Sprintf("%s_%s.%s%s", version, name, direction, ext)
			filename := filepath.Join(dir, basename)

			if err = createFile(filename); err != nil {
				return err
			}

			absPath, _ := filepath.Abs(filename)
			fmt.Println(absPath)
		}
		return nil
	},
}

func createFile(filename string) error {
	// create exclusive (fails if file already exists)
	// os.Create() specifies 0666 as the FileMode, so we're doing the same
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return err
	}

	return f.Close()
}

var errInvalidTimeFormat = errors.New("Time format may not be empty")

func timeVersion(startTime time.Time, format string) (version string, err error) {
	switch format {
	case "":
		err = errInvalidTimeFormat
	case "unix":
		version = strconv.FormatInt(startTime.Unix(), 10)
	case "unixNano":
		version = strconv.FormatInt(startTime.UnixNano(), 10)
	default:
		version = startTime.Format(format)
	}

	return
}

var name string

func init() {
	migrateCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.
	createCmd.Flags().StringVarP(&name, "name", "n", "", "name of the migration")
	createCmd.MarkFlagRequired("name")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
