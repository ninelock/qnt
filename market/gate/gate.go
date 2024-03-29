package gate

import (
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"

	"qnt-robot/entity"
	"qnt-robot/utils"
)

var baseURL = "wss://api.gateio.ws/ws/v4/"

func format(cp string) string {
	return strings.Replace(cp, "#", "_", -1)
}

type Service struct {
	wg *sync.WaitGroup

	conn *websocket.Conn
	subs map[string]chan *entity.Depth

	errs chan error
	done chan struct{}
}

func NewService() *Service {
	return &Service{
		wg:   &sync.WaitGroup{},
		subs: make(map[string]chan *entity.Depth),

		errs: make(chan error, 1),
		done: make(chan struct{}),
	}
}

func (s *Service) receive() {
	defer func() {
		for _, msg := range s.subs {
			close(msg)
		}
		s.wg.Done()
	}()

	for {
		select {
		case <-s.done:
			return
		default:
			t, msg, err := s.conn.ReadMessage()
			if err != nil {
				s.errs <- err
				return
			}
			switch t {
			case websocket.PingMessage:
				_ = s.conn.WriteMessage(websocket.PongMessage, nil)
				continue
			case websocket.TextMessage:
				dp := &Depth{}
				err = easyjson.Unmarshal(msg, dp)
				if err != nil {
					continue
				}

				asks := make([][]float64, 0)
				bids := make([][]float64, 0)
				for i := 0; i < min(len(dp.Result.Asks), len(dp.Result.Bids)); i++ {
					ask := dp.Result.Asks[i]
					asks = append(asks, []float64{
						utils.ToFloat64(ask[0]),
						utils.ToFloat64(ask[1]),
					})

					bid := dp.Result.Bids[i]
					bids = append(bids, []float64{
						utils.ToFloat64(bid[0]),
						utils.ToFloat64(bid[1]),
					})
				}

				// 推送数据
				cp := dp.Result.S
				sub, ok := s.subs[cp]
				if ok {
					sub <- &entity.Depth{
						Ts:   dp.TimeMs,
						Cp:   cp,
						Asks: asks,
						Bids: bids,
					}
				}
			}
		}
	}
}

func (s *Service) keepalive() {

}

func (s *Service) Close() {
	close(s.done)
	_ = s.conn.Close()
	s.wg.Wait()
}

func (s *Service) Connect() (<-chan error, error) {
	dia := websocket.DefaultDialer
	conn, _, err := dia.Dial(baseURL, nil)
	if err != nil {
		return s.errs, err
	}
	s.conn = conn

	// 有些情况需要读取一次数据
	for {
		break
	}

	s.wg.Add(1)
	go s.receive()

	return s.errs, nil
}

func (s *Service) Subscribe(cp string) (<-chan *entity.Depth, error) {
	cp = format(cp) // 格式化货币对
	msg := make(chan *entity.Depth, 2048)

	// 添加订阅
	s.subs[cp] = msg
	err := s.conn.WriteJSON(map[string]any{
		"time":    time.Now().Second(),
		"event":   "subscribe",
		"channel": "spot.order_book",
		"payload": []string{cp, "10", "100ms"},
	})
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (s *Service) Unsubscribe(cps ...string) error {
	for _, cp := range cps {
		cp = format(cp)
		err := s.conn.WriteJSON(map[string]any{
			"time":    time.Now().Second(),
			"event":   "unsubscribe",
			"channel": "spot.order_book",
			"payload": []string{cp, "10", "100ms"},
		})
		if err != nil {
			return err
		}

		// 删除订阅
		sub, ok := s.subs[cp]
		if ok {
			close(sub)
		}
	}
	return nil
}
