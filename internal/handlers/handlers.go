package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kothanzaw/bookings/internal/config"
	"github.com/kothanzaw/bookings/internal/models"
	"github.com/kothanzaw/bookings/internal/render"
)


var Repo *Repository
// Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates  anew repository
func NewRepo(a *config.AppConfig)*Repository{
	return &Repository{
		App:a,
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter,r *http.Request){
	remoteIP :=r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)

	render.RenderTemplate(w,r,"home.page.tmpl",&models.TemplateData{})

}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter,r *http.Request){
	


	render.RenderTemplate(w,r,"about.page.tmpl",&models.TemplateData{})	
}

///rooms/general-quarters
func (m *Repository) General(w http.ResponseWriter,r *http.Request){	
	stringMap:=make(map[string]string)
	stringMap["test"]="hello, again"

	render.RenderTemplate(w,r,"general.page.tmpl",&models.TemplateData{})	
}
///rooms/majors-suit
func (m *Repository) Majors(w http.ResponseWriter,r *http.Request){	



	render.RenderTemplate(w,r,"major.page.tmpl",&models.TemplateData{})	
}
///contact
func (m *Repository) Contact(w http.ResponseWriter,r *http.Request){	
	stringMap:=make(map[string]string)
	stringMap["test"]="hello, again"


	render.RenderTemplate(w,r,"contact.page.tmpl",&models.TemplateData{})	
}
func (m *Repository) MakeReservation(w http.ResponseWriter,r *http.Request){	
	stringMap:=make(map[string]string)
	stringMap["test"]="hello, again"


	render.RenderTemplate(w,r,"make-reservation.page.tmpl",&models.TemplateData{})	
}
func (m *Repository) Availability(w http.ResponseWriter,r *http.Request){	
	stringMap:=make(map[string]string)
	stringMap["test"]="hello, again"


	render.RenderTemplate(w,r,"search-availability.page.tmpl",&models.TemplateData{})	
}
func (m *Repository) PostAvailability(w http.ResponseWriter,r *http.Request){	

	start:=r.Form.Get("start")
	end:=r.Form.Get("end")



	
	w.Write([]byte(fmt.Sprintf("Start date is %s and End date is %s",start,end)))

	
}
type jsonResponse struct{
	OK bool `json:"ok"`
	Message string `json:"message"`
}
func (m *Repository) AvailabilityJSON(w http.ResponseWriter,r *http.Request){	
	resp :=jsonResponse{
		OK:true,
		Message: "Available",
	}
	out, err:=json.MarshalIndent(resp,"","     ")
	if err!=nil{
		log.Println(err)
	}
	
	w.Header().Set("Content-type","application/json")
	w.Write(out)
}