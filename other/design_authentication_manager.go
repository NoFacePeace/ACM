package other

type AuthenticationManager struct {
	timeToLive int
	m          map[string]int
}

func Constructor1(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive: timeToLive,
		m:          map[string]int{},
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.m[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if _, ok := this.m[tokenId]; !ok {
		return
	}
	old := this.m[tokenId]
	if old+this.timeToLive <= currentTime {
		return
	}
	this.m[tokenId] = currentTime
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	timeout := []string{}
	for k, v := range this.m {
		if v+this.timeToLive <= currentTime {
			timeout = append(timeout, k)
		}
	}
	for _, v := range timeout {
		delete(this.m, v)
	}
	return len(this.m)
}
