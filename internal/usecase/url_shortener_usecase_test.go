package usecase

import (
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
	"github.com/ridwanakf/url-shortener-service/internal/repo/cache/redis"
	"github.com/ridwanakf/url-shortener-service/internal/repo/db/postgres"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
	"github.com/stretchr/testify/assert"
)

var urlInputTest = entity.URL{
	ShortURL:  "aBcDe",
	LongURL:   "www.google.com",
	CreatedAt: time.Now(),
	CreatedBy: "anonymous",
}

var params = config.Params{
	ShortUrlLength: 6,
	ExpireDuration: 30,
}

func TestNewShortenerUsecase(t *testing.T) {
	t.Run("given nil param, should return not nil", func(t *testing.T) {
		got := NewShortenerUsecase(nil, nil, params.ShortUrlLength, params.ExpireDuration)
		assert.NotNil(t, got)
	})
	t.Run("given not nil param, should have it attached", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoDB := postgres.NewMockShortenerDBRepo(ctrl)
		repoCache := redis.NewMockShortenerCacheRepo(ctrl)

		got := NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
		assert.NotNil(t, got)
		assert.NotNil(t, got.db)
	})
}

func TestShortenerUsecase_CreateNewCustomShortURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *postgres.MockShortenerDBRepo
		repoCache *redis.MockShortenerCacheRepo
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = postgres.NewMockShortenerDBRepo(ctrl)
		repoCache = redis.NewMockShortenerCacheRepo(ctrl)
		unit = NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("if custom url already exist, should return entity.URL{} and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)

		got, err := unit.CreateNewCustomShortURL(urlInputTest.ShortURL, urlInputTest.LongURL)

		assert.Error(t, err)
		assert.Equal(t, "URL has already existed", err.Error())
		assert.Equal(t, entity.URL{}, got)
	})

	t.Run("given error when invoking db CreateNewShortURL, should return entity.URL{} and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(false)
		repoDB.EXPECT().CreateNewShortURL(gomock.Any()).Return(errors.New("there's something error"))

		got, err := unit.CreateNewCustomShortURL(urlInputTest.ShortURL, urlInputTest.LongURL)

		assert.Error(t, err)
		assert.Equal(t, "there's something error", err.Error())
		assert.Equal(t, entity.URL{}, got)
	})

	t.Run("given no error when invoking db CreateNewShortURL, should return data and nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(false)
		repoDB.EXPECT().CreateNewShortURL(gomock.Any()).Return(nil)

		got, err := unit.CreateNewCustomShortURL(urlInputTest.ShortURL, urlInputTest.LongURL)

		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotEqual(t, entity.URL{}, got)
		assert.Equal(t, "http://"+urlInputTest.LongURL, got.LongURL)
		assert.Equal(t, urlInputTest.ShortURL, got.ShortURL)
	})
}

func TestShortenerUsecase_CreateNewShortURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *postgres.MockShortenerDBRepo
		repoCache *redis.MockShortenerCacheRepo
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = postgres.NewMockShortenerDBRepo(ctrl)
		repoCache = redis.NewMockShortenerCacheRepo(ctrl)
		unit = NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("given error when invoking db CreateNewShortURL, should return entity.URL{} and error", func(t *testing.T) {
		begin(t)
		defer finish()

		//should try again until got new url that doesnt exist
		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(false)
		repoDB.EXPECT().CreateNewShortURL(gomock.Any()).Return(errors.New("there's something error"))

		got, err := unit.CreateNewShortURL(urlInputTest.LongURL)

		assert.Error(t, err)
		assert.Equal(t, "there's something error", err.Error())
		assert.Equal(t, entity.URL{}, got)
	})

	t.Run("given no error when invoking db CreateNewShortURL, should return data and nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		//should try again until got new url that doesnt exist
		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(false)
		repoDB.EXPECT().CreateNewShortURL(gomock.Any()).Return(nil)

		got, err := unit.CreateNewShortURL(urlInputTest.LongURL)

		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotEqual(t, entity.URL{}, got)
		assert.Equal(t, "http://"+urlInputTest.LongURL, got.LongURL)
	})
}

func TestShortenerUsecase_DeleteURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *postgres.MockShortenerDBRepo
		repoCache *redis.MockShortenerCacheRepo
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = postgres.NewMockShortenerDBRepo(ctrl)
		repoCache = redis.NewMockShortenerCacheRepo(ctrl)
		unit = NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("if short url does not exist, should return error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(false)

		err := unit.DeleteURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "URL does not exist", err.Error())
	})

	t.Run("given error when invoking db DeleteURL, should return error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().DeleteURL(gomock.Any()).Return(errors.New("there's something error"))

		err := unit.DeleteURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "there's something error", err.Error())
	})

	t.Run("given no error when invoking db DeleteURL and no cache, should return nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().DeleteURL(gomock.Any()).Return(nil)

		repoCache.EXPECT().IsSingleURLExist(gomock.Any()).Return(true, nil)
		repoCache.EXPECT().DeleteURL(gomock.Any()).Return(int64(1), nil)

		err := unit.DeleteURL(urlInputTest.ShortURL)

		assert.Nil(t, err)
		assert.NoError(t, err)
	})
}

func TestShortenerUsecase_GenerateShortURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *postgres.MockShortenerDBRepo
		repoCache *redis.MockShortenerCacheRepo
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = postgres.NewMockShortenerDBRepo(ctrl)
		repoCache = redis.NewMockShortenerCacheRepo(ctrl)
		unit = NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("the length of generated string must match desired shortUrlLength", func(t *testing.T) {
		begin(t)
		defer finish()

		got := unit.GenerateShortURL(params.ShortUrlLength)

		assert.Equal(t, unit.shortUrlLength, len(got))
	})
}

