package usecase

import (
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
	"github.com/ridwanakf/url-shortener-service/internal/repo/db"
	"github.com/stretchr/testify/assert"
)

var urlInputTest = entity.URL{
	ShortURL:  "aBcDe",
	LongURL:   "http://www.google.com",
	CreatedAt: time.Now(),
	CreatedBy: "anonymous",
}

var params = config.Params{
	ShortUrlLength: 6,
	ExpireDuration: 30,
}

func TestNewShortenerUsecase(t *testing.T) {
	t.Run("given nil param, should return not nil", func(t *testing.T) {
		got := NewShortenerUsecase(nil, params.ShortUrlLength, params.ExpireDuration)
		assert.NotNil(t, got)
	})
	t.Run("given not nil param, should have it attached", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoDB := db.NewMockShortenerDB(ctrl)

		got := NewShortenerUsecase(repoDB, params.ShortUrlLength, params.ExpireDuration)
		assert.NotNil(t, got)
		assert.NotNil(t, got.db)
	})
}

func TestShortenerUsecase_CreateNewCustomShortURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *db.MockShortenerDB
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = db.NewMockShortenerDB(ctrl)
		unit = NewShortenerUsecase(repoDB, params.ShortUrlLength, params.ExpireDuration)
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
		assert.Equal(t, urlInputTest.LongURL, got.LongURL)
		assert.Equal(t, urlInputTest.ShortURL, got.ShortURL)
	})
}

func TestShortenerUsecase_CreateNewShortURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *db.MockShortenerDB
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = db.NewMockShortenerDB(ctrl)
		unit = NewShortenerUsecase(repoDB, params.ShortUrlLength, params.ExpireDuration)
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
		assert.Equal(t, urlInputTest.LongURL, got.LongURL)
	})
}

func TestShortenerUsecase_DeleteURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *db.MockShortenerDB
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = db.NewMockShortenerDB(ctrl)
		unit = NewShortenerUsecase(repoDB, params.ShortUrlLength, params.ExpireDuration)
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

	t.Run("given no error when invoking db DeleteURL, should return nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().DeleteURL(gomock.Any()).Return(nil)

		err := unit.DeleteURL(urlInputTest.ShortURL)

		assert.Nil(t, err)
		assert.NoError(t, err)
	})
}

func TestShortenerUsecase_GenerateShortURL(t *testing.T) {
	var (
		ctrl   *gomock.Controller
		repoDB *db.MockShortenerDB
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = db.NewMockShortenerDB(ctrl)
		unit = NewShortenerUsecase(repoDB, params.ShortUrlLength, params.ExpireDuration)
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
		repoDB *db.MockShortenerDB
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = db.NewMockShortenerDB(ctrl)
		unit = NewShortenerUsecase(repoDB, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("given error when invoking db GetAllURL, should return []entity.URL{} and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().GetAllURL().Return([]entity.URL{}, errors.New("there's something error"))

		got, err := unit.GetAllURL()

		assert.Error(t, err)
		assert.Equal(t, "there's something error", err.Error())
		assert.Equal(t, []entity.URL{}, got)
	})

	t.Run("given no error when invoking db GetAllURL, should return data and nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().GetAllURL().Return([]entity.URL{urlInputTest}, nil)

		got, err := unit.GetAllURL()

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
		repoDB *db.MockShortenerDB
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = db.NewMockShortenerDB(ctrl)
		unit = NewShortenerUsecase(repoDB, params.ShortUrlLength, params.ExpireDuration)
	}
	finish := func() {
		ctrl.Finish()
	}

	t.Run("if short url does not exist, should return empty string and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(false)

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "URL does not exist", err.Error())
		assert.Equal(t, "", got)
	})

	t.Run("if short URL has expired and successfully delete the url, should return empty string and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().HasShortURLExpired(gomock.Any()).Return(true)
		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().DeleteURL(gomock.Any()).Return(nil)

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "URL has expired!", err.Error())
		assert.Equal(t, "", got)
	})

	t.Run("if short URL has expired but failed to delete the url, should return empty string and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().HasShortURLExpired(gomock.Any()).Return(true)
		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().DeleteURL(gomock.Any()).Return(errors.New("there is something error"))

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "URL has expired! but failed to delete it", err.Error())
		assert.Equal(t, "", got)
	})

	t.Run("given error when invoking db GetLongURL, should return empty string and error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().HasShortURLExpired(gomock.Any()).Return(false)
		repoDB.EXPECT().GetLongURL(gomock.Any()).Return("", errors.New("there's something error"))

		got, err := unit.GetLongURL(urlInputTest.ShortURL)

		assert.Error(t, err)
		assert.Equal(t, "there's something error", err.Error())
		assert.Equal(t, "", got)
	})

	t.Run("given no error when invoking db GetLongURL, should return data and nil error", func(t *testing.T) {
		begin(t)
		defer finish()

		repoDB.EXPECT().IsShortURLExist(gomock.Any()).Return(true)
		repoDB.EXPECT().HasShortURLExpired(gomock.Any()).Return(false)
		repoDB.EXPECT().GetLongURL(gomock.Any()).Return(urlInputTest.LongURL, nil)

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
		repoDB *db.MockShortenerDB
		unit   *ShortenerUsecase
	)
	begin := func(t *testing.T) {
		ctrl = gomock.NewController(t)
		repoDB = db.NewMockShortenerDB(ctrl)
		unit = NewShortenerUsecase(repoDB, params.ShortUrlLength, params.ExpireDuration)
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

		err := unit.UpdateShortURL(urlInputTest.ShortURL, urlInputTest.LongURL)

		assert.Nil(t, err)
		assert.NoError(t, err)
	})
}
