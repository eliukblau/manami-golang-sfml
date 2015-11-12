// +build appbundle

package main

import (
	"path/filepath"

	"github.com/kardianos/osext"
)

// ResourcePath devuelve una ruta absoluta a los recursos a partir
// de una ruta relativa pasada por el programador por parametros
// (hack necesario para la compilacion de App Bundles en Mac OSX)
func ResourcePath(elem ...string) string {
	folderAbsPath, _ := osext.ExecutableFolder()
	return filepath.Join(
		append([]string{folderAbsPath, "..", "Resources"}, elem...)...)
}
