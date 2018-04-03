package pubgo

type filterPlayer int
type shard string
type requestEndpoint string

const (
	pubgAPIBase    string = "https://api.playbattlegrounds.com"
	statusEndpoint string = "/status"

	pubgAPIShard    string = pubgAPIBase + "/shards/%s%s"
	playersEndpoint string = "/players"
	matchesEndpoint string = "/matches"

	FilterPlayerID   filterPlayer = iota
	FilterPlayerName filterPlayer = iota
)

var (
	shardList = []string{
		"xbox-as",
		"xbox-eu",
		"xbox-na",
		"xbox-oc",
		"pc-krjp",
		"pc-na",
		"pc-eu",
		"pc-oc",
		"pc-kakao",
		"pc-sea",
		"pc-sa",
		"pc-as",
	}
)
