package handlers

import (
	dto "dumbsound/dto/result"
	songdto "dumbsound/dto/song"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type handlerSong struct {
	SongRepository repositories.SongRepository
}

func HandlerSong(SongRepository repositories.SongRepository) *handlerSong {
	return &handlerSong{SongRepository}
}

func (h *handlerSong) FindSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	songs, err := h.SongRepository.FindSong()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range songs {
		imagePath := os.Getenv("PATH_FILE") + p.Image
		songs[i].Image = imagePath
	}
	for i, p := range songs {
		imagePath := os.Getenv("PATH_FILE") + p.Song
		songs[i].Song = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: songs}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSong) GetSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	song, err := h.SongRepository.GetSong(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	song.Image = os.Getenv("PATH_FILE") + song.Image
	song.Song = os.Getenv("PATH_FILE") + song.Song

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: song}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSong) CreateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	dataThumbnail := r.Context().Value("dataFile") // add this code
	fileImage := dataThumbnail.(string)            // add thImageode

	dataSong := r.Context().Value("songFile") // add this code
	fileSong := dataSong.(string)             // add this code

	year, _ := strconv.Atoi(r.FormValue("year"))
	artistId, _ := strconv.Atoi(r.FormValue("artist_id"))

	request := songdto.RequestSong{
		Title:    r.FormValue("title"),
		Year:     year,
		ArtistID: artistId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	song, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	thumbnail, err := cld.Upload.Upload(ctx, fileImage, uploader.UploadParams{Folder: "dumbsound"})
	music, err := song.Upload.Upload(ctx, fileSong, uploader.UploadParams{Folder: "dumbsound"})

	if err != nil {
		fmt.Println(err.Error())
	}

	dataMusic := models.Song{
		Title:    request.Title,
		Image:    thumbnail.SecureURL,
		Year:     request.Year,
		Song:     music.SecureURL,
		ArtistID: request.ArtistID,
	}

	data, err := h.SongRepository.CreateSong(dataMusic)
	if err != nil {
		w.Header().Set("Content-type", "aplication/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
