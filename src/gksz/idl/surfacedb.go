package idl

import (
	"io/ioutil"
	"strings"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

type (
	pkgSDB struct {
		surfaces map[string]*sdl.Surface
	}
)

var (
	thisSDB pkgSDB
)

// ----------> Package's functions <----------

func init() {
	thisSDB.surfaces = make(map[string]*sdl.Surface)
}

func Load(path string) error {
	surface, err := img.Load(path)
	if err == nil {
		thisSDB.surfaces[path] = surface
	}
	return err
}

func LoadFolder(path string) (uint, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}

	nbrLoaded := uint(0)
	for _, file := range files {
		if file.IsDir() == true {
			nbr, err := LoadFolder(file.Name())
			nbrLoaded = nbrLoaded + nbr
			if err != nil {
				return nbrLoaded, err
			}
		} else {
			err := Load(file.Name())
			if err != nil {
				nbrLoaded++
			}
		}
	}

	return nbrLoaded, nil
}

func Unload(path string) bool {
	_, ok := thisSDB.surfaces[path]
	if ok == true {
		thisSDB.surfaces[path].Free()
		delete(thisSDB.surfaces, path)
		return true
	}

	return false
}

func UnloadFolder(pathFolder string) uint {
	nbrUnloaded := uint(0)

	if pathFolder[len(pathFolder)] != '/' {
		pathFolder = pathFolder + "/"
	}

	for pathSurface := range thisSDB.surfaces {
		if strings.Index(pathSurface, pathFolder) == 0 {
			Unload(pathSurface)
			nbrUnloaded++
		}
	}

	return nbrUnloaded
}

func UnloadAll() {
	for pathSurface := range thisSDB.surfaces {
		Unload(pathSurface)
	}
}

func GetSurface(path string) *sdl.Surface {
	return thisSDB.surfaces[path]
}
