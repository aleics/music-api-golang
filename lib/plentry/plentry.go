package plentry

import (
        "fmt"
        "net/http"
        "net/url"
        "io/ioutil"
        "strings"
        "encoding/json"
        "os"
	"strconv"
	"github.com/aleics/gmusicgo/lib/gmusicjson"
)

type Plentry struct{
	Kind string `json:"kind"`
	Id string `json:"id"`
	ClientId string `json:"clientId"`
	PlaylistId string `json:"playlistId"`
	AbsolutePosition string `json:"absolutePosition"`
	TrackId string `json:"trackId"`
	CreationTimestamp string `json:"creationTimestamp"`
	LastModifiedTimestamp string `json:"lastModifiedTimestamp"`
	Deleted bool `json:"deleted"`
	Source string `json:"source"`
}
func Init() *Plentry{
        p := new(Plentry)
        return p
}
func ArrayInit() (p *[]Plentry){
        return p
}
func (p Plentry) GetKind() string {
        return p.Kind
}
func (p *Plentry) SetKind(kind string) {
        p.Kind = kind
}
func (p Plentry) GetId() string {
        return p.Id
}
func (p *Plentry) SetId(id string) {
        p.Id = id
}
func (p Plentry) GetClientId() string {
        return p.ClientId
}
func (p *Plentry) SetClientId(clientId string) {
        p.ClientId = clientId
}
func (p Plentry) GetPlaylistId() string {
        return p.PlaylistId
}
func (p *Plentry) SetPlaylistId(playlistId string) {
        p.PlaylistId = playlistId
}
func (p Plentry) GetAbsolutePosition() string {
        return p.AbsolutePosition
}
func (p *Plentry) SetAbsolutePosition(absolutePosition string) {
        p.AbsolutePosition = absolutePosition
}
func (p Plentry) GetTrackId() string {
        return p.TrackId
}
func (p *Plentry) SetTrackId(trackId string) {
        p.TrackId = trackId
}
func (p Plentry) GetCreationTimestamp() string {
        return p.CreationTimestamp
}
func (p *Plentry) SetCreationTimestamp(creationTimestamp string) {
        p.CreationTimestamp = creationTimestamp
}
func (p Plentry) GetLastModifiedTimestamp() string {
        return p.LastModifiedTimestamp
}
func (p *Plentry) SetLastModifiedTimestamp(lastModifiedTimestamp string) {
        p.LastModifiedTimestamp = lastModifiedTimestamp
}
func (p Plentry) GetDeleted() bool {
        return p.Deleted
}
func (p *Plentry) SetDeleted(deleted bool) {
        p.Deleted = deleted
}
func (p Plentry) GetSource() string {
        return p.Source
}
func (p *Plentry) SetSource(source string) {
        p.Source = source
}
func (p *Plentry) NewPlentry(kind string, id string, clientId string, playlistId string, absolutePosition string, trackId string, creationTimestamp string, lastModifiedTimestamp string, deleted bool, source string){

	p.Kind = kind
	p.Id = id
	p.ClientId = clientId
	p.PlaylistId = playlistId
	p.AbsolutePosition = absolutePosition
	p.TrackId = trackId
	p.CreationTimestamp = creationTimestamp
	p.LastModifiedTimestamp = lastModifiedTimestamp
	p.Deleted = deleted
	p.Source = source
}
func (p Plentry) ToMap() map[string]string {

        ret := make(map[string]string)

        ret["kind"] = p.Kind
        ret["id"] = p.Id
        ret["clientId"] = p.ClientId
        ret["playlistId"] = p.PlaylistId
        ret["absolutePosition"] = p.AbsolutePosition
	ret["trackId"] = p.TrackId
	ret["creationTimestamp"] = p.CreationTimestamp
	ret["lastModifiedTimestamp"] = p.LastModifiedTimestamp
	ret["deleted"] = strconv.FormatBool(p.Deleted)
        ret["source"] = p.Source
        
        return ret
}
func PlentryRequest(auth string, path string) []Plentry{

        hostname := "https://www.googleapis.com"
        resource := "/sj/v1beta1/plentries"

        u, _ := url.ParseRequestURI(hostname)
        u.Path = resource
        urlStr := fmt.Sprintf("%v",u)

        cl := &http.Client{}

        r, err := http.NewRequest("GET", urlStr, nil)
        if err != nil {
                os.Exit(1)
        }

        auth_header := "GoogleLogin auth=" + auth

        r.Header.Add("Authorization", auth_header)


        resp, err := cl.Do(r)
        if err != nil {
                os.Exit(1)
        }

        defer resp.Body.Close()

        b, err := ioutil.ReadAll(resp.Body) //Get the body of the response                                      

        if err != nil { //Error management
                os.Exit(1)
        }

        pa := []string{path,"plentries.json"}
        jsonpath := strings.Join(pa,"")
        
        var f interface{}
        json.Unmarshal(b, &f)

        m := f.(map[string]interface{})

        plentriesmap := m["data"]
        plentries := plentriesmap.(map[string]interface{})


        itemsmap := plentries["items"]
        items := itemsmap.([]interface{})

        var singleitem map[string]interface{}

        length := len(items)
        arrayplentries := make([]Plentry,length)

        for i := 0; i < length; i++ {

                singleitem = items[i].(map[string]interface{})


                arrayplentries[i].NewPlentry(singleitem["kind"].(string), singleitem["id"].(string),  singleitem["clientId"].(string), singleitem["playlistId"].(string), singleitem["absolutePosition"].(string), singleitem["trackId"].(string), singleitem["creationTimestamp"].(string), singleitem["lastModifiedTimestamp"].(string), singleitem["deleted"].(bool), singleitem["source"].(string))

        }

	_, err = gmusicjson.Export(arrayplentries, jsonpath)
        if err != nil {
                fmt.Println("Error exporting Plentries: ")
                fmt.Println(err)
        }

        return arrayplentries
}
func (p *Plentry) Print(){
        fmt.Print("kind: ")
        fmt.Println(p.Kind)
        fmt.Print("id: ")
        fmt.Println(p.Id)
	fmt.Print("clientId: ")
        fmt.Println(p.ClientId)
        fmt.Print("playlistId: ")
        fmt.Println(p.PlaylistId)
        fmt.Print("absolutePosition: ")
        fmt.Println(p.AbsolutePosition)
        fmt.Print("trackId: ")
        fmt.Println(p.TrackId)
        fmt.Print("creationTimestamp: ")
        fmt.Println(p.CreationTimestamp)
        fmt.Print("lastModifiedTimestamp: ")
        fmt.Println(p.LastModifiedTimestamp)
        fmt.Print("deleted: ")
        fmt.Println(p.Deleted)
        fmt.Print("source: ")
        fmt.Println(p.Source)
}
