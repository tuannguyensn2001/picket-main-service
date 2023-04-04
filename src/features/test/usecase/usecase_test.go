package test_usecase

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	test_repository "picket-main-service/src/features/test/repository"
	"sync"
	"testing"
)

func TestGetContent(t *testing.T) {
	db, _ := gorm.Open(postgres.Open("host=localhost port=5432 user=tuannguyen password='' dbname=picket"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	rd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	})

	testRepository := test_repository.New(db)
	usecase := New(testRepository, rd)

	testId := 1
	rd.Del(context.TODO(), fmt.Sprintf("test-content-%d", testId))
	var wg sync.WaitGroup
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			usecase.GetContent(context.TODO(), testId)
		}()
	}

	wg.Wait()

}
