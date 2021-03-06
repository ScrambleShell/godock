// Package dock wraps all backends and clients to start a dock.
package dock

import (
	// Globals and extras.
	"github.com/sqp/godock/libs/cdglobal"       // Global consts.
	"github.com/sqp/godock/libs/dock/confown"   // New dock own settings.
	"github.com/sqp/godock/libs/packages/build" // Pkgbuild config.

	// Dock frontend.
	"github.com/sqp/godock/libs/dock/eventmouse" // Mouse events callbacks.
	"github.com/sqp/godock/libs/dock/guibridge"  // GUI interface.
	"github.com/sqp/godock/libs/dock/maindock"   // Dock settings.
	"github.com/sqp/godock/libs/dock/mainmenu"   // Build menu callbacks.

	// Dock backend.
	"github.com/sqp/godock/libs/cdtype"           // Logger type.
	"github.com/sqp/godock/libs/gldi"             // Gldi access.
	"github.com/sqp/godock/libs/gldi/backendgui"  // GUI callbacks.
	"github.com/sqp/godock/libs/gldi/backendmenu" // Menu items.
	"github.com/sqp/godock/libs/gldi/globals"     // Dock globals.
	"github.com/sqp/godock/libs/gldi/mgrgldi"     // Internal go applets service.

	// Register applets services.
	_ "github.com/sqp/godock/services/allapps"

	// Other services.
	"github.com/sqp/godock/libs/net/websrv"       // Web server.
	"github.com/sqp/godock/libs/srvdbus"          // DBus own service.
	"github.com/sqp/godock/libs/srvdbus/dockpath" // hack dock dbus path

	// Help.
	"github.com/sqp/godock/libs/text/strhelp"  // String helpers.
	"github.com/sqp/godock/libs/text/versions" // Print API version.

	"errors"
	"fmt"
)

var (
	// CustomHacks defines developer optional custom settings launched during init.
	CustomHacks = func() {}
)

// Run starts dock routines and locks the main thread with gtk.
//
// It wraps all backends and clients to start a dock :
// Dbus server, applets manager, GUI, menu and gldi.
//
// Dbus service is enabled by default. Mandatory if enabled, to provide remote control.
// This will prevent launching a 2nd instance without the disable Dbus service option.
//
// You can add custom changes, launched before the start, with CustomHacks.
//
// Run returns true if the dock is able to start. This can be done with:
//   gldi.Lock()      // alias for gtk_main.
//   maindock.Clean() // may be better with defer, but cause confused panic messages.
//
func Run(log cdtype.Logger, getSettings func() maindock.DockSettings) bool {
	settings := getSettings()

	// Logger debug state.
	log.SetDebug(settings.Debug || confown.Current.OnStartDebug)
	maindock.SetLogger(log)

	// Dock init.
	settings.Init()

	// Load new config settings. New options are in an other file to keep the
	// original config file as compatible with the real dock as possible.
	file, e := globals.DirUserAppData(confown.GuiFilename)
	confown.Init(log, file, e)

	// Build settings.
	file, e = globals.DirUserAppData(cdglobal.FileBuildSource)
	build.Init(log, file, e)

	// Register go internal applets events.
	appmgr := mgrgldi.Register(log)

	// Start the polling loop for go internal applets (can be in DBus with other events).
	if settings.DisableDBus {
		log.GoTry(appmgr.StartLoop)

	} else {
		// DBus service is mandatory if enabled. This prevent double launch.
		dbus, e := serviceDbus(log)
		if log.Err(e, "start dbus service") {
			fmt.Println("restart the program with the -N flag if you really need a second instance")
			return false
		}
		dbus.SetManager(appmgr)
		log.GoTry(dbus.StartLoop)
	}

	// HTTP listener for the pprof debug.
	websrv.Init(log)
	websrv.Service.Host = settings.WebHost
	websrv.Service.Port = settings.WebPort

	if settings.WebMonitor || confown.Current.OnStartWebMon {
		websrv.Service.SetMonitored(true)
	}

	// Print useful packages versions.
	gtkX, gtkY, gtkZ := globals.GtkVersion()
	versions.Dock = []versions.Field{
		{K: "OpenGL", V: strhelp.YesNo(gldi.GLBackendIsUsed())},
		{K: "gldi", V: globals.Version()},
		{K: "GTK", V: fmt.Sprintf("%d.%d.%d", gtkX, gtkY, gtkZ)},
	}
	versions.SetName(cdglobal.AppName)
	versions.Print()

	// Custom calls added by devs for their own uses and tests.
	CustomHacks()

	// Register GUI events.
	backendgui.Register(guibridge.New(log))

	// Register mouse events.
	eventmouse.Register(log)

	// Register menus events.
	backendmenu.Register(log, mainmenu.BuildMenuContainer, mainmenu.BuildMenuIcon)

	// Finish startup.
	settings.Start()

	return true
}

// Start Loader.
//
func serviceDbus(log cdtype.Logger) (*srvdbus.Loader, error) {
	dockpath.DbusPathDock = "/org/cdc/Cdc" // TODO: improve to autodetect.

	loader := srvdbus.NewLoader(log)
	if loader == nil {
		return nil, errors.New("Dbus service failed to start")
	}

	active, e := loader.Connect()
	switch {
	case e != nil:
		return nil, e

	case !active:
		return nil, errors.New("service already active")
	}

	return loader, nil
}
