package httpx

import "net/http"

func Headers(strs ...string) http.Header {
	if len(strs)%2 != 0 {
		panic("httpx: uneven strings given")
	}

	h := http.Header{}
	var skip bool
	for i := 0; i < len(strs); i++ {
		if i == len(strs) {
			break
		}

		if skip {
			skip = false
			continue
		}

		x := i + 1
		h.Set(strs[i], strs[x])
		skip = true
	}

	return h
}
