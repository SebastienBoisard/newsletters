package server

import (
	"fmt"
	"github.com/SebastienBoisard/newsletters/internal/server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{"filename": "config", "error": err}).Fatal("Configuration file not found")
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	log.Printf("customHTTPErrorHandler - Begin")

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	log.Printf("customHTTPErrorHandler - err=%+v", err)
	log.Printf("customHTTPErrorHandler - code=%d", code)

	errorPage := fmt.Sprintf("/home/sen/Work/newsletters/web/%d.html", code)
	log.Printf("customHTTPErrorHandler - errorPage=%s", errorPage)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func handleGetNewsletter(c echo.Context) error {

	log.Printf("handleGetNewsletter - Begin")

	newsletterShortname := c.Param("newsletter_shortname")
	episodeIDAsString := c.Param("episode_id")
	subscriberKey := c.Param("subscriber_key")
	newsletterKey := c.Param("newsletter_key")

	episodeID, err := strconv.Atoi(episodeIDAsString)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "the episode for this newsletter was not found")
	}

	log.Printf("handleGetNewsletter - newsletterShortname=%s", newsletterShortname)
	log.Printf("handleGetNewsletter - episodeID=%d", episodeID)
	log.Printf("handleGetNewsletter - subscriberKey=%s", subscriberKey)
	log.Printf("handleGetNewsletter - newsletterKey=%s", newsletterKey)

	subscription, err := models.GetSubscription(db, newsletterShortname, episodeID, newsletterKey, subscriberKey)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "the newsletter, episode or subscriber was not found")
	}

	log.Printf("handleGetNewsletter - subscription=%+v", subscription)

	blocks, err := models.GetBlocks(db, subscription.NewsletterID, episodeID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Blocks for the episode were not found")
	}

	log.Printf("handleGetNewsletter - blocks=%+v", blocks)

	episode, err := models.GetEpisode(db, subscription.NewsletterID, episodeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Episode was not found")
	}
	log.Printf("handleGetNewsletter - episode=%+v", episode)

	header, err := models.GetHeader(db, episode.HeaderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Header for the episode was not found")
	}

	footer, err := models.GetFooter(db, episode.FooterID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Footer for the episode was not found")
	}

	var content strings.Builder
	content.WriteString(header.Content)
	for _, b := range *blocks {
		content.WriteString(b.Content)
	}
	content.WriteString(footer.Content)

	return c.HTML(http.StatusOK, content.String())
}

func Run(portNumber int, releaseMode bool) {

	LoadConfig()

	var err error
	db, err = InitDatabase()
	if err != nil {
		panic(err)
	}

	// Create a new instance of Echo
	e := echo.New()

	if releaseMode {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist([]string{"www." + ProjectDomainName, ProjectDomainName}...)
		e.AutoTLSManager.Cache = autocert.DirCache(os.Getenv("BACKEND_HOME") + "/cache")
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8080", "http://localhost:8000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderCookie},
		AllowCredentials: false,
		AllowMethods:     []string{http.MethodGet},
	}))

	e.Static("/static", "web/static")

	e.GET("/:newsletter_shortname/:episode_id/:subscriber_key/:newsletter_key", handleGetNewsletter)

	// Start as a web server
	if releaseMode {
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		e.Logger.Fatal(e.Start(":" + strconv.Itoa(portNumber)))
	}
}
