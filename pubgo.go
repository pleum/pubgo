package pubgo

func NewClient(apiKey string) (*Client, error) {
	client := &Client{}
	client.apiKey = apiKey
	return client, nil
}

func checkShard(shard string) bool {
	for _, v := range shardList {
		if v == shard {
			return true
		}
	}
	return false
}
