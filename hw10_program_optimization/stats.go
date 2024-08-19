package hw10programoptimization

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"
)

type User struct {
	Email string `json:"email"`
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)
	const maxCapacity = 64 * 1024 // 64kb
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	user := &User{}
	stats := make(DomainStat)
	for scanner.Scan() {
		line := scanner.Bytes()
		if err := json.Unmarshal(line, user); err != nil {
			return nil, err
		}

		afterAt := strings.SplitN(user.Email, "@", 2)[1]
		if strings.Contains(afterAt, "."+domain) {
			stats[strings.ToLower(afterAt)]++
		}
	}

	return stats, nil
}
