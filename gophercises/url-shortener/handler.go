package urlshortener

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathURLs []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathURLs)
	if err != nil {
		return nil, err
	}
	pathToUrls := map[string]string{}
	for _, pu := range pathURLs {
		pathToUrls[pu.Path] = pu.URL
	}
	return MapHandler(pathToUrls, fallback), nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
