package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/huyrun/counting-mouse/src/model"
	"github.com/huyrun/counting-mouse/src/test/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_handleCountingAliveMouses(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		originalAmount     int
		lifeSpan           int
		months             int
		aliveMouseAmount   int
		expectedStatusCode int
		expectedVerdict    string
	}{
		{"success", 1, 3, 4, 13, http.StatusOK, model.VerdictSuccess},
		{"invalid_parameter", 0, 0, 0, -1, http.StatusBadRequest, model.VerdictInvalidParameters},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			app, err := newTestApp(t, ctx)
			require.NoError(t, err)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecaseMock := mock.NewMockUseCase(ctrl)
			app.Usecase = usecaseMock

			if tc.expectedVerdict == model.VerdictSuccess {
				usecaseMock.EXPECT().CountAliveMouse(gomock.Any(), tc.originalAmount, tc.lifeSpan, tc.months).Times(1).Return(tc.aliveMouseAmount)
			}

			path := fmt.Sprintf("/v1/mouses/alive?originalAmount=%d&lifeSpan=%d&months=%d", tc.originalAmount, tc.lifeSpan, tc.months)
			request, err := http.NewRequest(http.MethodGet, path, nil)
			require.NoError(t, err)

			response := httptest.NewRecorder()

			app.Engine.ServeHTTP(response, request)
			require.Equal(t, tc.expectedStatusCode, response.Code)

			var respBody struct {
				Verdict string `json:"verdict"`
				Message string `json:"message"`
				Data    struct {
					AliveMouseAmount int `json:"alive_mouse_amount"`
				}
			}
			err = json.Unmarshal(response.Body.Bytes(), &respBody)
			require.NoError(t, err)

			require.Equal(t, tc.expectedVerdict, respBody.Verdict)

			if tc.expectedVerdict == model.VerdictSuccess {
				require.Equal(t, tc.aliveMouseAmount, respBody.Data.AliveMouseAmount)
			}
		})
	}
}
