package service

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
)

const (
	minSecretCode = 10
	maxSecretCode = 99
)

type SecureCodeService struct {
	dropCode string
	codeMu   sync.Mutex
}

func NewSecureCodeService() *SecureCodeService {
	return &SecureCodeService{}
}

func (s *SecureCodeService) GenerateDropCode() string {
	s.codeMu.Lock()
	defer s.codeMu.Unlock()

	hostID := strings.Split(getOutboundIP().String(), ".")[3]
	secretCode := strconv.Itoa(rand.Intn(maxSecretCode-minSecretCode+1) + minSecretCode)

	s.dropCode = hostID + secretCode

	return s.dropCode
}

func (s *SecureCodeService) CheckDropCode(dropCode string) bool {
	s.codeMu.Lock()
	defer s.codeMu.Unlock()

	return s.dropCode == dropCode
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
