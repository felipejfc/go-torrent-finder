package finder

import (
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

type Finder struct {
	Plugins []*Plugin
	Config  *viper.Viper
}

type SearchResult struct {
	Title      string
	TorrentUrl string
	MagnetLink string
	Provider   string
}

func GetTorrentFinder() *Finder {
	f := &Finder{}
	f.loadConfig()
	SetupLogger(f.Config.GetString("logger.level"))
	f.setupProvidersPlugins()
	return f
}

func (f *Finder) loadConfig() {
	f.Config = viper.New()
	c := f.Config
	c.SetConfigFile("./config/torrent_finder.yaml")
	c.SetConfigType("yaml")
	c.SetDefault("plugins.path", "./plugins")
	c.SetDefault("logger.level", "DEBUG")
	err := c.ReadInConfig()
	if err != nil {
		Logger.Fatal(err)
	}
}

func (f *Finder) setupProvidersPlugins() {
	files, err := ioutil.ReadDir(f.Config.GetString("plugins.path"))
	f.Plugins = []*Plugin{}
	c := f.Config
	if err != nil {
		Logger.Fatalf("could not read plugins path, err: %s", err)
	}
	for _, file := range files {
		plugin := &Plugin{
			PluginName: strings.Split(file.Name(), ".")[0],
			Config:     c,
		}
		f.Plugins = append(f.Plugins, plugin)
		Logger.Debugf("loaded plugin %s", plugin.PluginName)
	}
}

func (f *Finder) SearchTorrents(searchQuery string) []SearchResult {
	for _, plugin := range f.Plugins {
		plugin.ListTorrents(searchQuery)
	}
	return []SearchResult{}
}
