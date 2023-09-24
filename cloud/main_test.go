// Основные тесты сервера целиком
package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	main "github.com/thefrol/go-vue-recipe-blog/cloud"
)

// TODO
// аутентификацию по токену ещё бы сделать для тестирования

// Test_OpenPages тестирует основное поведение,
// что запускается основная страница, все основные страницы,
// а лишние и закрытые нет
func Test_PublicPages(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		route         string
		allowRedirect bool
		resultCode    int
	}{
		{
			name:          "Главная страница доступна",
			method:        "GET",
			route:         "/",
			allowRedirect: false,
			resultCode:    http.StatusOK,
		},
		{
			name:          "Логин страница доступна",
			method:        "GET",
			route:         "/login",
			allowRedirect: false,
			resultCode:    http.StatusOK,
		},
		{
			name:          "HTML редактирования недоступна",
			method:        "GET",
			route:         "/edit.html",
			allowRedirect: false,
			resultCode:    http.StatusNotFound,
		},
		{
			name:          "HTML логина недоступна",
			method:        "GET",
			route:         "/login.html",
			allowRedirect: false,
			resultCode:    http.StatusNotFound,
		},
		{
			name:          "Редактирование дает редирект",
			method:        "GET",
			route:         "/edit/22",
			allowRedirect: false,
			resultCode:    http.StatusFound, //todo how как проверить редирект при неавторизованности
		},
		{
			name:          "Редактирование дает редирект, и страница найдена в итоге",
			method:        "GET",
			route:         "/edit/22",
			allowRedirect: true,
			resultCode:    http.StatusOK, //todo how как проверить редирект при неавторизованности
		},
		// todo
		// ещё можно проверить куда в итоге редиректы заводяn
		//
		// а ещё можно проверить поведение, если у нас будут нужные куки
		// но я не знаю, можно ли писать такой бекдор для тестов?
	}

	// TODO
	// это тоже в сниппеты запулить
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := httptest.NewServer(main.Router) //todo тут бы адрес сервера другой поставить
			defer s.Close()

			client := resty.New()
			if !tt.allowRedirect {
				client.SetRedirectPolicy(
					resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
						return http.ErrUseLastResponse
					}),
				)
			}

			resp, err := client.R().Execute(tt.method, s.URL+tt.route)
			defer resp.RawBody().Close()

			assert.NoError(t, err)

			actualCode := resp.StatusCode()
			assert.Equalf(t, tt.resultCode, actualCode, "on route %v, Status Code expected %v, got %v", tt.route, tt.resultCode, actualCode)

		})
	}
}
