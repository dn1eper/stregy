package backtester

import (
	"net/http"
	"os"
	"stregy/internal/adapters/api"
	userapi "stregy/internal/adapters/api/user"
	"stregy/internal/domain/backtester"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	launchBacktestURL  = "/api/backtest"
	executeBacktestURL = "/api/execute_backtest"
)

type handler struct {
	backtesterService backtester.Service
	exgAccService     exgaccount.Service
	userService       user.Service
}

func NewHandler(
	backtesterService backtester.Service,
	userService user.Service,
) api.Handler {
	return &handler{
		backtesterService: backtesterService,
		userService:       userService,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	createSEHandler := handlers.JsonHandler(h.LaunchBacktestHandler, &BacktesterDTO{})
	userHandler := userapi.AuthenticationHandler(createSEHandler, h.userService)
	router.POST(launchBacktestURL, handlers.ToSimpleHandler(userHandler))

	userHandler = userapi.AuthenticationHandler(h.RunBacktestHandler, h.userService)
	router.POST(executeBacktestURL, handlers.ToSimpleHandler(userHandler))
}

func (h *handler) LaunchBacktestHandler(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
	args map[string]interface{},
) {
	logger := logging.GetLogger()
	// Parse and validate request.
	user := user.User{}
	mapstructure.Decode(args["user"], &user)
	apiDTO := BacktesterDTO{}
	mapstructure.Decode(args["json"], &apiDTO)

	// Create db record.
	startDate, _ := time.Parse("2006-01-02 15:04:05", apiDTO.StartDate)
	endDate, _ := time.Parse("2006-01-02 15:04:05", apiDTO.EndDate)
	domainDTO := backtester.BacktestDTO{
		StrategyName:        apiDTO.StrategyName,
		Timeframe:           apiDTO.Timeframe,
		Symbol:              apiDTO.Symbol,
		StartDate:           startDate,
		EndDate:             endDate,
		HighOrderResolution: apiDTO.HighOrderResolution,
		BarsNeeded:          apiDTO.BarsNeeded,
		ATRperiod:           apiDTO.ATRperiod,
	}
	btDomain, err := h.backtesterService.Create(domainDTO)
	if err != nil {
		logger.Error(err.Error())
		handlers.ReturnError(w, http.StatusInternalServerError, "")
		return
	}
	// set fields not saved to db
	btDomain.BarsNeeded = apiDTO.BarsNeeded
	btDomain.ATRperiod = apiDTO.ATRperiod

	// Execute.
	err = h.backtesterService.Launch(btDomain)
	if err != nil {
		logger.Error(err.Error())
		handlers.ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handlers.JsonResponseWriter(w, map[string]string{"backtest_id": btDomain.Id})
}

func (h *handler) RunBacktestHandler(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
	args map[string]interface{},
) {
	h.backtesterService.Run()
	os.Exit(0)
}
