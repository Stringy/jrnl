package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ServeCommand struct {
	name    string
	desc    string
	usage   string
	longuse string
}

var cmdServe = &ServeCommand{
	name:    "serve",
	desc:    "renders markdown entries as html and serves them",
	usage:   "",
	longuse: ``,
}

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

// defaults to port 1337 for now (later use args)
func (s *ServeCommand) Do(args []string) error {

	if _, err := os.Stat("./.jrnl"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Can't serve journal in uninitialised directory")
			return nil
		} else {
			return err
		}
	}

	exphandle := &RegexpHandler{}

	exphandle.HandleFunc(makeregex(`^/$`), serveTemplate)
	exphandle.HandleFunc(makeregex(`/[0-9]+/[0-9]+/[0-9]+/[a-zA-Z]+\.md`), serveMarkdown)
	exphandle.HandleFunc(makeregex(`/res/css/.+\.css`), serveCSS)
	exphandle.HandleFunc(makeregex(`/res/js/.+\.js`), serveJS)
	http.ListenAndServe(":1337", exphandle)
	return nil
}

func makeregex(str string) *regexp.Regexp {
	exp, err := regexp.Compile(str)
	handle(err)
	return exp
}

func serveMarkdown(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./.jrnl/res/base.html")
	handle(err)
	years := getDirStructure()
	entry, err := os.Open("." + r.URL.Path)
	handle(err)
	bmd, err := ioutil.ReadAll(entry)
	handle(err)
	md := blackfriday.MarkdownCommon(bmd)

	err = tmp.Execute(w, map[string]interface{}{
		"Years":    years,
		"Markdown": template.HTML(string(md)),
	})
	handle(err)
}

func serveJS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Header().Set("Content-Type", "text/javascript")
	http.ServeFile(w, r, "./.jrnl/"+r.URL.Path)
}

func serveCSS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "./.jrnl/"+r.URL.Path)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./.jrnl/res/base.html")
	handle(err)

	years := getDirStructure()

	err = tmp.Execute(w, map[string]interface{}{"Years": years})
	handle(err)

}

func getDirStructure() map[string]Year {
	years := make(map[string]Year)
	walk := func(pth string, info os.FileInfo, err error) error {
		if strings.HasPrefix(pth, ".jrnl") ||
			strings.HasPrefix(pth, ".") {
			return nil
		}
		if !info.IsDir() && strings.Contains(pth, "/") {
			dir, file := filepath.Split(pth)
			dir, day := filepath.Split(filepath.Clean(dir))
			dir, mon := filepath.Split(filepath.Clean(dir))
			_, year := filepath.Split(filepath.Clean(dir))

			if yr, ok := years[year]; ok {
				if mn, mok := yr.Months[mon]; mok {
					if d, dok := mn.Days[day]; dok {
						d.Entries[file] = Entry{
							Name: file,
						}
					} else {
						mn.Days[day] = Day{
							Name: day,
							Entries: map[string]Entry{
								file: {Name: file},
							},
						}
					}
				} else {
					yr.Months[mon] = Month{
						Name: mon,
						Days: map[string]Day{
							day: {Name: day, Entries: map[string]Entry{
								file: {Name: file},
							}},
						},
					}
				}
			} else {
				years[year] = Year{
					Name: year,
					Months: map[string]Month{
						mon: Month{
							Name: mon,
							Days: map[string]Day{
								day: Day{
									Name: day,
									Entries: map[string]Entry{
										file: Entry{Name: file},
									},
								},
							},
						},
					},
				}
			}
		}
		return nil
	}
	err := filepath.Walk(".", walk)
	if err != nil {
		panic(err)
	}
	return years
}

func (s *ServeCommand) Name() string    { return s.name }
func (s *ServeCommand) Desc() string    { return s.desc }
func (s *ServeCommand) Usage() string   { return s.usage }
func (s *ServeCommand) LongUse() string { return s.longuse }
