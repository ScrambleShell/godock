// +build NetActivity || all

package allapps

import "github.com/sqp/godock/services/NetActivity"

func init() {
	AddService("NetActivity", NetActivity.NewApplet)
}
