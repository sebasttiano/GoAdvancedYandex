package persistent

import (
	"errors"
	"persistent/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockStore(ctrl)

	// возвращаемая ошибка
	errEmptyKey := errors.New("Указан пустой ключ")

	// допишите код

	m.EXPECT().Get("").Return([]byte(""), errEmptyKey)
}
