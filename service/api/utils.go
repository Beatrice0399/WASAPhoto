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

// Function that checks if the string is valid
func validStringUsername(identifier string) bool {
	var trimmedId = strings.TrimSpace(identifier)
	return len(identifier) >= 3 && len(identifier) <= 16 && trimmedId != "" && !strings.ContainsAny(trimmedId, "?_")
}

// Function that extracts the bearer token from the authorization header
func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return strings.Trim(tokens[1], " ")
	}
	return ""
}

// Function that checks if a user is logged
func isNotLogged(auth string) bool {
	return auth == ""
}

// Function that checks if the user has a valid token for the operation
func validateRequestingUser(identifier string, bearerToken string) int {
	// If the user has an invalid token then respond with a fobidden status
	if isNotLogged(bearerToken) {
		return http.StatusForbidden
	}

	//  If the user's id is different than the one in the path then respond with a unathorized status.
	if identifier != bearerToken {
		return http.StatusUnauthorized
	}
	return 0
}

// Functiopn that checks if the photo's format is valid
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
