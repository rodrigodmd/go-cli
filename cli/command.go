package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Func func(arg []string)

type Handler interface {
	AddRoute(name, description string, fn Func)
	AddCommand(cmds ...*cobra.Command)
	SetDefault(fn Func)
	Run()
}

type handler struct {
	appname string
	rootCmd *cobra.Command
	fns     []Func
}

func New(appname, version string, fns ...Func) Handler {
	h := handler{
		appname: appname,
		rootCmd: &cobra.Command{
			Use:     appname,
			Version: version,
		},
		fns: fns,
	}
	h.Initialize()
	return &h
}

func (h *handler) Initialize() {
	h.rootCmd.Run = h.run
	cmds := []*cobra.Command{}

	cmds = append(cmds, &cobra.Command{
		Use:   "config [show]",
		Short: "Set basic config for " + h.appname,
		Long:  `config is used to prpmpt question needed for the basic configuration.`,
		Args:  cobra.MaximumNArgs(1),
		Run:   h.config,
	})

	//TODO add version
	//h.cmds = append(h.cmds, &cobra.Command{
	//	Use:   "version",
	//	Short: "Print the version number of " + commandNAme,
	//	Long:  `All software has versions. This is Hugo's`,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("v1.0")
	//	},
	//})

	h.AddCommand(cmds...)
}

func (h *handler) AddRoute(name, description string, fn Func) {
	command := &cobra.Command{
		Use:   name + " [param]",
		Short: description,
		Run: func(cmd *cobra.Command, args []string) {
			fn(args)
		},
	}
	h.AddCommand(command)
}

func (h *handler) AddCommand(cmds ...*cobra.Command) {
	h.rootCmd.AddCommand(cmds...)
}

func (h *handler) SetDefault(fn Func) {
	h.fns = []Func{fn}
}

func (h *handler) Run() {
	h.rootCmd.Execute()
}

/////////////////////////////////////////////
// Handler

func (h *handler) config(cmd *cobra.Command, args []string) {

	var confType string
	if len(args) == 0 {
		confType = "run"
	} else {
		confType = args[0]
	}

	switch confType {
	case "show":
		fmt.Println("TBD")
	default:
		fmt.Println("Available optinos TBD")
	}
}

func (h *handler) run(cmd *cobra.Command, args []string) {
	for _, fn := range h.fns {
		fn(args)
	}
}
