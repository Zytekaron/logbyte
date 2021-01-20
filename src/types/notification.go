package types

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

type Severity string

const (
	SeverityNone     Severity = "none"     // debug-like entries such as containers/doors being accessed
	SeverityLow               = "low"      // discord pings and other common, unimportant notifications
	SeverityMedium            = "medium"   // event reminders, semi-important messages, common warnings
	SeverityHigh              = "high"     // important messages, event reminders, security updates
	SeverityCritical          = "critical" // property damage, money loss, and other severe warnings
	SeverityFatal             = "fatal"    // potentially life-threatening events and warnings
)

//type jsonInt64 int64
//
//func (j *jsonInt64) MarshalJSON() ([]byte, error) {
//	str := strconv.FormatInt(int64(*j), 10)
//	return json.Marshal(str)
//}
//
//func (j *jsonInt64) UnmarshalJSON(b []byte) error {
//	i, err := strconv.ParseInt(string(b), 10, 64)
//	if err != nil {
//		return err
//	}
//	*j = jsonInt64(i)
//	return nil
//}

type Notification struct {
	// The notification id
	Id string `json:"id" bson:"_id"`

	// The service from which this notification originated
	Service string `json:"service" bson:"service"`

	// The severity of the notification
	Severity Severity `json:"severity" bson:"severity"`

	// The title of the notification
	Title string `json:"title" bson:"title"`

	// The content associated with the notification
	Content string `json:"content" bson:"content"`

	// The icon used to represent this notification
	IconURL string `json:"icon_url" bson:"icon_url"`

	// The time this notification was created
	CreatedAt int64 `json:"created_at" bson:"created_at"`
}

// Create a new Notification
func NewNotification(id string, service string, severity Severity, title string, content string, iconURL string, createdAt int64) *Notification {
	return &Notification{
		Id:        id,
		Service:   service,
		Severity:  severity,
		Title:     title,
		Content:   content,
		IconURL:   iconURL,
		CreatedAt: createdAt,
	}
}

// Return a path to the image using temporary files
//
// Fetches the image if it does not already exist
func (n *Notification) ImagePath() string {
	// look for an existing file with the image
	// specified in this notification's icon url
	sum := md5.Sum([]byte(n.IconURL))
	hash := fmt.Sprintf("%x", sum)

	imgPath := path.Join(os.TempDir(), "notif-img"+hash)
	if _, err := os.Stat(imgPath); !os.IsNotExist(err) {
		return imgPath
	}

	// otherwise make a request to the url, create
	// the file, and copy the body (image) into it
	res, err := http.Get(n.IconURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	file, err := os.Create(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return imgPath
}
