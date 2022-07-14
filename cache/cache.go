package cache

type NewCacheMap map[string]string

func TokenIsInMap(token string, cacheMap map[string]string) bool {
	//iterating through the map
	for _, value := range cacheMap {

		//check if present value is equal to the token string
		if value == token {
			//if so return true
			return true
		}
	}
	//if the token was not found return false
	return false
}

func AddToCacheMap(id string, token string, cm NewCacheMap) NewCacheMap {
	_, ok := cm[id]
	if !ok {
		cm[id] = token
	}
	return cm
}

func RemoveFromCacheMap(id string, cm NewCacheMap) NewCacheMap {
	_, ok := cm[id]
	if ok {
		delete(cm, id)
	}
	return cm
}

func IdIsInMap(id string, m NewCacheMap) bool {
	_, ok := m[id]
	if ok {
		return true
	} else {
		return false
	}
}

func GiveIdToToken(token string, cm NewCacheMap) string {
	for key, value := range cm {

		if value == token {
			return key
		}
	}
	return ""
}
