package models

type Project struct {
	ID           int      `json:"id"`
	Titre        string   `json:"title"`
	DateCreation string   `json:"date"`
	Description  string   `json:"description"`
	Technologie  string   `json:"technologie"`
	Explication  string   `json:"explication"`
	Probleme     string   `json:"probleme"`
	Solution     string   `json:"solution"`
	UrlSource    string   `json:"url_source"`
	Images       []string `json:"images"`
}

type ContactUpdate struct {
	ID        string `json:"id"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Linkedin  string `json:"linkedin"`
	Github    string `json:"github"`
}

type ProjetUpdate struct {
	ID          string   `json:"id"`
	Titre       string   `json:"title"`
	Description string   `json:"description"`
	Technologie string   `json:"technologie"`
	Explication string   `json:"explication"`
	Probleme    string   `json:"probleme"`
	Solution    string   `json:"solution"`
	UrlSource   string   `json:"url_source"`
	Images      []string `json:"images"`
}

type CorbeilleEntry struct {
	ID              int      `json:"id"`
	ProjectID       int      `json:"project_id"`
	Titre           string   `json:"titre"`
	DateSuppression string   `json:"date_suppression"`
	Images          []string `json:"images"`
}

type Technologie struct {
	ID         int    `json:"id"`
	Nom        string `json:"nom"`
	Icone      string `json:"icone,omitempty"`
	Url_source string `json:"url_source"`
}

type HomeData struct {
	Projects     []map[string]interface{}
	Technologies []map[string]interface{}
}

type CorbeilleTech struct {
	ID              int    `json:"id"`
	TechID          int    `json:"tech_id"`
	Nom             string `json:"nom"`
	Icone           string `json:"icone"`
	UrlSource       string `json:"url_source"`
	DateSuppression string `json:"date_suppression"`
}
