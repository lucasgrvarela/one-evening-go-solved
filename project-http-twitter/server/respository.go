package server

import "sync"

type tweetsRepository interface {
	AddTweet(t Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
	lock   sync.RWMutex
}

func (t *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.tweets, nil
}

func (t *TweetsMemoryRepository) AddTweet(tw Tweet) (int, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}
