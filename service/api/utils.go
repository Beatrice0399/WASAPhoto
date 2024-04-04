package api

import (
	"errors"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

func validStringUsername(identifier string) bool {
	var trimmedId = strings.TrimSpace(identifier)
	return len(identifier) >= 3 && len(identifier) <= 16 && trimmedId != "" && !strings.ContainsAny(trimmedId, "?_")
}

/*
func (rt *_router) getProfiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := rt.db.GetAllProfiles()
	if err != nil {
		return
	}
	w.Header().Set("Content-type", "application/json")
	for _, value := range rows {
		// str := fmt.Sprintf("id: %d, Name: %s, Follower: %d, Following: %d, Photo: %d\n", value.ID, value.Name, value.Followers, value.Following, value.NumberPhotos)
		// w.Write([]byte(str))
		_ = json.NewEncoder(w).Encode(value)

	}
}
*/

func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return strings.Trim(tokens[1], " ")
	}
	return ""
}

func isNotLogged(auth string) bool {
	return auth == ""
}

func validateRequestingUser(identifier string, bearerToken string) int {
	// If the requesting user has an invalid token then respond with a fobidden status
	if isNotLogged(bearerToken) {
		return http.StatusForbidden
	}

	//  If the requesting user's id is different than the one in the path then respond with a unathorized status.
	if identifier != bearerToken {
		return http.StatusUnauthorized
	}
	return 0
}

func checkFormatPhoto(body io.ReadCloser, newReader io.ReadCloser) error {

	_, errJpg := jpeg.Decode(body)
	if errJpg != nil {

		body = newReader
		_, errPng := png.Decode(body)
		if errPng != nil {
			return errors.New(FORMAT_ERROR_IMG)
		}
		return nil
	}
	return nil
}

// Function that returns the photo folder path for a given user
func getPhotoFolder(user_id string) (UserPhotoFoldrPath string, err error) {
	// Path of the photo dir "./media/user_id/photos/"
	photoPath := filepath.Join(photoFolder, user_id, "photos")
	return photoPath, nil
}