func TestShortenerUsecase_GetAllURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *postgres.MockShortenerDBRepo
		repoCache *redis.MockShortenerCacheRepo
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = postgres.NewMockShortenerDBRepo(ctrl)
		repoCache = redis.NewMockShortenerCacheRepo(ctrl)
		unit = NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("given error when invoking db GetAllURL, should return []entity.URL{} and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().GetAllURL("").Return([]entity.URL{}, errors.New("there's something error"))

		got, err := unit.GetAllURL("")

		assert.Error(t, err)
		assert.Equal(t, "there's something error", err.Error())
		assert.Equal(t, []entity.URL{}, got)
	})

	t.Run("given no error when invoking db GetAllURL, should return data and nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().GetAllURL("").Return([]entity.URL{urlInputTest}, nil)

		got, err := unit.GetAllURL("")

		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotEqual(t, []entity.URL{}, got)
		assert.Equal(t, urlInputTest.LongURL, got[0].LongURL)
		assert.Equal(t, urlInputTest.ShortURL, got[0].ShortURL)
		assert.Equal(t, 1, len(got))
	})
}

func TestShortenerUsecase_GetLongURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *postgres.MockShortenerDBRepo
		repoCache *redis.MockShortenerCacheRepo
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = postgres.NewMockShortenerDBRepo(ctrl)
		repoCache = redis.NewMockShortenerCacheRepo(ctrl)
		unit = NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("if short url does not exist in db, should return empty string and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoCache.EXPECT().IsSingleURLExist(gomock.Any()).Return(false, nil)

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(false)

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "URL does not exist", err.Error())
		assert.Equal(t, "", got)
	})

	t.Run("if short URL has expired in db and successfully delete the url, should return empty string and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoCache.EXPECT().IsSingleURLExist(gomock.Any()).Return(false, nil)

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().HasShortURLExpired(gomock.Any()).Return(true)
		repoDB.EXPECT().DeleteURL(gomock.Any()).Return(nil)

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "URL has expired!", err.Error())
		assert.Equal(t, "", got)
	})

	t.Run("if short URL has expired in db but failed to delete the url, should return empty string and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoCache.EXPECT().IsSingleURLExist(gomock.Any()).Return(false, nil)

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().HasShortURLExpired(gomock.Any()).Return(true)
		repoDB.EXPECT().DeleteURL(gomock.Any()).Return(errors.New("there is something error"))

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "URL has expired! but failed to delete it", err.Error())
		assert.Equal(t, "", got)
	})

	t.Run("given error when invoking db GetLongURL, should return empty string and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoCache.EXPECT().IsSingleURLExist(gomock.Any()).Return(false, nil)

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().HasShortURLExpired(gomock.Any()).Return(false)
		repoDB.EXPECT().GetURL(gomock.Any()).Return(entity.URL{}, errors.New("there's something error"))

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "there's something error", err.Error())
		assert.Equal(t, "", got)
	})

	t.Run("given no error when invoking db GetLongURL, should return data and nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoCache.EXPECT().IsSingleURLExist(gomock.Any()).Return(true, nil)
		repoCache.EXPECT().HasShortURLExpired(gomock.Any()).Return(true, nil)
		repoCache.EXPECT().DeleteURL(gomock.Any()).Return(int64(1), nil)

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().HasShortURLExpired(gomock.Any()).Return(false)
		repoDB.EXPECT().GetURL(gomock.Any()).Return(urlInputTest, nil)
		
		repoCache.EXPECT().SetURL(gomock.Any()).Return(nil)

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotEqual(t, "", got)
		assert.Equal(t, urlInputTest.LongURL, got)
	})
}

func TestShortenerUsecase_UpdateShortURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *postgres.MockShortenerDBRepo
		repoCache *redis.MockShortenerCacheRepo
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = postgres.NewMockShortenerDBRepo(ctrl)
		repoCache = redis.NewMockShortenerCacheRepo(ctrl)
		unit = NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("if short url does not exist, should return error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(false)

		err := unit.UpdateShortURL(urlInputTest.ShortURL, urlInputTest.LongURL)

		assert.Error(t, err)
		assert.Equal(t, "URL does not exist", err.Error())
	})

	t.Run("given error when invoking db UpdateShortURL, should return error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().UpdateShortURL(urlInputTest.ShortURL, urlInputTest.LongURL).Return(errors.New("there's something error"))

		err := unit.UpdateShortURL(urlInputTest.ShortURL, urlInputTest.LongURL)

		assert.Error(t, err)
		assert.Equal(t, "there's something error", err.Error())
	})

	t.Run("given no error when invoking db UpdateShortURL, should return nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().UpdateShortURL(urlInputTest.ShortURL, urlInputTest.LongURL).Return(nil)

		repoCache.EXPECT().IsSingleURLExist(gomock.Any()).Return(true, nil)
		repoCache.EXPECT().DeleteURL(gomock.Any()).Return(int64(1), nil)

		err := unit.UpdateShortURL(urlInputTest.ShortURL, urlInputTest.LongURL)

		assert.Nil(t, err)
		assert.NoError(t, err)
	})
}

func TestShortenerUsecase_IsValidURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *postgres.MockShortenerDBRepo
		repoCache *redis.MockShortenerCacheRepo
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = postgres.NewMockShortenerDBRepo(ctrl)
		repoCache = redis.NewMockShortenerCacheRepo(ctrl)
		unit = NewShortenerUsecase(repoDB, repoCache, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("if url doesnt have http or https schema, add it", func(t *testing.T) {
		begin(t)
		defer finish()

		got, err := unit.IsValidURL(urlInputTest.LongURL)

		assert.NoError(t, err)
		assert.Nil(t, err)
		assert.Equal(t, "http://"+urlInputTest.LongURL, got)
	})
}
