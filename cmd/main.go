package main

import (
	"github.com/a2htray/wlbdqm"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"strings"
)

var (
	debug              = false
	interval           = "12h"
	percentage float64 = 80
	to                 = "a2htray@outlook.com"
	env                = ""
)

var toList = make([]string, 0)

func RunTask() {
	output, err := wlbdqm.RunDiskQuota()
	if err != nil {
		panic(err)
	}

	dpOutput, err := wlbdqm.ParseDiskQuotaOutput(output)
	if err != nil {
		panic(err)
	}

	dPercentage := dpOutput.ByteOfBlocks() / dpOutput.ByteOfDQuota() * 100
	fPercentage := float64(dpOutput.NumOfFFiles()) / float64(dpOutput.NumOfFQuota()) * 100
	wlbdqm.DebugPrintln("dPercentage", dPercentage)
	wlbdqm.DebugPrintln("fPercentage", fPercentage)

	if dPercentage > percentage || fPercentage > percentage {
		message, err := wlbdqm.PrepareDiskInsufficientContent(wlbdqm.ContentItem{
			MaxPercentage: percentage,
			DPercentage:   dPercentage,
			FPercentage:   fPercentage,
			DPOutput:      dpOutput.HTMLTable(),
		})
		if err != nil {
			panic(err)
		}

		err = wlbdqm.SendEmail(toList, message)
		if err != nil {
			panic(err)
		}
		wlbdqm.InfoPrintln("email send")
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use: "wlbdqm [flags]",
		Run: func(cmd *cobra.Command, args []string) {
			wlbdqm.InfoPrintln("run command `wlbdqm`")

			if !debug {
				wlbdqm.SetAppMode(wlbdqm.AppModeProduction)
			}

			defer func() {
				if err := recover(); err != nil {
					wlbdqm.ErrorPrintln(err)
				}
			}()

			if env == "" {
				wlbdqm.ErrorPrintln("program needs an environment variables file, use -e [env file]")
				return
			}

			if err := godotenv.Load(env); err != nil {
				panic(err)
			}

			intervalStruct, err := wlbdqm.IntervalFromString(interval)
			if err != nil {
				panic(err)
			}
			wlbdqm.DebugPrintln("interval", intervalStruct)

			toList = strings.Split(to, "|")

			c := cron.New()
			entryID, err := c.AddFunc(intervalStruct.ToSpec(), RunTask)
			if err != nil {
				panic(err)
			}
			wlbdqm.InfoPrintln("task", entryID, "runs on", intervalStruct.ToSpec())

			c.Start()

			select {}
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", debug, "used in debug mode")
	rootCmd.PersistentFlags().StringVarP(&interval, "interval", "i", interval, "time interval")
	rootCmd.PersistentFlags().Float64VarP(&percentage, "percentage", "p", percentage, "exceeded percentage")
	rootCmd.PersistentFlags().StringVarP(&to, "to", "t", to, "a target email address or list, use separator \"|\"")
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", env, "environment variables file")

	if err := rootCmd.Execute(); err != nil {
		wlbdqm.ErrorPrintln(err.Error())
	}
}
