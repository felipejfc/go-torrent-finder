local http = require("http")
local PROVIDER_NAME = "eztv.ag"

function get_provider_name()
  return PROVIDER_NAME
end

function search_torrents(search_query)
  local base_url = "https://eztv.ag/search/"
  search_query = string.gsub(search_query, " ", "-")
  result = http.get(base_url .. search_query)
  print(result.body)
end
