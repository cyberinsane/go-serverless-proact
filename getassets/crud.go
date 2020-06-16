package getassets

//Asset struct
type Asset struct {
	id             string
	name           string
	coverImagePath string
	likes          int
}

//GetAssets list
func GetAssets() []Asset {
	var assets = []Asset{
		Asset{"someId1", "someName1", "www.google.com", 10},
		Asset{"someId2", "someName2", "www.google.com", 10},
	}
	return assets
}

// GetAsset single record
func GetAsset(assetID string) Asset {
	return Asset{"someId", "someName", "www.google.com", 10}
}
