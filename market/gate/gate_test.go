package gate

import (
	"fmt"
	"testing"
	"time"
)

var cps = []string{"MICE#USDT", "1CAT#USDT", "PEPE2#USDT", "AITECH#USDT", "AINN#USDT"}

func start(s *Service, cp string) {
	ms, err := s.Subscribe(cp)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case m := <-ms:
			fmt.Printf("%+v\n", m)
		}
	}
}

func TestService_Start(t *testing.T) {
	s := NewService()
	errs, err := s.Connect()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	for i := 0; i < len(cps); i++ {
		go start(s, cps[i])

		// 等待一段时间
		time.Sleep(time.Second)
	}

	for {
		select {
		case err = <-errs:
			s.Close()
			t.Fatalf("%+v", err)
			return
		}
	}
}
