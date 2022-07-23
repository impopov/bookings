package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/ivanpopov/bookings/internal/models"
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MainChan      chan models.MailData
}
