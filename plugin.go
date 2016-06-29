package finder

import (
	"net/http"

	"github.com/cjoudrey/gluahttp"
	"github.com/felipejfc/gluahttpscrape"
	"github.com/layeh/gopher-json"
	"github.com/spf13/viper"
	"github.com/topfreegames/mqttbot/logger"
	"github.com/yuin/gopher-lua"
)

type Plugin struct {
	PluginName string
	Config     *viper.Viper
}

func GetPlugin(pluginName string, config *viper.Viper) *Plugin {
	p := &Plugin{
		PluginName: pluginName,
		Config:     config,
	}
	return p
}

func (p *Plugin) loadModules(L *lua.LState) {
	L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	L.PreloadModule("json", json.Loader)
	L.PreloadModule("scrape", gluahttpscrape.NewHttpScrapeModule().Loader)
}

func (p *Plugin) ListTorrents(searchQuery string) []SearchResult {
	Logger.Debug("me is here")
	L := lua.NewState()
	p.loadModules(L)
	L.DoFile(p.Config.GetString("plugins.path") + "/" + p.PluginName + ".lua")
	defer L.Close()
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("search_torrents"),
		NRet:    2,
		Protect: true,
	}, lua.LString(searchQuery)); err != nil {
		logger.Logger.Error(err)
		return []SearchResult{}
	}
	//ret := L.Get(-1)
	retErr := L.Get(-2)
	L.Pop(2)
	if retErr != nil && retErr != lua.LNil {
		logger.Logger.Error(retErr.String())
		return []SearchResult{}
	}
	return []SearchResult{}
}
