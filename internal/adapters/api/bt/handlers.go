package bt

import (
	"net/http"
	"stregy/internal/adapters/api"
	userapi "stregy/internal/adapters/api/user"
	"stregy/internal/config"
	"stregy/internal/domain/btservice"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	backtesterURL = "/api/backtester"
)

type handler struct {
	backtesterService btservice.Service
	exgAccService     exgaccount.Service
	userService       user.Service
}

func NewHandler(
	backtesterService btservice.Service,
	userService user.Service,
) api.Handler {
	return &handler{
		backtesterService: backtesterService,
		userService:       userService,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	createSEHandler := handlers.JsonHandler(h.backtestHandler, &BacktesterDTO{})
	userHandler := userapi.AuthenticationHandler(createSEHandler, h.userService)
	router.POST(backtesterURL, handlers.ToSimpleHandler(userHandler))
}

func (h *handler) backtestHandler(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
	args map[string]interface{},
) {
	cfg := config.GetConfig()
	logger := logging.GetLogger()
	// Parse and validate request.
	user := user.User{}
	mapstructure.Decode(args["user"], &user)
	apiDTO := BacktesterDTO{}
	mapstructure.Decode(args["json"], &apiDTO)

	// Create db record.
	startDate, _ := time.Parse("2006-01-02 15:04:05", apiDTO.StartDate)
	endDate, _ := time.Parse("2006-01-02 15:04:05", apiDTO.EndDate)
	domainDTO := btservice.BacktestDTO{
		StrategyName: apiDTO.StrategyName,
		SymbolName:   apiDTO.Symbol,
		TimeframeSec: apiDTO.TimeframeSec,
		StartDate:    startDate,
		EndDate:      endDate,
	}
	btDomain, err := h.backtesterService.Create(domainDTO)
	if err != nil {
		logger.Error(err.Error())
		handlers.ReturnError(w, http.StatusInternalServerError, "")
		return
	}

	// Execute.
	if !*cfg.IsDebug {
		err = h.backtesterService.Launch(btDomain)
		if err != nil {
			logger.Error(err.Error())
			handlers.ReturnError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	handlers.JsonResponseWriter(w, map[string]string{"backtest_id": btDomain.ID})
}
