package idl

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"io/ioutil"
	"strings"
)

type (
	pkgTDB struct {
		textures map[string]*sdl.Texture
		renderer *sdl.Renderer
	}
)

var (
	thisTDB pkgTDB
)

// ----------> Package's functions <----------

func init() {
	thisTDB.textures = make(map[string]*sdl.Texture)
}

func SetRenderer(newRenderer *sdl.Renderer) {
	thisTDB.renderer = newRenderer
}

func Load(path string) error {
	surface, err := img.LoadTexture(thisTDB.renderer, path)
	if err == nil {
		thisTDB.textures[path] = surface
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
	_, ok := thisTDB.textures[path]
	if ok == true {
		thisTDB.textures[path].Destroy()
		delete(thisTDB.textures, path)
		return true
	}

	return false
}

func UnloadFolder(pathFolder string) uint {
	nbrUnloaded := uint(0)

	if pathFolder[len(pathFolder)] != '/' {
		pathFolder = pathFolder + "/"
	}

	for pathSurface := range thisTDB.textures {
		if strings.Index(pathSurface, pathFolder) == 0 {
			Unload(pathSurface)
			nbrUnloaded++
		}
	}

	return nbrUnloaded
}

func UnloadAll() {
	for pathSurface := range thisTDB.textures {
		Unload(pathSurface)
	}
}

func GetSurface(path string) *sdl.Texture {
	return thisTDB.textures[path]
}
